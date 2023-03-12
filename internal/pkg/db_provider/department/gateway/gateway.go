package gateway

import (
	"database/sql"
	"fmt"

	"github.com/Constantine27K/crnt-data-manager/internal/pkg/db_provider/department/models"
	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

type DepartmentGateway interface {
	Add(department *models.DepartmentRow) (int64, error)
	AddProject(id int64, projectID int64) (int64, error)
	RemoveProject(id int64, projectID int64) (int64, error)
	Get(filter *models.DepartmentFilter) ([]*models.DepartmentRow, error)
	GetByID(id int64) (*models.DepartmentRow, error)
	Update(department *models.DepartmentRow) (int64, error)
}

type gateway struct {
	db      *sql.DB
	builder sq.StatementBuilderType
}

func NewDepartmentGateway(db *sql.DB) DepartmentGateway {
	return &gateway{
		db:      db,
		builder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

const (
	table = "department"
)

var (
	columns = []string{"id", "name", "projects"}
)

func (g *gateway) Add(department *models.DepartmentRow) (int64, error) {
	values := []interface{}{
		department.Name, pq.Array(department.Projects),
	}

	query, args, err := g.builder.Insert(table).
		Columns(columns[1:]...).
		Values(values...).
		Suffix("returning id").ToSql()
	if err != nil {
		log.Error("Gateway.Add query error",
			zap.Any("department", department),
			zap.Error(err),
		)
		return 0, err
	}

	var id int64
	err = g.db.QueryRow(query, args...).Scan(&id)
	if err != nil {
		log.Error("Gateway.Add scan error",
			zap.Any("department", department),
			zap.Error(err),
		)
		return 0, err
	}

	return id, nil
}

func (g *gateway) AddProject(id int64, projectID int64) (int64, error) {
	query := "update department set projects = array_append(projects, $2) where id = $1"

	res, err := g.db.Exec(query, id, projectID)
	if err != nil {
		log.Error("Gateway.AddProject exec error",
			zap.Int64("departmentID", id),
			zap.Int64("projectID", projectID),
			zap.Error(err),
		)
		return 0, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		log.Error("Gateway.AddProject affected error",
			zap.Int64("departmentID", id),
			zap.Int64("projectID", projectID),
			zap.Error(err),
		)
		return 0, err
	}

	if affected == 0 {
		log.Error("Gateway.AddProject no rows affected",
			zap.Int64("departmentID", id),
			zap.Int64("projectID", projectID),
			zap.Error(err),
		)
		return 0, fmt.Errorf("no rows affected")
	}

	return id, nil
}

func (g *gateway) RemoveProject(id int64, projectID int64) (int64, error) {
	query := "update department set projects = array_remove(projects, $2) where id = $1"

	res, err := g.db.Exec(query, id, projectID)
	if err != nil {
		log.Error("Gateway.AddProject exec error",
			zap.Int64("departmentID", id),
			zap.Int64("projectID", projectID),
			zap.Error(err),
		)
		return 0, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		log.Error("Gateway.AddProject affected error",
			zap.Int64("departmentID", id),
			zap.Int64("projectID", projectID),
			zap.Error(err),
		)
		return 0, err
	}

	if affected == 0 {
		log.Error("Gateway.AddProject no rows affected",
			zap.Int64("departmentID", id),
			zap.Int64("projectID", projectID),
			zap.Error(err),
		)
		return 0, fmt.Errorf("no rows affected")
	}

	return id, nil
}

func (g *gateway) Get(filter *models.DepartmentFilter) ([]*models.DepartmentRow, error) {
	query := g.builder.Select(columns...).
		From(table)

	if filter != nil {
		query = filter.Apply(query)
	}

	result := make([]*models.DepartmentRow, 0)

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
		var row models.DepartmentRow
		err = rows.Scan(
			&row.ID,
			&row.Name,
			pq.Array(&row.Projects),
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

func (g *gateway) GetByID(id int64) (*models.DepartmentRow, error) {
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

	var row models.DepartmentRow
	err = g.db.QueryRow(query, args...).Scan(
		&row.ID,
		&row.Name,
		pq.Array(&row.Projects),
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

func (g *gateway) Update(department *models.DepartmentRow) (int64, error) {
	query := g.builder.Update(table).
		Where(sq.Eq{"id": department.ID})

	if len(department.Name) > 0 {
		query = query.Set("name", department.Name)
	}

	if len(department.Projects) > 0 {
		query = query.Set("projects", pq.Array(department.Projects))
	}

	stmt, args, err := query.ToSql()
	if err != nil {
		log.Error("Gateway.Update query error",
			zap.Any("sprint", department),
			zap.Error(err),
		)
		return 0, err
	}

	res, err := g.db.Exec(stmt, args...)
	if err != nil {
		log.Error("Gateway.Update exec error",
			zap.Any("sprint", department),
			zap.Error(err),
		)
		return 0, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		log.Error("Gateway.Update affected error",
			zap.Any("sprint", department),
			zap.Error(err),
		)
		return 0, err
	}

	if affected == 0 {
		log.Error("Gateway.Update no rows affected",
			zap.Any("sprint", department),
			zap.Error(err),
		)
		return 0, fmt.Errorf("no rows affected")
	}

	return department.ID, nil
}
