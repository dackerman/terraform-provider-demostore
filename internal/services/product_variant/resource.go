// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package product_variant

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/dackerman/demostore-private-go/v2"
	"github.com/dackerman/demostore-private-go/v2/option"
	"github.com/dackerman/demostore-private-go/v2/packages/param"
	"github.com/dackerman/terraform-provider-demostore/internal/apijson"
	"github.com/dackerman/terraform-provider-demostore/internal/importpath"
	"github.com/dackerman/terraform-provider-demostore/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*ProductVariantResource)(nil)
var _ resource.ResourceWithModifyPlan = (*ProductVariantResource)(nil)
var _ resource.ResourceWithImportState = (*ProductVariantResource)(nil)

func NewResource() resource.Resource {
	return &ProductVariantResource{}
}

// ProductVariantResource defines the resource implementation.
type ProductVariantResource struct {
	client *dackermanstore.Client
}

func (r *ProductVariantResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_product_variant"
}

func (r *ProductVariantResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	r.client = client
}

func (r *ProductVariantResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ProductVariantModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params := dackermanstore.ProductVariantNewParams{}

	if !data.OrgID.IsNull() {
		params.OrgID = param.NewOpt(data.OrgID.ValueString())
	}

	dataBytes, err := data.MarshalJSON()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	_, err = r.client.Products.Variants.New(
		ctx,
		data.ProductID.ValueString(),
		params,
		option.WithRequestBody("application/json", dataBytes),
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
	data.ID = data.VariantID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProductVariantResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ProductVariantModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *ProductVariantModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params := dackermanstore.ProductVariantUpdateParams{}

	if !data.OrgID.IsNull() {
		params.OrgID = param.NewOpt(data.OrgID.ValueString())
	}

	dataBytes, err := data.MarshalJSONForUpdate(*state)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	_, err = r.client.Products.Variants.Update(
		ctx,
		data.ProductID.ValueString(),
		data.VariantID.ValueString(),
		params,
		option.WithRequestBody("application/json", dataBytes),
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
	data.ID = data.VariantID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProductVariantResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ProductVariantModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params := dackermanstore.ProductVariantGetParams{}

	if !data.OrgID.IsNull() {
		params.OrgID = param.NewOpt(data.OrgID.ValueString())
	}

	res := new(http.Response)
	_, err := r.client.Products.Variants.Get(
		ctx,
		data.ProductID.ValueString(),
		data.VariantID.ValueString(),
		params,
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if res != nil && res.StatusCode == 404 {
		resp.Diagnostics.AddWarning("Resource not found", "The resource was not found on the server and will be removed from state.")
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &data)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data.ID = data.VariantID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProductVariantResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ProductVariantModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params := dackermanstore.ProductVariantDeleteParams{}

	if !data.OrgID.IsNull() {
		params.OrgID = param.NewOpt(data.OrgID.ValueString())
	}

	_, err := r.client.Products.Variants.Delete(
		ctx,
		data.ProductID.ValueString(),
		data.VariantID.ValueString(),
		params,
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	data.ID = data.VariantID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProductVariantResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data *ProductVariantModel = new(ProductVariantModel)

	path_org_id := ""
	path_product_id := ""
	path_variant_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<org_id>/<product_id>/<variant_id>",
		&path_org_id,
		&path_product_id,
		&path_variant_id,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.OrgID = types.StringValue(path_org_id)
	data.ProductID = types.StringValue(path_product_id)
	data.VariantID = types.StringValue(path_variant_id)

	res := new(http.Response)
	_, err := r.client.Products.Variants.Get(
		ctx,
		path_product_id,
		path_variant_id,
		dackermanstore.ProductVariantGetParams{
			OrgID: param.NewOpt(path_org_id),
		},
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &data)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data.ID = data.VariantID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProductVariantResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}
