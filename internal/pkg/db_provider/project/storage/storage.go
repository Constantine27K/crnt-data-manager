package storage

import (
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/project/gateway"
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/project/models"
	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/project"
)

type ProjectStorage interface {
	Add(project *desc.Project) (int64, error)
	AddResponsibleTeam(projectID, teamID int64) (int64, error)
	RemoveResponsibleTeam(projectID, teamID int64) (int64, error)
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
		ShortName:        project.GetShortName(),
		IsArchived:       false,
		ResponsibleTeams: project.GetResponsibleTeamIds(),
		Description:      project.GetDescription(),
		Department:       project.GetDepartment(),
		Responsible:      project.GetResponsible(),
	}

	return s.gw.Add(row)
}

func (s *storage) AddResponsibleTeam(projectID, teamID int64) (int64, error) {
	return s.gw.AddResponsibleTeam(projectID, teamID)
}

func (s *storage) RemoveResponsibleTeam(projectID, teamID int64) (int64, error) {
	return s.gw.RemoveResponsibleTeam(projectID, teamID)
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
			ShortName:          row.ShortName,
			IsArchived:         row.IsArchived,
			ResponsibleTeamIds: row.ResponsibleTeams,
			Description:        row.Description,
			Department:         row.Department,
			Responsible:        row.Responsible,
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
		ShortName:          row.ShortName,
		IsArchived:         row.IsArchived,
		ResponsibleTeamIds: row.ResponsibleTeams,
		Description:        row.Description,
		Department:         row.Department,
		Responsible:        row.Responsible,
	}

	return result, err
}

func (s *storage) Update(id int64, project *desc.Project) (int64, error) {
	row := &models.ProjectRow{
		ID:               id,
		Name:             project.GetName(),
		ShortName:        project.GetShortName(),
		IsArchived:       project.GetIsArchived(),
		ResponsibleTeams: project.GetResponsibleTeamIds(),
		Description:      project.GetDescription(),
		Department:       project.GetDepartment(),
		Responsible:      project.GetResponsible(),
	}

	return s.gw.Update(row)
}
