// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package products_test

import (
	"context"
	"testing"

	"github.com/dackerman/terraform-provider-demostore/internal/services/products"
	"github.com/dackerman/terraform-provider-demostore/internal/test_helpers"
)

func TestProductsModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*products.ProductsModel)(nil)
	schema := products.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
