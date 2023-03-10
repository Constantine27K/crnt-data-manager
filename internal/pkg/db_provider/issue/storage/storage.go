package storage

import (
	"fmt"
	"time"

	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/issue/gateway"
	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/issue/models"
	projectGateway "github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/project/gateway"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/comments"
	desc "github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/issue"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type IssueStorage interface {
	Add(issue *desc.Issue) (int64, error)
	CreateSubtask(parentID int64, child *desc.Issue) (int64, error)
	AddComment(id int64, comment *comments.Comment) (int64, error)
	Get(filter *models.IssueFilter) ([]*desc.Issue, error)
	GetInfo(filter *models.IssueFilter) ([]*desc.IssueInfo, error)
	GetByID(id int64) (*desc.Issue, error)
	GetInfoByID(id int64) (*desc.IssueInfo, error)
	Update(id int64, issue *desc.Issue) (int64, error)
}

type storage struct {
	gw        gateway.IssueGateway
	projectGw projectGateway.ProjectGateway
}

func NewIssueStorage(
	gw gateway.IssueGateway,
	projectGw projectGateway.ProjectGateway,
) IssueStorage {
	return &storage{
		gw:        gw,
		projectGw: projectGw,
	}
}

func (s *storage) Add(issue *desc.Issue) (int64, error) {
	comm, err := protojson.MarshalOptions{EmitUnpopulated: true}.Marshal(issue.GetComments())
	if err != nil {
		return 0, err
	}

	compositeName, err := s.generateCompositeName(issue)
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
		CompositeName: compositeName,
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

func (s *storage) generateCompositeName(issue *desc.Issue) (string, error) {
	projectShortName, err := s.projectGw.GetShortName(issue.GetProjectId())
	if err != nil {
		return "", err
	}

	lastProjectID, err := s.gw.GetProjectLastID(issue.GetProjectId())
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s-%v", projectShortName, lastProjectID+1), nil
}

func (s *storage) CreateSubtask(parentID int64, child *desc.Issue) (int64, error) {
	parent, err := s.GetByID(parentID)
	if err != nil {
		return 0, err
	}

	child.SprintId = parent.SprintId
	child.ProjectId = parent.ProjectId
	child.Template = parent.Template

	childID, err := s.Add(child)
	if err != nil {
		return 0, err
	}

	err = s.gw.AddChild(parentID, childID)
	if err != nil {
		return 0, err
	}

	return childID, nil
}

func (s *storage) AddComment(id int64, comment *comments.Comment) (int64, error) {
	issue, err := s.gw.GetByID(id)
	if err != nil {
		return 0, err
	}

	var oldComments comments.Comments
	err = protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal([]byte(issue.Comments), &oldComments)
	if err != nil {
		return 0, err
	}

	newComments := make([]*comments.Comment, 0, len(oldComments.GetItems())+1)
	newComments = append(newComments, oldComments.GetItems()...)
	newComments = append(newComments, comment)

	comms, err := protojson.MarshalOptions{EmitUnpopulated: true}.Marshal(&comments.Comments{Items: newComments})
	if err != nil {
		return 0, err
	}

	return s.gw.AddComment(id, string(comms))
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

func (s *storage) GetInfo(filter *models.IssueFilter) ([]*desc.IssueInfo, error) {
	issueInfoRows, err := s.gw.GetInfo(filter)
	if err != nil {
		return nil, err
	}

	result := make([]*desc.IssueInfo, 0, len(issueInfoRows))

	for _, row := range issueInfoRows {
		issueType := models.MapToProtoIssueType(row.IssueType)
		status := models.MapToProtoStatus(row.Status)
		priority := models.MapToProtoPriority(row.Priority)

		result = append(result, &desc.IssueInfo{
			Id:            row.ID,
			CompositeName: row.CompositeName,
			Name:          row.Name,
			Type:          issueType,
			Assigned:      row.Assigned,
			Status:        status,
			Priority:      priority,
			StoryPoints:   row.StoryPoints,
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

func (s *storage) GetInfoByID(id int64) (*desc.IssueInfo, error) {
	row, err := s.gw.GetInfoByID(id)
	if err != nil {
		return nil, err
	}

	issueType := models.MapToProtoIssueType(row.IssueType)
	priority := models.MapToProtoPriority(row.Priority)
	status := models.MapToProtoStatus(row.Status)

	return &desc.IssueInfo{
		Id:            row.ID,
		CompositeName: row.CompositeName,
		Name:          row.Name,
		Type:          issueType,
		Assigned:      row.Assigned,
		Status:        status,
		Priority:      priority,
		StoryPoints:   row.StoryPoints,
	}, nil
}
