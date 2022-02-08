package redishelper

import (
	"github.com/gomodule/redigo/redis"

	"github.com/sajjanjyothi/ziglunews/envloader"
	log "github.com/sirupsen/logrus"
)

type RedisHelper struct {
	connection redis.Conn
}

type URLS struct {
	Category string
	URL      string
}

func New(env *envloader.ENV) *RedisHelper {

	c, err := redis.Dial("tcp", env.RedisURL+":6379", redis.DialPassword(env.RedisPassword))
	if err != nil {
		log.Error(err)
		return nil
	}
	redisHelper := &RedisHelper{
		connection: c,
	}
	return redisHelper
}

func (redis *RedisHelper) SetValue(key string, value string) error {
	_, err := redis.connection.Do("SET", key, value)
	return err
}

func (redishelper *RedisHelper) GetValue(key string) (string, error) {
	value, err := redis.String(redishelper.connection.Do("GET", key))
	return value, err
}

func (redishelper *RedisHelper) Close() error {
	return redishelper.connection.Close()
}

func (redishelper *RedisHelper) GetAll() ([]URLS, error) {
	categories := []string{"uk", "technology"}
	urls := make([]URLS, 0)

	for _, category := range categories {
		url, _ := redishelper.GetValue(category)
		urls = append(urls, URLS{
			Category: category,
			URL:      url,
		})
	}
	return urls, nil
}
