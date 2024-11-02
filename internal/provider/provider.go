package provider

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/vvbogdanov87/terraform-provider-acme/internal/provider/common"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &crdProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &crdProvider{
			version: version,
		}
	}
}

// crdProvider is the provider implementation.
type crdProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// crdProviderModel maps provider schema data to a Go type.
type crdProviderModel struct {
	Namespace types.String `tfsdk:"namespace"`
}

// Metadata returns the provider type name.
func (p *crdProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "acme"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *crdProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"namespace": schema.StringAttribute{
				Optional: false,
				Required: true,
			},
		},
	}
}

// Configure prepares a Kubernetes API client for data sources and resources.
func (p *crdProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var model crdProviderModel
	diags := req.Config.Get(ctx, &model)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	home := homedir.HomeDir()
	kubeconfig := filepath.Join(home, ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		resp.Diagnostics.AddError(
			"read kubernetes client config",
			fmt.Sprintf("Error building Kubernetes client config from a master url or a kubeconfig filepath:\n%s", err.Error()),
		)
		return
	}

	clientset, err := dynamic.NewForConfig(config)
	if err != nil {
		resp.Diagnostics.AddError(
			"create kuberentes client",
			fmt.Sprintf("Error creating Kubernetes client from config:\n%s", err.Error()),
		)
		return
	}

	resp.ResourceData = common.ResourceData{
		Clientset: clientset,
		Namespace: model.Namespace.ValueString(),
	}
}

// DataSources defines the data sources implemented in the provider.
func (p *crdProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}
