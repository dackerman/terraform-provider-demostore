// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package internal

import (
	"context"
	"os"

	"github.com/dackerman/demostore-go"
	"github.com/dackerman/demostore-go/option"
	"github.com/dackerman/terraform-provider-demostore/internal/services/product"
	"github.com/dackerman/terraform-provider-demostore/internal/services/product_variant"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
	AuthToken types.String `tfsdk:"auth_token" json:"auth_token,optional"`
	OrgID     types.String `tfsdk:"org_id" json:"org_id,optional"`
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
				Optional: true,
			},
			"org_id": schema.StringAttribute{
				Optional: true,
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

	if !data.BaseURL.IsNull() && !data.BaseURL.IsUnknown() {
		opts = append(opts, option.WithBaseURL(data.BaseURL.ValueString()))
	} else if o, ok := os.LookupEnv("STAINLESS_STORE_BASE_URL"); ok {
		opts = append(opts, option.WithBaseURL(o))
	}

	if !data.AuthToken.IsNull() && !data.AuthToken.IsUnknown() {
		opts = append(opts, option.WithAuthToken(data.AuthToken.ValueString()))
	} else if o, ok := os.LookupEnv("DEMOSTORE_API_KEY"); ok {
		opts = append(opts, option.WithAuthToken(o))
	} else {
		resp.Diagnostics.AddAttributeError(
			path.Root("auth_token"),
			"Missing auth_token value",
			"The auth_token field is required. Set it in provider configuration or via the \"DEMOSTORE_API_KEY\" environment variable.",
		)
		return
	}

	if !data.OrgID.IsNull() && !data.OrgID.IsUnknown() {
		opts = append(opts, option.WithOrgID(data.OrgID.ValueString()))
	} else if o, ok := os.LookupEnv("DEMOSTORE_ORG_ID"); ok {
		opts = append(opts, option.WithOrgID(o))
	} else {
		resp.Diagnostics.AddAttributeError(
			path.Root("org_id"),
			"Missing org_id value",
			"The org_id field is required. Set it in provider configuration or via the \"DEMOSTORE_ORG_ID\" environment variable.",
		)
		return
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
		product.NewProductsDataSource,
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
