// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package product_test

import (
	"context"
	"testing"

	"github.com/dackerman/terraform-provider-demostore/internal/services/product"
	"github.com/dackerman/terraform-provider-demostore/internal/test_helpers"
)

func TestProductModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*product.ProductModel)(nil)
	schema := product.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
