package gateway

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/issue/models"
	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

type IssueGateway interface {
	Add(row *models.IssueRow) (int64, error)
	Update(row *models.IssueRow) (int64, error)
	Get(filter *models.IssueFilter) ([]*models.IssueRow, error)
	GetByID(id int64) (*models.IssueRow, error)
}

type gateway struct {
	db      *sql.DB
	builder sq.StatementBuilderType
}

func NewIssueGateWay(db *sql.DB) IssueGateway {
	return &gateway{
		db:      db,
		builder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

const (
	table = "issue"
)

var (
	columns = []string{"id", "composite_name", "name", "issue_type", "parent_id", "description",
		"comments", "author", "assigned", "qa", "reviewer", "template", "created_at", "updated_at",
		"deadline", "status", "priority", "sprint_id", "project_id", "components", "story_points", "children"}
)

func (g *gateway) Add(issueRow *models.IssueRow) (int64, error) {
	values := []interface{}{
		issueRow.CompositeName, issueRow.Name, issueRow.IssueType, issueRow.ParentID, issueRow.Description, issueRow.Comments,
		issueRow.Author, issueRow.Assigned, issueRow.QA, issueRow.Reviewer, issueRow.Template, issueRow.CreatedAt, issueRow.UpdatedAt,
		issueRow.Deadline, issueRow.Status, issueRow.Priority, issueRow.SprintID, issueRow.ProjectID,
		pq.Array(issueRow.Components), issueRow.StoryPoints, pq.Array(issueRow.Children),
	}

	query, args, err := g.builder.Insert(table).
		Columns(columns[1:]...).
		Values(values...).
		Suffix("returning id").ToSql()
	if err != nil {
		log.Error("Gateway.Add query error",
			zap.Any("issue", issueRow),
			zap.Error(err),
		)
		return 0, err
	}

	var id int64
	err = g.db.QueryRow(query, args...).Scan(&id)
	if err != nil {
		log.Error("Gateway.Add scan error",
			zap.Any("issue", issueRow),
			zap.Error(err),
		)
		return 0, err
	}

	return id, nil
}

func (g *gateway) Update(issueRow *models.IssueRow) (int64, error) {
	query := g.builder.Update(table).
		Where(sq.Eq{"id": issueRow.ID})

	if len(issueRow.Name) > 0 {
		query = query.Set("name", issueRow.Name)
	}

	if len(issueRow.Description) > 0 {
		query = query.Set("description", issueRow.Description)
	}

	if len(issueRow.Assigned) > 0 {
		query = query.Set("assigned", issueRow.Assigned)
	}

	if len(issueRow.QA) > 0 {
		query = query.Set("qa", issueRow.QA)
	}

	if len(issueRow.Reviewer) > 0 {
		query = query.Set("reviewer", issueRow.Reviewer)
	}

	if issueRow.Deadline.After(time.Now().UTC()) {
		query = query.Set("deadline", issueRow.Deadline)
	}

	if issueRow.Priority != models.PriorityUnknown {
		query = query.Set("priority", issueRow.Priority)
	}

	if issueRow.SprintID != 0 {
		query = query.Set("sprint_id", issueRow.SprintID)
	}

	query = query.Set("components", pq.Array(issueRow.Components)).
		Set("story_points", issueRow.StoryPoints).
		Set("updated_at", issueRow.UpdatedAt)

	stmt, args, err := query.ToSql()
	if err != nil {
		log.Error("Gateway.Update query error",
			zap.Any("issue", issueRow),
			zap.Error(err),
		)
		return 0, err
	}

	res, err := g.db.Exec(stmt, args...)
	if err != nil {
		log.Error("Gateway.Update exec error",
			zap.Any("issue", issueRow),
			zap.Error(err),
		)
		return 0, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		log.Error("Gateway.Update affected error",
			zap.Any("issue", issueRow),
			zap.Error(err),
		)
		return 0, err
	}

	if affected == 0 {
		log.Error("Gateway.Update no rows affected",
			zap.Any("issue", issueRow),
			zap.Error(err),
		)
		return 0, fmt.Errorf("no rows affected")
	}

	return issueRow.ID, nil
}

func (g *gateway) Get(filter *models.IssueFilter) ([]*models.IssueRow, error) {
	query := g.builder.Select(columns...).
		From(table)

	if filter != nil {
		query = filter.Apply(query)
	}

	result := make([]*models.IssueRow, 0)

	stmt, args, err := query.ToSql()
	if err != nil {
		log.Error("Gateway.Get query error",
			zap.Any("filter", filter),
			zap.Error(err),
		)
		return nil, err
	}

	rows, err := g.db.Query(stmt, args...)
	if err != nil {
		log.Error("Gateway.Get query error",
			zap.Any("filter", filter),
			zap.Error(err),
		)
		return nil, err
	}

	for rows.Next() {
		var issueRow models.IssueRow
		err = rows.Scan(
			&issueRow.ID,
			&issueRow.CompositeName,
			&issueRow.Name,
			&issueRow.IssueType,
			&issueRow.ParentID,
			&issueRow.Description,
			&issueRow.Comments,
			&issueRow.Author,
			&issueRow.Assigned,
			&issueRow.QA,
			&issueRow.Reviewer,
			&issueRow.Template,
			&issueRow.CreatedAt,
			&issueRow.UpdatedAt,
			&issueRow.Deadline,
			&issueRow.Status,
			&issueRow.Priority,
			&issueRow.SprintID,
			&issueRow.ProjectID,
			pq.Array(&issueRow.Components),
			&issueRow.StoryPoints,
			pq.Array(&issueRow.Children),
		)
		if err != nil {
			log.Error("Gateway.Get scan error",
				zap.Any("filter", filter),
				zap.Error(err),
			)
			return nil, err
		}
		result = append(result, &issueRow)
	}

	return result, nil
}

func (g *gateway) GetByID(id int64) (*models.IssueRow, error) {
	query, args, err := g.builder.Select(columns...).
		From(table).
		Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		log.Error("Gateway.GetByID query error",
			zap.Int64("id", id),
			zap.Error(err),
		)
		return nil, err
	}

	var issueRow models.IssueRow
	err = g.db.QueryRow(query, args...).Scan(&issueRow.ID,
		&issueRow.CompositeName,
		&issueRow.Name,
		&issueRow.IssueType,
		&issueRow.ParentID,
		&issueRow.Description,
		&issueRow.Comments,
		&issueRow.Author,
		&issueRow.Assigned,
		&issueRow.QA,
		&issueRow.Reviewer,
		&issueRow.Template,
		&issueRow.CreatedAt,
		&issueRow.UpdatedAt,
		&issueRow.Deadline,
		&issueRow.Status,
		&issueRow.Priority,
		&issueRow.SprintID,
		&issueRow.ProjectID,
		pq.Array(&issueRow.Components),
		&issueRow.StoryPoints,
		pq.Array(&issueRow.Children))
	if err != nil {
		log.Error("Gateway.GetByID scan error",
			zap.Int64("id", id),
			zap.Error(err),
		)
		return nil, err
	}

	return &issueRow, nil
}
