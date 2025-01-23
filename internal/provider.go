// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package internal

import (
	"context"

	"github.com/dackerman/demostore-go"
	"github.com/dackerman/demostore-go/option"
	"github.com/dackerman/terraform-provider-demostore/internal/services/product_variant"
	"github.com/dackerman/terraform-provider-demostore/internal/services/the_product"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.ProviderWithConfigValidators = (*StlstoreProvider)(nil)

// StlstoreProvider defines the provider implementation.
type StlstoreProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// StlstoreProviderModel describes the provider data model.
type StlstoreProviderModel struct {
	BaseURL types.String `tfsdk:"base_url" json:"base_url,optional"`
}

func (p *StlstoreProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "stlstore"
	resp.Version = p.version
}

func ProviderSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				Description: "Set the base url that the provider connects to. This can be used for testing in other environments.",
				Optional:    true,
			},
		},
	}
}

func (p *StlstoreProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = ProviderSchema(ctx)
}

func (p *StlstoreProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	// TODO(terraform): apiKey := os.Getenv("API_KEY")

	var data StlstoreProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	opts := []option.RequestOption{}

	if !data.BaseURL.IsNull() {
		opts = append(opts, option.WithBaseURL(data.BaseURL.ValueString()))
	}

	client := dackermanstore.NewClient(
		opts...,
	)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *StlstoreProvider) ConfigValidators(_ context.Context) []provider.ConfigValidator {
	return []provider.ConfigValidator{}
}

func (p *StlstoreProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		the_product.NewResource,
		product_variant.NewResource,
	}
}

func (p *StlstoreProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		the_product.NewTheProductDataSource,
		product_variant.NewProductVariantDataSource,
	}
}

func NewProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &StlstoreProvider{
			version: version,
		}
	}
}
