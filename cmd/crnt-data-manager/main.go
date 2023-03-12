package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"

	issueService "github.com/Constantine27K/crnt-data-manager/internal/app/crnt-data-manager/issue"
	projectService "github.com/Constantine27K/crnt-data-manager/internal/app/crnt-data-manager/project"
	sprintService "github.com/Constantine27K/crnt-data-manager/internal/app/crnt-data-manager/sprint"
	teamService "github.com/Constantine27K/crnt-data-manager/internal/app/crnt-data-manager/team"
	issueGateway "github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/issue/gateway"
	issueStorage "github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/issue/storage"
	projectGateway "github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/project/gateway"
	projectStorage "github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/project/storage"
	sprintGateway "github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/sprint/gateway"
	sprintStorage "github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/sprint/storage"
	teamGateway "github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/team/gateway"
	teamStorage "github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/team/storage"
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/infrastructure/postgres"
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/validate"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/project"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/sprint"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/issue"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/team"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	loadEnv()
	setLogger()

	go createGrpcServer()
	createHttpServer()
}

func loadEnv() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found",
			zap.Error(err),
		)
	}
}

func setLogger() {
	logLevel, _ := strconv.ParseInt(os.Getenv("LOG_LEVEL"), 10, 32)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.SetLevel(log.Level(logLevel))
}

func createGrpcServer() {
	port := os.Getenv("GRPC_PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Error(fmt.Sprintf("failed to listen localhost:%v", port),
			zap.Error(err),
		)
	}

	grpcServer := grpc.NewServer()

	db, err := postgres.NewPostgres(postgres.Options{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DBNAME"),
	})
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}

	validator := validate.NewValidator()

	sprintGw := sprintGateway.NewSprintGateway(db)
	sprintStore := sprintStorage.NewSprintStorage(sprintGw)

	teamGw := teamGateway.NewTeamGateway(db)
	teamStore := teamStorage.NewTeamStorage(teamGw)

	projectGw := projectGateway.NewProjectGateway(db)
	projectStore := projectStorage.NewProjectStorage(projectGw)

	issueGw := issueGateway.NewIssueGateWay(db)
	issueStore := issueStorage.NewIssueStorage(issueGw, projectGw)

	issue.RegisterIssueRegistryServer(grpcServer, issueService.NewService(validator, issueStore))
	sprint.RegisterSprintRegistryServer(grpcServer, sprintService.NewService(sprintStore, validator))
	team.RegisterTeamRegistryServer(grpcServer, teamService.NewService(validator, teamStore))
	project.RegisterProjectRegistryServer(grpcServer, projectService.NewService(validator, projectStore))

	log.Infof("grpc service started on port %s", port)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Error("error during serving GRPC",
			zap.Error(err),
		)
	}
}

func createHttpServer() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	grpcPort := os.Getenv("GRPC_PORT")
	// dial the gRPC server above to make a client connection
	conn, err := grpc.Dial(fmt.Sprintf(":%s", grpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Error(fmt.Sprintf("failed to dial localhost:%s", grpcPort),
			zap.Error(err),
		)
	}
	defer conn.Close()

	// create an HTTP router using the client connection above
	// and register it with the service client
	rmux := runtime.NewServeMux()
	clientIssue := issue.NewIssueRegistryClient(conn)
	err = issue.RegisterIssueRegistryHandlerClient(ctx, rmux, clientIssue)
	if err != nil {
		log.Error("failed to register task handler client",
			zap.Error(err),
		)
	}
	clientSprint := sprint.NewSprintRegistryClient(conn)
	err = sprint.RegisterSprintRegistryHandlerClient(ctx, rmux, clientSprint)
	if err != nil {
		log.Error("failed to register sprint handler client",
			zap.Error(err),
		)
	}
	clientTeam := team.NewTeamRegistryClient(conn)
	err = team.RegisterTeamRegistryHandlerClient(ctx, rmux, clientTeam)
	if err != nil {
		log.Error("failed to register team handler client",
			zap.Error(err),
		)
	}
	clientProject := project.NewProjectRegistryClient(conn)
	err = project.RegisterProjectRegistryHandlerClient(ctx, rmux, clientProject)
	if err != nil {
		log.Error("failed to register project handler client",
			zap.Error(err),
		)
	}

	// creating swagger
	mux := http.NewServeMux()
	// mount the gRPC HTTP gateway to the root
	mux.Handle("/", rmux)
	fs := http.FileServer(http.Dir("./swagger"))
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	log.Info("swagger created")

	httpPort := os.Getenv("HTTP_PORT")
	log.Infof("http service started on port %s", httpPort)

	// start a standard HTTP server with the router
	err = http.ListenAndServe(fmt.Sprintf("localhost:%s", httpPort), mux)
	if err != nil {
		log.Error("error during listening and serving HTTP",
			zap.Error(err),
		)
	}
}
