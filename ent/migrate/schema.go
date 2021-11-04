// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
	}
	// EquipmentColumns holds the columns for the "equipment" table.
	EquipmentColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "condition", Type: field.TypeString},
		{Name: "category_equipment", Type: field.TypeInt, Nullable: true},
		{Name: "location_equipment", Type: field.TypeInt, Nullable: true},
	}
	// EquipmentTable holds the schema information for the "equipment" table.
	EquipmentTable = &schema.Table{
		Name:       "equipment",
		Columns:    EquipmentColumns,
		PrimaryKey: []*schema.Column{EquipmentColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "equipment_categories_equipment",
				Columns:    []*schema.Column{EquipmentColumns[5]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "equipment_locations_equipment",
				Columns:    []*schema.Column{EquipmentColumns[6]},
				RefColumns: []*schema.Column{LocationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ImplementsColumns holds the columns for the "implements" table.
	ImplementsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "category_implement", Type: field.TypeInt, Nullable: true},
		{Name: "location_implement", Type: field.TypeInt, Nullable: true},
	}
	// ImplementsTable holds the schema information for the "implements" table.
	ImplementsTable = &schema.Table{
		Name:       "implements",
		Columns:    ImplementsColumns,
		PrimaryKey: []*schema.Column{ImplementsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "implements_categories_implement",
				Columns:    []*schema.Column{ImplementsColumns[4]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "implements_locations_implement",
				Columns:    []*schema.Column{ImplementsColumns[5]},
				RefColumns: []*schema.Column{LocationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// LocationsColumns holds the columns for the "locations" table.
	LocationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "zone", Type: field.TypeInt32, Unique: true},
		{Name: "category_location", Type: field.TypeInt, Nullable: true},
	}
	// LocationsTable holds the schema information for the "locations" table.
	LocationsTable = &schema.Table{
		Name:       "locations",
		Columns:    LocationsColumns,
		PrimaryKey: []*schema.Column{LocationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "locations_categories_location",
				Columns:    []*schema.Column{LocationsColumns[5]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ToolsColumns holds the columns for the "tools" table.
	ToolsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "powered", Type: field.TypeBool, Default: false},
		{Name: "category_tool", Type: field.TypeInt, Nullable: true},
		{Name: "location_tool", Type: field.TypeInt, Nullable: true},
	}
	// ToolsTable holds the schema information for the "tools" table.
	ToolsTable = &schema.Table{
		Name:       "tools",
		Columns:    ToolsColumns,
		PrimaryKey: []*schema.Column{ToolsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tools_categories_tool",
				Columns:    []*schema.Column{ToolsColumns[5]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "tools_locations_tool",
				Columns:    []*schema.Column{ToolsColumns[6]},
				RefColumns: []*schema.Column{LocationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// VehiclesColumns holds the columns for the "vehicles" table.
	VehiclesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "make", Type: field.TypeString},
		{Name: "model", Type: field.TypeString},
		{Name: "hours", Type: field.TypeInt64, Default: 0},
		{Name: "year", Type: field.TypeInt64, Nullable: true},
		{Name: "active", Type: field.TypeBool, Default: true},
		{Name: "power", Type: field.TypeString, Nullable: true, Default: "GAS"},
		{Name: "category_vehicle", Type: field.TypeInt, Nullable: true},
		{Name: "location_vehicle", Type: field.TypeInt, Nullable: true},
	}
	// VehiclesTable holds the schema information for the "vehicles" table.
	VehiclesTable = &schema.Table{
		Name:       "vehicles",
		Columns:    VehiclesColumns,
		PrimaryKey: []*schema.Column{VehiclesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "vehicles_categories_vehicle",
				Columns:    []*schema.Column{VehiclesColumns[9]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "vehicles_locations_vehicle",
				Columns:    []*schema.Column{VehiclesColumns[10]},
				RefColumns: []*schema.Column{LocationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CategoriesTable,
		EquipmentTable,
		ImplementsTable,
		LocationsTable,
		ToolsTable,
		VehiclesTable,
	}
)

func init() {
	EquipmentTable.ForeignKeys[0].RefTable = CategoriesTable
	EquipmentTable.ForeignKeys[1].RefTable = LocationsTable
	ImplementsTable.ForeignKeys[0].RefTable = CategoriesTable
	ImplementsTable.ForeignKeys[1].RefTable = LocationsTable
	LocationsTable.ForeignKeys[0].RefTable = CategoriesTable
	ToolsTable.ForeignKeys[0].RefTable = CategoriesTable
	ToolsTable.ForeignKeys[1].RefTable = LocationsTable
	VehiclesTable.ForeignKeys[0].RefTable = CategoriesTable
	VehiclesTable.ForeignKeys[1].RefTable = LocationsTable
}
