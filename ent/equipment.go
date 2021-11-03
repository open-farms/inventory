// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/open-farms/inventory/ent/equipment"
)

// Equipment is the model entity for the Equipment schema.
type Equipment struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Condition holds the value of the "condition" field.
	Condition string `json:"condition,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Equipment) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case equipment.FieldID:
			values[i] = new(sql.NullInt64)
		case equipment.FieldName, equipment.FieldCondition:
			values[i] = new(sql.NullString)
		case equipment.FieldCreateTime, equipment.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Equipment", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Equipment fields.
func (e *Equipment) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case equipment.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = int(value.Int64)
		case equipment.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				e.CreateTime = value.Time
			}
		case equipment.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				e.UpdateTime = value.Time
			}
		case equipment.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				e.Name = value.String
			}
		case equipment.FieldCondition:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field condition", values[i])
			} else if value.Valid {
				e.Condition = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Equipment.
// Note that you need to call Equipment.Unwrap() before calling this method if this Equipment
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Equipment) Update() *EquipmentUpdateOne {
	return (&EquipmentClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the Equipment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Equipment) Unwrap() *Equipment {
	tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Equipment is not a transactional entity")
	}
	e.config.driver = tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Equipment) String() string {
	var builder strings.Builder
	builder.WriteString("Equipment(")
	builder.WriteString(fmt.Sprintf("id=%v", e.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(e.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(e.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(e.Name)
	builder.WriteString(", condition=")
	builder.WriteString(e.Condition)
	builder.WriteByte(')')
	return builder.String()
}

// EquipmentSlice is a parsable slice of Equipment.
type EquipmentSlice []*Equipment

func (e EquipmentSlice) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
