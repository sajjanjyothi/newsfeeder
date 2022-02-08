package newsfetcher

import (
	"testing"

	"github.com/sajjanjyothi/ziglunews/envloader"
	"github.com/stretchr/testify/assert"
)

func Test_GetFeed(t *testing.T) {
	env := &envloader.ENV{}
	envloader.LoadConfig(env)
	news := New(env)
	_, err := news.GetFeed("https://feeds.skynews.com/feeds/rss/uk.xml")
	assert.NoError(t, err, "Getfeed failed")
}
