// Code generated by entc, DO NOT EDIT.

package vehicle

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the vehicle type in the database.
	Label = "vehicle"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMake holds the string denoting the make field in the database.
	FieldMake = "make"
	// FieldModel holds the string denoting the model field in the database.
	FieldModel = "model"
	// FieldYear holds the string denoting the year field in the database.
	FieldYear = "year"
	// FieldActive holds the string denoting the active field in the database.
	FieldActive = "active"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// FieldCondition holds the string denoting the condition field in the database.
	FieldCondition = "condition"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the vehicle in the database.
	Table = "vehicles"
)

// Columns holds all SQL columns for vehicle fields.
var Columns = []string{
	FieldID,
	FieldMake,
	FieldModel,
	FieldYear,
	FieldActive,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(uint64) error
)

// Condition defines the type for the "condition" enum field.
type Condition string

// Condition values.
const (
	ConditionUNSPECIFIED Condition = "UNSPECIFIED"
	ConditionMINT        Condition = "MINT"
	ConditionGOOD        Condition = "GOOD"
	ConditionPOOR        Condition = "POOR"
	ConditionBROKEN      Condition = "BROKEN"
)

func (c Condition) String() string {
	return string(c)
}

// ConditionValidator is a validator for the "condition" field enum values. It is called by the builders before save.
func ConditionValidator(c Condition) error {
	switch c {
	case ConditionUNSPECIFIED, ConditionMINT, ConditionGOOD, ConditionPOOR, ConditionBROKEN:
		return nil
	default:
		return fmt.Errorf("vehicle: invalid enum value for condition field: %q", c)
	}
}
