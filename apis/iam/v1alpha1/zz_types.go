/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
)

// +kubebuilder:skipversion
type AccessDetail struct {
	LastAuthenticatedTime *metav1.Time `json:"lastAuthenticatedTime,omitempty"`

	Region *string `json:"region,omitempty"`
}

// +kubebuilder:skipversion
type AccessKey struct {
	CreateDate *metav1.Time `json:"createDate,omitempty"`
}

// +kubebuilder:skipversion
type AccessKeyLastUsed struct {
	LastUsedDate *metav1.Time `json:"lastUsedDate,omitempty"`

	Region *string `json:"region,omitempty"`

	ServiceName *string `json:"serviceName,omitempty"`
}

// +kubebuilder:skipversion
type AccessKeyMetadata struct {
	CreateDate *metav1.Time `json:"createDate,omitempty"`
}

// +kubebuilder:skipversion
type AttachedPermissionsBoundary struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	PermissionsBoundaryARN *string `json:"permissionsBoundaryARN,omitempty"`

	PermissionsBoundaryType *string `json:"permissionsBoundaryType,omitempty"`
}

// +kubebuilder:skipversion
type AttachedPolicy struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	PolicyARN *string `json:"policyARN,omitempty"`
}

// +kubebuilder:skipversion
type EntityDetails struct {
	LastAuthenticated *metav1.Time `json:"lastAuthenticated,omitempty"`
}

// +kubebuilder:skipversion
type EntityInfo struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	ARN *string `json:"arn,omitempty"`

	ID *string `json:"id,omitempty"`

	Path *string `json:"path,omitempty"`
}

// +kubebuilder:skipversion
type ErrorDetails struct {
	Code *string `json:"code,omitempty"`

	Message *string `json:"message,omitempty"`
}

// +kubebuilder:skipversion
type Group struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	ARN *string `json:"arn,omitempty"`

	CreateDate *metav1.Time `json:"createDate,omitempty"`

	GroupID *string `json:"groupID,omitempty"`

	GroupName *string `json:"groupName,omitempty"`

	Path *string `json:"path,omitempty"`
}

// +kubebuilder:skipversion
type GroupDetail struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	ARN *string `json:"arn,omitempty"`

	CreateDate *metav1.Time `json:"createDate,omitempty"`

	GroupID *string `json:"groupID,omitempty"`

	GroupName *string `json:"groupName,omitempty"`

	Path *string `json:"path,omitempty"`
}

// +kubebuilder:skipversion
type InstanceProfile_SDK struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	ARN *string `json:"arn,omitempty"`

	CreateDate *metav1.Time `json:"createDate,omitempty"`

	InstanceProfileID *string `json:"instanceProfileID,omitempty"`

	InstanceProfileName *string `json:"instanceProfileName,omitempty"`

	Path *string `json:"path,omitempty"`
	// Contains a list of IAM roles.
	//
	// This data type is used as a response element in the ListRoles operation.
	Roles []*Role `json:"roles,omitempty"`

	Tags []*Tag `json:"tags,omitempty"`
}

// +kubebuilder:skipversion
type LoginProfile struct {
	CreateDate *metav1.Time `json:"createDate,omitempty"`

	PasswordResetRequired *bool `json:"passwordResetRequired,omitempty"`
}

// +kubebuilder:skipversion
type MFADevice struct {
	EnableDate *metav1.Time `json:"enableDate,omitempty"`
}

// +kubebuilder:skipversion
type ManagedPolicyDetail struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	ARN *string `json:"arn,omitempty"`

	CreateDate *metav1.Time `json:"createDate,omitempty"`

	IsAttachable *bool `json:"isAttachable,omitempty"`

	PolicyID *string `json:"policyID,omitempty"`

	UpdateDate *metav1.Time `json:"updateDate,omitempty"`
}

// +kubebuilder:skipversion
type OpenIDConnectProviderListEntry struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	ARN *string `json:"arn,omitempty"`
}

// +kubebuilder:skipversion
type OrganizationsDecisionDetail struct {
	AllowedByOrganizations *bool `json:"allowedByOrganizations,omitempty"`
}

// +kubebuilder:skipversion
type PasswordPolicy struct {
	AllowUsersToChangePassword *bool `json:"allowUsersToChangePassword,omitempty"`

	ExpirePasswords *bool `json:"expirePasswords,omitempty"`

	RequireLowercaseCharacters *bool `json:"requireLowercaseCharacters,omitempty"`

	RequireNumbers *bool `json:"requireNumbers,omitempty"`

	RequireSymbols *bool `json:"requireSymbols,omitempty"`

	RequireUppercaseCharacters *bool `json:"requireUppercaseCharacters,omitempty"`
}

