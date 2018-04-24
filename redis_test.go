package testinfra_test

import (
	"testing"
	"time"

	"bitbucket.org/qiscus/testinfra"
	"github.com/go-redis/redis"
)

func TestRedis(t *testing.T) {
	addr, close := testinfra.Redis("4")
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	if client == nil {
		t.Fatal("got nil client")
	}

	if err := client.Set("key", "value", time.Hour).Err(); err != nil {
		t.Error("error while setting value")
	}
	val, err := client.Get("key").Result()
	if err != nil {
		t.Error("error while getting value")
	}
	if val != "value" {
		t.Error("value missmatch")
	}

	close()

	if err := client.Ping().Err(); err == nil {
		t.Error("should get error client disconnected")
	}
}
