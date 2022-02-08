//Responsible for handling newsfeeder server interface
package newsservice

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/sajjanjyothi/ziglunews/envloader"
	"github.com/sajjanjyothi/ziglunews/newsfeeder"
	"github.com/sajjanjyothi/ziglunews/newsfetcher"
	"github.com/sajjanjyothi/ziglunews/redishelper"
	log "github.com/sirupsen/logrus"
)

type NewsService struct {
	env         *envloader.ENV
	newsfetcher newsfetcher.NewsFetcher
	redis       redishelper.RedisHelperImpl
}

func New(env *envloader.ENV) *NewsService {
	return &NewsService{
		env:         env,
		newsfetcher: newsfetcher.New(env),
		redis:       redishelper.New(env),
	}
}

func (service *NewsService) Getnewsbycategory(ctx echo.Context, category string) error {
	newsfetcher := newsfetcher.New(service.env)
	url, err := service.redis.GetValue(category)
	if err != nil {
		log.Error(err)
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	news, err := newsfetcher.GetFeed(url)
	if err != nil {
		log.Error(err)
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, news)
}

func (service *NewsService) Updatenewsurl(ctx echo.Context) error {
	updateBody := &newsfeeder.UpdatenewsurlJSONBody{}
	if err := ctx.Bind(updateBody); err != nil {
		log.Error(err)
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	if err := service.redis.SetValue(*updateBody.Category, *updateBody.Url); err != nil {
		log.Error(err)
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.String(http.StatusOK, "updated")
}

func (service *NewsService) Geturls(ctx echo.Context) error {
	all, err := service.redis.GetAll()
	if err != nil {
		log.Error(err)
		return ctx.String(http.StatusNotFound, err.Error())
	}

	return ctx.JSON(http.StatusOK, all)
}
