// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/open-farms/inventory/ent/tool"
)

// Tool is the model entity for the Tool schema.
type Tool struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Powered holds the value of the "powered" field.
	Powered bool `json:"powered,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Tool) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case tool.FieldPowered:
			values[i] = new(sql.NullBool)
		case tool.FieldID:
			values[i] = new(sql.NullInt64)
		case tool.FieldName:
			values[i] = new(sql.NullString)
		case tool.FieldCreateTime, tool.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Tool", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Tool fields.
func (t *Tool) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tool.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case tool.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				t.CreateTime = value.Time
			}
		case tool.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				t.UpdateTime = value.Time
			}
		case tool.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case tool.FieldPowered:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field powered", values[i])
			} else if value.Valid {
				t.Powered = value.Bool
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Tool.
// Note that you need to call Tool.Unwrap() before calling this method if this Tool
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Tool) Update() *ToolUpdateOne {
	return (&ToolClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Tool entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Tool) Unwrap() *Tool {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Tool is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Tool) String() string {
	var builder strings.Builder
	builder.WriteString("Tool(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(t.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(t.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(t.Name)
	builder.WriteString(", powered=")
	builder.WriteString(fmt.Sprintf("%v", t.Powered))
	builder.WriteByte(')')
	return builder.String()
}

// Tools is a parsable slice of Tool.
type Tools []*Tool

func (t Tools) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}