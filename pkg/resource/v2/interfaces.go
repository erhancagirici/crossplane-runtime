package v2

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type TypedRequiredProviderConfigReferencer interface {
	GetProviderConfigTypedReference() xpv1.TypedReference
	SetProviderConfigTypedReference(p xpv1.TypedReference)
}

// A TypedProviderConfigReferencer is the same as ProviderConfigReferencer,
// but the ProviderConfig is typed.
type TypedProviderConfigReferencer interface {
	GetProviderConfigTypedReference() *xpv1.TypedReference
	SetProviderConfigTypedReference(p *xpv1.TypedReference)
}

type LocalConnectionSecretWriterTo interface {
	SetWriteConnectionSecretToLocalReference(r *xpv1.LocalSecretReference)
	GetWriteConnectionSecretToLocalReference() *xpv1.LocalSecretReference
}

type RequiredTypedProviderConfigReferencer interface {
	GetProviderConfigTypedReference() xpv1.TypedReference
	SetProviderConfigTypedReference(p xpv1.TypedReference)
}

// A Managed is a Kubernetes object representing a concrete managed
// resource (e.g. a CloudSQL instance).
type Managed interface { //nolint:interfacebloat // This interface has to be big.
	resource.Object

	TypedProviderConfigReferencer
	resource.LocalConnectionSecretWriterTo
	resource.Manageable

	resource.Conditioned
}

// A ManagedList is a list of managed resources.
type ManagedList interface {
	client.ObjectList

	// GetItems returns the list of managed resources.
	GetItems() []Managed
}

// A ProviderConfigUsage indicates a usage of a Crossplane provider config.
type ProviderConfigUsage interface {
	resource.Object

	RequiredTypedProviderConfigReferencer
	resource.RequiredTypedResourceReferencer
}

// A ProviderConfigUsageList is a list of provider config usages.
type ProviderConfigUsageList interface {
	client.ObjectList

	// GetItems returns the list of provider config usages.
	GetItems() []ProviderConfigUsage
}
