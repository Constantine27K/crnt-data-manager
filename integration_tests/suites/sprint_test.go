package suites

import (
	"time"

	"github.com/Constantine27K/crnt-data-manager/integration_tests/fixtures"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/project"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/sprint"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/state/status"
	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/issue"
	"github.com/stretchr/testify/require"
)

var (
	now = time.Now()
)

func (s *CrntDMSuite) TestSprint_Create() {
	sp := fixtures.CreateSprint(
		fixtures.WithDates(now, now.Add(time.Hour)),
	)

	respCreateSprint, err := s.sprintService.CreateSprint(s.ctx, &sprint.SprintCreateRequest{Sprint: sp})
	require.NoError(s.T(), err)
	sprintID := respCreateSprint.GetId()

	require.Greater(s.T(), sprintID, int64(0))

	respGetSprint, err := s.sprintService.GetSprintByID(s.ctx, &sprint.SprintGetByIDRequest{Id: sprintID})
	require.NoError(s.T(), err)
	gotSprint := respGetSprint.GetSprint()

	s.assertSprint(sp, gotSprint)
}

func (s *CrntDMSuite) TestSprint_Update() {
	sp := fixtures.CreateSprint(
		fixtures.WithDates(now, now.Add(time.Hour)),
	)

	respCreateSprint, err := s.sprintService.CreateSprint(s.ctx, &sprint.SprintCreateRequest{Sprint: sp})
	require.NoError(s.T(), err)
	sprintID := respCreateSprint.GetId()

	require.Greater(s.T(), sprintID, int64(0))

	respGetSprint, err := s.sprintService.GetSprintByID(s.ctx, &sprint.SprintGetByIDRequest{Id: sprintID})
	require.NoError(s.T(), err)
	gotSprint := respGetSprint.GetSprint()

	s.assertSprint(sp, gotSprint)

	updSprint := fixtures.CreateSprint(
		fixtures.WithStatus(status.Sprint_STATUS_SPRINT_ACTIVE),
		fixtures.WithDates(now, now.Add(time.Hour)),
	)
	respUpdateSprint, err := s.sprintService.UpdateSprint(s.ctx, &sprint.SprintUpdateRequest{Id: sprintID, Sprint: updSprint})
	require.NoError(s.T(), err)
	sprintUpdatedID := respUpdateSprint.GetId()

	require.Equal(s.T(), sprintID, sprintUpdatedID)

	respGetUpdatedSprint, err := s.sprintService.GetSprintByID(s.ctx, &sprint.SprintGetByIDRequest{Id: sprintID})
	require.NoError(s.T(), err)
	gotUpdatedSprint := respGetUpdatedSprint.GetSprint()

	s.assertSprint(updSprint, gotUpdatedSprint)
}

func (s *CrntDMSuite) TestSprint_AddIssue() {
	sp := fixtures.CreateSprint(
		fixtures.WithDates(now, now.Add(time.Hour)),
	)

	respCreateSprint, err := s.sprintService.CreateSprint(s.ctx, &sprint.SprintCreateRequest{Sprint: sp})
	require.NoError(s.T(), err)
	sprintID := respCreateSprint.GetId()

	require.Greater(s.T(), sprintID, int64(0))

	respGetSprint, err := s.sprintService.GetSprintByID(s.ctx, &sprint.SprintGetByIDRequest{Id: sprintID})
	require.NoError(s.T(), err)
	gotSprint := respGetSprint.GetSprint()

	s.assertSprint(sp, gotSprint)

	proj := fixtures.CreateProject(
		fixtures.WithResponsibleTeams(1),
	)

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

	respAddIssueToSprint, err := s.sprintService.AddIssue(s.ctx, &sprint.AddIssueRequest{SprintId: sprintID, IssueId: issueID})
	require.NoError(s.T(), err)
	sprintAddedID := respAddIssueToSprint.GetSprintId()

	require.Equal(s.T(), sprintID, sprintAddedID)

	respGetAddedSprint, err := s.sprintService.GetSprintByID(s.ctx, &sprint.SprintGetByIDRequest{Id: sprintID})
	require.NoError(s.T(), err)
	gotAddedSprint := respGetAddedSprint.GetSprint()

	require.Equal(s.T(), 1, len(gotAddedSprint.GetIssues()))
	require.Equal(s.T(), issueID, gotAddedSprint.GetIssues()[0])
}

