// Code generated by entc, DO NOT EDIT.

package http

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/mailru/easyjson"
	"github.com/open-farms/inventory/ent"
	"github.com/open-farms/inventory/ent/category"
	"github.com/open-farms/inventory/ent/equipment"
	"github.com/open-farms/inventory/ent/vehicle"
	"go.uber.org/zap"
)

// Update updates a given ent.Category and saves the changes to the database.
func (h CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Update"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be an integer")
		return
	}
	// Get the post data.
	var d CategoryUpdateRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		BadRequest(w, "invalid json string")
		return
	}
	// Save the data.
	b := h.client.Category.UpdateOneID(id)
	if d.Name != nil {
		b.SetName(*d.Name)
	}
	// Store in database.
	e, err := b.Save(r.Context())
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
			l.Error("could-not-update-category", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	// Reload entry.
	q := h.client.Category.Query().Where(category.ID(e.ID))
	e, err = q.Only(r.Context())
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
			l.Error("could-not-read-category", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("category rendered", zap.Int("id", id))
	easyjson.MarshalToHTTPResponseWriter(NewCategory4094953247View(e), w)
}

// Update updates a given ent.Equipment and saves the changes to the database.
func (h EquipmentHandler) Update(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Update"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be an integer")
		return
	}
	// Get the post data.
	var d EquipmentUpdateRequest
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
	b := h.client.Equipment.UpdateOneID(id)
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
	// Store in database.
	e, err := b.Save(r.Context())
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
			l.Error("could-not-update-equipment", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	// Reload entry.
	q := h.client.Equipment.Query().Where(equipment.ID(e.ID))
	e, err = q.Only(r.Context())
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
			l.Error("could-not-read-equipment", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("equipment rendered", zap.Int("id", id))
	easyjson.MarshalToHTTPResponseWriter(NewEquipment3958372643View(e), w)
}

// Update updates a given ent.Vehicle and saves the changes to the database.
func (h VehicleHandler) Update(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Update"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be an integer")
		return
	}
	// Get the post data.
	var d VehicleUpdateRequest
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
	if d.Condition != nil {
		if err := vehicle.ConditionValidator(*d.Condition); err != nil {
			errs["condition"] = strings.TrimPrefix(err.Error(), "vehicle: ")
		}
	}
	if len(errs) > 0 {
		l.Info("validation failed", zapFields(errs)...)
		BadRequest(w, errs)
		return
	}
	// Save the data.
	b := h.client.Vehicle.UpdateOneID(id)
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
	if d.Miles != nil {
		b.SetMiles(*d.Miles)
	}
	if d.Mpg != nil {
		b.SetMpg(*d.Mpg)
	}
	if d.Owner != nil {
		b.SetOwner(*d.Owner)
	}
	if d.Year != nil {
		b.SetYear(*d.Year)
	}
	if d.Active != nil {
		b.SetActive(*d.Active)
	}
	if d.Condition != nil {
		b.SetCondition(*d.Condition)
	}
	// Store in database.
	e, err := b.Save(r.Context())
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
			l.Error("could-not-update-vehicle", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	// Reload entry.
	q := h.client.Vehicle.Query().Where(vehicle.ID(e.ID))
	e, err = q.Only(r.Context())
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
			l.Error("could-not-read-vehicle", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("vehicle rendered", zap.Int("id", id))
	easyjson.MarshalToHTTPResponseWriter(NewVehicle2530256765View(e), w)
}
