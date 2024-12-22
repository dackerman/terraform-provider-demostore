// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package the_product

import (
	"github.com/dackerman/terraform-provider-demostore/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TheProductModel struct {
	ID          types.String `tfsdk:"id" json:"-,computed"`
	ProductID   types.String `tfsdk:"product_id" json:"product_id,computed"`
	Description types.String `tfsdk:"description" json:"description,required"`
	ImageURL    types.String `tfsdk:"image_url" json:"image_url,required"`
	Name        types.String `tfsdk:"name" json:"name,required"`
	Price       types.Int64  `tfsdk:"price" json:"price,required"`
}

func (m TheProductModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m TheProductModel) MarshalJSONForUpdate(state TheProductModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
