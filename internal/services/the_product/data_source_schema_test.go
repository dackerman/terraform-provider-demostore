// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package the_product_test

import (
	"context"
	"testing"

	"github.com/dackerman/terraform-provider-demostore/internal/services/the_product"
	"github.com/dackerman/terraform-provider-demostore/internal/test_helpers"
)

func TestTheProductDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*the_product.TheProductDataSourceModel)(nil)
	schema := the_product.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
