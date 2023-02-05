package models

type TeamRow struct {
	ID            int64    `db:"id"`
	Name          string   `db:"name"`
	Members       []string `db:"members"`
	Projects      []int64  `db:"projects"`
	TechOwner     string   `db:"tech_owner"`
	BusinessOwner string   `db:"business_owner"`
	Department    string   `db:"department"`
}
