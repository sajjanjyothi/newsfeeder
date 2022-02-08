//Responsible for handling newsfeeder server interface
package newsservice

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/sajjanjyothi/ziglunews/newsfetcher"
	log "github.com/sirupsen/logrus"
)

type NewsService struct {
}

func (service *NewsService) Getnewsbycategory(ctx echo.Context, category string) error {
	newsfetcher := newsfetcher.New()
	news, err := newsfetcher.GetFeed("https://feeds.skynews.com/feeds/rss/technology.xml")
	if err != nil {
		log.Error(err)
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, news)
}
