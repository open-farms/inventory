package main

import (
	"flag"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	equipment "github.com/open-farms/inventory/api/equipment/service/v1"
	vehicles "github.com/open-farms/inventory/api/vehicles/service/v1"
	"github.com/open-farms/inventory/internal/service"
	"github.com/open-farms/inventory/internal/settings"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "inventory"
	// Version is the version of the compiled software.
	Version string = "x.y.z"

	// ID is the hostname identifier
	id, _ = os.Hostname()

	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "./config", "config path, eg: -conf config.yaml")
}

func setupConfig() (*settings.Server, *settings.Storage) {
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc settings.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	return bc.Server, bc.Storage
}

func main() {
	flag.Parse()

	server, _ := setupConfig()

	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"service.name", Name,
		"service.version", Version,
	)

	httpSrv := http.NewServer(
		http.Address(server.Http.Addr),
		http.Middleware(
			logging.Server(logger),
			recovery.Recovery(),
		),
	)

	grpcSrv := grpc.NewServer(
		grpc.Address(server.Grpc.Addr),
		grpc.Middleware(
			logging.Server(logger),
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
		kratos.Metadata(map[string]string{
			"organization": "github.com/open-farms",
			"service":      "inventory",
		}),
		kratos.Logger(logger),
		kratos.Server(
			httpSrv,
			grpcSrv,
		),
	)

	if err := app.Run(); err != nil {
		logger.Log(log.LevelFatal, err)
	}
}
