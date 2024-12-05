// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package product_variant_test

import (
	"context"
	"testing"

	"github.com/dackerman/terraform-provider-demostore/internal/services/product_variant"
	"github.com/dackerman/terraform-provider-demostore/internal/test_helpers"
)

func TestProductVariantDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*product_variant.ProductVariantDataSourceModel)(nil)
	schema := product_variant.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
