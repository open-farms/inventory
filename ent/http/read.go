// Code generated by entc, DO NOT EDIT.

package http

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mailru/easyjson"
	"github.com/open-farms/inventory/ent"
	equipment "github.com/open-farms/inventory/ent/equipment"
	"github.com/open-farms/inventory/ent/vehicle"
	"go.uber.org/zap"
)

// Read fetches the ent.Equipment identified by a given url-parameter from the
// database and renders it to the client.
func (h *EquipmentHandler) Read(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Read"))
	// ID is URL parameter.
	id64, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 0)
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be an integer greater zero")
		return
	}
	id := int64(id64)
	// Create the query to fetch the Equipment
	q := h.client.Equipment.Query().Where(equipment.ID(id))
	e, err := q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int64("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int64("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could not read equipment", zap.Error(err), zap.Int64("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("equipment rendered", zap.Int64("id", id))
	easyjson.MarshalToHTTPResponseWriter(NewEquipment2075188150View(e), w)
}

// Read fetches the ent.Vehicle identified by a given url-parameter from the
// database and renders it to the client.
func (h *VehicleHandler) Read(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Read"))
	// ID is URL parameter.
	id64, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 0)
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be an integer greater zero")
		return
	}
	id := int64(id64)
	// Create the query to fetch the Vehicle
	q := h.client.Vehicle.Query().Where(vehicle.ID(id))
	e, err := q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int64("id", id))
			NotFound(w, msg)
		case ent.IsNotSingular(err):
			msg := stripEntError(err)
			l.Error(msg, zap.Error(err), zap.Int64("id", id))
			BadRequest(w, msg)
		default:
			l.Error("could not read vehicle", zap.Error(err), zap.Int64("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("vehicle rendered", zap.Int64("id", id))
	easyjson.MarshalToHTTPResponseWriter(NewVehicle2848838632View(e), w)
}
