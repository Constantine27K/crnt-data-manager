//go:build integration
// +build integration

package suites

import (
	"context"
	"fmt"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/comments"
	"os"
	"testing"
	"time"

	"github.com/Constantine27K/crnt-data-manager/integration_tests/fixtures"
	"github.com/Constantine27K/crnt-data-manager/integration_tests/helper"
	issueService "github.com/Constantine27K/crnt-data-manager/internal/app/crnt-data-manager/issue"
	projectService "github.com/Constantine27K/crnt-data-manager/internal/app/crnt-data-manager/project"
	issueGateway "github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/issue/gateway"
	issueStorage "github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/issue/storage"
	projectGateway "github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/project/gateway"
	projectStorage "github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/project/storage"
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/infrastructure/postgres"
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/validate"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/project"
	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/issue"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type IssueSuite struct {
	suite.Suite
	ctx      context.Context
	cancelFn context.CancelFunc

	issueService   *issueService.Implementation
	projectService *projectService.Implementation
	helper         helper.Helper
}

func TestRunner(t *testing.T) {
	suite.Run(t, new(IssueSuite))
}

func (s *IssueSuite) SetupSuite() {
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
}

func (s *IssueSuite) TearDownSuite() {
	s.cancelFn()
}

func (s *IssueSuite) TearDownTest() {
	require.NoError(s.T(), s.helper.FlushAllTables())
}

func (s *IssueSuite) TestIssues_Dummy() {
	proj := fixtures.CreateProject()

	respCreateProject, err := s.projectService.CreateProject(s.ctx, &project.ProjectCreateRequest{Project: proj})
	require.NoError(s.T(), err)
	projectID := respCreateProject.GetId()

	require.Greater(s.T(), projectID, int64(0))

	respGetProject, err := s.projectService.GetProjectByID(s.ctx, &project.ProjectGetByIDRequest{Id: projectID})
	require.NoError(s.T(), err)
	gotProject := respGetProject.GetProject()

	s.assertProject(proj, gotProject)

	issue := fixtures.CreateIssue(
		fixtures.WithProject(respCreateProject.GetId()),
	)

	respCreateIssue, err := s.issueService.CreateIssue(s.ctx, &desc.IssueCreateRequest{Issue: issue})
	require.NoError(s.T(), err)
	issueID := respCreateIssue.GetId()

	require.Greater(s.T(), issueID, int64(0))

	respGetIssue, err := s.issueService.GetIssueByID(s.ctx, &desc.IssueGetByIDRequest{Id: issueID})
	require.NoError(s.T(), err)
	gotIssue := respGetIssue.GetIssue()

	s.assertIssues(issue, gotIssue)
	s.assertCompositeName(proj, gotIssue, 1)
}

func (s *IssueSuite) TestIssues_Comments() {
	proj := fixtures.CreateProject()

	respCreateProject, err := s.projectService.CreateProject(s.ctx, &project.ProjectCreateRequest{Project: proj})
	require.NoError(s.T(), err)
	projectID := respCreateProject.GetId()

	require.Greater(s.T(), projectID, int64(0))

	respGetProject, err := s.projectService.GetProjectByID(s.ctx, &project.ProjectGetByIDRequest{Id: projectID})
	require.NoError(s.T(), err)
	gotProject := respGetProject.GetProject()

	s.assertProject(proj, gotProject)

	issue := fixtures.CreateIssue(
		fixtures.WithProject(respCreateProject.GetId()),
		fixtures.WithComment("Bob", "Any dummy comment"),
	)

	respCreateIssue, err := s.issueService.CreateIssue(s.ctx, &desc.IssueCreateRequest{Issue: issue})
	require.NoError(s.T(), err)
	issueID := respCreateIssue.GetId()

	require.Greater(s.T(), issueID, int64(0))

	respGetIssue, err := s.issueService.GetIssueByID(s.ctx, &desc.IssueGetByIDRequest{Id: issueID})
	require.NoError(s.T(), err)
	gotIssue := respGetIssue.GetIssue()

	s.assertIssues(issue, gotIssue)
	s.assertCompositeName(proj, gotIssue, 1)
	s.assertComments(issue.GetComments(), gotIssue.GetComments())
}

