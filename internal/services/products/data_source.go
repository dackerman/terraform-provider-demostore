// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package products

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/dackerman/demostore-go"
	"github.com/dackerman/demostore-go/option"
	"github.com/dackerman/terraform-provider-demostore/internal/apijson"
	"github.com/dackerman/terraform-provider-demostore/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type ProductsDataSource struct {
	client *dackermanstore.Client
}

var _ datasource.DataSourceWithConfigure = (*ProductsDataSource)(nil)

func NewProductsDataSource() datasource.DataSource {
	return &ProductsDataSource{}
}

func (d *ProductsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_products"
}

func (d *ProductsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ProductsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *ProductsDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	_, err := d.client.Products.Get(
		ctx,
		data.ProductID.ValueString(),
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
