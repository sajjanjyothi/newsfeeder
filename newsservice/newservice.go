//Responsible for handling newsfeeder server interface
package newsservice

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/sajjanjyothi/ziglunews/envloader"
	"github.com/sajjanjyothi/ziglunews/newsfetcher"
	log "github.com/sirupsen/logrus"
)

type NewsService struct {
	env *envloader.ENV
}

func New(env *envloader.ENV) *NewsService {
	return &NewsService{
		env: env,
	}
}

func (service *NewsService) Getnewsbycategory(ctx echo.Context, category string) error {
	newsfetcher := newsfetcher.New(service.env)
	news, err := newsfetcher.GetFeed(category)
	if err != nil {
		log.Error(err)
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, news)
}
