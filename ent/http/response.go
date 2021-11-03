// Code generated by entc, DO NOT EDIT.

package http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/mailru/easyjson"
	"github.com/open-farms/inventory/ent"
	"github.com/open-farms/inventory/ent/vehicle"
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
	// Equipment822375389View represents the data serialized for the following serialization group combinations:
	// []
	Equipment822375389View struct {
		ID         int       `json:"id,omitempty"`
		Name       string    `json:"name,omitempty"`
		Condition  string    `json:"condition,omitempty"`
		CreateTime time.Time `json:"create_time,omitempty"`
		UpdateTime time.Time `json:"update_time,omitempty"`
	}
	Equipment822375389Views []*Equipment822375389View
)

func NewEquipment822375389View(e *ent.Equipment) *Equipment822375389View {
	if e == nil {
		return nil
	}
	return &Equipment822375389View{
		ID:         e.ID,
		Name:       e.Name,
		Condition:  e.Condition,
		CreateTime: e.CreateTime,
		UpdateTime: e.UpdateTime,
	}
}

func NewEquipment822375389Views(es []*ent.Equipment) Equipment822375389Views {
	if len(es) == 0 {
		return nil
	}
	r := make(Equipment822375389Views, len(es))
	for i, e := range es {
		r[i] = NewEquipment822375389View(e)
	}
	return r
}

type (
	// Vehicle2848838632View represents the data serialized for the following serialization group combinations:
	// []
	Vehicle2848838632View struct {
		ID         int64             `json:"id,omitempty"`
		Make       string            `json:"make,omitempty"`
		Model      string            `json:"model,omitempty"`
		Miles      int64             `json:"miles,omitempty"`
		Mpg        int64             `json:"mpg,omitempty"`
		Owner      string            `json:"owner,omitempty"`
		Year       string            `json:"year,omitempty"`
		Active     bool              `json:"active,omitempty"`
		Tags       []string          `json:"tags,omitempty"`
		Condition  vehicle.Condition `json:"condition,omitempty"`
		CreateTime time.Time         `json:"create_time,omitempty"`
		UpdateTime time.Time         `json:"update_time,omitempty"`
	}
	Vehicle2848838632Views []*Vehicle2848838632View
)

func NewVehicle2848838632View(e *ent.Vehicle) *Vehicle2848838632View {
	if e == nil {
		return nil
	}
	return &Vehicle2848838632View{
		ID:         e.ID,
		Make:       e.Make,
		Model:      e.Model,
		Miles:      e.Miles,
		Mpg:        e.Mpg,
		Owner:      e.Owner,
		Year:       e.Year,
		Active:     e.Active,
		Tags:       e.Tags,
		Condition:  e.Condition,
		CreateTime: e.CreateTime,
		UpdateTime: e.UpdateTime,
	}
}

func NewVehicle2848838632Views(es []*ent.Vehicle) Vehicle2848838632Views {
	if len(es) == 0 {
		return nil
	}
	r := make(Vehicle2848838632Views, len(es))
	for i, e := range es {
		r[i] = NewVehicle2848838632View(e)
	}
	return r
}
