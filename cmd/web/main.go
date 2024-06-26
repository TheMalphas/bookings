package main

import (
	"fmt"
	"github.com/TheMalphas/bookings/pkg/config"
	"github.com/TheMalphas/bookings/pkg/handlers"
	"github.com/TheMalphas/bookings/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const PortNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {
	app.InProduction = false

	// variable scoping
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Listening on port %s\n", PortNumber)
	srv := &http.Server{
		Addr:    PortNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
