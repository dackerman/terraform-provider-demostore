// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package product

import (
	"context"

	"github.com/dackerman/demostore-go/v2"
	"github.com/dackerman/demostore-go/v2/packages/param"
	"github.com/dackerman/terraform-provider-demostore/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ProductsDataListDataSourceEnvelope struct {
	Data customfield.NestedObjectList[ProductsItemsDataSourceModel] `json:"data,computed"`
}

type ProductsDataSourceModel struct {
	OrgID    types.String                                               `tfsdk:"org_id" path:"org_id,optional"`
	MaxItems types.Int64                                                `tfsdk:"max_items"`
	Items    customfield.NestedObjectList[ProductsItemsDataSourceModel] `tfsdk:"items"`
}

func (m *ProductsDataSourceModel) toListParams(_ context.Context) (params dackermanstore.ProductListParams, diags diag.Diagnostics) {
	params = dackermanstore.ProductListParams{}

	if !m.OrgID.IsNull() {
		params.OrgID = param.NewOpt(m.OrgID.ValueString())
	}

	return
}

type ProductsItemsDataSourceModel struct {
	Description types.String `tfsdk:"description" json:"description,computed"`
	ImageURL    types.String `tfsdk:"image_url" json:"image_url,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
	Price       types.Int64  `tfsdk:"price" json:"price,computed"`
	ProductID   types.String `tfsdk:"product_id" json:"product_id,computed"`
}
