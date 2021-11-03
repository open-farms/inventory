package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/open-farms/inventory/api/equipment/service/v1"
	"github.com/open-farms/inventory/ent"
	"github.com/open-farms/inventory/ent/equipment"
	"github.com/open-farms/inventory/internal/biz"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EquipmentService struct {
	logger   log.Logger
	dbDriver string
	dbURI    string

	pb.UnimplementedEquipmentServiceServer
}

func NewEquipmentService(logger log.Logger, configPath string) *EquipmentService {
	// Initialize configuration
	cfg, err := biz.Configure(configPath)
	if err != nil {
		panic(err)
	}

	dbDriver, err := cfg.Value("storage.database.driver").String()
	if err != nil {
		panic(err)
	}

	dbURI, err := cfg.Value("storage.database.source").String()
	if err != nil {
		panic(err)
	}

	return &EquipmentService{
		logger:   logger,
		dbDriver: dbDriver,
		dbURI:    dbURI,
	}
}

func (s *EquipmentService) CreateEquipment(ctx context.Context, req *pb.CreateEquipmentRequest) (*pb.CreateEquipmentResponse, error) {
	client, err := ent.Open(s.dbDriver, s.dbURI)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	resp, err := biz.CreateEquipment(ctx, client, ent.Equipment{
		Name:      req.Equipment.Name,
		Tags:      req.Equipment.Tags,
		Condition: equipment.Condition(req.Equipment.Condition.String()),
	})

	if err != nil {
		return nil, err
	}

	condition := pb.Equipment_Condition(pb.Equipment_Condition_value[resp.Condition.String()])
	return &pb.CreateEquipmentResponse{
		Equipment: &pb.Equipment{
			Id:        resp.ID,
			Name:      resp.Name,
			Tags:      resp.Tags,
			Condition: condition,
			CreatedAt: timestamppb.New(resp.CreatedAt),
			UpdatedAt: timestamppb.New(resp.UpdatedAt),
		},
	}, nil
}

func (s *EquipmentService) UpdateEquipment(ctx context.Context, req *pb.UpdateEquipmentRequest) (*pb.UpdateEquipmentResponse, error) {
	client, err := ent.Open(s.dbDriver, s.dbURI)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	equipment := ent.Equipment{
		Name:      req.Equipment.Name,
		Tags:      req.Equipment.Tags,
		Condition: equipment.Condition(req.Equipment.Condition.String()),
	}

	resp, err := biz.UpdateEquipment(ctx, client, req.Id, equipment)
	if err != nil {
		return nil, err
	}

	condition := pb.Equipment_Condition(pb.Equipment_Condition_value[resp.Condition.String()])
	return &pb.UpdateEquipmentResponse{
		Equipment: &pb.Equipment{
			Id:        resp.ID,
			Name:      resp.Name,
			Tags:      resp.Tags,
			Condition: condition,
			CreatedAt: timestamppb.New(resp.CreatedAt),
			UpdatedAt: timestamppb.New(resp.UpdatedAt),
		},
	}, nil
}

func (s *EquipmentService) DeleteEquipment(ctx context.Context, req *pb.DeleteEquipmentRequest) (*pb.DeleteEquipmentResponse, error) {
	client, err := ent.Open(s.dbDriver, s.dbURI)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	resp, err := biz.DeleteEquipment(ctx, client, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteEquipmentResponse{
		Id: resp.ID,
	}, nil
}

func (s *EquipmentService) GetEquipment(ctx context.Context, req *pb.GetEquipmentRequest) (*pb.GetEquipmentResponse, error) {
	client, err := ent.Open(s.dbDriver, s.dbURI)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	resp, err := biz.GetEquipment(ctx, client, req.Id)
	if err != nil {
		return nil, err
	}

	condition := pb.Equipment_Condition(pb.Equipment_Condition_value[resp.Condition.String()])
	return &pb.GetEquipmentResponse{
		Equipment: &pb.Equipment{
			Id:        resp.ID,
			Name:      resp.Name,
			Tags:      resp.Tags,
			Condition: condition,
			CreatedAt: timestamppb.New(resp.CreatedAt),
			UpdatedAt: timestamppb.New(resp.UpdatedAt),
		},
	}, nil
}

func (s *EquipmentService) ListEquipment(ctx context.Context, req *pb.ListEquipmentRequest) (*pb.ListEquipmentResponse, error) {
	client, err := ent.Open(s.dbDriver, s.dbURI)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	resp, err := biz.ListEquipments(ctx, client)
	if err != nil {
		return nil, err
	}

	rv := make([]*pb.Equipment, 0)
	for _, p := range resp {
		condition := pb.Equipment_Condition(pb.Equipment_Condition_value[p.Condition.String()])
		rv = append(rv, &pb.Equipment{
			Id:        p.ID,
			Name:      p.Name,
			Condition: condition,
			Tags:      p.Tags,
			CreatedAt: timestamppb.New(p.CreatedAt),
			UpdatedAt: timestamppb.New(p.UpdatedAt),
		})
	}
	return &pb.ListEquipmentResponse{
		Equipment: rv,
	}, nil
}
