// Code generated by entc, DO NOT EDIT.

package http

import (
	"net/http"
	"strings"

	"github.com/mailru/easyjson"
	"github.com/open-farms/inventory/ent"
	"github.com/open-farms/inventory/ent/category"
	"github.com/open-farms/inventory/ent/equipment"
	"github.com/open-farms/inventory/ent/implement"
	"github.com/open-farms/inventory/ent/location"
	"github.com/open-farms/inventory/ent/tool"
	"github.com/open-farms/inventory/ent/vehicle"
	"go.uber.org/zap"
)

// Create creates a new ent.Category and stores it in the database.
func (h CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Create"))
	// Get the post data.
	var d CategoryCreateRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		BadRequest(w, "invalid json string")
		return
	}
	// Save the data.
	b := h.client.Category.Create()
	if d.CreateTime != nil {
		b.SetCreateTime(*d.CreateTime)
	}
	if d.UpdateTime != nil {
		b.SetUpdateTime(*d.UpdateTime)
	}
	if d.Name != nil {
		b.SetName(*d.Name)
	}
	e, err := b.Save(r.Context())
	if err != nil {
		switch {
		default:
			l.Error("could not create category", zap.Error(err))
			InternalServerError(w, nil)
		}
		return
	}
	// Store id of fresh entity to log errors for the reload.
	id := e.ID
	// Reload entry.
	q := h.client.Category.Query().Where(category.ID(e.ID))
	ret, err := q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could not read category", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("category rendered", zap.Int("id", id))
	easyjson.MarshalToHTTPResponseWriter(NewCategory1462705340View(ret), w)
}

// Create creates a new ent.Equipment and stores it in the database.
func (h EquipmentHandler) Create(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Create"))
	// Get the post data.
	var d EquipmentCreateRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		BadRequest(w, "invalid json string")
		return
	}
	// Validate the data.
	errs := make(map[string]string)
	if d.Name == nil {
		errs["name"] = `missing required field: "name"`
	}
	if d.Condition == nil {
		errs["condition"] = `missing required field: "condition"`
	} else if err := equipment.ConditionValidator(*d.Condition); err != nil {
		errs["condition"] = strings.TrimPrefix(err.Error(), "equipment: ")
	}
	if len(errs) > 0 {
		l.Info("validation failed", zapFields(errs)...)
		BadRequest(w, errs)
		return
	}
	// Save the data.
	b := h.client.Equipment.Create()
	if d.CreateTime != nil {
		b.SetCreateTime(*d.CreateTime)
	}
	if d.UpdateTime != nil {
		b.SetUpdateTime(*d.UpdateTime)
	}
	if d.Name != nil {
		b.SetName(*d.Name)
	}
	if d.Condition != nil {
		b.SetCondition(*d.Condition)
	}
	e, err := b.Save(r.Context())
	if err != nil {
		switch {
		default:
			l.Error("could not create equipment", zap.Error(err))
			InternalServerError(w, nil)
		}
		return
	}
	// Store id of fresh entity to log errors for the reload.
	id := e.ID
	// Reload entry.
	q := h.client.Equipment.Query().Where(equipment.ID(e.ID))
	ret, err := q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could not read equipment", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("equipment rendered", zap.Int("id", id))
	easyjson.MarshalToHTTPResponseWriter(NewEquipment3958372643View(ret), w)
}

// Create creates a new ent.Implement and stores it in the database.
func (h ImplementHandler) Create(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Create"))
	// Get the post data.
	var d ImplementCreateRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		BadRequest(w, "invalid json string")
		return
	}
	// Save the data.
	b := h.client.Implement.Create()
	if d.CreateTime != nil {
		b.SetCreateTime(*d.CreateTime)
	}
	if d.UpdateTime != nil {
		b.SetUpdateTime(*d.UpdateTime)
	}
	if d.Name != nil {
		b.SetName(*d.Name)
	}
	e, err := b.Save(r.Context())
	if err != nil {
		switch {
		default:
			l.Error("could not create implement", zap.Error(err))
			InternalServerError(w, nil)
		}
		return
	}
	// Store id of fresh entity to log errors for the reload.
	id := e.ID
	// Reload entry.
	q := h.client.Implement.Query().Where(implement.ID(e.ID))
	ret, err := q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could not read implement", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("implement rendered", zap.Int("id", id))
	easyjson.MarshalToHTTPResponseWriter(NewImplement1296325875View(ret), w)
}

