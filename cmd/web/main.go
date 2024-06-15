package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/motoshi-suzuki-hp/go-cource/pkg/config"
	"github.com/motoshi-suzuki-hp/go-cource/pkg/handlers"
	"github.com/motoshi-suzuki-hp/go-cource/pkg/render"
)

const portnumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Starting application on port %s\n", portnumber)

	srv := &http.Server {
		Addr: portnumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}