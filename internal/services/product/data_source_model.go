// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package product

import (
	"context"

	"github.com/dackerman/demostore-private-go/v2"
	"github.com/dackerman/demostore-private-go/v2/packages/param"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ProductDataSourceModel struct {
	OrgID       types.String `tfsdk:"org_id" path:"org_id,required"`
	ProductID   types.String `tfsdk:"product_id" path:"product_id,required"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	ImageURL    types.String `tfsdk:"image_url" json:"image_url,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
	Price       types.Int64  `tfsdk:"price" json:"price,computed"`
}

func (m *ProductDataSourceModel) toReadParams(_ context.Context) (params dackermanstore.ProductGetParams, diags diag.Diagnostics) {
	params = dackermanstore.ProductGetParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = param.NewOpt(m.OrgID.ValueString())
	}

	return
}
