// Code generated by entc, DO NOT EDIT.

package http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/mailru/easyjson"
	"github.com/open-farms/inventory/ent"
)

// Basic HTTP Error Response
type ErrResponse struct {
	Code   int         `json:"code"`             // http response status code
	Status string      `json:"status"`           // user-level status message
	Errors interface{} `json:"errors,omitempty"` // application-level error
}

func (e ErrResponse) MarshalToHTTPResponseWriter(w http.ResponseWriter) (int, error) {
	d, err := easyjson.Marshal(e)
	if err != nil {
		return 0, err
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(d)))
	w.WriteHeader(e.Code)
	return w.Write(d)
}

func BadRequest(w http.ResponseWriter, msg interface{}) (int, error) {
	return ErrResponse{
		Code:   http.StatusBadRequest,
		Status: http.StatusText(http.StatusBadRequest),
		Errors: msg,
	}.MarshalToHTTPResponseWriter(w)
}

func Conflict(w http.ResponseWriter, msg interface{}) (int, error) {
	return ErrResponse{
		Code:   http.StatusConflict,
		Status: http.StatusText(http.StatusConflict),
		Errors: msg,
	}.MarshalToHTTPResponseWriter(w)
}

func Forbidden(w http.ResponseWriter, msg interface{}) (int, error) {
	return ErrResponse{
		Code:   http.StatusForbidden,
		Status: http.StatusText(http.StatusForbidden),
		Errors: msg,
	}.MarshalToHTTPResponseWriter(w)
}

func InternalServerError(w http.ResponseWriter, msg interface{}) (int, error) {
	return ErrResponse{
		Code:   http.StatusInternalServerError,
		Status: http.StatusText(http.StatusInternalServerError),
		Errors: msg,
	}.MarshalToHTTPResponseWriter(w)
}

func NotFound(w http.ResponseWriter, msg interface{}) (int, error) {
	return ErrResponse{
		Code:   http.StatusNotFound,
		Status: http.StatusText(http.StatusNotFound),
		Errors: msg,
	}.MarshalToHTTPResponseWriter(w)
}

func Unauthorized(w http.ResponseWriter, msg interface{}) (int, error) {
	return ErrResponse{
		Code:   http.StatusUnauthorized,
		Status: http.StatusText(http.StatusUnauthorized),
		Errors: msg,
	}.MarshalToHTTPResponseWriter(w)
}

type (
	// Category1462705340View represents the data serialized for the following serialization group combinations:
	// []
	Category1462705340View struct {
		ID         int       `json:"id,omitempty"`
		CreateTime time.Time `json:"create_time,omitempty"`
		UpdateTime time.Time `json:"update_time,omitempty"`
		Name       string    `json:"name,omitempty"`
	}
	Category1462705340Views []*Category1462705340View
)

func NewCategory1462705340View(e *ent.Category) *Category1462705340View {
	if e == nil {
		return nil
	}
	return &Category1462705340View{
		ID:         e.ID,
		CreateTime: e.CreateTime,
		UpdateTime: e.UpdateTime,
		Name:       e.Name,
	}
}

func NewCategory1462705340Views(es []*ent.Category) Category1462705340Views {
	if len(es) == 0 {
		return nil
	}
	r := make(Category1462705340Views, len(es))
	for i, e := range es {
		r[i] = NewCategory1462705340View(e)
	}
	return r
}

type (
	// Equipment3958372643View represents the data serialized for the following serialization group combinations:
	// []
	Equipment3958372643View struct {
		ID         int       `json:"id,omitempty"`
		CreateTime time.Time `json:"create_time,omitempty"`
		UpdateTime time.Time `json:"update_time,omitempty"`
		Name       string    `json:"name,omitempty"`
		Condition  string    `json:"condition,omitempty"`
	}
	Equipment3958372643Views []*Equipment3958372643View
)

func NewEquipment3958372643View(e *ent.Equipment) *Equipment3958372643View {
	if e == nil {
		return nil
	}
	return &Equipment3958372643View{
		ID:         e.ID,
		CreateTime: e.CreateTime,
		UpdateTime: e.UpdateTime,
		Name:       e.Name,
		Condition:  e.Condition,
	}
}

func NewEquipment3958372643Views(es []*ent.Equipment) Equipment3958372643Views {
	if len(es) == 0 {
		return nil
	}
	r := make(Equipment3958372643Views, len(es))
	for i, e := range es {
		r[i] = NewEquipment3958372643View(e)
	}
	return r
}

