package main

import (
	"flag"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/open-farms/inventory/ent"
	elk "github.com/open-farms/inventory/ent/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-kratos/kratos/v2"

	transgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/open-farms/inventory/ent/proto/entpb"
	"github.com/open-farms/inventory/internal/logging"
	"github.com/open-farms/inventory/internal/settings"
	"go.uber.org/zap"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "inventory"
	// Version is the version of the compiled software.
	Version string = "x.y.z"

	logger *zap.Logger
)

func setup() (*settings.Config, error) {
	flag.Parse()
	logger = zap.NewExample(
		zap.Fields(zap.String("service", Name)),
	)

	c, err := settings.Configure()
	if err != nil {
		return nil, err
	}

	if c.Storage.Migrate {
		err = settings.Migrate(false)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func grpcService(client *ent.Client, address string, timeout time.Duration) *transgrpc.Server {
	equipmentSvc := entpb.NewEquipmentService(client)
	vehicleSvc := entpb.NewVehicleService(client)
	toolsSvc := entpb.NewToolService(client)
	implementSvc := entpb.NewImplementService(client)

	grpcServer := transgrpc.NewServer(transgrpc.Address(address), transgrpc.Timeout(timeout))
	entpb.RegisterEquipmentServiceServer(grpcServer, equipmentSvc)
	entpb.RegisterVehicleServiceServer(grpcServer, vehicleSvc)
	entpb.RegisterToolServiceServer(grpcServer, toolsSvc)
	entpb.RegisterImplementServiceServer(grpcServer, implementSvc)

	return grpcServer
}

func httpService(c *ent.Client, address string, timeout time.Duration, auth settings.Auth) *transhttp.Server {
	r := chi.NewRouter()
	r.Use(logging.Logger(logger))
	r.Use(middleware.StripSlashes)
	r.Use(middleware.Timeout(timeout))
	if auth.Enabled {
		r.Use(middleware.BasicAuth("Open Farms Inventory Database", map[string]string{"token": auth.Token}))
	}
	r.Route("/v1", func(r chi.Router) {
		elk.NewEquipmentHandler(c, logger).MountRoutes(r)
		elk.NewVehicleHandler(c, logger).MountRoutes(r)
		elk.NewToolHandler(c, logger).MountRoutes(r)
		elk.NewImplementHandler(c, logger).MountRoutes(r)
		elk.NewLocationHandler(c, logger).MountRoutes(r)
		elk.NewCategoryHandler(c, logger).MountRoutes(r)
	})

	httpServer := transhttp.NewServer(transhttp.Address(address))
	httpServer.HandlePrefix("/", r)
	return httpServer
}

func main() {

	// Initialize configs
	// and parse flags
	cfg, err := setup()
	if err != nil {
		logger.Fatal(err.Error())
	}

	driver := cfg.Storage.Driver
	source := cfg.Storage.URI
	httpAddress := cfg.HTTP.Addr
	httpTimeout := cfg.HTTP.Timeout
	grpcAddress := cfg.GRPC.Addr
	grpcTimeout := cfg.HTTP.Timeout

	// Connect to database
	c, err := ent.Open(driver, source)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer c.Close()

	// Create servers with context and graceful shutdown timer
	httpServer := httpService(c, httpAddress, httpTimeout, cfg.HTTP.Auth)
	grpcServer := grpcService(c, grpcAddress, grpcTimeout)
	app := kratos.New(
		kratos.Name("inventory"),
		kratos.Server(httpServer, grpcServer),
		kratos.Version(Version),
		kratos.Logger(
			logging.NewZapLogger(
				zap.NewProductionEncoderConfig(),
				zap.NewProductionConfig().Level,
			),
		),
	)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
