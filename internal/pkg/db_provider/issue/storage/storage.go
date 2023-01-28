package storage

import (
	"time"

	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/issue/gateway"
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/issue/models"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/comments"
	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/issue"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type IssueStorage interface {
	Add(issue *desc.Issue) (int64, error)
	Update(id int64, issue *desc.Issue) (int64, error)
	Get(filter *models.IssueFilter) ([]*desc.Issue, error)
	GetByID(id int64) (*desc.Issue, error)
}

type storage struct {
	gw gateway.IssueGateway
}

func NewIssueStorage(gw gateway.IssueGateway) IssueStorage {
	return &storage{gw: gw}
}

func (s *storage) Add(issue *desc.Issue) (int64, error) {
	comm, err := protojson.MarshalOptions{EmitUnpopulated: true}.Marshal(issue.GetComments())
	if err != nil {
		return 0, err
	}

	childrenIDs := make([]int64, 0, len(issue.GetChildren()))
	for _, child := range issue.GetChildren() {
		childrenIDs = append(childrenIDs, child.GetId())
	}

	componentIDs := make([]int64, 0, len(issue.GetComponentIds()))
	for _, compID := range issue.GetComponentIds() {
		componentIDs = append(componentIDs, compID)
	}

	issueType := models.MapFromProtoIssueType(issue.GetType())
	template := models.MapFromProtoTemplate(issue.GetTemplate())
	status := models.MapFromProtoStatus(issue.GetStatus())
	priority := models.MapFromProtoPriority(issue.GetPriority())

	now := time.Now().UTC()

	model := &models.IssueRow{
		CompositeName: issue.GetCompositeName(),
		Name:          issue.GetName(),
		IssueType:     issueType,
		ParentID:      issue.GetParentId(),
		Description:   issue.GetDescription(),
		Comments:      string(comm),
		Author:        issue.GetAuthor(),
		Assigned:      issue.GetAssigned(),
		QA:            issue.GetQa(),
		Reviewer:      issue.GetReviewer(),
		CreatedAt:     now,
		UpdatedAt:     now,
		Template:      template,
		Deadline:      issue.GetDeadline().AsTime().UTC(),
		Status:        status,
		Priority:      priority,
		SprintID:      issue.GetSprintId(),
		ProjectID:     issue.GetProjectId(),
		Components:    componentIDs,
		StoryPoints:   issue.GetStoryPoints(),
		Children:      childrenIDs,
	}

	return s.gw.Add(model)
}

func (s *storage) Update(id int64, issue *desc.Issue) (int64, error) {
	componentIDs := make([]int64, 0, len(issue.GetComponentIds()))
	for _, compID := range issue.GetComponentIds() {
		componentIDs = append(componentIDs, compID)
	}

	priority := models.MapFromProtoPriority(issue.GetPriority())

	now := time.Now().UTC()

	model := &models.IssueRow{
		ID:          id,
		Name:        issue.GetName(),
		Description: issue.GetDescription(),
		Assigned:    issue.GetAssigned(),
		QA:          issue.GetQa(),
		Reviewer:    issue.GetReviewer(),
		UpdatedAt:   now,
		Deadline:    issue.GetDeadline().AsTime().UTC(),
		Priority:    priority,
		SprintID:    issue.GetSprintId(),
		Components:  componentIDs,
		StoryPoints: issue.GetStoryPoints(),
	}

	return s.gw.Update(model)
}

func (s *storage) Get(filter *models.IssueFilter) ([]*desc.Issue, error) {
	issueRows, err := s.gw.Get(filter)
	if err != nil {
		return nil, err
	}

	result := make([]*desc.Issue, 0, len(issueRows))

	for _, row := range issueRows {
		issueType := models.MapToProtoIssueType(row.IssueType)
		template := models.MapToProtoTemplate(row.Template)
		priority := models.MapToProtoPriority(row.Priority)
		status := models.MapToProtoStatus(row.Status)

		var comm comments.Comments
		err = protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(row.Comments), &comm)
		if err != nil {
			return nil, err
		}

		var children []*desc.Issue
		var chErr error
		if len(row.Children) > 0 {
			children, chErr = s.Get(&models.IssueFilter{IDs: row.Children})
			if chErr != nil {
				return nil, chErr
			}
		}

		result = append(result, &desc.Issue{
			Id:            row.ID,
			CompositeName: row.CompositeName,
			Name:          row.Name,
			Type:          issueType,
			ParentId:      row.ParentID,
			Description:   row.Description,
			Comments:      &comm,
			Author:        row.Author,
			Assigned:      row.Assigned,
			Qa:            row.QA,
			Reviewer:      row.Reviewer,
			Template:      template,
			CreatedAt:     timestamppb.New(row.CreatedAt),
			UpdatedAt:     timestamppb.New(row.UpdatedAt),
			Deadline:      timestamppb.New(row.Deadline),
			Status:        status,
			Priority:      priority,
			SprintId:      row.SprintID,
			ProjectId:     row.ProjectID,
			ComponentIds:  row.Components,
			StoryPoints:   row.StoryPoints,
			Children:      children,
		})
	}

	return result, nil
}

func (s *storage) GetByID(id int64) (*desc.Issue, error) {
	row, err := s.gw.GetByID(id)
	if err != nil {
		return nil, err
	}

	issueType := models.MapToProtoIssueType(row.IssueType)
	template := models.MapToProtoTemplate(row.Template)
	priority := models.MapToProtoPriority(row.Priority)
	status := models.MapToProtoStatus(row.Status)

	var comm comments.Comments
	err = protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(row.Comments), &comm)
	if err != nil {
		return nil, err
	}

	var children []*desc.Issue
	var chErr error
	if len(row.Children) > 0 {
		children, chErr = s.Get(&models.IssueFilter{IDs: row.Children})
		if chErr != nil {
			return nil, chErr
		}
	}

	return &desc.Issue{
		Id:            row.ID,
		CompositeName: row.CompositeName,
		Name:          row.Name,
		Type:          issueType,
		ParentId:      row.ParentID,
		Description:   row.Description,
		Comments:      &comm,
		Author:        row.Author,
		Assigned:      row.Assigned,
		Qa:            row.QA,
		Reviewer:      row.Reviewer,
		Template:      template,
		CreatedAt:     timestamppb.New(row.CreatedAt),
		UpdatedAt:     timestamppb.New(row.UpdatedAt),
		Deadline:      timestamppb.New(row.Deadline),
		Status:        status,
		Priority:      priority,
		SprintId:      row.SprintID,
		ProjectId:     row.ProjectID,
		ComponentIds:  row.Components,
		StoryPoints:   row.StoryPoints,
		Children:      children,
	}, nil
}