package models

type ProjectRow struct {
	ID               int64   `db:"id"`
	Name             string  `db:"name"`
	ShortName        string  `db:"short_name"`
	IsArchived       bool    `db:"is_archived"`
	ResponsibleTeams []int64 `db:"responsible_teams"`
	Description      string  `db:"description"`
	Department       int64   `db:"department"`
	Responsible      string  `db:"responsible"`
}
