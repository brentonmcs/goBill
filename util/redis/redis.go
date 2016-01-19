package redis

import r "gopkg.in/redis.v3"

var redisClient = r.NewClient(&r.Options{
	Addr: "localhost:6379"})

//NewRedis - returns a single redis client
func NewRedis() *r.Client {
	return redisClient
}
