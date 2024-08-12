package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	/* Define a new command-line flag with the name 'addr', a default value of ":4000"
	Esto lo q hace es definir el port via linea de comandos para no tener que hardcodearlo a mano,
	en caso de no usar la bandera se iniciará automáticamente en el port :4000*/
	addr := flag.String("addr", ":4000", "HTTP Network Address")
	flag.Parse()

	// Logger de info y error
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// Start a new web server
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr) // Information message
	err := srv.ListenAndServe()
	errorLog.Fatal(err) // Error message
}
