package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/vvbogdanov87/terraform-provider-acme/internal/provider/cloud_acme_local_bucket_v1"
)

// Resources defines the resources implemented in the provider.
func (p *crdProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		cloud_acme_local_bucket_v1.NewTFResource,
	}
}
