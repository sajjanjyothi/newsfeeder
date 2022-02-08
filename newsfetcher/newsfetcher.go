//This will fetch news from newfeed
package newsfetcher

import (
	"github.com/mmcdole/gofeed"
	log "github.com/sirupsen/logrus"
)

//testable interface
type NewsFatcher interface {
	GetFeeds(URL string) (string, error)
}
type newsfetcher struct {
}

func New() *newsfetcher {
	return &newsfetcher{}
}

func (fetcher *newsfetcher) GetFeed(URL string) ([]*gofeed.Item, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(URL)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return feed.Items, nil
}
