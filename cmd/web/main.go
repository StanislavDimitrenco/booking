package main

import (
	"fmt"
	"github.com/StanislavDimitrenko/booking/pkg/config"
	"github.com/StanislavDimitrenko/booking/pkg/handler"
	"github.com/StanislavDimitrenko/booking/pkg/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {

	//change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	//get the template cache from the app config
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Can't get templates")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handler.NewRepo(&app)
	handler.NewHandler(repo)

	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    config.PORT,
		Handler: routs(&app),
	}

	fmt.Println(fmt.Sprintf("Staring application on port http://localhost%s/", config.PORT))
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
