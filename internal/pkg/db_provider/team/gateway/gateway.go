package gateway

import (
	"database/sql"
	"fmt"

	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/team/models"
	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

type TeamGateway interface {
	Add(team *models.TeamRow) (int64, error)
	Get(filter *models.TeamFilter) ([]*models.TeamRow, error)
	GetByID(id int64) (*models.TeamRow, error)
	Update(team *models.TeamRow) (int64, error)
}

type gateway struct {
	db      *sql.DB
	builder sq.StatementBuilderType
}

func NewTeamGateway(db *sql.DB) TeamGateway {
	return &gateway{
		db:      db,
		builder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

const (
	table = "team"
)

var (
	columns = []string{"id", "name", "members", "projects", "tech_owner", "business_owner", "department"}
)

func (g *gateway) Add(teamRow *models.TeamRow) (int64, error) {
	values := []interface{}{
		teamRow.Name, pq.Array(teamRow.Members), pq.Array(teamRow.Projects), teamRow.TechOwner, teamRow.BusinessOwner, teamRow.Department,
	}

	query, args, err := g.builder.Insert(table).
		Columns(columns[1:]...).
		Values(values...).
		Suffix("returning id").ToSql()
	if err != nil {
		log.Error("Gateway.Add query error",
			zap.Any("team", teamRow),
			zap.Error(err),
		)
		return 0, err
	}

	var id int64
	err = g.db.QueryRow(query, args...).Scan(&id)
	if err != nil {
		log.Error("Gateway.Add scan error",
			zap.Any("team", teamRow),
			zap.Error(err),
		)
		return 0, err
	}

	return id, nil
}

func (g *gateway) Get(filter *models.TeamFilter) ([]*models.TeamRow, error) {
	query := g.builder.Select(columns...).
		From(table)

	if filter != nil {
		query = filter.Apply(query)
	}

	result := make([]*models.TeamRow, 0)

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
		var teamRow models.TeamRow
		err = rows.Scan(
			&teamRow.ID,
			&teamRow.Name,
			pq.Array(&teamRow.Members),
			pq.Array(&teamRow.Projects),
			&teamRow.TechOwner,
			&teamRow.BusinessOwner,
			&teamRow.Department,
		)
		if err != nil {
			log.Error("Gateway.Get scan error",
				zap.Any("filter", filter),
				zap.Error(err),
			)
			return nil, err
		}
		result = append(result, &teamRow)
	}

	return result, nil
}

func (g *gateway) GetByID(id int64) (*models.TeamRow, error) {
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

	var teamRow models.TeamRow
	err = g.db.QueryRow(query, args...).Scan(
		&teamRow.ID,
		&teamRow.Name,
		pq.Array(&teamRow.Members),
		pq.Array(&teamRow.Projects),
		&teamRow.TechOwner,
		&teamRow.BusinessOwner,
		&teamRow.Department,
	)
	if err != nil {
		log.Error("Gateway.GetByID scan error",
			zap.Int64("id", id),
			zap.Error(err),
		)
		return nil, err
	}

	return &teamRow, nil
}

func (g *gateway) Update(teamRow *models.TeamRow) (int64, error) {
	query := g.builder.Update(table).
		Where(sq.Eq{"id": teamRow.ID})

	if len(teamRow.Name) > 0 {
		query = query.Set("name", teamRow.Name)
	}

	if len(teamRow.Members) > 0 {
		query = query.Set("members", pq.Array(teamRow.Members))
	}

	if len(teamRow.Projects) > 0 {
		query = query.Set("projects", pq.Array(teamRow.Projects))
	}

	if len(teamRow.TechOwner) > 0 {
		query = query.Set("tech_owner", teamRow.TechOwner)
	}

	if len(teamRow.BusinessOwner) > 0 {
		query = query.Set("business_owner", teamRow.BusinessOwner)
	}

	if len(teamRow.Department) > 0 {
		query = query.Set("department", teamRow.Department)
	}

	stmt, args, err := query.ToSql()
	if err != nil {
		log.Error("Gateway.Update query error",
			zap.Any("team", teamRow),
			zap.Error(err),
		)
		return 0, err
	}

	res, err := g.db.Exec(stmt, args...)
	if err != nil {
		log.Error("Gateway.Update exec error",
			zap.Any("team", teamRow),
			zap.Error(err),
		)
		return 0, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		log.Error("Gateway.Update affected error",
			zap.Any("team", teamRow),
			zap.Error(err),
		)
		return 0, err
	}

	if affected == 0 {
		log.Error("Gateway.Update no rows affected",
			zap.Any("team", teamRow),
			zap.Error(err),
		)
		return 0, fmt.Errorf("no rows affected")
	}

	return teamRow.ID, nil
}