// +kubebuilder:skipversion
type PermissionsBoundaryDecisionDetail struct {
	AllowedByPermissionsBoundary *bool `json:"allowedByPermissionsBoundary,omitempty"`
}

// +kubebuilder:skipversion
type Policy struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	ARN *string `json:"arn,omitempty"`

	CreateDate *metav1.Time `json:"createDate,omitempty"`

	IsAttachable *bool `json:"isAttachable,omitempty"`

	PolicyID *string `json:"policyID,omitempty"`

	Tags []*Tag `json:"tags,omitempty"`

	UpdateDate *metav1.Time `json:"updateDate,omitempty"`
}

// +kubebuilder:skipversion
type PolicyDetail struct {
	PolicyDocument *string `json:"policyDocument,omitempty"`
}

// +kubebuilder:skipversion
type PolicyGrantingServiceAccess struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	PolicyARN *string `json:"policyARN,omitempty"`
}

// +kubebuilder:skipversion
type PolicyGroup struct {
	GroupID *string `json:"groupID,omitempty"`

	GroupName *string `json:"groupName,omitempty"`
}

// +kubebuilder:skipversion
type PolicyRole struct {
	RoleID *string `json:"roleID,omitempty"`

	RoleName *string `json:"roleName,omitempty"`
}

// +kubebuilder:skipversion
type PolicyUser struct {
	UserID *string `json:"userID,omitempty"`
}

// +kubebuilder:skipversion
type PolicyVersion struct {
	CreateDate *metav1.Time `json:"createDate,omitempty"`

	Document *string `json:"document,omitempty"`

	IsDefaultVersion *bool `json:"isDefaultVersion,omitempty"`
}

// +kubebuilder:skipversion
type Role struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	ARN *string `json:"arn,omitempty"`

	AssumeRolePolicyDocument *string `json:"assumeRolePolicyDocument,omitempty"`

	CreateDate *metav1.Time `json:"createDate,omitempty"`

	Description *string `json:"description,omitempty"`

	MaxSessionDuration *int64 `json:"maxSessionDuration,omitempty"`

	Path *string `json:"path,omitempty"`
	// Contains information about an attached permissions boundary.
	//
	// An attached permissions boundary is a managed policy that has been attached
	// to a user or role to set the permissions boundary.
	//
	// For more information about permissions boundaries, see Permissions boundaries
	// for IAM identities (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html)
	// in the IAM User Guide.
	PermissionsBoundary *AttachedPermissionsBoundary `json:"permissionsBoundary,omitempty"`

	RoleID *string `json:"roleID,omitempty"`
	// Contains information about the last time that an IAM role was used. This
	// includes the date and time and the Region in which the role was last used.
	// Activity is only reported for the trailing 400 days. This period can be shorter
	// if your Region began supporting these features within the last year. The
	// role might have been used more than 400 days ago. For more information, see
	// Regions where data is tracked (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html#access-advisor_tracking-period)
	// in the IAM User Guide.
	//
	// This data type is returned as a response element in the GetRole and GetAccountAuthorizationDetails
	// operations.
	RoleLastUsed *RoleLastUsed `json:"roleLastUsed,omitempty"`

	RoleName *string `json:"roleName,omitempty"`

	Tags []*Tag `json:"tags,omitempty"`
}

// +kubebuilder:skipversion
type RoleDetail struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	ARN *string `json:"arn,omitempty"`

	AssumeRolePolicyDocument *string `json:"assumeRolePolicyDocument,omitempty"`

	CreateDate *metav1.Time `json:"createDate,omitempty"`
	// Contains a list of instance profiles.
	InstanceProfileList []*InstanceProfile_SDK `json:"instanceProfileList,omitempty"`

	Path *string `json:"path,omitempty"`
	// Contains information about an attached permissions boundary.
	//
	// An attached permissions boundary is a managed policy that has been attached
	// to a user or role to set the permissions boundary.
	//
	// For more information about permissions boundaries, see Permissions boundaries
	// for IAM identities (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html)
	// in the IAM User Guide.
	PermissionsBoundary *AttachedPermissionsBoundary `json:"permissionsBoundary,omitempty"`

	RoleID *string `json:"roleID,omitempty"`
	// Contains information about the last time that an IAM role was used. This
	// includes the date and time and the Region in which the role was last used.
	// Activity is only reported for the trailing 400 days. This period can be shorter
	// if your Region began supporting these features within the last year. The
	// role might have been used more than 400 days ago. For more information, see
	// Regions where data is tracked (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html#access-advisor_tracking-period)
	// in the IAM User Guide.
	//
	// This data type is returned as a response element in the GetRole and GetAccountAuthorizationDetails
	// operations.
	RoleLastUsed *RoleLastUsed `json:"roleLastUsed,omitempty"`

	RoleName *string `json:"roleName,omitempty"`

	Tags []*Tag `json:"tags,omitempty"`
}

