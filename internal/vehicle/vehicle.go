package vehicle

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

func (s *VehicleService) CreateVehicle(ctx context.Context, req *pb.CreateVehicleRequest) (*pb.CreateVehicleReply, error) {
	return &pb.CreateVehicleReply{}, nil
}
func (s *VehicleService) UpdateVehicle(ctx context.Context, req *pb.UpdateVehicleRequest) (*pb.UpdateVehicleReply, error) {
	return &pb.UpdateVehicleReply{}, nil
}
func (s *VehicleService) DeleteVehicle(ctx context.Context, req *pb.DeleteVehicleRequest) (*pb.DeleteVehicleReply, error) {
	return &pb.DeleteVehicleReply{}, nil
}
func (s *VehicleService) GetVehicle(ctx context.Context, req *pb.GetVehicleRequest) (*pb.GetVehicleReply, error) {
	return &pb.GetVehicleReply{}, nil
}
func (s *VehicleService) ListVehicle(ctx context.Context, req *pb.ListVehicleRequest) (*pb.ListVehicleReply, error) {
	return &pb.ListVehicleReply{
		Vehicles: []*pb.Vehicle{
			{Id: 25, Make: "toyota", Model: "4Runner", Year: "2020"},
		},
	}, nil
}
