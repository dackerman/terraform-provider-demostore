// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package products

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*ProductsResource)(nil)

func (r *ProductsResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
