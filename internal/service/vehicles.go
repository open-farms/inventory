package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/open-farms/inventory/api/vehicles/service/v1"
	"github.com/open-farms/inventory/ent"
	"github.com/open-farms/inventory/ent/vehicle"
	"github.com/open-farms/inventory/internal/biz"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type VehicleService struct {
	logger   log.Logger
	dbDriver string
	dbURI    string

	pb.UnimplementedVehicleServiceServer
}

func NewVehicleService(logger log.Logger, configPath string) *VehicleService {
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

	return &VehicleService{
		logger:   logger,
		dbDriver: dbDriver,
		dbURI:    dbURI,
	}
}

func (s *VehicleService) CreateVehicle(ctx context.Context, req *pb.CreateVehicleRequest) (*pb.CreateVehicleResponse, error) {
	vehicle := ent.Vehicle{
		Make:      req.Vehicle.Make,
		Model:     req.Vehicle.Model,
		Year:      req.Vehicle.Year,
		Condition: vehicle.Condition(req.Vehicle.Condition.String()),
		Active:    req.Vehicle.Active,
		Tags:      req.Vehicle.Tags,
	}

	client, err := ent.Open(s.dbDriver, s.dbURI)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	resp, err := biz.CreateVehicle(ctx, client, vehicle)
	if err != nil {
		return nil, err
	}

	condition := pb.Vehicle_Condition(pb.Vehicle_Condition_value[resp.Condition.String()])
	return &pb.CreateVehicleResponse{
		Vehicle: &pb.Vehicle{
			Id:        resp.ID,
			Make:      resp.Make,
			Model:     resp.Model,
			Year:      resp.Year,
			Condition: condition,
			Active:    resp.Active,
			Tags:      resp.Tags,
			CreatedAt: timestamppb.New(resp.CreatedAt),
			UpdatedAt: timestamppb.New(resp.UpdatedAt),
		},
	}, nil
}

func (s *VehicleService) UpdateVehicle(ctx context.Context, req *pb.UpdateVehicleRequest) (*pb.UpdateVehicleResponse, error) {
	client, err := ent.Open(s.dbDriver, s.dbURI)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	vehicle := ent.Vehicle{
		ID:        req.Id,
		Make:      req.Vehicle.Make,
		Model:     req.Vehicle.Model,
		Year:      req.Vehicle.Year,
		Condition: vehicle.Condition(req.Vehicle.Condition.String()),
		Active:    req.Vehicle.Active,
		Tags:      req.Vehicle.Tags,
	}

	resp, err := biz.UpdateVehicle(ctx, client, req.Id, vehicle)
	if err != nil {
		return nil, err
	}

	condition := pb.Vehicle_Condition(pb.Vehicle_Condition_value[resp.Condition.String()])
	return &pb.UpdateVehicleResponse{
		Vehicle: &pb.Vehicle{
			Id:        resp.ID,
			Make:      resp.Make,
			Model:     resp.Model,
			Year:      resp.Year,
			Condition: condition,
			Active:    resp.Active,
			Tags:      resp.Tags,
			CreatedAt: timestamppb.New(resp.CreatedAt),
			UpdatedAt: timestamppb.New(resp.UpdatedAt),
		},
	}, nil
}

func (s *VehicleService) DeleteVehicle(ctx context.Context, req *pb.DeleteVehicleRequest) (*pb.DeleteVehicleResponse, error) {
	client, err := ent.Open(s.dbDriver, s.dbURI)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	resp, err := biz.DeleteVehicle(ctx, client, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteVehicleResponse{
		Id: resp.ID,
	}, nil
}

func (s *VehicleService) GetVehicle(ctx context.Context, req *pb.GetVehicleRequest) (*pb.GetVehicleResponse, error) {
	client, err := ent.Open(s.dbDriver, s.dbURI)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	resp, err := biz.GetVehicle(ctx, client, req.Id)
	if err != nil {
		return nil, err
	}

	condition := pb.Vehicle_Condition(pb.Vehicle_Condition_value[resp.Condition.String()])
	return &pb.GetVehicleResponse{
		Vehicle: &pb.Vehicle{
			Id:        resp.ID,
			Make:      resp.Make,
			Model:     resp.Model,
			Year:      resp.Year,
			Condition: condition,
			Active:    resp.Active,
			Tags:      resp.Tags,
			CreatedAt: timestamppb.New(resp.CreatedAt),
			UpdatedAt: timestamppb.New(resp.UpdatedAt),
		},
	}, nil
}

func (s *VehicleService) ListVehicle(ctx context.Context, req *pb.ListVehicleRequest) (*pb.ListVehicleResponse, error) {
	client, err := ent.Open(s.dbDriver, s.dbURI)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	resp, err := biz.ListVehicles(ctx, client)
	if err != nil {
		return nil, err
	}

	rv := make([]*pb.Vehicle, 0)
	for _, p := range resp {
		condition := pb.Vehicle_Condition(pb.Vehicle_Condition_value[p.Condition.String()])
		rv = append(rv, &pb.Vehicle{
			Id:        p.ID,
			Make:      p.Make,
			Model:     p.Model,
			Year:      p.Year,
			Condition: condition,
			Active:    p.Active,
			Tags:      p.Tags,
			CreatedAt: timestamppb.New(p.CreatedAt),
			UpdatedAt: timestamppb.New(p.UpdatedAt),
		})
	}

	return &pb.ListVehicleResponse{Vehicles: rv}, nil
}
