// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package the_product

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*TheProductResource)(nil)

func (r *TheProductResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
