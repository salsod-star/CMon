package models

import "time"

type Contribution struct {
	ID          uint64    `json:"id"`
	Amount      int       `json:"amount"`
	Outstanding int       `json:"outstanding"`
	LastPaid    time.Time `json:"last_paid"`
	Frequency   string    `json:"frequency"`
	Status      string    `json:"status"`
	Notes       string    `json:"notes"`
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
