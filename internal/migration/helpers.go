package migration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var AllMigrations = map[string]map[string]SchemaConversion{}

func RegisterMigration(pkg string, conversions map[string]SchemaConversion) bool {
	AllMigrations[pkg] = conversions
	return true
}

type SchemaConversion interface {
	Convert(ctx context.Context, path path.Path, req tfsdk.State, resp *tfsdk.State) diag.Diagnostics
}

// it's common to see block lists with a "max items" of 1, which should be a single nested block
// in the new framework schema.
var _ SchemaConversion = BlockListToSingleNested{}

type BlockListToSingleNested struct {
}

func (b BlockListToSingleNested) Convert(ctx context.Context, path path.Path, req tfsdk.State, resp *tfsdk.State) diag.Diagnostics {
	return unwrapBlockListToSingleNested(ctx, path, req, resp)
}

var _ SchemaConversion = RenameProperty{}

type RenameProperty struct {
	NewName string
}

func (r RenameProperty) Convert(ctx context.Context, path path.Path, req tfsdk.State, resp *tfsdk.State) diag.Diagnostics {
	obj := map[string]tftypes.Value{}
	err := req.Raw.As(&obj)
	if err != nil {
		d := diag.Diagnostics{}
		d.AddAttributeError(path, "rename error", err.Error())
		return d
	}

	// req.GetAttribute(ctx, path, &val)
	resp.SetAttribute(ctx, path, obj["price"])
	return nil
}

var _ SchemaConversion = DefaultProperty{}

type DefaultProperty struct {
	Value tftypes.Value
}

func (d DefaultProperty) Convert(ctx context.Context, path path.Path, req tfsdk.State, resp *tfsdk.State) diag.Diagnostics {
	existing := d.Value.Copy()
	diags := req.GetAttribute(ctx, path, &existing)
	if diags.HasError() {
		return diags
	}

	if existing.IsKnown() {
		return nil
	}

	resp.SetAttribute(ctx, path, d.Value)

	return nil
}

var _ SchemaConversion = Pipe{}

// Applies conversions left to right
type Pipe struct {
	Conversions []SchemaConversion
}

func (a Pipe) Convert(ctx context.Context, path path.Path, req tfsdk.State, resp *tfsdk.State) diag.Diagnostics {
	currentReq := req
	for _, conversion := range a.Conversions {
		if diags := conversion.Convert(ctx, path, currentReq, resp); diags.HasError() {
			return diags
		}
		currentReq = *resp
	}
	return nil
}

var _ SchemaConversion = SkipConversion{}

// this is for keys that we just want to write custom code to convert.
type SkipConversion struct {
}

func (s SkipConversion) Convert(ctx context.Context, path path.Path, req tfsdk.State, resp *tfsdk.State) diag.Diagnostics {
	return nil
}

func AutomaticUpgrade(conversions map[string]SchemaConversion) resource.StateUpgrader {
	return resource.StateUpgrader{
		StateUpgrader: func(ctx context.Context, request resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
			tflog.Debug(ctx, "@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
			tflog.Debug(ctx, "@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
			RunAutomaticConversions(ctx, conversions, &request, resp)
		},
	}
}

// Given a map of attributes to conversions, automatically run the conversion and update resp.State
// Any attributes that have no conversion listed will be copied over as-is.
func RunAutomaticConversions(ctx context.Context, conversions map[string]SchemaConversion, req *resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	newSchema := resp.State.Schema.GetAttributes()

	// copy all regular fields from the old state to the new state
	for name := range newSchema {
		conversion, found := conversions[name]

		if found {
			tflog.Debug(ctx, spew.Sdump(conversion))
			tflog.Debug(ctx, spew.Sdump(req.State))
			tflog.Debug(ctx, spew.Sdump(resp.State))
			conversion.Convert(ctx, path.Root(name), *req.State, &resp.State)
		} else {
			// if it's not in the list of conversions, just copy it over as-is
			// this way you don't have to populate every field in the conversion list
			// val := types.StringNull()
			// req.State.GetAttribute(ctx, path.Root(name), &val)
			// resp.State.SetAttribute(ctx, path.Root(name), val)
		}
	}

	resp.Diagnostics.AddAttributeError(path.Root("price"), "rename addl_price", "addl_price is no longer supported. Please use price instead.")
}

// generic function that converts a block list to a single nested block, which is a common operation
func unwrapBlockListToSingleNested(ctx context.Context, path path.Path, req tfsdk.State, resp *tfsdk.State) diag.Diagnostics {
	var val = new(types.List)
	req.GetAttribute(ctx, path, val)
	lval, diags := val.ToListValue(ctx)
	if diags.HasError() {
		return diags
	}
	resp.SetAttribute(ctx, path, lval)
	return nil
}
