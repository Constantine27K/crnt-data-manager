package suites

import (
	"context"
	"os"
	"testing"

	"github.com/Constantine27K/crnt-data-manager/integration_tests/helper"
	departmentService "github.com/Constantine27K/crnt-data-manager/internal/app/crnt-data-manager/department"
	issueService "github.com/Constantine27K/crnt-data-manager/internal/app/crnt-data-manager/issue"
	projectService "github.com/Constantine27K/crnt-data-manager/internal/app/crnt-data-manager/project"
	sprintService "github.com/Constantine27K/crnt-data-manager/internal/app/crnt-data-manager/sprint"
	teamService "github.com/Constantine27K/crnt-data-manager/internal/app/crnt-data-manager/team"
	departmentGateway "github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/department/gateway"
	departmentStorage "github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/department/storage"
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
	"github.com/Constantine27K/crnt-sdk/pkg/authorization"
	"github.com/Constantine27K/crnt-sdk/pkg/token"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type CrntDMSuite struct {
	suite.Suite
	ctx      context.Context
	cancelFn context.CancelFunc

	issueService      *issueService.Implementation
	projectService    *projectService.Implementation
	teamService       *teamService.Implementation
	sprintService     *sprintService.Implementation
	departmentService *departmentService.Implementation

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

	departmentGw := departmentGateway.NewDepartmentGateway(db)
	departmentStore := departmentStorage.NewDepartmentStorage(departmentGw)

	dbHelper := helper.NewDBHelper(db)

	tokenMaker, err := token.NewMaker(os.Getenv("TOKEN_SYMMETRIC_KEY"))
	require.NoError(s.T(), err)
	authorizer := authorization.NewAuthorizer(tokenMaker)

	s.ctx = ctx
	s.cancelFn = cancelFn
	s.helper = dbHelper

	s.issueService = issueService.NewService(issueStore, validator)
	s.projectService = projectService.NewService(projectStore, sprintStore, validator, authorizer)
	s.teamService = teamService.NewService(teamStore, validator, authorizer)
	s.sprintService = sprintService.NewService(sprintStore, validator)
	s.departmentService = departmentService.NewService(departmentStore, validator, authorizer)
}

func (s *CrntDMSuite) TearDownSuite() {
	s.cancelFn()
}

func (s *CrntDMSuite) TearDownTest() {
	require.NoError(s.T(), s.helper.FlushAllTables())
}
