// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package product_variant

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ProductVariantDataSourceModel struct {
	ProductID types.String `tfsdk:"product_id" path:"product_id,computed"`
	VariantID types.String `tfsdk:"variant_id" path:"variant_id,computed"`
	ImageURL  types.String `tfsdk:"image_url" json:"image_url,computed"`
	Name      types.String `tfsdk:"name" json:"name,computed"`
	Price     types.Int64  `tfsdk:"price" json:"price,computed"`
	Type      types.String `tfsdk:"type" json:"type,computed"`
}
