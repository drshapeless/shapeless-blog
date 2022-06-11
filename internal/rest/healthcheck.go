package rest

import "net/http"

func (app *Application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"status":  "available",
		"version": app.Version,
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
