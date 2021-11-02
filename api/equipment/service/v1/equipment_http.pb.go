// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.0

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type EquipmentServiceHTTPServer interface {
	CreateEquipment(context.Context, *CreateEquipmentRequest) (*CreateEquipmentReply, error)
	DeleteEquipment(context.Context, *DeleteEquipmentRequest) (*DeleteEquipmentReply, error)
	GetEquipment(context.Context, *GetEquipmentRequest) (*GetEquipmentReply, error)
	ListEquipment(context.Context, *ListEquipmentRequest) (*ListEquipmentReply, error)
	UpdateEquipment(context.Context, *UpdateEquipmentRequest) (*UpdateEquipmentReply, error)
}

func RegisterEquipmentServiceHTTPServer(s *http.Server, srv EquipmentServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/equipment", _EquipmentService_CreateEquipment0_HTTP_Handler(srv))
	r.PUT("/equipment/{id}", _EquipmentService_UpdateEquipment0_HTTP_Handler(srv))
	r.DELETE("/equipment/{id}", _EquipmentService_DeleteEquipment0_HTTP_Handler(srv))
	r.GET("/equipment/{id}", _EquipmentService_GetEquipment0_HTTP_Handler(srv))
	r.GET("/equipment", _EquipmentService_ListEquipment0_HTTP_Handler(srv))
}

func _EquipmentService_CreateEquipment0_HTTP_Handler(srv EquipmentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateEquipmentRequest
		if err := ctx.Bind(&in.Equipment); err != nil {
			return err
		}
		http.SetOperation(ctx, "/equipment.service.v1.EquipmentService/CreateEquipment")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateEquipment(ctx, req.(*CreateEquipmentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateEquipmentReply)
		return ctx.Result(200, reply)
	}
}

func _EquipmentService_UpdateEquipment0_HTTP_Handler(srv EquipmentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateEquipmentRequest
		if err := ctx.Bind(&in.Equipment); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/equipment.service.v1.EquipmentService/UpdateEquipment")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateEquipment(ctx, req.(*UpdateEquipmentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateEquipmentReply)
		return ctx.Result(200, reply)
	}
}

func _EquipmentService_DeleteEquipment0_HTTP_Handler(srv EquipmentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteEquipmentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/equipment.service.v1.EquipmentService/DeleteEquipment")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteEquipment(ctx, req.(*DeleteEquipmentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteEquipmentReply)
		return ctx.Result(200, reply)
	}
}

func _EquipmentService_GetEquipment0_HTTP_Handler(srv EquipmentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetEquipmentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/equipment.service.v1.EquipmentService/GetEquipment")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetEquipment(ctx, req.(*GetEquipmentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetEquipmentReply)
		return ctx.Result(200, reply)
	}
}

func _EquipmentService_ListEquipment0_HTTP_Handler(srv EquipmentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListEquipmentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/equipment.service.v1.EquipmentService/ListEquipment")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListEquipment(ctx, req.(*ListEquipmentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListEquipmentReply)
		return ctx.Result(200, reply)
	}
}

type EquipmentServiceHTTPClient interface {
	CreateEquipment(ctx context.Context, req *CreateEquipmentRequest, opts ...http.CallOption) (rsp *CreateEquipmentReply, err error)
	DeleteEquipment(ctx context.Context, req *DeleteEquipmentRequest, opts ...http.CallOption) (rsp *DeleteEquipmentReply, err error)
	GetEquipment(ctx context.Context, req *GetEquipmentRequest, opts ...http.CallOption) (rsp *GetEquipmentReply, err error)
	ListEquipment(ctx context.Context, req *ListEquipmentRequest, opts ...http.CallOption) (rsp *ListEquipmentReply, err error)
	UpdateEquipment(ctx context.Context, req *UpdateEquipmentRequest, opts ...http.CallOption) (rsp *UpdateEquipmentReply, err error)
}

type EquipmentServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewEquipmentServiceHTTPClient(client *http.Client) EquipmentServiceHTTPClient {
	return &EquipmentServiceHTTPClientImpl{client}
}

func (c *EquipmentServiceHTTPClientImpl) CreateEquipment(ctx context.Context, in *CreateEquipmentRequest, opts ...http.CallOption) (*CreateEquipmentReply, error) {
	var out CreateEquipmentReply
	pattern := "/equipment"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/equipment.service.v1.EquipmentService/CreateEquipment"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.Equipment, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *EquipmentServiceHTTPClientImpl) DeleteEquipment(ctx context.Context, in *DeleteEquipmentRequest, opts ...http.CallOption) (*DeleteEquipmentReply, error) {
	var out DeleteEquipmentReply
	pattern := "/equipment/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/equipment.service.v1.EquipmentService/DeleteEquipment"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *EquipmentServiceHTTPClientImpl) GetEquipment(ctx context.Context, in *GetEquipmentRequest, opts ...http.CallOption) (*GetEquipmentReply, error) {
	var out GetEquipmentReply
	pattern := "/equipment/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/equipment.service.v1.EquipmentService/GetEquipment"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *EquipmentServiceHTTPClientImpl) ListEquipment(ctx context.Context, in *ListEquipmentRequest, opts ...http.CallOption) (*ListEquipmentReply, error) {
	var out ListEquipmentReply
	pattern := "/equipment"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/equipment.service.v1.EquipmentService/ListEquipment"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *EquipmentServiceHTTPClientImpl) UpdateEquipment(ctx context.Context, in *UpdateEquipmentRequest, opts ...http.CallOption) (*UpdateEquipmentReply, error) {
	var out UpdateEquipmentReply
	pattern := "/equipment/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/equipment.service.v1.EquipmentService/UpdateEquipment"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Equipment, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
