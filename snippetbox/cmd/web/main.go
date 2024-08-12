package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	/* Define a new command-line flag with the name 'addr', a default value of ":4000"
	Esto lo q hace es definir el port via linea de comandos para no tener que hardcodearlo a mano,
	en caso de no usar la bandera se iniciará automáticamente en el port :4000*/
	addr := flag.String("addr", ":4000", "HTTP Network Address")
	flag.Parse()

	/* Start a new router (servemux).
	Hacerlo de esta manera (y no vía DefaultServeMux) para evitar handlers maliciosos de third-party packages */
	mux := http.NewServeMux()

	// Create a static file server for delivering static files like imgs or css
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Usar la Handle func para que responda a todos los paths q empiecen con /static/
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Handlers
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Start a new web server
	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
