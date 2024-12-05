// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package product_variant

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ProductVariantDataSourceModel struct {
	VariantID types.String  `tfsdk:"variant_id" path:"variant_id,required"`
	ProductID types.String  `tfsdk:"product_id" path:"product_id,computed"`
	ID        types.String  `tfsdk:"id" json:"id,computed"`
	Name      types.String  `tfsdk:"name" json:"name,computed"`
	Price     types.Float64 `tfsdk:"price" json:"price,computed"`
}