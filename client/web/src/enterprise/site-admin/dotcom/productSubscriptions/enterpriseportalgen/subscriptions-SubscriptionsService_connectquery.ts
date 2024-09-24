// @generated by protoc-gen-connect-query v1.4.1 with parameter "target=ts"
// @generated from file subscriptions.proto (package enterpriseportal.subscriptions.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { MethodIdempotency, MethodKind } from "@bufbuild/protobuf";
import { ArchiveEnterpriseSubscriptionRequest, ArchiveEnterpriseSubscriptionResponse, CreateEnterpriseSubscriptionLicenseRequest, CreateEnterpriseSubscriptionLicenseResponse, CreateEnterpriseSubscriptionRequest, CreateEnterpriseSubscriptionResponse, GetEnterpriseSubscriptionRequest, GetEnterpriseSubscriptionResponse, ListEnterpriseSubscriptionLicensesRequest, ListEnterpriseSubscriptionLicensesResponse, ListEnterpriseSubscriptionsRequest, ListEnterpriseSubscriptionsResponse, RevokeEnterpriseSubscriptionLicenseRequest, RevokeEnterpriseSubscriptionLicenseResponse, UpdateEnterpriseSubscriptionMembershipRequest, UpdateEnterpriseSubscriptionMembershipResponse, UpdateEnterpriseSubscriptionRequest, UpdateEnterpriseSubscriptionResponse } from "./subscriptions_pb.js";

/**
 * GetEnterpriseSubscription retrieves an exact match on an Enterprise subscription.
 *
 * @generated from rpc enterpriseportal.subscriptions.v1.SubscriptionsService.GetEnterpriseSubscription
 */
export const getEnterpriseSubscription = {
  localName: "getEnterpriseSubscription",
  name: "GetEnterpriseSubscription",
  kind: MethodKind.Unary,
  I: GetEnterpriseSubscriptionRequest,
  O: GetEnterpriseSubscriptionResponse,
      idempotency: MethodIdempotency.NoSideEffects,
  service: {
    typeName: "enterpriseportal.subscriptions.v1.SubscriptionsService"
  }
} as const;

/**
 * ListEnterpriseSubscriptions queries for Enterprise subscriptions.
 *
 * @generated from rpc enterpriseportal.subscriptions.v1.SubscriptionsService.ListEnterpriseSubscriptions
 */
export const listEnterpriseSubscriptions = {
  localName: "listEnterpriseSubscriptions",
  name: "ListEnterpriseSubscriptions",
  kind: MethodKind.Unary,
  I: ListEnterpriseSubscriptionsRequest,
  O: ListEnterpriseSubscriptionsResponse,
      idempotency: MethodIdempotency.NoSideEffects,
  service: {
    typeName: "enterpriseportal.subscriptions.v1.SubscriptionsService"
  }
} as const;

/**
 * ListEnterpriseSubscriptionLicenses queries for licenses associated with
 * Enterprise subscription licenses, with the ability to list licenses across
 * all subscriptions, or just a specific subscription.
 *
 * Each subscription owns a collection of licenses, typically a series of
 * licenses with the most recent one being a subscription's active license.
 *
 * @generated from rpc enterpriseportal.subscriptions.v1.SubscriptionsService.ListEnterpriseSubscriptionLicenses
 */
export const listEnterpriseSubscriptionLicenses = {
  localName: "listEnterpriseSubscriptionLicenses",
  name: "ListEnterpriseSubscriptionLicenses",
  kind: MethodKind.Unary,
  I: ListEnterpriseSubscriptionLicensesRequest,
  O: ListEnterpriseSubscriptionLicensesResponse,
      idempotency: MethodIdempotency.NoSideEffects,
  service: {
    typeName: "enterpriseportal.subscriptions.v1.SubscriptionsService"
  }
} as const;

/**
 * CreateEnterpriseSubscription creates license for an Enterprise subscription.
 *
 * Not idempotent - we could implement https://google.aip.dev/155 for
 * optional idempotency in the future.
 *
 * @generated from rpc enterpriseportal.subscriptions.v1.SubscriptionsService.CreateEnterpriseSubscriptionLicense
 */
