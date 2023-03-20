// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/nanaminakano/snippeter/models/ent/schema"
	"github.com/nanaminakano/snippeter/models/ent/snippet"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	snippetFields := schema.Snippet{}.Fields()
	_ = snippetFields
	// snippetDescContent is the schema descriptor for content field.
	snippetDescContent := snippetFields[1].Descriptor()
	// snippet.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	snippet.ContentValidator = snippetDescContent.Validators[0].(func(string) error)
	// snippetDescCreatedAt is the schema descriptor for created_at field.
	snippetDescCreatedAt := snippetFields[2].Descriptor()
	// snippet.DefaultCreatedAt holds the default value on creation for the created_at field.
	snippet.DefaultCreatedAt = snippetDescCreatedAt.Default.(func() time.Time)
}