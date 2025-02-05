// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package products

import (
	"github.com/dackerman/terraform-provider-demostore/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ProductsModel struct {
	ID              types.String `tfsdk:"id" json:"-,computed"`
	ProductID       types.String `tfsdk:"product_id" json:"product_id,computed"`
	Description     types.String `tfsdk:"description" json:"description,required"`
	ImageURL        types.String `tfsdk:"image_url" json:"image_url,required"`
	Name            types.String `tfsdk:"name" json:"name,required"`
	Price           types.Int64  `tfsdk:"price" json:"price,required"`
	LongDescription types.String `tfsdk:"long_description" json:"long_description,optional"`
}

func (m ProductsModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ProductsModel) MarshalJSONForUpdate(state ProductsModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
