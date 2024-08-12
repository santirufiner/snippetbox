package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	/* Start a new router (servemux).
	Hacerlo de esta manera (y no vía DefaultServeMux) para evitar handlers maliciosos de third-party packages */
	mux := http.NewServeMux()

	// Create a static file server for delivering static files like imgs or css
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Usar la Handle func para que responda a todos los paths q empiecen con /static/
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Handlers
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return mux
}
