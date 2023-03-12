package suites

import (
	"github.com/Constantine27K/crnt-data-manager/integration_tests/fixtures"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/project"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/sprint"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/state/status"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *CrntDMSuite) TestProject_Create() {
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

	respGetDefaultSprint, err := s.sprintService.GetSprint(s.ctx, &sprint.SprintGetRequest{Projects: []int64{projectID}})
	require.NoError(s.T(), err)
	require.Equal(s.T(), 1, len(respGetDefaultSprint.GetSprints()))

	sp := respGetDefaultSprint.GetSprints()[0]
	require.Greater(s.T(), sp.GetId(), int64(0))
	require.Equal(s.T(), "Backlog", sp.GetName())
	require.Equal(s.T(), status.Sprint_STATUS_SPRINT_ACTIVE, sp.GetStatus())

	// should meet constraint unique_backlog
	_, err = s.sprintService.CreateSprint(s.ctx,
		&sprint.SprintCreateRequest{
			Sprint: &sprint.Sprint{
				Project:    projectID,
				Name:       "Backlog",
				StartedAt:  timestamppb.Now(),
				FinishedAt: timestamppb.Now(),
			},
		})
	require.Error(s.T(), err)
}

func (s *CrntDMSuite) TestProject_Update() {
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

	updProject := fixtures.CreateProject(
		fixtures.WithName("NewName"),
		fixtures.WithShortName("NewShortName"),
		fixtures.WithArchived(true),
		fixtures.WithResponsibleTeams(1, 2, 3),
	)

	respUpdateProject, err := s.projectService.UpdateProject(s.ctx, &project.ProjectUpdateRequest{Id: projectID, Project: updProject})
	require.NoError(s.T(), err)
	projectUpdatedID := respUpdateProject.GetId()

	require.Equal(s.T(), projectID, projectUpdatedID)

	respGetUpdatedProject, err := s.projectService.GetProjectByID(s.ctx, &project.ProjectGetByIDRequest{Id: projectID})
	require.NoError(s.T(), err)
	gotUpdatedProject := respGetUpdatedProject.GetProject()

	s.assertProject(updProject, gotUpdatedProject)
}

func (s *CrntDMSuite) TestProject_AddTeam() {
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

	respAddTeamProject, err := s.projectService.AddResponsibleTeam(s.ctx, &project.ProjectAddTeamRequest{ProjectId: projectID, TeamId: 2})
	require.NoError(s.T(), err)
	projectAddTeamID := respAddTeamProject.GetProjectId()

	require.Equal(s.T(), projectID, projectAddTeamID)

	respGetUpdatedProject, err := s.projectService.GetProjectByID(s.ctx, &project.ProjectGetByIDRequest{Id: projectID})
	require.NoError(s.T(), err)
	gotUpdatedProject := respGetUpdatedProject.GetProject()

	require.Equal(s.T(), []int64{1, 2}, gotUpdatedProject.GetResponsibleTeamIds())
}

func (s *CrntDMSuite) assertProject(expected, actual *project.Project) {
	require.Equal(s.T(), expected.GetName(), actual.GetName())
	require.Equal(s.T(), expected.GetShortName(), actual.GetShortName())
	require.Equal(s.T(), expected.GetIsArchived(), actual.GetIsArchived())
	require.Equal(s.T(), expected.GetResponsibleTeamIds(), actual.GetResponsibleTeamIds())
}
