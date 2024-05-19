package models

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Contributions ContributionModel
}

func NewModels(db *sql.DB) Models {

	return Models{
		Contributions: ContributionModel{DB: db},
	}
}
