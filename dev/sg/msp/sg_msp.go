//go:build msp
// +build msp

package msp

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/sourcegraph/sourcegraph/dev/managedservicesplatform"
	"github.com/sourcegraph/sourcegraph/dev/managedservicesplatform/googlesecretsmanager"
	"github.com/sourcegraph/sourcegraph/dev/managedservicesplatform/spec"
	"github.com/sourcegraph/sourcegraph/dev/managedservicesplatform/terraformcloud"
	"github.com/sourcegraph/sourcegraph/dev/sg/internal/secrets"
	"github.com/sourcegraph/sourcegraph/dev/sg/internal/std"
	"github.com/sourcegraph/sourcegraph/dev/sg/msp/schema"
	"github.com/sourcegraph/sourcegraph/lib/errors"
	"github.com/sourcegraph/sourcegraph/lib/output"
	"github.com/sourcegraph/sourcegraph/lib/pointers"
)

// This file is only built when '-tags=msp' is passed to go build while 'sg msp'
// is experimental, as the introduction of this command currently increases the
// the binary size of 'sg' by ~20%.
//
// To install a variant of 'sg' with 'sg msp' enabled, run:
//
//   go build -tags=msp -o=./sg ./dev/sg && ./sg install -f -p=false
//
// To work with msp in VS Code, add the following to your VS Code configuration:
//
//  "gopls": {
//     "build.buildFlags": ["-tags=msp"]
//  }