// +kubebuilder:skipversion
type RoleLastUsed struct {
	LastUsedDate *metav1.Time `json:"lastUsedDate,omitempty"`

	Region *string `json:"region,omitempty"`
}

// +kubebuilder:skipversion
type SAMLProviderListEntry struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	ARN *string `json:"arn,omitempty"`

	CreateDate *metav1.Time `json:"createDate,omitempty"`

	ValidUntil *metav1.Time `json:"validUntil,omitempty"`
}

// +kubebuilder:skipversion
type SSHPublicKey struct {
	UploadDate *metav1.Time `json:"uploadDate,omitempty"`
}

// +kubebuilder:skipversion
type SSHPublicKeyMetadata struct {
	UploadDate *metav1.Time `json:"uploadDate,omitempty"`
}

// +kubebuilder:skipversion
type ServerCertificate struct {
	Tags []*Tag `json:"tags,omitempty"`
}

// +kubebuilder:skipversion
type ServerCertificateMetadata struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	ARN *string `json:"arn,omitempty"`

	Expiration *metav1.Time `json:"expiration,omitempty"`

	Path *string `json:"path,omitempty"`

	ServerCertificateID *string `json:"serverCertificateID,omitempty"`

	UploadDate *metav1.Time `json:"uploadDate,omitempty"`
}

// +kubebuilder:skipversion
type ServiceLastAccessed struct {
	LastAuthenticated *metav1.Time `json:"lastAuthenticated,omitempty"`
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	LastAuthenticatedEntity *string `json:"lastAuthenticatedEntity,omitempty"`

	LastAuthenticatedRegion *string `json:"lastAuthenticatedRegion,omitempty"`
}

// +kubebuilder:skipversion
type ServiceSpecificCredential struct {
	CreateDate *metav1.Time `json:"createDate,omitempty"`
}

// +kubebuilder:skipversion
type ServiceSpecificCredentialMetadata struct {
	CreateDate *metav1.Time `json:"createDate,omitempty"`
}

// +kubebuilder:skipversion
type SigningCertificate struct {
	UploadDate *metav1.Time `json:"uploadDate,omitempty"`
}

// +kubebuilder:skipversion
type Tag struct {
	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`
}

// +kubebuilder:skipversion
type TrackedActionLastAccessed struct {
	ActionName *string `json:"actionName,omitempty"`
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	LastAccessedEntity *string `json:"lastAccessedEntity,omitempty"`

	LastAccessedRegion *string `json:"lastAccessedRegion,omitempty"`

	LastAccessedTime *metav1.Time `json:"lastAccessedTime,omitempty"`
}

// +kubebuilder:skipversion
type User struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	ARN *string `json:"arn,omitempty"`

	CreateDate *metav1.Time `json:"createDate,omitempty"`

	PasswordLastUsed *metav1.Time `json:"passwordLastUsed,omitempty"`

	Path *string `json:"path,omitempty"`
	// Contains information about an attached permissions boundary.
	//
	// An attached permissions boundary is a managed policy that has been attached
	// to a user or role to set the permissions boundary.
	//
	// For more information about permissions boundaries, see Permissions boundaries
	// for IAM identities (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html)
	// in the IAM User Guide.
	PermissionsBoundary *AttachedPermissionsBoundary `json:"permissionsBoundary,omitempty"`

	Tags []*Tag `json:"tags,omitempty"`

	UserID *string `json:"userID,omitempty"`
}

// +kubebuilder:skipversion
type UserDetail struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	ARN *string `json:"arn,omitempty"`

	CreateDate *metav1.Time `json:"createDate,omitempty"`

	Path *string `json:"path,omitempty"`
	// Contains information about an attached permissions boundary.
	//
	// An attached permissions boundary is a managed policy that has been attached
	// to a user or role to set the permissions boundary.
	//
	// For more information about permissions boundaries, see Permissions boundaries
	// for IAM identities (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html)
	// in the IAM User Guide.
	PermissionsBoundary *AttachedPermissionsBoundary `json:"permissionsBoundary,omitempty"`

	Tags []*Tag `json:"tags,omitempty"`

	UserID *string `json:"userID,omitempty"`
}

// +kubebuilder:skipversion
type VirtualMFADevice struct {
	EnableDate *metav1.Time `json:"enableDate,omitempty"`

	Tags []*Tag `json:"tags,omitempty"`
}
