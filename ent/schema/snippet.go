package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Snippet holds the schema definition for the Snippet entity.
type Snippet struct {
	ent.Schema
}

// Fields of the Snippet.
func (Snippet) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()),
		field.String("content").NotEmpty(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Snippet.
func (Snippet) Edges() []ent.Edge {
	return nil
}
