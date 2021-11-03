// Code generated by entc, DO NOT EDIT.

package http

import (
	"time"

	equipment "github.com/open-farms/inventory/ent/equipment"
	"github.com/open-farms/inventory/ent/vehicle"
)

// Payload of a ent.Equipment create request.
type EquipmentCreateRequest struct {
	Name       *string              `json:"name"`
	Tags       *[]string            `json:"tags"`
	Condition  *equipment.Condition `json:"condition"`
	CreateTime *time.Time           `json:"create_time"`
	UpdateTime *time.Time           `json:"update_time"`
}

// Payload of a ent.Equipment update request.
type EquipmentUpdateRequest struct {
	Name       *string              `json:"name"`
	Tags       *[]string            `json:"tags"`
	Condition  *equipment.Condition `json:"condition"`
	UpdateTime *time.Time           `json:"update_time"`
}

// Payload of a ent.Vehicle create request.
type VehicleCreateRequest struct {
	Make       *string            `json:"make"`
	Model      *string            `json:"model"`
	Miles      *int64             `json:"miles"`
	Mpg        *int64             `json:"mpg"`
	Owner      *string            `json:"owner"`
	Year       *string            `json:"year"`
	Active     *bool              `json:"active"`
	Tags       *[]string          `json:"tags"`
	Condition  *vehicle.Condition `json:"condition"`
	CreateTime *time.Time         `json:"create_time"`
	UpdateTime *time.Time         `json:"update_time"`
}

// Payload of a ent.Vehicle update request.
type VehicleUpdateRequest struct {
	Make       *string            `json:"make"`
	Model      *string            `json:"model"`
	Miles      *int64             `json:"miles"`
	Mpg        *int64             `json:"mpg"`
	Owner      *string            `json:"owner"`
	Year       *string            `json:"year"`
	Active     *bool              `json:"active"`
	Tags       *[]string          `json:"tags"`
	Condition  *vehicle.Condition `json:"condition"`
	UpdateTime *time.Time         `json:"update_time"`
}
