package models

type DepartmentRow struct {
	ID       int64    `db:"id"`
	Name     string   `db:"name"`
	Projects []int64  `db:"projects"`
	Members  []string `db:"members"`
}
