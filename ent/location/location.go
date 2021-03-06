// Code generated by entc, DO NOT EDIT.

package location

import (
	"time"
)

const (
	// Label holds the string label denoting the location type in the database.
	Label = "location"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldZone holds the string denoting the zone field in the database.
	FieldZone = "zone"
	// EdgeVehicle holds the string denoting the vehicle edge name in mutations.
	EdgeVehicle = "vehicle"
	// EdgeTool holds the string denoting the tool edge name in mutations.
	EdgeTool = "tool"
	// EdgeImplement holds the string denoting the implement edge name in mutations.
	EdgeImplement = "implement"
	// EdgeEquipment holds the string denoting the equipment edge name in mutations.
	EdgeEquipment = "equipment"
	// EdgeCategory holds the string denoting the category edge name in mutations.
	EdgeCategory = "category"
	// Table holds the table name of the location in the database.
	Table = "locations"
	// VehicleTable is the table that holds the vehicle relation/edge.
	VehicleTable = "vehicles"
	// VehicleInverseTable is the table name for the Vehicle entity.
	// It exists in this package in order to avoid circular dependency with the "vehicle" package.
	VehicleInverseTable = "vehicles"
	// VehicleColumn is the table column denoting the vehicle relation/edge.
	VehicleColumn = "location_vehicle"
	// ToolTable is the table that holds the tool relation/edge.
	ToolTable = "tools"
	// ToolInverseTable is the table name for the Tool entity.
	// It exists in this package in order to avoid circular dependency with the "tool" package.
	ToolInverseTable = "tools"
	// ToolColumn is the table column denoting the tool relation/edge.
	ToolColumn = "location_tool"
	// ImplementTable is the table that holds the implement relation/edge.
	ImplementTable = "implements"
	// ImplementInverseTable is the table name for the Implement entity.
	// It exists in this package in order to avoid circular dependency with the "implement" package.
	ImplementInverseTable = "implements"
	// ImplementColumn is the table column denoting the implement relation/edge.
	ImplementColumn = "location_implement"
	// EquipmentTable is the table that holds the equipment relation/edge.
	EquipmentTable = "equipment"
	// EquipmentInverseTable is the table name for the Equipment entity.
	// It exists in this package in order to avoid circular dependency with the "equipment" package.
	EquipmentInverseTable = "equipment"
	// EquipmentColumn is the table column denoting the equipment relation/edge.
	EquipmentColumn = "location_equipment"
	// CategoryTable is the table that holds the category relation/edge.
	CategoryTable = "locations"
	// CategoryInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	CategoryInverseTable = "categories"
	// CategoryColumn is the table column denoting the category relation/edge.
	CategoryColumn = "category_location"
)

// Columns holds all SQL columns for location fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldName,
	FieldZone,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "locations"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"category_location",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
)
