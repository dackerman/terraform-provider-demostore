// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	tfjson "github.com/hashicorp/terraform-json"
)

// Result represents the structured output of the breaking changes analysis
type Result struct {
	AddedResources    []string                   `json:"added_resources,omitempty"`
	RemovedResources  []string                   `json:"removed_resources,omitempty"`
	ChangedResources  map[string]ResourceChanges `json:"changed_resources,omitempty"`
	AddedProviders    []string                   `json:"added_providers,omitempty"`
	HasBreakingChange bool                       `json:"-"`
}

// ResourceChanges represents changes to a resource
type ResourceChanges struct {
	BreakingChanges    []AttributeChange `json:"breaking_changes,omitempty"`
	NonBreakingChanges []AttributeChange `json:"non_breaking_changes,omitempty"`
}

// AttributeChange represents a change to an attribute
type AttributeChange struct {
	Name   string           `json:"name"`
	Type   string           `json:"type"` // "added", "removed", "updated"
	Before *AttributeConfig `json:"before,omitempty"`
	After  *AttributeConfig `json:"after,omitempty"`
}

// AttributeConfig represents the configurability of an attribute
type AttributeConfig struct {
	AttributeType string `json:"attribute_type"`
	Required      bool   `json:"required"`
	Computed      bool   `json:"computed"`
}

func main() {
	var previousVersion string
	var currentVersion string

	flag.StringVar(&previousVersion, "previous-version", "", "Path to the previous version JSON file")
	flag.StringVar(&currentVersion, "current-version", "", "Path to the current version JSON file")
	flag.Parse()

	if previousVersion == "" || currentVersion == "" {
		fmt.Println("Both --previous-version and --current-version flags are required")
		os.Exit(1)
	}

	oldData, err := os.ReadFile(previousVersion)
	if err != nil {
		fmt.Printf("Error reading previousVersion: %v\n", err)
		os.Exit(1)
	}

	oldSchema := tfjson.ProviderSchemas{}
	if err := json.Unmarshal(oldData, &oldSchema); err != nil {
		fmt.Printf("Error unmarshaling previous version schema: %v\n", err)
		os.Exit(1)
	}

	newData, err := os.ReadFile(currentVersion)
	if err != nil {
		fmt.Printf("Error reading currentVersion: %v\n", err)
		os.Exit(1)
	}

	newSchema := tfjson.ProviderSchemas{}
	if err := json.Unmarshal(newData, &newSchema); err != nil {
		fmt.Printf("Error unmarshaling current version schema: %v\n", err)
		os.Exit(1)
	}

	result := Result{
		ChangedResources: make(map[string]ResourceChanges),
	}

	// Process each provider schema
	for providerName, oldProvider := range oldSchema.Schemas {
		newProvider, exists := newSchema.Schemas[providerName]
		if !exists {
			panic(fmt.Sprintf("Expected same providers in before and after schema but %s doesn't exist in the new schema.", providerName))
		}

		// Check resources in old schema
		for resourceName, oldResource := range oldProvider.ResourceSchemas {
			newResource, exists := newProvider.ResourceSchemas[resourceName]
			if !exists {
				result.RemovedResources = append(result.RemovedResources, resourceName)
				result.HasBreakingChange = true
				continue
			}

			resourceChanges := ResourceChanges{}
			hasChanges := false

			// Check attributes in new resource that might not be in old resource (added attributes)
			for attrName, newAttr := range newResource.Block.Attributes {
				oldAttr, exists := oldResource.Block.Attributes[attrName]
				if !exists {
					// New attribute
					change := AttributeChange{
						Name:  attrName,
						Type:  "added",
						After: attributeConfig(newAttr),
					}

					// If the attribute is required, it's a breaking change
					if newAttr.Required {
						resourceChanges.BreakingChanges = append(resourceChanges.BreakingChanges, change)
						result.HasBreakingChange = true
					} else {
						// Non-breaking new attribute
						resourceChanges.NonBreakingChanges = append(resourceChanges.NonBreakingChanges, change)
					}
					hasChanges = true
				} else {
					// Check for breaking configurability changes
					if breakingChangeMsg := isBreakingConfigurabilityChange(oldAttr, newAttr); breakingChangeMsg != "" {
						// Breaking change to existing attribute
						change := AttributeChange{
							Name:   attrName,
							Type:   "updated",
							Before: attributeConfig(oldAttr),
							After:  attributeConfig(newAttr),
						}
						resourceChanges.BreakingChanges = append(resourceChanges.BreakingChanges, change)
						hasChanges = true
						result.HasBreakingChange = true
					} else if !attributesEqual(oldAttr, newAttr) {
						// Non-breaking change to existing attribute
						change := AttributeChange{
							Name:   attrName,
							Type:   "updated",
							Before: attributeConfig(oldAttr),
							After:  attributeConfig(newAttr),
						}
						resourceChanges.NonBreakingChanges = append(resourceChanges.NonBreakingChanges, change)
						hasChanges = true
					}
				}
			}

			// Check attributes in old resource that might not be in new resource (removed attributes)
			for attrName := range oldResource.Block.Attributes {
				_, exists := newResource.Block.Attributes[attrName]
				if !exists {
					// Removed attribute (always breaking)
					change := AttributeChange{
						Name: attrName,
						Type: "removed",
					}
					resourceChanges.BreakingChanges = append(resourceChanges.BreakingChanges, change)
					hasChanges = true
					result.HasBreakingChange = true
				}
			}

			// Add to changed resources if there are any changes
			if hasChanges {
				result.ChangedResources[resourceName] = resourceChanges
			}
		}

		// Check for new resources in new schema that weren't in old schema
		for resourceName := range newProvider.ResourceSchemas {
			_, exists := oldProvider.ResourceSchemas[resourceName]
			if !exists {
				result.AddedResources = append(result.AddedResources, resourceName)
			}
		}
	}

	// Check for new providers in new schema
	for providerName := range newSchema.Schemas {
		_, exists := oldSchema.Schemas[providerName]
		if !exists {
			result.AddedProviders = append(result.AddedProviders, providerName)
		}
	}

	// Output the result as JSON
	jsonOutput, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling result to JSON: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(jsonOutput))

	if result.HasBreakingChange {
		os.Exit(1)
	}
}

