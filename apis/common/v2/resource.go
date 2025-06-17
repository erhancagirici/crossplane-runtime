/*
Copyright 2019 The Crossplane Authors.

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

package v2

import (
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// A LocalSecretKeySelector is a reference to a secret key in an arbitrary namespace.
type LocalSecretKeySelector struct {
	xpv1.LocalSecretReference `json:",inline"`

	// The key to select.
	Key string `json:"key"`
}

// Policy represents the Resolve and Resolution policies of Reference instance.
type Policy struct {
	// Resolve specifies when this reference should be resolved. The default
	// is 'IfNotPresent', which will attempt to resolve the reference only when
	// the corresponding field is not present. Use 'Always' to resolve the
	// reference on every reconcile.
	// +optional
	// +kubebuilder:validation:Enum=Always;IfNotPresent
	Resolve *xpv1.ResolvePolicy `json:"resolve,omitempty"`

	// Resolution specifies whether resolution of this reference is required.
	// The default is 'Required', which means the reconcile will fail if the
	// reference cannot be resolved. 'Optional' means this reference will be
	// a no-op if it cannot be resolved.
	// +optional
	// +kubebuilder:default=Required
	// +kubebuilder:validation:Enum=Required;Optional
	Resolution *xpv1.ResolutionPolicy `json:"resolution,omitempty"`
}

// IsResolutionPolicyOptional checks whether the resolution policy of relevant reference is Optional.
func (p *Policy) IsResolutionPolicyOptional() bool {
	if p == nil || p.Resolution == nil {
		return false
	}
	return *p.Resolution == xpv1.ResolutionPolicyOptional
}

// IsResolvePolicyAlways checks whether the resolution policy of relevant reference is Always.
func (p *Policy) IsResolvePolicyAlways() bool {
	if p == nil || p.Resolve == nil {
		return false
	}
	return *p.Resolve == xpv1.ResolvePolicyAlways
}

// A Reference to a named object.
type Reference struct {
	// Name of the referenced object.
	Name string `json:"name"`

	// Policies for referencing.
	// +optional
	Policy *Policy `json:"policy,omitempty"`
}

type TypedNamespacedReference struct {
	xpv1.TypedReference `json:",inline"`
	Namespace           string `json:"namespace"`
}

// SetGroupVersionKind sets the Kind and APIVersion of a TypedReference.
func (obj *TypedNamespacedReference) SetGroupVersionKind(gvk schema.GroupVersionKind) {
	obj.APIVersion, obj.Kind = gvk.ToAPIVersionAndKind()
}

// GroupVersionKind gets the GroupVersionKind of a TypedReference.
func (obj *TypedNamespacedReference) GroupVersionKind() schema.GroupVersionKind {
	return schema.FromAPIVersionAndKind(obj.APIVersion, obj.Kind)
}

// GetObjectKind get the ObjectKind of a TypedReference.
func (obj *TypedNamespacedReference) GetObjectKind() schema.ObjectKind { return obj }

// A ManagedResourceSpec defines the desired state of a managed resource.
type ManagedResourceSpec struct {
	// WriteConnectionSecretToReference specifies the namespace and name of a
	// Secret to which any connection details for this managed resource should
	// be written. Connection details frequently include the endpoint, username,
	// and password required to connect to the managed resource.
	// +optional
	WriteConnectionSecretToReference *xpv1.LocalSecretReference `json:"writeConnectionSecretToRef,omitempty"`

	// ProviderConfigReference specifies how the provider that will be used to
	// create, observe, update, and delete this managed resource should be
	// configured.
	// +kubebuilder:default={"name": "default"}
	ProviderConfigReference *xpv1.TypedReference `json:"providerConfigRef,omitempty"`

	// THIS IS A BETA FIELD. It is on by default but can be opted out
	// through a Crossplane feature flag.
	// ManagementPolicies specify the array of actions Crossplane is allowed to
	// take on the managed and external resources.
	// See the design doc for more information: https://github.com/crossplane/crossplane/blob/499895a25d1a1a0ba1604944ef98ac7a1a71f197/design/design-doc-observe-only-resources.md?plain=1#L223
	// and this one: https://github.com/crossplane/crossplane/blob/444267e84783136daa93568b364a5f01228cacbe/design/one-pager-ignore-changes.md
	// +optional
	// +kubebuilder:default={"*"}
	ManagementPolicies xpv1.ManagementPolicies `json:"managementPolicies,omitempty"`
}

// ResourceStatus represents the observed state of a managed resource.
//type ResourceStatus struct {
//	xpv1.ConditionedStatus `json:",inline"`
//	ObservedStatus    `json:",inline"`
//}

// A CredentialsSource is a source from which provider credentials may be
// acquired.
type CredentialsSource string

const (
	// CredentialsSourceNone indicates that a provider does not require
	// credentials.
	CredentialsSourceNone CredentialsSource = "None"

	// CredentialsSourceSecret indicates that a provider should acquire
	// credentials from a secret.
	CredentialsSourceSecret CredentialsSource = "Secret"

	// CredentialsSourceInjectedIdentity indicates that a provider should use
	// credentials via its (pod's) identity; i.e. via IRSA for AWS,
	// Workload Identity for GCP, Pod Identity for Azure, or in-cluster
	// authentication for the Kubernetes API.
	CredentialsSourceInjectedIdentity CredentialsSource = "InjectedIdentity"

	// CredentialsSourceEnvironment indicates that a provider should acquire
	// credentials from an environment variable.
	CredentialsSourceEnvironment CredentialsSource = "Environment"

	// CredentialsSourceFilesystem indicates that a provider should acquire
	// credentials from the filesystem.
	CredentialsSourceFilesystem CredentialsSource = "Filesystem"
)

// A TypedProviderConfigUsage is a record that a particular managed resource is using
// a particular provider configuration.
type TypedProviderConfigUsage struct {
	// ProviderConfigReference to the provider config being used.
	ProviderConfigReference xpv1.TypedReference `json:"providerConfigRef"`

	// ResourceReference to the managed resource using the provider config.
	ResourceReference xpv1.TypedReference `json:"resourceRef"`
}

type TypedNamespacedProviderConfigUsage struct {
	// ProviderConfigReference to the provider config being used.
	ProviderConfigReference TypedNamespacedReference `json:"providerConfigRef"`

	// ResourceReference to the managed resource using the provider config.
	ResourceReference xpv1.TypedReference `json:"resourceRef"`
}
