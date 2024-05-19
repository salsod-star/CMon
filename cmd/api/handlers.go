package main

import (
	"errors"
	"fmt"
	"net/http"

	"cmon.salsod.dev/internal/models"
	"cmon.salsod.dev/internal/validator"
)

func (app *application) createContributionHandler(w http.ResponseWriter, r *http.Request) {
	var newContribution struct {
		TotalAmount       float64 `json:"total_amount"`
		CurrentPaidAmount float64 `json:"current_paid_amount"`
		Outstanding       float64 `json:"outstanding"`
		Status            string  `json:"status"`
		Interval          string  `json:"interval"`
	}

	err := app.readJSON(w, r, &newContribution)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	contrib := &models.Contribution{
		TotalAmount:       newContribution.TotalAmount,
		CurrentPaidAmount: newContribution.CurrentPaidAmount,
		Outstanding:       newContribution.Outstanding,
		Status:            newContribution.Status,
		Interval:          newContribution.Interval,
	}

	v := validator.New()

	models.ValidateContribution(v, contrib)

	if !v.Valid() {
		app.failedValidationErrorResponse(w, r, v.Errors)
		return
	}

	err = app.models.Contributions.Insert(contrib)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)

	headers.Set("Location", fmt.Sprintf("/api/v1/contributions/%d", contrib.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"contribution": contrib}, headers)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showContributionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	contribution, err := app.models.Contributions.Get(id)

	if err != nil {
		switch {
		case errors.Is(err, models.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"contribution": contribution}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) updateContributionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	contribution, err := app.models.Contributions.Get(id)

	if err != nil {
		switch {
		case errors.Is(err, models.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var newContribution struct {
		TotalAmount       float64 `json:"total_amount"`
		CurrentPaidAmount float64 `json:"current_paid_amount"`
		Status            string  `json:"status"`
		Interval          string  `json:"interval"`
	}

	err = app.readJSON(w, r, &newContribution)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	contribution.TotalAmount = newContribution.TotalAmount
	contribution.CurrentPaidAmount = newContribution.CurrentPaidAmount
	contribution.Status = newContribution.Status
	contribution.Interval = newContribution.Interval

	v := validator.New()

	models.ValidateContribution(v, contribution)

	if !v.Valid() {
		app.failedValidationErrorResponse(w, r, v.Errors)
		return
	}

	err = app.models.Contributions.Update(contribution)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"contribution": contribution}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteContributionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Contributions.Delete(id)

	if err != nil {
		switch {
		case errors.Is(err, models.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "contribution successfully deleted"}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
