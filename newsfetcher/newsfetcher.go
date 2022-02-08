//This will fetch news from newfeed
package newsfetcher

import (
	"github.com/mmcdole/gofeed"
	"github.com/sajjanjyothi/ziglunews/envloader"
	"github.com/sajjanjyothi/ziglunews/redishelper"
	log "github.com/sirupsen/logrus"
)

//testable interface
type NewsFetcher interface {
	GetFeed(URL string) ([]*gofeed.Item, error)
}
type newsfetcher struct {
	env *envloader.ENV
}

func New(env *envloader.ENV) *newsfetcher {
	return &newsfetcher{
		env: env,
	}
}

func (fetcher *newsfetcher) GetFeed(category string) ([]*gofeed.Item, error) {
	redis := redishelper.New(fetcher.env)
	defer redis.Close()
	URL, err := redis.GetValue(category)
	if err != nil {
		return nil, err
	}
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(URL)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return feed.Items, nil
}
