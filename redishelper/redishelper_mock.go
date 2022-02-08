package redishelper

import (
	"github.com/stretchr/testify/mock"
)

type RedisMock struct {
	mock.Mock
}

func NewMock() *RedisMock {

	redismock := &RedisMock{}

	redismock.On("SetValue", "uk", "").Return(nil)
	redismock.On("GetValue", "uk").Return("http://feeds.bbci.co.uk/news/uk/rss.xml", nil)
	return redismock
}

func (redishelper *RedisMock) SetValue(key string, value string) error {
	args := redishelper.Called(key, value)
	return args.Error(0)
}

func (redishelper *RedisMock) GetValue(key string) (string, error) {
	args := redishelper.Called(key)
	return args.String(0), args.Error(1)
}

func (redishelper *RedisMock) Close() error {
	args := redishelper.Called()
	return args.Error(0)
}

func (redishelper *RedisMock) GetAll() ([]URLS, error) {
	urls := []URLS{
		{
			Category: "uk",
			URL:      "https://www.google.com/url?sa=D&q=http://feeds.skynews.com/feeds/rss/uk.xml",
		},
		{
			Category: "technology",
			URL:      "https://feeds.skynews.com/feeds/rss/technology.xml",
		},
	}

	return urls, nil
}