type (
	// Implement1296325875View represents the data serialized for the following serialization group combinations:
	// []
	Implement1296325875View struct {
		ID         int       `json:"id,omitempty"`
		CreateTime time.Time `json:"create_time,omitempty"`
		UpdateTime time.Time `json:"update_time,omitempty"`
		Name       string    `json:"name,omitempty"`
	}
	Implement1296325875Views []*Implement1296325875View
)

func NewImplement1296325875View(e *ent.Implement) *Implement1296325875View {
	if e == nil {
		return nil
	}
	return &Implement1296325875View{
		ID:         e.ID,
		CreateTime: e.CreateTime,
		UpdateTime: e.UpdateTime,
		Name:       e.Name,
	}
}

func NewImplement1296325875Views(es []*ent.Implement) Implement1296325875Views {
	if len(es) == 0 {
		return nil
	}
	r := make(Implement1296325875Views, len(es))
	for i, e := range es {
		r[i] = NewImplement1296325875View(e)
	}
	return r
}

type (
	// Location948740745View represents the data serialized for the following serialization group combinations:
	// []
	Location948740745View struct {
		ID         int       `json:"id,omitempty"`
		CreateTime time.Time `json:"create_time,omitempty"`
		UpdateTime time.Time `json:"update_time,omitempty"`
		Name       string    `json:"name,omitempty"`
		Zone       int32     `json:"zone,omitempty"`
	}
	Location948740745Views []*Location948740745View
)

func NewLocation948740745View(e *ent.Location) *Location948740745View {
	if e == nil {
		return nil
	}
	return &Location948740745View{
		ID:         e.ID,
		CreateTime: e.CreateTime,
		UpdateTime: e.UpdateTime,
		Name:       e.Name,
		Zone:       e.Zone,
	}
}

func NewLocation948740745Views(es []*ent.Location) Location948740745Views {
	if len(es) == 0 {
		return nil
	}
	r := make(Location948740745Views, len(es))
	for i, e := range es {
		r[i] = NewLocation948740745View(e)
	}
	return r
}

type (
	// Tool178816486View represents the data serialized for the following serialization group combinations:
	// []
	Tool178816486View struct {
		ID         int       `json:"id,omitempty"`
		CreateTime time.Time `json:"create_time,omitempty"`
		UpdateTime time.Time `json:"update_time,omitempty"`
		Name       string    `json:"name,omitempty"`
		Powered    bool      `json:"powered,omitempty"`
	}
	Tool178816486Views []*Tool178816486View
)

func NewTool178816486View(e *ent.Tool) *Tool178816486View {
	if e == nil {
		return nil
	}
	return &Tool178816486View{
		ID:         e.ID,
		CreateTime: e.CreateTime,
		UpdateTime: e.UpdateTime,
		Name:       e.Name,
		Powered:    e.Powered,
	}
}

func NewTool178816486Views(es []*ent.Tool) Tool178816486Views {
	if len(es) == 0 {
		return nil
	}
	r := make(Tool178816486Views, len(es))
	for i, e := range es {
		r[i] = NewTool178816486View(e)
	}
	return r
}

type (
	// Vehicle1702989761View represents the data serialized for the following serialization group combinations:
	// []
	Vehicle1702989761View struct {
		ID         int       `json:"id,omitempty"`
		CreateTime time.Time `json:"create_time,omitempty"`
		UpdateTime time.Time `json:"update_time,omitempty"`
		Make       string    `json:"make,omitempty"`
		Model      string    `json:"model,omitempty"`
		Hours      int64     `json:"hours,omitempty"`
		Year       int64     `json:"year,omitempty"`
		Active     bool      `json:"active,omitempty"`
		Power      string    `json:"power,omitempty"`
	}
	Vehicle1702989761Views []*Vehicle1702989761View
)

func NewVehicle1702989761View(e *ent.Vehicle) *Vehicle1702989761View {
	if e == nil {
		return nil
	}
	return &Vehicle1702989761View{
		ID:         e.ID,
		CreateTime: e.CreateTime,
		UpdateTime: e.UpdateTime,
		Make:       e.Make,
		Model:      e.Model,
		Hours:      e.Hours,
		Year:       e.Year,
		Active:     e.Active,
		Power:      e.Power,
	}
}

func NewVehicle1702989761Views(es []*ent.Vehicle) Vehicle1702989761Views {
	if len(es) == 0 {
		return nil
	}
	r := make(Vehicle1702989761Views, len(es))
	for i, e := range es {
		r[i] = NewVehicle1702989761View(e)
	}
	return r
}
