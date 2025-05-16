// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package product_variant

import (
	"context"
	"dackerman/demostore-go/v2"
	"dackerman/demostore-go/v2/packages/param"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ProductVariantDataSourceModel struct {
	OrgID     types.String `tfsdk:"org_id" path:"org_id,required"`
	ProductID types.String `tfsdk:"product_id" path:"product_id,required"`
	VariantID types.String `tfsdk:"variant_id" path:"variant_id,required"`
	ImageURL  types.String `tfsdk:"image_url" json:"image_url,computed"`
	Name      types.String `tfsdk:"name" json:"name,computed"`
	Price     types.Int64  `tfsdk:"price" json:"price,computed"`
}

func (m *ProductVariantDataSourceModel) toReadParams(_ context.Context) (params dackermanstore.ProductVariantGetParams, diags diag.Diagnostics) {
	params = dackermanstore.ProductVariantGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = param.NewOpt(m.OrgID.ValueString())
	}

	return
}
