// Code generated by entc, DO NOT EDIT.

package http

import (
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/open-farms/inventory/ent"
	"go.uber.org/zap"
)

// NewHandler returns a ready to use handler with all generated endpoints mounted.
func NewHandler(c *ent.Client, l *zap.Logger) chi.Router {
	r := chi.NewRouter()
	MountRoutes(c, l, r)
	return r
}

// MountRoutes mounts all generated routes on the given router.
func MountRoutes(c *ent.Client, l *zap.Logger, r chi.Router) {
	NewCategoryHandler(c, l).MountRoutes(r)
	NewEquipmentHandler(c, l).MountRoutes(r)
	NewVehicleHandler(c, l).MountRoutes(r)
}

// CategoryHandler handles http crud operations on ent.Category.
type CategoryHandler struct {
	client *ent.Client
	log    *zap.Logger
}

func NewCategoryHandler(c *ent.Client, l *zap.Logger) *CategoryHandler {
	return &CategoryHandler{
		client: c,
		log:    l.With(zap.String("handler", "CategoryHandler")),
	}
}
func (h *CategoryHandler) MountCreateRoute(r chi.Router) *CategoryHandler {
	r.Post("/categories", h.Create)
	return h
}
func (h *CategoryHandler) MountReadRoute(r chi.Router) *CategoryHandler {
	r.Get("/categories/{id}", h.Read)
	return h
}
func (h *CategoryHandler) MountUpdateRoute(r chi.Router) *CategoryHandler {
	r.Patch("/categories/{id}", h.Update)
	return h
}
func (h *CategoryHandler) MountDeleteRoute(r chi.Router) *CategoryHandler {
	r.Delete("/categories/{id}", h.Delete)
	return h
}
func (h *CategoryHandler) MountListRoute(r chi.Router) *CategoryHandler {
	r.Get("/categories", h.List)
	return h
}
func (h *CategoryHandler) MountRoutes(r chi.Router) {
	h.MountCreateRoute(r).MountReadRoute(r).MountUpdateRoute(r).MountDeleteRoute(r).MountListRoute(r)
}

// EquipmentHandler handles http crud operations on ent.Equipment.
type EquipmentHandler struct {
	client *ent.Client
	log    *zap.Logger
}

func NewEquipmentHandler(c *ent.Client, l *zap.Logger) *EquipmentHandler {
	return &EquipmentHandler{
		client: c,
		log:    l.With(zap.String("handler", "EquipmentHandler")),
	}
}
func (h *EquipmentHandler) MountCreateRoute(r chi.Router) *EquipmentHandler {
	r.Post("/equipment", h.Create)
	return h
}
func (h *EquipmentHandler) MountReadRoute(r chi.Router) *EquipmentHandler {
	r.Get("/equipment/{id}", h.Read)
	return h
}
func (h *EquipmentHandler) MountUpdateRoute(r chi.Router) *EquipmentHandler {
	r.Patch("/equipment/{id}", h.Update)
	return h
}
func (h *EquipmentHandler) MountDeleteRoute(r chi.Router) *EquipmentHandler {
	r.Delete("/equipment/{id}", h.Delete)
	return h
}
func (h *EquipmentHandler) MountListRoute(r chi.Router) *EquipmentHandler {
	r.Get("/equipment", h.List)
	return h
}
func (h *EquipmentHandler) MountRoutes(r chi.Router) {
	h.MountCreateRoute(r).MountReadRoute(r).MountUpdateRoute(r).MountDeleteRoute(r).MountListRoute(r)
}

// VehicleHandler handles http crud operations on ent.Vehicle.
type VehicleHandler struct {
	client *ent.Client
	log    *zap.Logger
}

func NewVehicleHandler(c *ent.Client, l *zap.Logger) *VehicleHandler {
	return &VehicleHandler{
		client: c,
		log:    l.With(zap.String("handler", "VehicleHandler")),
	}
}
func (h *VehicleHandler) MountCreateRoute(r chi.Router) *VehicleHandler {
	r.Post("/vehicles", h.Create)
	return h
}
func (h *VehicleHandler) MountReadRoute(r chi.Router) *VehicleHandler {
	r.Get("/vehicles/{id}", h.Read)
	return h
}
func (h *VehicleHandler) MountUpdateRoute(r chi.Router) *VehicleHandler {
	r.Patch("/vehicles/{id}", h.Update)
	return h
}
func (h *VehicleHandler) MountDeleteRoute(r chi.Router) *VehicleHandler {
	r.Delete("/vehicles/{id}", h.Delete)
	return h
}
func (h *VehicleHandler) MountListRoute(r chi.Router) *VehicleHandler {
	r.Get("/vehicles", h.List)
	return h
}
func (h *VehicleHandler) MountRoutes(r chi.Router) {
	h.MountCreateRoute(r).MountReadRoute(r).MountUpdateRoute(r).MountDeleteRoute(r).MountListRoute(r)
}

func stripEntError(err error) string {
	return strings.TrimPrefix(err.Error(), "ent: ")
}

func zapFields(errs map[string]string) []zap.Field {
	if errs == nil || len(errs) == 0 {
		return nil
	}
	r := make([]zap.Field, 0)
	for k, v := range errs {
		r = append(r, zap.String(k, v))
	}
	return r
}
