// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package internal

import (
	"context"
	"os"

	"github.com/dackerman/terraform-provider-demostore/internal/services/product"
	"github.com/dackerman/terraform-provider-demostore/internal/services/product_variant"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/dackerman-store-go"
	"github.com/stainless-sdks/dackerman-store-go/option"
)

var _ provider.ProviderWithConfigValidators = (*DemostoreProvider)(nil)

// DemostoreProvider defines the provider implementation.
type DemostoreProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// DemostoreProviderModel describes the provider data model.
type DemostoreProviderModel struct {
	BaseURL   types.String `tfsdk:"base_url" json:"base_url,optional"`
	AuthToken types.String `tfsdk:"auth_token" json:"auth_token,required"`
}

func (p *DemostoreProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "demostore"
	resp.Version = p.version
}

func ProviderSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				Description: "Set the base url that the provider connects to. This can be used for testing in other environments.",
				Optional:    true,
			},
			"auth_token": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (p *DemostoreProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = ProviderSchema(ctx)
}

func (p *DemostoreProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	var data DemostoreProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	opts := []option.RequestOption{}

	if !data.BaseURL.IsNull() {
		opts = append(opts, option.WithBaseURL(data.BaseURL.ValueString()))
	}
	if o, ok := os.LookupEnv("DEMOSTORE_API_KEY"); ok {
		opts = append(opts, option.WithAuthToken(o))
	}
	if !data.AuthToken.IsNull() {
		opts = append(opts, option.WithAuthToken(data.AuthToken.ValueString()))
	}

	client := dackermanstore.NewClient(
		opts...,
	)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *DemostoreProvider) ConfigValidators(_ context.Context) []provider.ConfigValidator {
	return []provider.ConfigValidator{}
}

func (p *DemostoreProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		product.NewResource,
		product_variant.NewResource,
	}
}

func (p *DemostoreProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		product.NewProductDataSource,
		product_variant.NewProductVariantDataSource,
	}
}

func NewProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &DemostoreProvider{
			version: version,
		}
	}
}