// Create creates a new ent.Location and stores it in the database.
func (h LocationHandler) Create(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Create"))
	// Get the post data.
	var d LocationCreateRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		BadRequest(w, "invalid json string")
		return
	}
	// Save the data.
	b := h.client.Location.Create()
	if d.CreateTime != nil {
		b.SetCreateTime(*d.CreateTime)
	}
	if d.UpdateTime != nil {
		b.SetUpdateTime(*d.UpdateTime)
	}
	if d.Name != nil {
		b.SetName(*d.Name)
	}
	if d.Zone != nil {
		b.SetZone(*d.Zone)
	}
	if d.Vehicle != nil {
		b.AddVehicleIDs(d.Vehicle...)
	}
	e, err := b.Save(r.Context())
	if err != nil {
		switch {
		default:
			l.Error("could not create location", zap.Error(err))
			InternalServerError(w, nil)
		}
		return
	}
	// Store id of fresh entity to log errors for the reload.
	id := e.ID
	// Reload entry.
	q := h.client.Location.Query().Where(location.ID(e.ID))
	ret, err := q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could not read location", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("location rendered", zap.Int("id", id))
	easyjson.MarshalToHTTPResponseWriter(NewLocation948740745View(ret), w)
}

// Create creates a new ent.Tool and stores it in the database.
func (h ToolHandler) Create(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Create"))
	// Get the post data.
	var d ToolCreateRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		BadRequest(w, "invalid json string")
		return
	}
	// Save the data.
	b := h.client.Tool.Create()
	if d.CreateTime != nil {
		b.SetCreateTime(*d.CreateTime)
	}
	if d.UpdateTime != nil {
		b.SetUpdateTime(*d.UpdateTime)
	}
	if d.Name != nil {
		b.SetName(*d.Name)
	}
	if d.Powered != nil {
		b.SetPowered(*d.Powered)
	}
	e, err := b.Save(r.Context())
	if err != nil {
		switch {
		default:
			l.Error("could not create tool", zap.Error(err))
			InternalServerError(w, nil)
		}
		return
	}
	// Store id of fresh entity to log errors for the reload.
	id := e.ID
	// Reload entry.
	q := h.client.Tool.Query().Where(tool.ID(e.ID))
	ret, err := q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could not read tool", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("tool rendered", zap.Int("id", id))
	easyjson.MarshalToHTTPResponseWriter(NewTool178816486View(ret), w)
}

// Create creates a new ent.Vehicle and stores it in the database.
func (h VehicleHandler) Create(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Create"))
	// Get the post data.
	var d VehicleCreateRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		BadRequest(w, "invalid json string")
		return
	}
	// Validate the data.
	errs := make(map[string]string)
	if d.Make == nil {
		errs["make"] = `missing required field: "make"`
	}
	if d.Model == nil {
		errs["model"] = `missing required field: "model"`
	}
	if d.Hours != nil {
		if err := vehicle.HoursValidator(*d.Hours); err != nil {
			errs["hours"] = strings.TrimPrefix(err.Error(), "vehicle: ")
		}
	}
	if d.Power != nil {
		if err := vehicle.PowerValidator(*d.Power); err != nil {
			errs["power"] = strings.TrimPrefix(err.Error(), "vehicle: ")
		}
	}
	if d.Location == nil {
		errs["location"] = `missing required edge: "location"`
	}
	if len(errs) > 0 {
		l.Info("validation failed", zapFields(errs)...)
		BadRequest(w, errs)
		return
	}
	// Save the data.
	b := h.client.Vehicle.Create()
	if d.CreateTime != nil {
		b.SetCreateTime(*d.CreateTime)
	}
	if d.UpdateTime != nil {
		b.SetUpdateTime(*d.UpdateTime)
	}
	if d.Make != nil {
		b.SetMake(*d.Make)
	}
	if d.Model != nil {
		b.SetModel(*d.Model)
	}
	if d.Hours != nil {
		b.SetHours(*d.Hours)
	}
	if d.Year != nil {
		b.SetYear(*d.Year)
	}
	if d.Active != nil {
		b.SetActive(*d.Active)
	}
	if d.Power != nil {
		b.SetPower(*d.Power)
	}
	if d.Location != nil {
		b.SetLocationID(*d.Location)
	}
	e, err := b.Save(r.Context())
	if err != nil {
		switch {
		default:
			l.Error("could not create vehicle", zap.Error(err))
			InternalServerError(w, nil)
		}
		return
	}
	// Store id of fresh entity to log errors for the reload.
	id := e.ID
	// Reload entry.
	q := h.client.Vehicle.Query().Where(vehicle.ID(e.ID))
	ret, err := q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could not read vehicle", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("vehicle rendered", zap.Int("id", id))
	easyjson.MarshalToHTTPResponseWriter(NewVehicle1702989761View(ret), w)
}
