package schema

import (
	"github.com/rancher/norman/types"
	"github.com/rancher/types/factory"

	inetworkingv1alpha3 "github.com/istio/api/networking/v1alpha3"
)

var (
	Version = types.APIVersion{
		Version: "v1alpha3",
		Group:   "networking.istio.io",
		Path:    "/v1alpha3",
	}

	Schemas = factory.Schemas(&Version).
		Init(istioNetworkingTypes)
)

func istioNetworkingTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.MustImport(&Version, inetworkingv1alpha3.Gateway{}).
		MustImport(&Version, inetworkingv1alpha3.VirtualService{}).
		MustImport(&Version, inetworkingv1alpha3.DestinationRule{}).
		MustImport(&Version, inetworkingv1alpha3.ServiceEntry{}).
		MustImport(&Version, inetworkingv1alpha3.EnvoyFilter{})
}