func (s *IssueSuite) TestIssues_Children() {
	proj := fixtures.CreateProject()

	respCreateProject, err := s.projectService.CreateProject(s.ctx, &project.ProjectCreateRequest{Project: proj})
	require.NoError(s.T(), err)
	projectID := respCreateProject.GetId()

	require.Greater(s.T(), projectID, int64(0))

	respGetProject, err := s.projectService.GetProjectByID(s.ctx, &project.ProjectGetByIDRequest{Id: projectID})
	require.NoError(s.T(), err)
	gotProject := respGetProject.GetProject()

	s.assertProject(proj, gotProject)

	issue := fixtures.CreateIssue(
		fixtures.WithProject(respCreateProject.GetId()),
		fixtures.WithComment("Bob", "Any dummy comment"),
	)

	respCreateIssue, err := s.issueService.CreateIssue(s.ctx, &desc.IssueCreateRequest{Issue: issue})
	require.NoError(s.T(), err)
	issueID := respCreateIssue.GetId()

	require.Greater(s.T(), issueID, int64(0))

	child := fixtures.CreateChildIssue(
		fixtures.WithProject(respCreateProject.GetId()),
	)

	respCreateChild, err := s.issueService.CreateSubtask(s.ctx, &desc.IssueCreateSubtaskRequest{ParentId: issueID, Child: child})
	require.NoError(s.T(), err)
	childID := respCreateChild.GetId()

	require.Greater(s.T(), childID, int64(0))

	respGetIssue, err := s.issueService.GetIssueByID(s.ctx, &desc.IssueGetByIDRequest{Id: issueID})
	require.NoError(s.T(), err)
	gotIssue := respGetIssue.GetIssue()

	s.assertIssues(issue, gotIssue)
	s.assertCompositeName(proj, gotIssue, 1)
	s.assertComments(issue.GetComments(), gotIssue.GetComments())

	require.Equal(s.T(), 1, len(gotIssue.GetChildren()))
	gotChild := gotIssue.GetChildren()[0]
	s.assertIssues(child, gotChild)
	s.assertCompositeName(proj, gotChild, 2)
	s.assertComments(child.GetComments(), gotChild.GetComments())
}

func (s *IssueSuite) assertIssues(expected, actual *desc.Issue) {
	require.Equal(s.T(), expected.GetName(), actual.GetName())
	require.Equal(s.T(), expected.GetParentId(), actual.GetParentId())
	require.Equal(s.T(), expected.GetType(), actual.GetType())
	require.Equal(s.T(), expected.GetDescription(), actual.GetDescription())
	require.Equal(s.T(), expected.GetAuthor(), actual.GetAuthor())
	require.Equal(s.T(), expected.GetAssigned(), actual.GetAssigned())
	require.Equal(s.T(), expected.GetQa(), actual.GetQa())
	require.Equal(s.T(), expected.GetReviewer(), actual.GetReviewer())
	require.Equal(s.T(), expected.GetTemplate(), actual.GetTemplate())
	require.WithinDuration(s.T(), expected.GetCreatedAt().AsTime(), actual.GetCreatedAt().AsTime(), 1*time.Second)
	require.WithinDuration(s.T(), expected.GetDeadline().AsTime(), actual.GetDeadline().AsTime(), 1*time.Second)
	require.Equal(s.T(), expected.GetStatus(), actual.GetStatus())
	require.Equal(s.T(), expected.GetPriority(), actual.GetPriority())
	require.Equal(s.T(), expected.GetStoryPoints(), actual.GetStoryPoints())
	require.Equal(s.T(), expected.GetSprintId(), actual.GetSprintId())
	require.Equal(s.T(), expected.GetProjectId(), actual.GetProjectId())
}

func (s *IssueSuite) assertCompositeName(proj *project.Project, issue *desc.Issue, number int) {
	require.Equal(s.T(), fmt.Sprintf("%s-%v", proj.GetShortName(), number), issue.GetCompositeName())
}

func (s *IssueSuite) assertComments(expected, actual *comments.Comments) {
	expComms, actComms := expected.GetItems(), actual.GetItems()
	require.Equal(s.T(), len(expComms), len(actComms))

	for i := 0; i < len(expComms); i++ {
		require.Equal(s.T(), expComms[i].GetAuthor(), actComms[i].GetAuthor())
		require.Equal(s.T(), expComms[i].GetText(), actComms[i].GetText())
		require.WithinDuration(s.T(), expComms[i].GetWrittenAt().AsTime(), actComms[i].GetWrittenAt().AsTime(), 1*time.Second)
	}
}

func (s *IssueSuite) assertProject(expected, actual *project.Project) {
	require.Equal(s.T(), expected.GetName(), actual.GetName())
	require.Equal(s.T(), expected.GetShortName(), actual.GetShortName())
	require.Equal(s.T(), expected.GetIsArchived(), actual.GetIsArchived())
	require.Equal(s.T(), expected.GetResponsibleTeamIds(), actual.GetResponsibleTeamIds())
}
