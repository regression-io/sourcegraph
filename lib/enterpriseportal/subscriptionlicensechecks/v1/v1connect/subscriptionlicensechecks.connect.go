// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: subscriptionlicensechecks.proto

package v1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/sourcegraph/sourcegraph/lib/enterpriseportal/subscriptionlicensechecks/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// SubscriptionLicenseChecksServiceName is the fully-qualified name of the
	// SubscriptionLicenseChecksService service.
	SubscriptionLicenseChecksServiceName = "enterpriseportal.subscriptionlicensechecks.v1.SubscriptionLicenseChecksService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SubscriptionLicenseChecksServiceCheckLicenseKeyProcedure is the fully-qualified name of the
	// SubscriptionLicenseChecksService's CheckLicenseKey RPC.
	SubscriptionLicenseChecksServiceCheckLicenseKeyProcedure = "/enterpriseportal.subscriptionlicensechecks.v1.SubscriptionLicenseChecksService/CheckLicenseKey"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	subscriptionLicenseChecksServiceServiceDescriptor               = v1.File_subscriptionlicensechecks_proto.Services().ByName("SubscriptionLicenseChecksService")
	subscriptionLicenseChecksServiceCheckLicenseKeyMethodDescriptor = subscriptionLicenseChecksServiceServiceDescriptor.Methods().ByName("CheckLicenseKey")
)

// SubscriptionLicenseChecksServiceClient is a client for the
// enterpriseportal.subscriptionlicensechecks.v1.SubscriptionLicenseChecksService service.
type SubscriptionLicenseChecksServiceClient interface {
	// CheckLicenseKey returns the status of a license key.
	CheckLicenseKey(context.Context, *connect.Request[v1.CheckLicenseKeyRequest]) (*connect.Response[v1.CheckLicenseKeyResponse], error)
}

// NewSubscriptionLicenseChecksServiceClient constructs a client for the
// enterpriseportal.subscriptionlicensechecks.v1.SubscriptionLicenseChecksService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSubscriptionLicenseChecksServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) SubscriptionLicenseChecksServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &subscriptionLicenseChecksServiceClient{
		checkLicenseKey: connect.NewClient[v1.CheckLicenseKeyRequest, v1.CheckLicenseKeyResponse](
			httpClient,
			baseURL+SubscriptionLicenseChecksServiceCheckLicenseKeyProcedure,
			connect.WithSchema(subscriptionLicenseChecksServiceCheckLicenseKeyMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// subscriptionLicenseChecksServiceClient implements SubscriptionLicenseChecksServiceClient.
type subscriptionLicenseChecksServiceClient struct {
	checkLicenseKey *connect.Client[v1.CheckLicenseKeyRequest, v1.CheckLicenseKeyResponse]
}

// CheckLicenseKey calls
// enterpriseportal.subscriptionlicensechecks.v1.SubscriptionLicenseChecksService.CheckLicenseKey.
func (c *subscriptionLicenseChecksServiceClient) CheckLicenseKey(ctx context.Context, req *connect.Request[v1.CheckLicenseKeyRequest]) (*connect.Response[v1.CheckLicenseKeyResponse], error) {
	return c.checkLicenseKey.CallUnary(ctx, req)
}

// SubscriptionLicenseChecksServiceHandler is an implementation of the
// enterpriseportal.subscriptionlicensechecks.v1.SubscriptionLicenseChecksService service.
type SubscriptionLicenseChecksServiceHandler interface {
	// CheckLicenseKey returns the status of a license key.
	CheckLicenseKey(context.Context, *connect.Request[v1.CheckLicenseKeyRequest]) (*connect.Response[v1.CheckLicenseKeyResponse], error)
}

// NewSubscriptionLicenseChecksServiceHandler builds an HTTP handler from the service
// implementation. It returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSubscriptionLicenseChecksServiceHandler(svc SubscriptionLicenseChecksServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	subscriptionLicenseChecksServiceCheckLicenseKeyHandler := connect.NewUnaryHandler(
		SubscriptionLicenseChecksServiceCheckLicenseKeyProcedure,
		svc.CheckLicenseKey,
		connect.WithSchema(subscriptionLicenseChecksServiceCheckLicenseKeyMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/enterpriseportal.subscriptionlicensechecks.v1.SubscriptionLicenseChecksService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SubscriptionLicenseChecksServiceCheckLicenseKeyProcedure:
			subscriptionLicenseChecksServiceCheckLicenseKeyHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSubscriptionLicenseChecksServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSubscriptionLicenseChecksServiceHandler struct{}

func (UnimplementedSubscriptionLicenseChecksServiceHandler) CheckLicenseKey(context.Context, *connect.Request[v1.CheckLicenseKeyRequest]) (*connect.Response[v1.CheckLicenseKeyResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("enterpriseportal.subscriptionlicensechecks.v1.SubscriptionLicenseChecksService.CheckLicenseKey is not implemented"))
}