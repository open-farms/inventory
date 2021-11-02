package main

import (
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	equipment "github.com/open-farms/inventory/api/equipment/service/v1"
	vehicles "github.com/open-farms/inventory/api/vehicles/service/v1"
	"github.com/open-farms/inventory/internal/service"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "inventory"
	// Version is the version of the compiled software.
	Version string = "x.y.z"

	id, _ = os.Hostname()
)

func main() {

	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)

	httpSrv := http.NewServer(
		http.Address(":8000"),
		http.Middleware(
			recovery.Recovery(),
		),
	)

	grpcSrv := grpc.NewServer(
		grpc.Address(":9000"),
		grpc.Middleware(
			recovery.Recovery(),
		),
	)

	equipmentSvc := service.NewEquipmentService()
	equipment.RegisterEquipmentServiceServer(grpcSrv, equipmentSvc)
	equipment.RegisterEquipmentServiceHTTPServer(httpSrv, equipmentSvc)

	vehicleSvc := service.NewVehicleService()
	vehicles.RegisterVehicleServiceServer(grpcSrv, vehicleSvc)
	vehicles.RegisterVehicleServiceHTTPServer(httpSrv, vehicleSvc)

	app := kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			httpSrv,
			grpcSrv,
		),
	)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
