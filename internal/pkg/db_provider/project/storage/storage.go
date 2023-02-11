package storage

import (
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/project/gateway"
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/project/models"
	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/project"
)

type ProjectStorage interface {
	Add(project *desc.Project) (int64, error)
	AddResponsibleTeam(projectID, teamID int64) (int64, error)
	Get(filter *models.ProjectFilter) ([]*desc.Project, error)
	GetByID(id int64) (*desc.Project, error)
	Update(id int64, project *desc.Project) (int64, error)
}

type storage struct {
	gw gateway.ProjectGateway
}

func NewProjectStorage(gw gateway.ProjectGateway) ProjectStorage {
	return &storage{
		gw: gw,
	}
}

func (s *storage) Add(project *desc.Project) (int64, error) {
	row := &models.ProjectRow{
		Name:             project.GetName(),
		IsArchived:       false,
		ResponsibleTeams: project.GetResponsibleTeamIds(),
	}

	return s.gw.Add(row)
}

func (s *storage) AddResponsibleTeam(projectID, teamID int64) (int64, error) {
	err := s.gw.AddResponsibleTeam(projectID, teamID)
	if err != nil {
		return 0, err
	}

	return teamID, err
}

func (s *storage) Get(filter *models.ProjectFilter) ([]*desc.Project, error) {
	result := make([]*desc.Project, 0)

	rows, err := s.gw.Get(filter)
	if err != nil {
		return nil, err
	}

	for _, row := range rows {
		result = append(result, &desc.Project{
			Id:                 row.ID,
			Name:               row.Name,
			IsArchived:         row.IsArchived,
			ResponsibleTeamIds: row.ResponsibleTeams,
		})
	}

	return result, err
}

func (s *storage) GetByID(id int64) (*desc.Project, error) {
	row, err := s.gw.GetByID(id)
	if err != nil {
		return nil, err
	}

	result := &desc.Project{
		Id:                 row.ID,
		Name:               row.Name,
		IsArchived:         row.IsArchived,
		ResponsibleTeamIds: row.ResponsibleTeams,
	}

	return result, err
}

func (s *storage) Update(id int64, project *desc.Project) (int64, error) {
	row := &models.ProjectRow{
		ID:               id,
		Name:             project.GetName(),
		IsArchived:       false,
		ResponsibleTeams: project.GetResponsibleTeamIds(),
	}

	return s.gw.Update(row)
}
