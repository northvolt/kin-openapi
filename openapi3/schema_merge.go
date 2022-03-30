package openapi3

import (
	"errors"
	"fmt"

	"github.com/mitchellh/copystructure"
)

// merAllOf helps validation of AllOf.
// To start simple:
// - no support for nested AllOf
// - no mix of AnyOf, OneOf, Not
// - all schemas of the same type
// lots of todos
func mergeAllOf(refs SchemaRefs) (*Schema, error) {
	if len(refs) < 2 {
		return nil, errors.New("less than 2 schemas to merge")
	}

	x, err := copystructure.Copy(refs[0].Value)
	if err != nil {
		return nil, err
	}
	merged := x.(*Schema)

	// initialize pointer slice and map
	if merged.Extensions == nil {
		merged.Extensions = make(map[string]interface{})
	}

	if merged.Description != "" && refs[0].Ref != "" {
		merged.Description = fmt.Sprintf("%s: %s.", refs[0].Ref, refs[0].Value.Description)
	}

	if merged.Enum == nil {
		merged.Enum = make([]interface{}, 0)
	}

	if merged.Required == nil {
		merged.Required = make([]string, 0)
	}

	if merged.Properties == nil {
		merged.Properties = make(Schemas)
	}

	for _, ref := range refs[1:] {
		if ref.Value != nil {
			if merged.Type == "" && ref.Value.Type != "" {
				merged.Type = ref.Value.Type
			}

			if merged.Title == "" && ref.Value.Title != "" {
				merged.Title = ref.Value.Title
			}

			if merged.Format == "" && ref.Value.Format != "" {
				merged.Format = ref.Value.Format
			}

			merged.Description += fmt.Sprintf("%s: %s.", ref.Ref, ref.Value.Description)

			merged.Enum = append(merged.Enum, ref.Value.Enum...)

			// merge default and example, return error if type is not object

			merged.Required = append(merged.Required, ref.Value.Required...)

			for k, v := range ref.Value.Properties {
				if _, ok := merged.Properties[k]; ok {
					return nil, fmt.Errorf("duplicated properties: %q", k)
				}

				merged.Properties[k] = v
			}
		}
	}

	return merged, nil
}
