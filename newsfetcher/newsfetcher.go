//This will fetch news from newfeed
package newsfetcher

import (
	"github.com/mmcdole/gofeed"
	"github.com/sajjanjyothi/ziglunews/envloader"
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

func (fetcher *newsfetcher) GetFeed(url string) ([]*gofeed.Item, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return feed.Items, nil
}
