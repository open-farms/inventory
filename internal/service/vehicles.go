package service

import (
	"context"

	pb "github.com/open-farms/inventory/api/vehicles/service/v1"
)

type VehicleService struct {
	pb.UnimplementedVehicleServiceServer
}

func NewVehicleService() *VehicleService {
	return &VehicleService{}
}

func (s *VehicleService) CreateVehicle(ctx context.Context, req *pb.CreateVehicleRequest) (*pb.CreateVehicleResponse, error) {
	return &pb.CreateVehicleResponse{}, nil
}

func (s *VehicleService) UpdateVehicle(ctx context.Context, req *pb.UpdateVehicleRequest) (*pb.UpdateVehicleResponse, error) {
	return &pb.UpdateVehicleResponse{}, nil
}

func (s *VehicleService) DeleteVehicle(ctx context.Context, req *pb.DeleteVehicleRequest) (*pb.DeleteVehicleResponse, error) {
	return &pb.DeleteVehicleResponse{}, nil
}

func (s *VehicleService) GetVehicle(ctx context.Context, req *pb.GetVehicleRequest) (*pb.GetVehicleResponse, error) {
	return &pb.GetVehicleResponse{}, nil
}

func (s *VehicleService) ListVehicle(ctx context.Context, req *pb.ListVehicleRequest) (*pb.ListVehicleResponse, error) {
	return &pb.ListVehicleResponse{}, nil
}
