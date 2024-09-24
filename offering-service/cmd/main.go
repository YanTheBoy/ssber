package main

import (
	"context"
	"errors"
	"net/http"
	"offering-service/internal/app"
	"offering-service/internal/config"
	gen "offering-service/internal/generated"
	"offering-service/internal/logger"
	"offering-service/internal/otel"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	middleware "github.com/oapi-codegen/nethttp-middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/riandyrn/otelchi"
	"github.com/yarlson/chiprom"
	"offering-service/internal/db"
	"offering-service/internal/handlers"
	mw "offering-service/internal/middleware"
)

func main() {
	log := logger.New()

	cfg, err := config.Parse()
	if err != nil {
		log.WithError(err, "error parse ENV")
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	serviceName := cfg.ServiceName
	serviceVersion := os.Getenv("SERVICE_VERSION")

	otelShudown, err := otel.SetupOTelSDK(ctx, serviceName, serviceVersion, true)
	if err != nil {
		log.WithError(err, "error setup opentel")
		return
	}
	defer func() {
		err = errors.Join(err, otelShudown(context.Background()))
	}()

	conn, err := pgxpool.New(ctx, cfg.PostgresConn)
	if err != nil {
		log.WithError(err, "err db connection")
		os.Exit(1)
	}
	defer conn.Close()

	// TODO SetUp connection TO AuthService
	/*	grpcConn, err := grpc.DialContext(
			context.Background(),
			"GRPC:ADDDR",
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
			grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
		)

	*/

	flatsRepository, err := db.NewClient(&cfg)
	if err != nil {
		log.WithError(err, "err init db")
	}
	flatsManagerService := app.NewFlatsService(flatsRepository)

	handle := handlers.New(log, flatsManagerService)

	r := chi.NewRouter()
	swagger, err := gen.GetSwagger()
	if err != nil {
		log.WithError(err, "get swagger")
		os.Exit(1)
	}

	r.Use(mw.RateLimiterMiddleware(2, 5, 10*time.Minute))
	r.Use(middleware.OapiRequestValidator(swagger))
	r.Use(chimiddleware.Recoverer)
	r.Use(chiprom.NewMiddleware(serviceName))
	r.Use(otelchi.Middleware(serviceName, otelchi.WithChiRoutes(r)))

	baseRouter := chi.NewRouter()
	baseRouter.Handle("/healthcheck", http.HandlerFunc(healthCheck))
	baseRouter.Handle("/metrics", promhttp.Handler())

	gen.HandlerFromMux(handle, r)
	baseRouter.Mount("/", r)

	// TODO pop address
	s := &http.Server{
		Handler: baseRouter,
		Addr:    cfg.ServerAddress,
	}

	go func() {
		log.Info("Listen and Serve on", "address", cfg.ServerAddress)
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.WithError(err, "stop listening")
			stop()
		}
	}()
	<-ctx.Done()
	log.Info("Listen stopped")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := s.Shutdown(ctx); err != nil {
		log.Error("Shutdown err", "err", err.Error())
		os.Exit(1)
	}
	log.Info("Shutdown completed")

}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"APP RUNNING"}`))
	//TODO add FlatsRepository status
}
