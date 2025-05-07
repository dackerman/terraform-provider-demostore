// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package product_test

import (
	"context"
	"testing"

	"github.com/dackerman/terraform-provider-demostore/internal/services/product"
	"github.com/dackerman/terraform-provider-demostore/internal/test_helpers"
)

func TestProductsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*product.ProductsDataSourceModel)(nil)
	schema := product.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