func (s *CrntDMSuite) TestSprint_RemoveIssue() {
	sp := fixtures.CreateSprint(
		fixtures.WithDates(now, now.Add(time.Hour)),
	)

	respCreateSprint, err := s.sprintService.CreateSprint(s.ctx, &sprint.SprintCreateRequest{Sprint: sp})
	require.NoError(s.T(), err)
	sprintID := respCreateSprint.GetId()

	require.Greater(s.T(), sprintID, int64(0))

	respGetSprint, err := s.sprintService.GetSprintByID(s.ctx, &sprint.SprintGetByIDRequest{Id: sprintID})
	require.NoError(s.T(), err)
	gotSprint := respGetSprint.GetSprint()

	s.assertSprint(sp, gotSprint)

	proj := fixtures.CreateProject(
		fixtures.WithResponsibleTeams(1),
	)

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

	respAddIssueToSprint, err := s.sprintService.AddIssue(s.ctx, &sprint.AddIssueRequest{SprintId: sprintID, IssueId: issueID})
	require.NoError(s.T(), err)
	sprintAddedID := respAddIssueToSprint.GetSprintId()

	require.Equal(s.T(), sprintID, sprintAddedID)

	respGetAddedSprint, err := s.sprintService.GetSprintByID(s.ctx, &sprint.SprintGetByIDRequest{Id: sprintID})
	require.NoError(s.T(), err)
	gotAddedSprint := respGetAddedSprint.GetSprint()

	require.Equal(s.T(), 1, len(gotAddedSprint.GetIssues()))
	require.Equal(s.T(), issueID, gotAddedSprint.GetIssues()[0])

	respRemoveIssueFromSprint, err := s.sprintService.RemoveIssue(s.ctx, &sprint.RemoveIssueRequest{SprintId: sprintID, IssueId: issueID})
	require.NoError(s.T(), err)
	sprintRemoveID := respRemoveIssueFromSprint.GetSprintId()

	require.Equal(s.T(), sprintID, sprintRemoveID)

	respGetRemovedSprint, err := s.sprintService.GetSprintByID(s.ctx, &sprint.SprintGetByIDRequest{Id: sprintID})
	require.NoError(s.T(), err)
	gotRemovedSprint := respGetRemovedSprint.GetSprint()

	require.Equal(s.T(), 0, len(gotRemovedSprint.GetIssues()))
}

func (s *CrntDMSuite) assertSprint(expected, actual *sprint.Sprint) {
	require.Equal(s.T(), expected.GetName(), actual.GetName())
	require.Equal(s.T(), expected.GetProject(), actual.GetProject())
	require.WithinDuration(s.T(), expected.GetStartedAt().AsTime(), actual.GetStartedAt().AsTime(), 1*time.Second)
	require.WithinDuration(s.T(), expected.GetFinishedAt().AsTime(), actual.GetFinishedAt().AsTime(), 1*time.Second)
	require.Equal(s.T(), expected.GetStatus(), actual.GetStatus())
	require.Equal(s.T(), expected.GetStoriesBacklog(), actual.GetStoriesBacklog())
	require.Equal(s.T(), expected.GetStoriesInProgress(), actual.GetStoriesInProgress())
	require.Equal(s.T(), expected.GetStoriesDone(), actual.GetStoriesDone())
	require.Equal(s.T(), expected.GetIssues(), actual.GetIssues())
}
