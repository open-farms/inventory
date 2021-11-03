package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/open-farms/inventory/ent"
	elk "github.com/open-farms/inventory/ent/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/open-farms/inventory/internal/settings"
	"go.uber.org/zap"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "inventory"
	// Version is the version of the compiled software.
	Version string = "x.y.z"

	// flagconf is the config flag.
	flagconf string

	logger *zap.Logger
)

func init() {
	flag.StringVar(&flagconf, "config", "./config", "config path, eg: -config config.yaml")
}

func setup() (config.Config, error) {
	flag.Parse()
	logger = zap.NewExample(
		zap.Fields(zap.String("service", Name)),
	)
	cfg, err := settings.Configure(flagconf)
	if err != nil {
		return nil, err
	}

	migrate, err := cfg.Value("storage.database.migrate").Bool()
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	if migrate {
		err = settings.Migrate(flagconf, false)
		if err != nil {
			return nil, err
		}
	}

	return cfg, nil
}

func service(c *ent.Client) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.StripSlashes)
	r.Route("/v1", func(r chi.Router) {
		elk.NewEquipmentHandler(c, logger).MountRoutes(r)
		elk.NewVehicleHandler(c, logger).MountRoutes(r)
	})

	return r
}

func main() {

	// Initialize configs
	// and parse flags
	cfg, err := setup()
	if err != nil {
		logger.Fatal(err.Error())
	}

	driver, _ := cfg.Value("storage.database.driver").String()
	source, _ := cfg.Value("storage.database.source").String()
	address, _ := cfg.Value("server.http.addr").String()
	timeout, _ := cfg.Value("server.http.timeout").Duration()

	// Connect to database
	c, err := ent.Open(driver, source)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer c.Close()

	// Create server with context and graceful shutdown timer
	server := &http.Server{Addr: address, Handler: service(c)}
	serverCtx, serverStopCtx := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig
		shutdownCtx, _ := context.WithTimeout(serverCtx, timeout)
		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				logger.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	logger.Info("Server running")
	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Fatal(err.Error())
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
	logger.Info("Server stopped")
}
