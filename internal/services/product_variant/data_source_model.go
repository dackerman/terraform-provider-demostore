// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package product_variant

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ProductVariantDataSourceModel struct {
	ProductID types.String  `tfsdk:"product_id" path:"product_id,computed"`
	VariantID types.String  `tfsdk:"variant_id" path:"variant_id,computed"`
	AddlPrice types.Float64 `tfsdk:"addl_price" json:"addl_price,computed"`
	ImageURL  types.String  `tfsdk:"image_url" json:"image_url,computed"`
	Name      types.String  `tfsdk:"name" json:"name,computed"`
}
