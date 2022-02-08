package main

import (
	echo "github.com/labstack/echo/v4"
	"github.com/sajjanjyothi/ziglunews/envloader"
	"github.com/sajjanjyothi/ziglunews/newsfeeder"
	"github.com/sajjanjyothi/ziglunews/newsservice"
	log "github.com/sirupsen/logrus"
)

func main() {

	env := envloader.ENV{}

	err := envloader.LoadConfig(&env)
	if err != nil {
		panic(err)
	}

	newsService := &newsservice.NewsService{}
	e := echo.New()

	newsfeeder.RegisterHandlers(e, newsService)
	log.Info("Web server starting")
	log.Fatal(e.Start(":" + env.ServicePort))
}