// attributesEqual checks if two attributes are equal (ignoring non-breaking changes)
func attributesEqual(oldAttr, newAttr *tfjson.SchemaAttribute) bool {
	// Check basic properties
	if oldAttr.Required != newAttr.Required ||
		oldAttr.Optional != newAttr.Optional ||
		oldAttr.Computed != newAttr.Computed {
		return false
	}

	// Check types
	if (oldAttr.AttributeNestedType == nil) != (newAttr.AttributeNestedType == nil) {
		return false
	}

	if oldAttr.AttributeNestedType == nil && newAttr.AttributeNestedType == nil {
		if !oldAttr.AttributeType.Equals(newAttr.AttributeType) {
			return false
		}
	}

	// For nested types, we'd need more detailed checking
	// This is a simplified version
	if oldAttr.AttributeNestedType != nil && newAttr.AttributeNestedType != nil {
		if oldAttr.AttributeNestedType.NestingMode != newAttr.AttributeNestedType.NestingMode {
			return false
		}
	}

	return true
}

// isBreakingConfigurabilityChange checks if the attribute's configurability changed in a breaking way
// Returns a descriptive message if there's a breaking change, or empty string if no breaking change
//
// Breaking changes:
// - Optional -> Required
// - Optional -> Computed (but not Optional)
// - Required -> Computed (but not Optional)
// - Computed Optional -> Required
//
// Non-breaking changes (allowed):
// - Required -> Optional
// - Required -> Computed Optional
// - Optional -> Computed Optional
func isBreakingConfigurabilityChange(oldAttr, newAttr *tfjson.SchemaAttribute) string {
	// Was optional, is now required (breaking)
	if !oldAttr.Required && newAttr.Required {
		return "changed from optional to required"
	}

	// Was optional, now computed-only (not optional anymore) (breaking)
	if oldAttr.Optional && !oldAttr.Computed && !newAttr.Optional && newAttr.Computed {
		return "changed from user-configurable to computed-only"
	}

	// Was required, now computed-only (not optional) (breaking)
	if oldAttr.Required && !newAttr.Required && !newAttr.Optional && newAttr.Computed {
		return "changed from required to computed-only"
	}

	// Was computed+optional, now required (breaking)
	if oldAttr.Computed && oldAttr.Optional && newAttr.Required {
		return "changed from computed+optional to required"
	}

	// Also check for type changes - handle both AttributeType and AttributeNestedType

	// First case: One uses AttributeType, another uses AttributeNestedType
	if (oldAttr.AttributeNestedType == nil && newAttr.AttributeNestedType != nil) ||
		(oldAttr.AttributeNestedType != nil && newAttr.AttributeNestedType == nil) {
		return "changed attribute structure type"
	}

	// Second case: Both use AttributeType, but types are different
	if oldAttr.AttributeNestedType == nil && newAttr.AttributeNestedType == nil &&
		!oldAttr.AttributeType.Equals(newAttr.AttributeType) {
		return fmt.Sprintf("type changed from %s to %s",
			oldAttr.AttributeType.FriendlyName(), newAttr.AttributeType.FriendlyName())
	}

	// Third case: Both use AttributeNestedType, but they're different
	if oldAttr.AttributeNestedType != nil && newAttr.AttributeNestedType != nil {
		// Simple check for different nested type kinds
		if oldAttr.AttributeNestedType.Attributes != nil != (newAttr.AttributeNestedType.Attributes != nil) ||
			oldAttr.AttributeNestedType.NestingMode != newAttr.AttributeNestedType.NestingMode {
			return "changed nested structure type"
		}
	}

	return ""
}

// attributeConfig converts a SchemaAttribute to an AttributeConfig
func attributeConfig(attr *tfjson.SchemaAttribute) *AttributeConfig {
	if attr == nil {
		return nil
	}

	var attrType string

	if attr.AttributeNestedType == nil {
		// Primitive type
		attrType = attr.AttributeType.FriendlyName()
	} else {
		// Collection or nested type
		switch attr.AttributeNestedType.NestingMode {
		case tfjson.SchemaNestingModeList:
			attrType = "List"
			if attr.AttributeNestedType.Attributes != nil {
				attrType += "[Object]" // List of objects
			} else {
				attrType += "[" + attr.AttributeType.FriendlyName() + "]" // List of primitives
			}
		case tfjson.SchemaNestingModeSet:
			attrType = "Set"
			if attr.AttributeNestedType.Attributes != nil {
				attrType += "[Object]" // Set of objects
			} else {
				attrType += "[" + attr.AttributeType.FriendlyName() + "]" // Set of primitives
			}
		case tfjson.SchemaNestingModeMap:
			attrType = "Map"
			if attr.AttributeNestedType.Attributes != nil {
				attrType += "[Object]" // Map of objects
			} else {
				attrType += "[" + attr.AttributeType.FriendlyName() + "]" // Map of primitives
			}
		case tfjson.SchemaNestingModeSingle:
			attrType = "Object"
		default:
			attrType = "Complex"
		}
	}

	return &AttributeConfig{
		AttributeType: attrType,
		Required:      attr.Required,
		Computed:      attr.Computed,
	}
}
