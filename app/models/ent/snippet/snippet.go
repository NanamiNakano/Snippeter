// Code generated by ent, DO NOT EDIT.

package snippet

import (
	"time"
)

const (
	// Label holds the string label denoting the snippet type in the database.
	Label = "snippet"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the snippet in the database.
	Table = "snippets"
)

// Columns holds all SQL columns for snippet fields.
var Columns = []string{
	FieldID,
	FieldContent,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// ContentValidator is a validator for the "content" field. It is called by the builders before save.
	ContentValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)