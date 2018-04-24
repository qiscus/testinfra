package testinfra

import (
	"log"

	"github.com/go-redis/redis"
)

func Redis(version string) (string, func()) {
	if version == "" {
		version = "latest"
	}

	resource, err := StartContainer("redis", version, nil)
	if err != nil {
		log.Panicln("can't start redis, err:", err)
	}

	addr := GetIPPort(resource, "6379/tcp")
	err = waitContainer(func() error {
		return RedisCheckFunc(addr)
	})
	if err != nil {
		log.Panicln("failed to wait redis to be ready, err:", err)
	}

	return addr, func() {
		pool.Purge(resource)
	}
}

func RedisCheckFunc(addr string) error {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return client.Ping().Err()
}
