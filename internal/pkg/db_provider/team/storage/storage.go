package storage

import (
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/team/gateway"
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/team/models"
	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/team"
)

type TeamStorage interface {
	Add(team *desc.Team) (int64, error)
	Get(filter *models.TeamFilter) ([]*desc.Team, error)
	GetByID(id int64) (*desc.Team, error)
	Update(id int64, team *desc.Team) (int64, error)
}

type storage struct {
	gw gateway.TeamGateway
}

func NewTeamStorage(gw gateway.TeamGateway) TeamStorage {
	return &storage{
		gw: gw,
	}
}

func (s *storage) Add(team *desc.Team) (int64, error) {
	row := &models.TeamRow{
		Name:          team.GetName(),
		Members:       team.GetMembers(),
		Projects:      team.GetProjects(),
		TechOwner:     team.GetTechOwner(),
		BusinessOwner: team.GetBusinessOwner(),
		Department:    team.GetDepartment(),
	}

	return s.gw.Add(row)
}

func (s *storage) Get(filter *models.TeamFilter) ([]*desc.Team, error) {
	rows, err := s.gw.Get(filter)
	if err != nil {
		return nil, err
	}

	result := make([]*desc.Team, 0, len(rows))

	for _, row := range rows {
		result = append(result, &desc.Team{
			Id:            row.ID,
			Name:          row.Name,
			Members:       row.Members,
			Projects:      row.Projects,
			TechOwner:     row.TechOwner,
			BusinessOwner: row.BusinessOwner,
			Department:    row.Department,
		})
	}

	return result, nil
}

func (s *storage) GetByID(id int64) (*desc.Team, error) {
	row, err := s.gw.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &desc.Team{
		Id:            row.ID,
		Name:          row.Name,
		Members:       row.Members,
		Projects:      row.Projects,
		TechOwner:     row.TechOwner,
		BusinessOwner: row.BusinessOwner,
		Department:    row.Department,
	}, nil
}

func (s *storage) Update(id int64, team *desc.Team) (int64, error) {
	row := &models.TeamRow{
		ID:            id,
		Name:          team.GetName(),
		Members:       team.GetMembers(),
		Projects:      team.GetProjects(),
		TechOwner:     team.GetTechOwner(),
		BusinessOwner: team.GetBusinessOwner(),
		Department:    team.GetDepartment(),
	}

	return s.gw.Update(row)
}
