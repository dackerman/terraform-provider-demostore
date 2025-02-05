// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package products_test

import (
	"context"
	"testing"

	"github.com/dackerman/terraform-provider-demostore/internal/services/products"
	"github.com/dackerman/terraform-provider-demostore/internal/test_helpers"
)

func TestProductsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*products.ProductsDataSourceModel)(nil)
	schema := products.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
