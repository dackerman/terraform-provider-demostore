// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package product_variant

import (
	"context"

	"github.com/dackerman/terraform-provider-demostore/internal/migration"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*ProductVariantResource)(nil)

var conversions = migration.RegisterMigration("stlstore_product_variant", map[string]migration.SchemaConversion{
	"price":      migration.SkipConversion{},
	"addl_price": migration.RenameProperty{NewName: "price"},
})

func (r *ProductVariantResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: migration.AutomaticUpgrade(conversions),
	}
}
