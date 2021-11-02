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

type VehicleServiceHTTPServer interface {
	CreateVehicle(context.Context, *CreateVehicleRequest) (*CreateVehicleResponse, error)
	DeleteVehicle(context.Context, *DeleteVehicleRequest) (*DeleteVehicleResponse, error)
	GetVehicle(context.Context, *GetVehicleRequest) (*GetVehicleResponse, error)
	ListVehicle(context.Context, *ListVehicleRequest) (*ListVehicleResponse, error)
	UpdateVehicle(context.Context, *UpdateVehicleRequest) (*UpdateVehicleResponse, error)
}

func RegisterVehicleServiceHTTPServer(s *http.Server, srv VehicleServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/vehicles", _VehicleService_CreateVehicle0_HTTP_Handler(srv))
	r.PUT("/v1/vehicles/{id}", _VehicleService_UpdateVehicle0_HTTP_Handler(srv))
	r.DELETE("/v1/vehicles/{id}", _VehicleService_DeleteVehicle0_HTTP_Handler(srv))
	r.GET("/v1/vehicles/{id}", _VehicleService_GetVehicle0_HTTP_Handler(srv))
	r.GET("/v1/vehicles/", _VehicleService_ListVehicle0_HTTP_Handler(srv))
}

func _VehicleService_CreateVehicle0_HTTP_Handler(srv VehicleServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateVehicleRequest
		if err := ctx.Bind(&in.Vehicle); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.vehicles.service.v1.VehicleService/CreateVehicle")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateVehicle(ctx, req.(*CreateVehicleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateVehicleResponse)
		return ctx.Result(200, reply)
	}
}

func _VehicleService_UpdateVehicle0_HTTP_Handler(srv VehicleServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateVehicleRequest
		if err := ctx.Bind(&in.Vehicle); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.vehicles.service.v1.VehicleService/UpdateVehicle")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateVehicle(ctx, req.(*UpdateVehicleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateVehicleResponse)
		return ctx.Result(200, reply)
	}
}

func _VehicleService_DeleteVehicle0_HTTP_Handler(srv VehicleServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteVehicleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.vehicles.service.v1.VehicleService/DeleteVehicle")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteVehicle(ctx, req.(*DeleteVehicleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteVehicleResponse)
		return ctx.Result(200, reply)
	}
}

func _VehicleService_GetVehicle0_HTTP_Handler(srv VehicleServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetVehicleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.vehicles.service.v1.VehicleService/GetVehicle")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetVehicle(ctx, req.(*GetVehicleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetVehicleResponse)
		return ctx.Result(200, reply)
	}
}

func _VehicleService_ListVehicle0_HTTP_Handler(srv VehicleServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListVehicleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.vehicles.service.v1.VehicleService/ListVehicle")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListVehicle(ctx, req.(*ListVehicleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListVehicleResponse)
		return ctx.Result(200, reply)
	}
}

type VehicleServiceHTTPClient interface {
	CreateVehicle(ctx context.Context, req *CreateVehicleRequest, opts ...http.CallOption) (rsp *CreateVehicleResponse, err error)
	DeleteVehicle(ctx context.Context, req *DeleteVehicleRequest, opts ...http.CallOption) (rsp *DeleteVehicleResponse, err error)
	GetVehicle(ctx context.Context, req *GetVehicleRequest, opts ...http.CallOption) (rsp *GetVehicleResponse, err error)
	ListVehicle(ctx context.Context, req *ListVehicleRequest, opts ...http.CallOption) (rsp *ListVehicleResponse, err error)
	UpdateVehicle(ctx context.Context, req *UpdateVehicleRequest, opts ...http.CallOption) (rsp *UpdateVehicleResponse, err error)
}

type VehicleServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewVehicleServiceHTTPClient(client *http.Client) VehicleServiceHTTPClient {
	return &VehicleServiceHTTPClientImpl{client}
}

func (c *VehicleServiceHTTPClientImpl) CreateVehicle(ctx context.Context, in *CreateVehicleRequest, opts ...http.CallOption) (*CreateVehicleResponse, error) {
	var out CreateVehicleResponse
	pattern := "/v1/vehicles"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.vehicles.service.v1.VehicleService/CreateVehicle"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.Vehicle, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *VehicleServiceHTTPClientImpl) DeleteVehicle(ctx context.Context, in *DeleteVehicleRequest, opts ...http.CallOption) (*DeleteVehicleResponse, error) {
	var out DeleteVehicleResponse
	pattern := "/v1/vehicles/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.vehicles.service.v1.VehicleService/DeleteVehicle"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *VehicleServiceHTTPClientImpl) GetVehicle(ctx context.Context, in *GetVehicleRequest, opts ...http.CallOption) (*GetVehicleResponse, error) {
	var out GetVehicleResponse
	pattern := "/v1/vehicles/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.vehicles.service.v1.VehicleService/GetVehicle"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *VehicleServiceHTTPClientImpl) ListVehicle(ctx context.Context, in *ListVehicleRequest, opts ...http.CallOption) (*ListVehicleResponse, error) {
	var out ListVehicleResponse
	pattern := "/v1/vehicles/"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.vehicles.service.v1.VehicleService/ListVehicle"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *VehicleServiceHTTPClientImpl) UpdateVehicle(ctx context.Context, in *UpdateVehicleRequest, opts ...http.CallOption) (*UpdateVehicleResponse, error) {
	var out UpdateVehicleResponse
	pattern := "/v1/vehicles/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.vehicles.service.v1.VehicleService/UpdateVehicle"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Vehicle, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
