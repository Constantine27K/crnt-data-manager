package helper

import (
	"database/sql"
)

type Helper interface {
	FlushAllTables() error
}

type helper struct {
	db *sql.DB
}

func NewDBHelper(db *sql.DB) Helper {
	return &helper{
		db: db,
	}
}

func (h *helper) FlushAllTables() error {
	_, err := h.db.Exec("truncate table issue")
	if err != nil {
		return err
	}
	_, err = h.db.Exec("truncate table project")
	if err != nil {
		return err
	}
	_, err = h.db.Exec("truncate table sprint")
	if err != nil {
		return err
	}
	_, err = h.db.Exec("truncate table team")
	return err
}
