package main

import (
	"fmt"
	"net/http"
)

func (app *application) createContributionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create a new contribution")
}

func (app *application) showContributionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil || id < 1 {
		app.logger.Println(err)
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of the contribution for entry %d", id)
}
