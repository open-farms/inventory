package service

import (
	"context"

	pb "github.com/open-farms/inventory/api/equipment/service/v1"
)

type EquipmentService struct {
	pb.UnimplementedEquipmentServiceServer
}

func NewEquipmentService() *EquipmentService {
	return &EquipmentService{}
}

func (s *EquipmentService) CreateEquipment(ctx context.Context, req *pb.CreateEquipmentRequest) (*pb.CreateEquipmentResponse, error) {
	return &pb.CreateEquipmentResponse{}, nil
}

func (s *EquipmentService) UpdateEquipment(ctx context.Context, req *pb.UpdateEquipmentRequest) (*pb.UpdateEquipmentResponse, error) {
	return &pb.UpdateEquipmentResponse{}, nil
}

func (s *EquipmentService) DeleteEquipment(ctx context.Context, req *pb.DeleteEquipmentRequest) (*pb.DeleteEquipmentResponse, error) {
	return &pb.DeleteEquipmentResponse{}, nil
}

func (s *EquipmentService) GetEquipment(ctx context.Context, req *pb.GetEquipmentRequest) (*pb.GetEquipmentResponse, error) {
	return &pb.GetEquipmentResponse{}, nil
}

func (s *EquipmentService) ListEquipment(ctx context.Context, req *pb.ListEquipmentRequest) (*pb.ListEquipmentResponse, error) {
	return &pb.ListEquipmentResponse{}, nil
}
