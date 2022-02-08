package main

import (
	"errors"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sajjanjyothi/ziglunews/envloader"
	"github.com/sajjanjyothi/ziglunews/newsfeeder"
	"github.com/sajjanjyothi/ziglunews/newsservice"
	"github.com/sajjanjyothi/ziglunews/redishelper"
	log "github.com/sirupsen/logrus"
)

const (
	TECHNOLOGY_URL = "http://feeds.bbci.co.uk/news/technology/rss.xml"
	UK_URL         = "http://feeds.bbci.co.uk/news/uk/rss.xml"
)

func setInitialURLs(env *envloader.ENV) error {
	redis := redishelper.New(env)
	if redis == nil {
		return errors.New("redis connection failed")
	}
	defer redis.Close()
	if err := redis.SetValue("uk", UK_URL); err != nil {
		return err
	}

	if err := redis.SetValue("technology", TECHNOLOGY_URL); err != nil {
		return err
	}
	return nil
}

func main() {

	env := envloader.ENV{}

	err := envloader.LoadConfig(&env)
	if err != nil {
		panic(err)
	}

	if err := setInitialURLs(&env); err != nil {
		panic(err)
	}
	newsService := newsservice.New(&env)
	e := echo.New()
	e.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == env.NewsFedderToken, nil
	}))
	newsfeeder.RegisterHandlers(e, newsService)
	log.Info("Web server starting")
	log.Fatal(e.Start(":" + env.ServicePort))
}
