package storage

import (
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/department/gateway"
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/department/models"
	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/department"
)

type DepartmentStorage interface {
	Add(department *desc.Department) (int64, error)
	AddProject(id int64, projectID int64) (int64, error)
	RemoveProject(id int64, projectID int64) (int64, error)
	Get(filter *models.DepartmentFilter) ([]*desc.Department, error)
	GetByID(id int64) (*desc.Department, error)
	Update(id int64, department *desc.Department) (int64, error)
}

type storage struct {
	gw gateway.DepartmentGateway
}

func NewDepartmentStorage(gw gateway.DepartmentGateway) DepartmentStorage {
	return &storage{
		gw: gw,
	}
}

func (s *storage) Add(department *desc.Department) (int64, error) {
	row := &models.DepartmentRow{
		Name:     department.GetName(),
		Projects: department.GetProjects(),
	}

	return s.gw.Add(row)
}

func (s *storage) AddProject(id int64, projectID int64) (int64, error) {
	return s.gw.AddProject(id, projectID)
}

func (s *storage) RemoveProject(id int64, projectID int64) (int64, error) {
	return s.gw.RemoveProject(id, projectID)
}
func (s *storage) Get(filter *models.DepartmentFilter) ([]*desc.Department, error) {
	rows, err := s.gw.Get(filter)
	if err != nil {
		return nil, err
	}

	result := make([]*desc.Department, 0)

	for _, row := range rows {
		result = append(result, &desc.Department{
			Id:       row.ID,
			Name:     row.Name,
			Projects: row.Projects,
		})
	}

	return result, nil
}

func (s *storage) GetByID(id int64) (*desc.Department, error) {
	row, err := s.gw.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &desc.Department{
		Id:       row.ID,
		Name:     row.Name,
		Projects: row.Projects,
	}, nil
}

func (s *storage) Update(id int64, department *desc.Department) (int64, error) {
	row := &models.DepartmentRow{
		ID:       id,
		Name:     department.GetName(),
		Projects: department.GetProjects(),
	}

	return s.gw.Update(row)
}
