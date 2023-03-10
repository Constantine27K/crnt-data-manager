package gateway

import (
	"database/sql"
	"fmt"

	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/project/models"
	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

type ProjectGateway interface {
	Add(row *models.ProjectRow) (int64, error)
	AddResponsibleTeam(projectID, teamID int64) (int64, error)
	Get(filter *models.ProjectFilter) ([]*models.ProjectRow, error)
	GetByID(id int64) (*models.ProjectRow, error)
	GetShortName(id int64) (string, error)
	Update(row *models.ProjectRow) (int64, error)
}

type gateway struct {
	db      *sql.DB
	builder sq.StatementBuilderType
}

func NewProjectGateway(db *sql.DB) ProjectGateway {
	return &gateway{
		db:      db,
		builder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

const (
	table = "project"
)

var (
	columns = []string{"id", "name", "short_name", "is_archived", "responsible_teams"}
)

func (g *gateway) Add(row *models.ProjectRow) (int64, error) {
	values := []interface{}{
		row.Name, row.ShortName, row.IsArchived, pq.Array(row.ResponsibleTeams),
	}

	query, args, err := g.builder.Insert(table).
		Columns(columns[1:]...).
		Values(values...).
		Suffix("returning id").ToSql()
	if err != nil {
		log.Error("Gateway.Add query error",
			zap.Any("project", row),
			zap.Error(err),
		)
		return 0, err
	}

	var id int64
	err = g.db.QueryRow(query, args...).Scan(&id)
	if err != nil {
		log.Error("Gateway.Add scan error",
			zap.Any("project", row),
			zap.Error(err),
		)
		return 0, err
	}

	return id, nil
}

func (g *gateway) AddResponsibleTeam(projectID, teamID int64) (int64, error) {
	query := "update project set responsible_teams = array_append(responsible_teams, $2) where id = $1"

	res, err := g.db.Exec(query, projectID, teamID)
	if err != nil {
		log.Error("Gateway.AddResponsibleTeam exec error",
			zap.Int64("projectID", projectID),
			zap.Int64("teamID", teamID),
			zap.Error(err),
		)
		return 0, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		log.Error("Gateway.AddResponsibleTeam affected error",
			zap.Int64("projectID", projectID),
			zap.Int64("teamID", teamID),
			zap.Error(err),
		)
		return 0, err
	}

	if affected == 0 {
		log.Error("Gateway.AddResponsibleTeam no rows affected",
			zap.Int64("projectID", projectID),
			zap.Int64("teamID", teamID),
			zap.Error(err),
		)
		return 0, fmt.Errorf("no rows affected")
	}

	return teamID, nil
}

func (g *gateway) Get(filter *models.ProjectFilter) ([]*models.ProjectRow, error) {
	query := g.builder.Select(columns...).
		From(table)

	if filter != nil {
		query = filter.Apply(query)
	}

	result := make([]*models.ProjectRow, 0)

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
		var row models.ProjectRow
		err = rows.Scan(
			&row.ID,
			&row.Name,
			&row.ShortName,
			&row.IsArchived,
			pq.Array(&row.ResponsibleTeams),
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

func (g *gateway) GetByID(id int64) (*models.ProjectRow, error) {
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

	var row models.ProjectRow
	err = g.db.QueryRow(query, args...).Scan(
		&row.ID,
		&row.Name,
		&row.ShortName,
		&row.IsArchived,
		pq.Array(&row.ResponsibleTeams),
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

func (g *gateway) GetShortName(id int64) (string, error) {
	query, args, err := g.builder.Select("short_name").
		From(table).
		Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		log.Error("Gateway.GetShortName query error",
			zap.Int64("id", id),
			zap.Error(err),
		)
		return "", err
	}

	var shortName string
	err = g.db.QueryRow(query, args...).Scan(&shortName)
	if err != nil {
		log.Error("Gateway.GetShortName scan error",
			zap.Int64("id", id),
			zap.Error(err),
		)
		return "", err
	}

	return shortName, nil
}

func (g *gateway) Update(row *models.ProjectRow) (int64, error) {
	query := g.builder.Update(table).
		Where(sq.Eq{"id": row.ID})

	if len(row.Name) > 0 {
		query = query.Set("name", row.Name)
	}

	if len(row.ShortName) > 0 {
		query = query.Set("short_name", row.ShortName)
	}

	if len(row.ResponsibleTeams) > 0 {
		query = query.Set("responsible_teams", pq.Array(row.ResponsibleTeams))
	}

	query = query.Set("is_archived", row.IsArchived)

	stmt, args, err := query.ToSql()
	if err != nil {
		log.Error("Gateway.Update query error",
			zap.Any("project", row),
			zap.Error(err),
		)
		return 0, err
	}

	res, err := g.db.Exec(stmt, args...)
	if err != nil {
		log.Error("Gateway.Update exec error",
			zap.Any("project", row),
			zap.Error(err),
		)
		return 0, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		log.Error("Gateway.Update affected error",
			zap.Any("project", row),
			zap.Error(err),
		)
		return 0, err
	}

	if affected == 0 {
		log.Error("Gateway.Update no rows affected",
			zap.Any("project", row),
			zap.Error(err),
		)
		return 0, fmt.Errorf("no rows affected")
	}

	return row.ID, nil
}
