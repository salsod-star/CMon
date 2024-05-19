package models

import (
	"database/sql"
	"errors"
	"time"

	"cmon.salsod.dev/internal/validator"
)

type Contribution struct {
	ID                int64     `json:"id"`
	TotalAmount       float64   `json:"total_amount"`
	CurrentPaidAmount float64   `json:"current_paid_amount"`
	Outstanding       float64   `json:"outstanding"`
	Status            string    `json:"status"`
	Interval          string    `json:"interval"`
	TimePaid          time.Time `json:"time_paid"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type ContributionModel struct {
	DB *sql.DB
}

func (cm ContributionModel) Insert(contribution *Contribution) error {
	query := `INSERT INTO contributions (total_amount, current_paid_amount, outstanding, status, interval)
			VALUES ($1, $2, $3, $4, $5) RETURNING id, status, time_paid`

	args := []interface{}{contribution.TotalAmount, contribution.CurrentPaidAmount, contribution.Outstanding, contribution.Status, contribution.Interval}

	return cm.DB.QueryRow(query, args...).Scan(&contribution.ID, &contribution.Status, &contribution.TimePaid)
}

func (cm ContributionModel) Get(id int64) (*Contribution, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `SELECT id, total_amount, current_paid_amount, outstanding, status, interval, time_paid, updated_at
			FROM contributions WHERE id=$1`

	var contribution Contribution

	err := cm.DB.QueryRow(query, id).Scan(
		&contribution.ID,
		&contribution.TotalAmount,
		&contribution.CurrentPaidAmount,
		&contribution.Outstanding,
		&contribution.Status,
		&contribution.Interval,
		&contribution.TimePaid,
		&contribution.UpdatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &contribution, nil
}

func (cm ContributionModel) Update(contribution *Contribution) error {
	query := `
	UPDATE contributions
	SET total_amount=$1, current_paid_amount=$2, outstanding= total_amount - $2, status=$3, interval=$4 WHERE id=$5
	RETURNING status, outstanding`

	args := []interface{}{
		contribution.TotalAmount,
		contribution.CurrentPaidAmount,
		contribution.Status,
		contribution.Interval,
		contribution.ID,
	}

	return cm.DB.QueryRow(query, args...).Scan(&contribution.Status, &contribution.Outstanding)
}

func (cm ContributionModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
			DELETE from contributions
			WHERE id=$1`

	result, err := cm.DB.Exec(query, id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

// Contribution Table

// ContributionID (Primary Key)
// UserID (Foreign Key referencing UserID)
// GroupID (Foreign Key referencing GroupID)
// Amount
// ContributionDate
// Frequency (e.g., weekly, monthly, daily)
// Status (e.g., paid, pending, cancelled)
// PaymentReferenceNumber
// LastUpdateDate
// Notes

func ValidateContribution(v *validator.Validator, c *Contribution) {

	v.Check(c.TotalAmount >= 0, "amount", "must be a positive integer")
	v.Check(c.TimePaid.Before(time.Now()), "lastPaid", "")
}
