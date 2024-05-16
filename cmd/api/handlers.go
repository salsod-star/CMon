package main

import (
	"fmt"
	"net/http"
	"time"

	"cmon.salsod.dev/internal/models"
	"cmon.salsod.dev/internal/validator"
)

func (app *application) createContributionHandler(w http.ResponseWriter, r *http.Request) {
	var newContribution struct {
		Amount      int       `json:"amount"`
		Outstanding int       `json:"outstanding"`
		LastPaid    time.Time `json:"last_paid"`
		Frequency   string    `json:"frequency"`
		Status      string    `json:"status"`
		Notes       string    `json:"notes"`
	}

	err := app.readJSON(w, r, &newContribution)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	contrib := &models.Contribution{
		Amount:      newContribution.Amount,
		Outstanding: newContribution.Outstanding,
		LastPaid:    newContribution.LastPaid,
		Frequency:   newContribution.Frequency,
		Status:      newContribution.Status,
		Notes:       newContribution.Notes,
	}

	v := validator.New()

	models.ValidateContribution(v, contrib)

	if !v.Valid() {
		app.failedValidationErrorResponse(w, r, err)
		return
	}

	fmt.Fprintf(w, "%+v", newContribution)
}

func (app *application) showContributionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	contribution := models.Contribution{
		ID:          uint64(id),
		Amount:      1400,
		Outstanding: 0,
		LastPaid:    time.Now(),
		Frequency:   "weekly",
		Status:      "paid",
		Notes:       "none",
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"contribution": contribution}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
