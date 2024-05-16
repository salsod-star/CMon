package main

import (
	"fmt"
	"net/http"
	"time"

	"cmon.salsod.dev/internal/models"
)

func (app *application) createContributionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create a new contribution")
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