export const createEnterpriseSubscriptionLicense = {
  localName: "createEnterpriseSubscriptionLicense",
  name: "CreateEnterpriseSubscriptionLicense",
  kind: MethodKind.Unary,
  I: CreateEnterpriseSubscriptionLicenseRequest,
  O: CreateEnterpriseSubscriptionLicenseResponse,
  service: {
    typeName: "enterpriseportal.subscriptions.v1.SubscriptionsService"
  }
} as const;

/**
 * RevokeEnterpriseSubscriptionLicense revokes an existing license for an
 * Enterprise subscription, permanently disabling its use for features
 * managed by Sourcegraph. Revocation cannot be undone.
 *
 * @generated from rpc enterpriseportal.subscriptions.v1.SubscriptionsService.RevokeEnterpriseSubscriptionLicense
 */
export const revokeEnterpriseSubscriptionLicense = {
  localName: "revokeEnterpriseSubscriptionLicense",
  name: "RevokeEnterpriseSubscriptionLicense",
  kind: MethodKind.Unary,
  I: RevokeEnterpriseSubscriptionLicenseRequest,
  O: RevokeEnterpriseSubscriptionLicenseResponse,
      idempotency: MethodIdempotency.Idempotent,
  service: {
    typeName: "enterpriseportal.subscriptions.v1.SubscriptionsService"
  }
} as const;

/**
 * UpdateEnterpriseSubscription updates an existing enterprise subscription.
 *
 * @generated from rpc enterpriseportal.subscriptions.v1.SubscriptionsService.UpdateEnterpriseSubscription
 */
export const updateEnterpriseSubscription = {
  localName: "updateEnterpriseSubscription",
  name: "UpdateEnterpriseSubscription",
  kind: MethodKind.Unary,
  I: UpdateEnterpriseSubscriptionRequest,
  O: UpdateEnterpriseSubscriptionResponse,
      idempotency: MethodIdempotency.Idempotent,
  service: {
    typeName: "enterpriseportal.subscriptions.v1.SubscriptionsService"
  }
} as const;

/**
 * ArchiveEnterpriseSubscriptionRequest archives an existing Enterprise
 * subscription. This is a permanent operation, and cannot be undone.
 *
 * Archiving a subscription also immediately and permanently revokes all
 * associated licenses.
 *
 * @generated from rpc enterpriseportal.subscriptions.v1.SubscriptionsService.ArchiveEnterpriseSubscription
 */
export const archiveEnterpriseSubscription = {
  localName: "archiveEnterpriseSubscription",
  name: "ArchiveEnterpriseSubscription",
  kind: MethodKind.Unary,
  I: ArchiveEnterpriseSubscriptionRequest,
  O: ArchiveEnterpriseSubscriptionResponse,
      idempotency: MethodIdempotency.Idempotent,
  service: {
    typeName: "enterpriseportal.subscriptions.v1.SubscriptionsService"
  }
} as const;

/**
 * CreateEnterpriseSubscription creates an Enterprise subscription.
 *
 * Not idempotent - we could implement https://google.aip.dev/155 for
 * optional idempotency in the future.
 *
 * @generated from rpc enterpriseportal.subscriptions.v1.SubscriptionsService.CreateEnterpriseSubscription
 */
export const createEnterpriseSubscription = {
  localName: "createEnterpriseSubscription",
  name: "CreateEnterpriseSubscription",
  kind: MethodKind.Unary,
  I: CreateEnterpriseSubscriptionRequest,
  O: CreateEnterpriseSubscriptionResponse,
  service: {
    typeName: "enterpriseportal.subscriptions.v1.SubscriptionsService"
  }
} as const;

/**
 * UpdateEnterpriseSubscriptionMembership updates an enterprise subscription
 * membership in an authoritative manner.
 *
 * @generated from rpc enterpriseportal.subscriptions.v1.SubscriptionsService.UpdateEnterpriseSubscriptionMembership
 */
export const updateEnterpriseSubscriptionMembership = {
  localName: "updateEnterpriseSubscriptionMembership",
  name: "UpdateEnterpriseSubscriptionMembership",
  kind: MethodKind.Unary,
  I: UpdateEnterpriseSubscriptionMembershipRequest,
  O: UpdateEnterpriseSubscriptionMembershipResponse,
      idempotency: MethodIdempotency.Idempotent,
  service: {
    typeName: "enterpriseportal.subscriptions.v1.SubscriptionsService"
  }
} as const;