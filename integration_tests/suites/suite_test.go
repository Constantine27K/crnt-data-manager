package suites

import (
	"context"
	"os"
	"testing"

	"github.com/Constantine27K/crnt-data-manager/integration_tests/helper"
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
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type CrntDMSuite struct {
	suite.Suite
	ctx      context.Context
	cancelFn context.CancelFunc

	issueService   *issueService.Implementation
	projectService *projectService.Implementation
	teamService    *teamService.Implementation
	sprintService  *sprintService.Implementation

	helper helper.Helper
}

func TestRunner(t *testing.T) {
	suite.Run(t, new(CrntDMSuite))
}

func (s *CrntDMSuite) SetupSuite() {
	ctx, cancelFn := context.WithCancel(context.Background())

	err := godotenv.Load("../../.env")
	require.NoError(s.T(), err)

	db, err := postgres.NewPostgres(postgres.Options{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DBNAME"),
	})
	require.NoError(s.T(), err)

	validator := validate.NewValidator()

	sprintGw := sprintGateway.NewSprintGateway(db)
	sprintStore := sprintStorage.NewSprintStorage(sprintGw)

	teamGw := teamGateway.NewTeamGateway(db)
	teamStore := teamStorage.NewTeamStorage(teamGw)

	issueGw := issueGateway.NewIssueGateWay(db)
	projectGw := projectGateway.NewProjectGateway(db)

	projectStore := projectStorage.NewProjectStorage(projectGw)
	issueStore := issueStorage.NewIssueStorage(issueGw, projectGw)

	dbHelper := helper.NewDBHelper(db)

	s.ctx = ctx
	s.cancelFn = cancelFn
	s.helper = dbHelper

	s.issueService = issueService.NewService(validator, issueStore)
	s.projectService = projectService.NewService(validator, projectStore)
	s.teamService = teamService.NewService(validator, teamStore)
	s.sprintService = sprintService.NewService(sprintStore, validator)
}

func (s *CrntDMSuite) TearDownSuite() {
	s.cancelFn()
}

func (s *CrntDMSuite) TearDownTest() {
	require.NoError(s.T(), s.helper.FlushAllTables())
}
