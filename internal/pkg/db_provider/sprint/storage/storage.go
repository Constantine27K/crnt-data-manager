package storage

import (
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/sprint/gateway"
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/sprint/models"
	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/sprint"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/state/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SprintStorage interface {
	Add(sprint *desc.Sprint) (int64, error)
	AddIssue(sprintID, issueID int64) (int64, error)
	RemoveIssue(sprintID, issueID int64) (int64, error)
	Get(filter *models.SprintFilter) ([]*desc.Sprint, error)
	GetByID(id int64) (*desc.Sprint, error)
	Update(id int64, sprint *desc.Sprint) (int64, error)
	ChangeStatus(id int64, stat status.Sprint) (int64, error)
}

type storage struct {
	gw gateway.SprintGateway
}

func NewSprintStorage(gw gateway.SprintGateway) SprintStorage {
	return &storage{
		gw: gw,
	}
}

func (s *storage) Add(sprint *desc.Sprint) (int64, error) {
	mappedStatus := models.MapFromProtoStatus(sprint.GetStatus())

	row := &models.SprintRow{
		Name:       sprint.GetName(),
		Project:    sprint.GetProject(),
		StartedAt:  sprint.GetStartedAt().AsTime().UTC(),
		FinishedAt: sprint.GetFinishedAt().AsTime().UTC(),
		Status:     mappedStatus,
		InBacklog:  sprint.GetStoriesBacklog(),
		InProgress: sprint.GetStoriesInProgress(),
		InDone:     sprint.GetStoriesDone(),
		Issues:     sprint.GetIssues(),
	}

	return s.gw.Add(row)
}

func (s *storage) Get(filter *models.SprintFilter) ([]*desc.Sprint, error) {
	result := make([]*desc.Sprint, 0)

	rows, err := s.gw.Get(filter)
	if err != nil {
		return nil, err
	}

	for _, row := range rows {
		mappedStatus := models.MapToProtoStatus(row.Status)

		result = append(result, &desc.Sprint{
			Id:                row.ID,
			Name:              row.Name,
			Project:           row.Project,
			StartedAt:         timestamppb.New(row.StartedAt),
			FinishedAt:        timestamppb.New(row.FinishedAt),
			Status:            mappedStatus,
			StoriesBacklog:    row.InBacklog,
			StoriesInProgress: row.InProgress,
			StoriesDone:       row.InDone,
			Issues:            row.Issues,
		})
	}

	return result, err
}

func (s *storage) GetByID(id int64) (*desc.Sprint, error) {
	row, err := s.gw.GetByID(id)
	if err != nil {
		return nil, err
	}

	mappedStatus := models.MapToProtoStatus(row.Status)

	result := &desc.Sprint{
		Id:                row.ID,
		Name:              row.Name,
		Project:           row.Project,
		StartedAt:         timestamppb.New(row.StartedAt),
		FinishedAt:        timestamppb.New(row.FinishedAt),
		Status:            mappedStatus,
		StoriesBacklog:    row.InBacklog,
		StoriesInProgress: row.InProgress,
		StoriesDone:       row.InDone,
		Issues:            row.Issues,
	}

	return result, err
}

func (s *storage) Update(id int64, sprint *desc.Sprint) (int64, error) {
	mappedStatus := models.MapFromProtoStatus(sprint.GetStatus())

	row := &models.SprintRow{
		ID:         id,
		Name:       sprint.GetName(),
		Project:    sprint.GetProject(),
		StartedAt:  sprint.GetStartedAt().AsTime().UTC(),
		FinishedAt: sprint.GetFinishedAt().AsTime().UTC(),
		Status:     mappedStatus,
		InBacklog:  sprint.GetStoriesBacklog(),
		InProgress: sprint.GetStoriesInProgress(),
		InDone:     sprint.GetStoriesDone(),
		Issues:     sprint.GetIssues(),
	}

	return s.gw.Update(row)
}

func (s *storage) ChangeStatus(id int64, stat status.Sprint) (int64, error) {
	mappedStatus := models.MapFromProtoStatus(stat)

	return s.gw.ChangeStatus(id, mappedStatus)
}

func (s *storage) AddIssue(sprintID, issueID int64) (int64, error) {
	return s.gw.AddIssue(sprintID, issueID)
}

func (s *storage) RemoveIssue(sprintID, issueID int64) (int64, error) {
	return s.gw.RemoveIssue(sprintID, issueID)
}
