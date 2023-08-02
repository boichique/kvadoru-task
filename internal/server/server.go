package server

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/boichique/kvadoru_task/internal/config"
	"github.com/boichique/kvadoru_task/internal/errlog"
	"github.com/boichique/kvadoru_task/internal/logger"
	"github.com/boichique/kvadoru_task/internal/modules/books"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	dbConnectTimeout = 10 * time.Second // Timeout for connecting to the database
)

// Server is a struct that contains the configuration, gRPC server, closers, and listener
type Server struct {
	cfg        *config.Config
	grpcServer *grpc.Server
	closers    []func() error
	listener   net.Listener
}

// NewServer creates a new Server with the provided context and configuration
func NewServer(ctx context.Context, cfg *config.Config) (*Server, error) {
	// Create a new gRPC server with interceptors for logging and error handling
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logger.UnaryServerInterceptor(),
			errlog.UnaryServerInterceptor(),
		),
	)

	// Connect to the database and add a closer for it
	var closers []func() error
	db, err := getDB(ctx, cfg.DBUrl)
	if err != nil {
		return nil, fmt.Errorf("connect to database: %w", err)
	}
	closers = append(closers, func() error {
		db.Close()
		return nil
	})

	// Create a new book server and register it with the gRPC server
	mysql := books.NewBookServer(db)
	books.RegisterBooksServiceV1Server(grpcServer, books.NewBookServer(mysql.Repository.DB))
	reflection.Register(grpcServer)

	return &Server{
		grpcServer: grpcServer,
		cfg:        cfg,
		closers:    closers,
	}, nil
}

// Start starts the Server
func (s *Server) Start() error {
	// Setup the logger
	logger, err := logger.SetupLogger(s.cfg.Local, s.cfg.LogLevel)
	if err != nil {
		return fmt.Errorf("setup logger: %s", err)
	}
	slog.SetDefault(logger)

	// Listen on the configured port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.Port))
	if err != nil {
		return fmt.Errorf("listen: %s", err)
	}

	// Store the listener and add a closer for it
	s.listener = lis
	s.closers = append(s.closers, lis.Close)
	logger.Info("starting server", "port", s.cfg.Port)

	// Start the gRPC server
	return s.grpcServer.Serve(lis)
}

// Stop stops the Server
func (s *Server) Stop(ctx context.Context) error {
	// Create a channel to signal when the server has stopped
	stopped := make(chan struct{})

	// Stop the server gracefully in a goroutine
	go func() {
		s.grpcServer.GracefulStop()
		close(stopped)
	}()

	// Wait for the server to stop or the context to be done
	select {
	case <-ctx.Done():
		s.grpcServer.Stop()
	case <-stopped:
	}

	// Close all closers
	return withClosers(s.closers, nil)
}

// Port returns the port the Server is listening on
func (s *Server) Port() (int, error) {
	if s.listener == nil || s.listener.Addr() == nil {
		return 0, errors.New("server is not started")
	}

	return s.listener.Addr().(*net.TCPAddr).Port, nil
}

// withClosers closes all closers and returns any errors
func withClosers(closers []func() error, err error) error {
	errs := []error{err}

	for i := len(closers) - 1; i >= 0; i-- {
		if err = closers[i](); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

// getDB connects to the database and returns it
func getDB(ctx context.Context, connString string) (*sql.DB, error) {
	_, cancel := context.WithTimeout(ctx, dbConnectTimeout)
	defer cancel()

	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping DB: %w", err)
	}

	return db, nil
}
