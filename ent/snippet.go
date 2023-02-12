// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/nanaminakano/snippeter/ent/snippet"
)

// Snippet is the model entity for the Snippet schema.
type Snippet struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Snippet) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case snippet.FieldContent:
			values[i] = new(sql.NullString)
		case snippet.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case snippet.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Snippet", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Snippet fields.
func (s *Snippet) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case snippet.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case snippet.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				s.Content = value.String
			}
		case snippet.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Snippet.
// Note that you need to call Snippet.Unwrap() before calling this method if this Snippet
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Snippet) Update() *SnippetUpdateOne {
	return NewSnippetClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Snippet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Snippet) Unwrap() *Snippet {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Snippet is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Snippet) String() string {
	var builder strings.Builder
	builder.WriteString("Snippet(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("content=")
	builder.WriteString(s.Content)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Snippets is a parsable slice of Snippet.
type Snippets []*Snippet

func (s Snippets) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
