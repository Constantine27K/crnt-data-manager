package suites

import (
	"github.com/Constantine27K/crnt-data-manager/integration_tests/fixtures"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/team"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
)

func (s *CrntDMSuite) TestTeam_Create() {
	t := fixtures.CreateTeam()

	respTeamCreate, err := s.teamService.CreateTeam(s.ctx, &team.TeamCreateRequest{Team: t})
	require.NoError(s.T(), err)
	teamID := respTeamCreate.GetId()

	require.Greater(s.T(), teamID, int64(0))

	respGetTeam, err := s.teamService.GetTeamByID(s.ctx, &team.TeamGetByIDRequest{Id: teamID})
	require.NoError(s.T(), err)
	gotTeam := respGetTeam.GetTeam()

	s.assertTeams(t, gotTeam)
}

func (s *CrntDMSuite) TestTeam_Update() {
	t := fixtures.CreateTeam()

	respTeamCreate, err := s.teamService.CreateTeam(s.ctx, &team.TeamCreateRequest{Team: t})
	require.NoError(s.T(), err)
	teamID := respTeamCreate.GetId()

	require.Greater(s.T(), teamID, int64(0))

	respGetTeam, err := s.teamService.GetTeamByID(s.ctx, &team.TeamGetByIDRequest{Id: teamID})
	require.NoError(s.T(), err)
	gotTeam := respGetTeam.GetTeam()

	s.assertTeams(t, gotTeam)

	updTeam := fixtures.CreateTeam(
		fixtures.WithMembers(gofakeit.Name(), gofakeit.Name(), gofakeit.Name()),
	)

	respUpdateTeam, err := s.teamService.UpdateTeam(s.ctx, &team.TeamUpdateRequest{Id: teamID, Team: updTeam})
	require.NoError(s.T(), err)
	updTeamID := respUpdateTeam.GetId()

	require.Equal(s.T(), teamID, updTeamID)

	respGetUpdatedTeam, err := s.teamService.GetTeamByID(s.ctx, &team.TeamGetByIDRequest{Id: teamID})
	require.NoError(s.T(), err)
	gotUpdatedTeam := respGetUpdatedTeam.GetTeam()

	s.assertTeams(updTeam, gotUpdatedTeam)
}

func (s *CrntDMSuite) assertTeams(expected, actual *team.Team) {
	require.Equal(s.T(), expected.GetName(), actual.GetName())
	require.Equal(s.T(), expected.GetMembers(), actual.GetMembers())
	require.Equal(s.T(), expected.GetProjects(), actual.GetProjects())
	require.Equal(s.T(), expected.GetTechOwner(), actual.GetTechOwner())
	require.Equal(s.T(), expected.GetBusinessOwner(), actual.GetBusinessOwner())
	require.Equal(s.T(), expected.GetDepartment(), actual.GetDepartment())
}
