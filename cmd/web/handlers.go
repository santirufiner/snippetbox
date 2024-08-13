package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/santirufiner/snippetbox/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	// 404 error if not match home
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	snippets, err := app.snippets.Latest()

	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, snippet := range snippets {
		fmt.Fprintf(w, "%v\n", snippet)
	}

	/*

		// archivos html que le voy a pasar a la biblioteca template para que los renderice
		files := []string{
			"./ui/html/base.html",
			"./ui/html/pages/home.html",
			"./ui/html/partials/nav.html",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.serverError(w, err)
			return
		}

		err = ts.ExecuteTemplate(w, "base", nil)
		if err != nil {
			app.serverError(w, err)
			return
		}

	*/
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	// Write the snippet data as a plain-text HTTP response body.
	fmt.Fprintf(w, "%+v", snippet)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {

	// 405 error if method is not POST
	if r.Method != http.MethodPost {

		// permite ver en la respuesta los métodos permitidos
		w.Header().Set("Allow", http.MethodPost)

		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	// Create some variables holding dummy data. We'll remove these later on
	// during the build.
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Redirect a un view del ingreso del snippet
	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
