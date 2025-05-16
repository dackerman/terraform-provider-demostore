// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package product_variant

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/dackerman/demostore-go/v2"
	"github.com/dackerman/demostore-go/v2/option"
	"github.com/dackerman/terraform-provider-demostore/internal/apijson"
	"github.com/dackerman/terraform-provider-demostore/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type ProductVariantDataSource struct {
	client *dackermanstore.Client
}

var _ datasource.DataSourceWithConfigure = (*ProductVariantDataSource)(nil)

func NewProductVariantDataSource() datasource.DataSource {
	return &ProductVariantDataSource{}
}

func (d *ProductVariantDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_product_variant"
}

func (d *ProductVariantDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*dackermanstore.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *dackermanstore.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *ProductVariantDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *ProductVariantDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params, diags := data.toReadParams(ctx)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	_, err := d.client.Products.Variants.Get(
		ctx,
		data.ProductID.ValueString(),
		data.VariantID.ValueString(),
		params,
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.UnmarshalComputed(bytes, &data)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