func init() {
	// Override no-op implementation with our real implementation.
	Command.Hidden = false
	Command.Action = nil
	Command.Subcommands = []*cli.Command{
		{
			Name:        "init",
			ArgsUsage:   "<service ID>",
			Description: "Initialize a template Managed Services Platform service spec",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "output",
					Aliases: []string{"o"},
					Usage:   "Output directory for generated spec file",
					Value:   "services",
				},
			},
			Action: func(c *cli.Context) error {
				if c.Args().Len() != 1 {
					return errors.New("exactly 1 argument required: service ID")
				}
				exampleSpec, err := (spec.Spec{
					Service: spec.ServiceSpec{
						ID: c.Args().First(),
					},
					Build: spec.BuildSpec{
						Image: "index.docker.io/sourcegraph/" + c.Args().First(),
					},
					Environments: []spec.EnvironmentSpec{
						{
							ID: "dev",
							// For dev deployment, specify category 'test'.
							Category: pointers.Ptr(spec.EnvironmentCategoryTest),

							Deploy: spec.EnvironmentDeploySpec{
								Type: "manual",
								Manual: &spec.EnvironmentDeployManualSpec{
									Tag: "insiders",
								},
							},
							Domain: spec.EnvironmentDomainSpec{
								Type: "cloudflare",
								Cloudflare: &spec.EnvironmentDomainCloudflareSpec{
									Subdomain: c.Args().First(),
									Zone:      "sgdev.org",
									Required:  false,
								},
							},
							Instances: spec.EnvironmentInstancesSpec{
								Resources: spec.EnvironmentInstancesResourcesSpec{
									CPU:    1,
									Memory: "512Mi",
								},
								Scaling: spec.EnvironmentInstancesScalingSpec{
									MaxCount: pointers.Ptr(1),
								},
							},
							StatupProbe: &spec.EnvironmentStartupProbeSpec{
								// Disable startup probes by default, as it is
								// prone to causing the entire initial Terraform
								// apply to fail.
								Disabled: pointers.Ptr(true),
							},
							Env: map[string]string{
								"SRC_LOG_LEVEL": "info",
							},
						},
					},
				}).MarshalYAML()
				if err != nil {
					return err
				}

				outputPath := filepath.Join(
					c.String("output"), c.Args().First(), "service.yaml")

				_ = os.MkdirAll(filepath.Dir(outputPath), 0755)
				if err := os.WriteFile(outputPath, exampleSpec, 0644); err != nil {
					return err
				}

				std.Out.WriteSuccessf("Rendered template spec in %s", outputPath)
				return nil
			},
		},
		{
			Name:        "generate",
			ArgsUsage:   "<service spec file> <environment ID>",
			Description: "Generate Terraform assets for a Managed Services Platform service spec.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "output",
					Aliases: []string{"o"},
					Usage:   "Output directory for generated Terraform assets, relative to service spec",
					Value:   "terraform",
				},
				&cli.BoolFlag{
					Name:  "tfc",
					Usage: "Generate infrastructure stacks with Terraform Cloud backends",
					Value: true,
				},
			},
			Action: func(c *cli.Context) error {
				if c.Args().Len() != 2 {
					return errors.New("exactly 2 arguments required: service spec file and environment ID")
				}

				// Load specification
				serviceSpecPath, err := getYAMLPathArg(c, 0)
				if err != nil {
					return err
				}
				serviceSpecData, err := os.ReadFile(serviceSpecPath)
				if err != nil {
					return err
				}
				service, err := spec.Parse(serviceSpecData)
				if err != nil {
					return err
				}
				deployEnv := service.GetEnvironment(c.Args().Get(1))
				if deployEnv == nil {
					return errors.Newf("environment %q not found in service spec", c.Args().Get(1))
				}

				renderer := managedservicesplatform.Renderer{
					OutputDir: filepath.Join(filepath.Dir(serviceSpecPath), c.String("output"), deployEnv.ID),
					GCP:       managedservicesplatform.GCPOptions{},
					TFC: managedservicesplatform.TerraformCloudOptions{
						Enabled: c.Bool("tfc"),
					},
				}

				// CDKTF needs the output dir to exist ahead of time, even for
				// rendering. If it doesn't exist yet, create it
				if f, err := os.Lstat(renderer.OutputDir); err != nil {
					if !os.IsNotExist(err) {
						return errors.Wrap(err, "check output directory")
					}
					if err := os.MkdirAll(renderer.OutputDir, 0755); err != nil {
						return errors.Wrap(err, "prepare output directory")
					}
				} else if !f.IsDir() {
					return errors.Newf("output directory %q is not a directory", renderer.OutputDir)
				}

				// Render environment
				cdktf, err := renderer.RenderEnvironment(service.Service, service.Build, *deployEnv)
				if err != nil {
					return err
				}

				pending := std.Out.Pending(output.Styledf(output.StylePending,
					"Generating Terraform assets in %q...", renderer.OutputDir))
				if err := cdktf.Synthesize(); err != nil {
					pending.Destroy()
					return err
				}
				pending.Complete(
					output.Styledf(output.StyleSuccess, "Terraform assets generated in %q!", renderer.OutputDir))
				return nil
			},
		},
		{
			Name:        "terraform-cloud",
			Aliases:     []string{"tfc"},
			Description: "Manage Terraform Cloud workspaces for a service",
			Subcommands: []*cli.Command{
				{
					Name:        "sync",
					Description: "Create or update all required Terraform Cloud workspaces for a service",
					ArgsUsage:   "<service spec file>",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "workspace-run-mode",
							Usage: "One of 'vcs', 'cli'",
							Value: "vcs",
						},
						&cli.BoolFlag{
							Name:  "delete",
							Usage: "Delete workspaces and projects - does NOT apply a teardown run",
							Value: false,
						},
					},
					Action: func(c *cli.Context) error {
						serviceSpecPath, err := getYAMLPathArg(c, 0)
						if err != nil {
							return err
						}
						serviceSpecData, err := os.ReadFile(serviceSpecPath)
						if err != nil {
							return err
						}
						service, err := spec.Parse(serviceSpecData)
						if err != nil {
							return err
						}

						secretStore, err := secrets.FromContext(c.Context)
						if err != nil {
							return err
						}
						tfcAccessToken, err := secretStore.GetExternal(c.Context, secrets.ExternalSecret{
							Name:    googlesecretsmanager.SecretTFCAccessToken,
							Project: googlesecretsmanager.ProjectID,
						})
						if err != nil {
							return errors.Wrap(err, "get AccessToken")
						}
						tfcOAuthClient, err := secretStore.GetExternal(c.Context, secrets.ExternalSecret{
							Name:    googlesecretsmanager.SecretTFCOAuthClientID,
							Project: googlesecretsmanager.ProjectID,
						})
						if err != nil {
							return errors.Wrap(err, "get TFC OAuth client ID")
						}

						tfcClient, err := terraformcloud.NewClient(tfcAccessToken, tfcOAuthClient,
							terraformcloud.WorkspaceConfig{
								RunMode: terraformcloud.WorkspaceRunMode(c.String("workspace-run-mode")),
							})
						if err != nil {
							return errors.Wrap(err, "init Terraform Cloud client")
						}

						for _, deployEnv := range service.Environments {
							if os.TempDir() == "" {
								return errors.New("no temp dir available")
							}
							renderer := &managedservicesplatform.Renderer{
								// Even though we're not synthesizing we still
								// need an output dir or CDKTF will not work
								OutputDir: filepath.Join(os.TempDir(), fmt.Sprintf("msp-tfc-%s-%s-%d",
									service.Service.ID, deployEnv.ID, time.Now().Unix())),
								GCP: managedservicesplatform.GCPOptions{},
								TFC: managedservicesplatform.TerraformCloudOptions{},
							}
							defer os.RemoveAll(renderer.OutputDir)

							cdktf, err := renderer.RenderEnvironment(service.Service, service.Build, deployEnv)
							if err != nil {
								return err
							}

							if c.Bool("delete") {
								std.Out.Promptf("Deleting workspaces for environment %q - are you sure? (y/N) ", deployEnv.ID)
								var input string
								if _, err := fmt.Scan(&input); err != nil {
									return err
								}
								if input != "y" {
									return errors.New("aborting")
								}

								if err := tfcClient.DeleteWorkspaces(c.Context, service.Service, deployEnv, cdktf.Stacks()); err != nil {
									return err
								}
							}

							workspaces, err := tfcClient.SyncWorkspaces(c.Context, service.Service, deployEnv, cdktf.Stacks())
							if err != nil {
								return errors.Wrapf(err, "env %q: sync Terraform Cloud workspace", deployEnv.ID)
							}
							std.Out.WriteSuccessf("Prepared Terraform Cloud workspaces for environment %q", deployEnv.ID)
							var summary strings.Builder
							for _, ws := range workspaces {
								summary.WriteString(fmt.Sprintf("- %s: %s", ws.Name, ws.URL()))
								if ws.Created {
									summary.WriteString(" (created)")
								} else {
									summary.WriteString(" (updated)")
								}
								summary.WriteString("\n")
							}
							std.Out.WriteMarkdown(summary.String())
						}

						return nil
					},
				},
			},
		},
		{
			Name:        "schema",
			Description: "Generate JSON schema definition for service specification",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "output",
					Aliases: []string{"o"},
					Usage:   "Output path for generated schema",
				},
			},
			Action: func(c *cli.Context) error {
				jsonSchema, err := schema.Render()
				if err != nil {
					return err
				}
				if output := c.String("output"); output != "" {
					_ = os.Remove(output)
					if err := os.WriteFile(output, jsonSchema, 0644); err != nil {
						return err
					}
					std.Out.WriteSuccessf("Rendered service spec JSON schema in %s", output)
					return nil
				}
				// Otherwise render it for reader
				return std.Out.WriteCode("json", string(jsonSchema))
			},
		},
	}
}

func getYAMLPathArg(c *cli.Context, n int) (string, error) {
	v := c.Args().Get(n)
	if strings.HasSuffix(v, ".yaml") || strings.HasSuffix(v, ".yml") {
		return v, nil
	}
	return v, errors.Newf("expected argument %d %q to be a path to a YAML file", n, v)
}