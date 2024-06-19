// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/sysadminsmedia/homebox/backend/internal/data/ent/attachment"
	"github.com/sysadminsmedia/homebox/backend/internal/data/ent/document"
	"github.com/sysadminsmedia/homebox/backend/internal/data/ent/item"
)

// Attachment is the model entity for the Attachment schema.
type Attachment struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Type holds the value of the "type" field.
	Type attachment.Type `json:"type,omitempty"`
	// Primary holds the value of the "primary" field.
	Primary bool `json:"primary,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AttachmentQuery when eager-loading is set.
	Edges                AttachmentEdges `json:"edges"`
	document_attachments *uuid.UUID
	item_attachments     *uuid.UUID
	selectValues         sql.SelectValues
}

// AttachmentEdges holds the relations/edges for other nodes in the graph.
type AttachmentEdges struct {
	// Item holds the value of the item edge.
	Item *Item `json:"item,omitempty"`
	// Document holds the value of the document edge.
	Document *Document `json:"document,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ItemOrErr returns the Item value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AttachmentEdges) ItemOrErr() (*Item, error) {
	if e.loadedTypes[0] {
		if e.Item == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: item.Label}
		}
		return e.Item, nil
	}
	return nil, &NotLoadedError{edge: "item"}
}

// DocumentOrErr returns the Document value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AttachmentEdges) DocumentOrErr() (*Document, error) {
	if e.loadedTypes[1] {
		if e.Document == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: document.Label}
		}
		return e.Document, nil
	}
	return nil, &NotLoadedError{edge: "document"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Attachment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case attachment.FieldPrimary:
			values[i] = new(sql.NullBool)
		case attachment.FieldType:
			values[i] = new(sql.NullString)
		case attachment.FieldCreatedAt, attachment.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case attachment.FieldID:
			values[i] = new(uuid.UUID)
		case attachment.ForeignKeys[0]: // document_attachments
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case attachment.ForeignKeys[1]: // item_attachments
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Attachment fields.
func (a *Attachment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case attachment.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case attachment.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case attachment.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case attachment.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				a.Type = attachment.Type(value.String)
			}
		case attachment.FieldPrimary:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field primary", values[i])
			} else if value.Valid {
				a.Primary = value.Bool
			}
		case attachment.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field document_attachments", values[i])
			} else if value.Valid {
				a.document_attachments = new(uuid.UUID)
				*a.document_attachments = *value.S.(*uuid.UUID)
			}
		case attachment.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field item_attachments", values[i])
			} else if value.Valid {
				a.item_attachments = new(uuid.UUID)
				*a.item_attachments = *value.S.(*uuid.UUID)
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Attachment.
// This includes values selected through modifiers, order, etc.
func (a *Attachment) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryItem queries the "item" edge of the Attachment entity.
func (a *Attachment) QueryItem() *ItemQuery {
	return NewAttachmentClient(a.config).QueryItem(a)
}

// QueryDocument queries the "document" edge of the Attachment entity.
func (a *Attachment) QueryDocument() *DocumentQuery {
	return NewAttachmentClient(a.config).QueryDocument(a)
}

// Update returns a builder for updating this Attachment.
// Note that you need to call Attachment.Unwrap() before calling this method if this Attachment
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Attachment) Update() *AttachmentUpdateOne {
	return NewAttachmentClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Attachment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Attachment) Unwrap() *Attachment {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Attachment is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Attachment) String() string {
	var builder strings.Builder
	builder.WriteString("Attachment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", a.Type))
	builder.WriteString(", ")
	builder.WriteString("primary=")
	builder.WriteString(fmt.Sprintf("%v", a.Primary))
	builder.WriteByte(')')
	return builder.String()
}

// Attachments is a parsable slice of Attachment.
type Attachments []*Attachment
