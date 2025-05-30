// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package product_variant

import (
	"github.com/dackerman/terraform-provider-demostore/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ProductVariantModel struct {
	ID        types.String `tfsdk:"id" json:"-,computed"`
	VariantID types.String `tfsdk:"variant_id" json:"variant_id,computed"`
	ProductID types.String `tfsdk:"product_id" path:"product_id,required"`
	OrgID     types.String `tfsdk:"org_id" path:"org_id,optional"`
	ImageURL  types.String `tfsdk:"image_url" json:"image_url,required"`
	Name      types.String `tfsdk:"name" json:"name,required"`
	Price     types.Int64  `tfsdk:"price" json:"price,required"`
}

func (m ProductVariantModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ProductVariantModel) MarshalJSONForUpdate(state ProductVariantModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
