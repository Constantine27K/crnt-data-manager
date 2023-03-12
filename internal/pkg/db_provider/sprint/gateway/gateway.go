package gateway

import (
	"database/sql"
	"fmt"

	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/sprint/models"
	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

type SprintGateway interface {
	Add(row *models.SprintRow) (int64, error)
	AddIssue(sprintID, issueID int64) (int64, error)
	RemoveIssue(sprintID, issueID int64) (int64, error)
	Get(filter *models.SprintFilter) ([]*models.SprintRow, error)
	GetByID(id int64) (*models.SprintRow, error)
	Update(row *models.SprintRow) (int64, error)
	ChangeStatus(id int64, status models.Status) (int64, error)
}

const (
	table = "sprint"
)

var (
	columns = []string{"id", "name", "project", "started_at", "finished_at", "status",
		"in_backlog", "in_progress", "in_done", "issues"}
)

type gateway struct {
	db      *sql.DB
	builder sq.StatementBuilderType
}

func NewSprintGateway(db *sql.DB) SprintGateway {
	return &gateway{
		db:      db,
		builder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (g *gateway) Add(row *models.SprintRow) (int64, error) {
	values := []interface{}{
		row.Name, row.Project, row.StartedAt, row.FinishedAt, row.Status,
		row.InBacklog, row.InProgress, row.InDone, pq.Array(row.Issues),
	}

	query, args, err := g.builder.Insert(table).
		Columns(columns[1:]...).
		Values(values...).
		Suffix("returning id").ToSql()
	if err != nil {
		log.Error("Gateway.Add query error",
			zap.Any("sprint", row),
			zap.Error(err),
		)
		return 0, err
	}

	var id int64
	err = g.db.QueryRow(query, args...).Scan(&id)
	if err != nil {
		log.Error("Gateway.Add scan error",
			zap.Any("sprint", row),
			zap.Error(err),
		)
		return 0, err
	}

	return id, nil
}

func (g *gateway) AddIssue(sprintID, issueID int64) (int64, error) {
	query := "update sprint set issues = array_append(issues, $2) where id = $1"

	res, err := g.db.Exec(query, sprintID, issueID)
	if err != nil {
		log.Error("Gateway.AddIssue exec error",
			zap.Int64("sprintID", sprintID),
			zap.Int64("issueID", issueID),
			zap.Error(err),
		)
		return 0, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		log.Error("Gateway.AddIssue affected error",
			zap.Int64("sprintID", sprintID),
			zap.Int64("issueID", issueID),
			zap.Error(err),
		)
		return 0, err
	}

	if affected == 0 {
		log.Error("Gateway.AddIssue no rows affected",
			zap.Int64("sprintID", sprintID),
			zap.Int64("issueID", issueID),
			zap.Error(err),
		)
		return 0, fmt.Errorf("no rows affected")
	}

	return sprintID, nil
}

func (g *gateway) RemoveIssue(sprintID, issueID int64) (int64, error) {
	query := "update sprint set issues = array_remove(issues, $2) where id = $1"

	res, err := g.db.Exec(query, sprintID, issueID)
	if err != nil {
		log.Error("Gateway.RemoveIssue exec error",
			zap.Int64("sprintID", sprintID),
			zap.Int64("issueID", issueID),
			zap.Error(err),
		)
		return 0, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		log.Error("Gateway.RemoveIssue affected error",
			zap.Int64("sprintID", sprintID),
			zap.Int64("issueID", issueID),
			zap.Error(err),
		)
		return 0, err
	}

	if affected == 0 {
		log.Error("Gateway.RemoveIssue no rows affected",
			zap.Int64("sprintID", sprintID),
			zap.Int64("issueID", issueID),
			zap.Error(err),
		)
		return 0, fmt.Errorf("no rows affected")
	}

	return sprintID, nil
}

func (g *gateway) Get(filter *models.SprintFilter) ([]*models.SprintRow, error) {
	query := g.builder.Select(columns...).
		From(table)

	if filter != nil {
		query = filter.Apply(query)
	}

	result := make([]*models.SprintRow, 0)

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
		var row models.SprintRow
		err = rows.Scan(
			&row.ID,
			&row.Name,
			&row.Project,
			&row.StartedAt,
			&row.FinishedAt,
			&row.Status,
			&row.InBacklog,
			&row.InProgress,
			&row.InDone,
			pq.Array(&row.Issues),
		)
		if err != nil {
			log.Error("Gateway.Get scan error",
				zap.Any("filter", filter),
				zap.Error(err),
			)
			return nil, err
		}
		result = append(result, &row)
	}

	return result, nil
}

func (g *gateway) GetByID(id int64) (*models.SprintRow, error) {
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

	var row models.SprintRow
	err = g.db.QueryRow(query, args...).Scan(
		&row.ID,
		&row.Name,
		&row.Project,
		&row.StartedAt,
		&row.FinishedAt,
		&row.Status,
		&row.InBacklog,
		&row.InProgress,
		&row.InDone,
		pq.Array(&row.Issues),
	)
	if err != nil {
		log.Error("Gateway.GetByID scan error",
			zap.Int64("id", id),
			zap.Error(err),
		)
		return nil, err
	}

	return &row, nil
}

func (g *gateway) Update(row *models.SprintRow) (int64, error) {
	query := g.builder.Update(table).
		Where(sq.Eq{"id": row.ID})

	if len(row.Name) > 0 {
		query = query.Set("name", row.Name)
	}

	if row.Project != 0 {
		query = query.Set("project", row.Project)
	}

	if row.StartedAt.Unix() != 0 {
		query = query.Set("started_at", row.StartedAt)
	}

	if row.FinishedAt.Unix() != 0 {
		query = query.Set("finished_at", row.FinishedAt)
	}

	if row.Status != 0 {
		query = query.Set("status", row.Status)
	}

	if len(row.Issues) > 0 {
		query = query.Set("issues", pq.Array(row.Issues))
	}

	query = query.
		Set("in_backlog", row.InBacklog).
		Set("in_progress", row.InProgress).
		Set("in_done", row.InDone)

	stmt, args, err := query.ToSql()
	if err != nil {
		log.Error("Gateway.Update query error",
			zap.Any("sprint", row),
			zap.Error(err),
		)
		return 0, err
	}

	res, err := g.db.Exec(stmt, args...)
	if err != nil {
		log.Error("Gateway.Update exec error",
			zap.Any("sprint", row),
			zap.Error(err),
		)
		return 0, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		log.Error("Gateway.Update affected error",
			zap.Any("sprint", row),
			zap.Error(err),
		)
		return 0, err
	}

	if affected == 0 {
		log.Error("Gateway.Update no rows affected",
			zap.Any("sprint", row),
			zap.Error(err),
		)
		return 0, fmt.Errorf("no rows affected")
	}

	return row.ID, nil
}

func (g *gateway) ChangeStatus(id int64, status models.Status) (int64, error) {
	query := g.builder.Update(table).
		Where(sq.Eq{"id": id}).
		Set("status", status)

	stmt, args, err := query.ToSql()
	if err != nil {
		log.Error("Gateway.ChangeStatus query error",
			zap.Any("id", id),
			zap.Error(err),
		)
		return 0, err
	}

	res, err := g.db.Exec(stmt, args...)
	if err != nil {
		log.Error("Gateway.ChangeStatus exec error",
			zap.Any("id", id),
			zap.Error(err),
		)
		return 0, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		log.Error("Gateway.ChangeStatus affected error",
			zap.Any("id", id),
			zap.Error(err),
		)
		return 0, err
	}

	if affected == 0 {
		log.Error("Gateway.ChangeStatus no rows affected",
			zap.Any("id", id),
			zap.Error(err),
		)
		return 0, fmt.Errorf("no rows affected")
	}

	return id, nil
}
