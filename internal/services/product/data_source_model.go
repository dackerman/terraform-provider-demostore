// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package product

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ProductDataSourceModel struct {
	ProductID   types.String  `tfsdk:"product_id" path:"product_id,required"`
	Description types.String  `tfsdk:"description" json:"description,computed"`
	ID          types.String  `tfsdk:"id" json:"id,computed"`
	ImageURL    types.String  `tfsdk:"image_url" json:"image_url,computed"`
	Name        types.String  `tfsdk:"name" json:"name,computed"`
	Price       types.Float64 `tfsdk:"price" json:"price,computed"`
}
