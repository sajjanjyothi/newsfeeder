package redishelper

import (
	"github.com/gomodule/redigo/redis"

	"github.com/sajjanjyothi/ziglunews/envloader"
	log "github.com/sirupsen/logrus"
)

type RedisHelper struct {
	connection redis.Conn
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
