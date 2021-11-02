// Code generated by entc, DO NOT EDIT.

package equipment

import (
	"fmt"
)

const (
	// Label holds the string label denoting the equipment type in the database.
	Label = "equipment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// FieldCondition holds the string denoting the condition field in the database.
	FieldCondition = "condition"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the equipment in the database.
	Table = "equipment"
)

// Columns holds all SQL columns for equipment fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldTags,
	FieldCondition,
	FieldCreatedAt,
	FieldUpdatedAt,
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
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(uint64) error
)

// Condition defines the type for the "condition" enum field.
type Condition string

// Condition values.
const (
	ConditionCONDITION_UNSPECIFIED Condition = "CONDITION_UNSPECIFIED"
	ConditionCONDITION_MINT        Condition = "CONDITION_MINT"
	ConditionCONDITION_GOOD        Condition = "CONDITION_GOOD"
	ConditionCONDITION_POOR        Condition = "CONDITION_POOR"
	ConditionCONDITION_BROKEN      Condition = "CONDITION_BROKEN"
)

func (c Condition) String() string {
	return string(c)
}

// ConditionValidator is a validator for the "condition" field enum values. It is called by the builders before save.
func ConditionValidator(c Condition) error {
	switch c {
	case ConditionCONDITION_UNSPECIFIED, ConditionCONDITION_MINT, ConditionCONDITION_GOOD, ConditionCONDITION_POOR, ConditionCONDITION_BROKEN:
		return nil
	default:
		return fmt.Errorf("equipment: invalid enum value for condition field: %q", c)
	}
}