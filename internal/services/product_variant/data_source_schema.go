// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package product_variant

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*ProductVariantDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"product_id": schema.StringAttribute{
				Computed: true,
			},
			"variant_id": schema.StringAttribute{
				Computed: true,
			},
			"image_url": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"price": schema.Int64Attribute{
				Computed: true,
			},
		},
	}
}

func (d *ProductVariantDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ProductVariantDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
