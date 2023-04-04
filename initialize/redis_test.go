package initialize

import (
	"context"
	"github.com/go-redis/redis/v8"
	"testing"
)

type loginState struct {
	Ip    string
	Token string
}

type loginStateMap map[string]loginState

func TestRedisMap(t *testing.T) {
	var Rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	Rdb.MSet(context.Background(), loginStateMap{
		"root": loginState{
			Ip:    "12314414",
			Token: "adgagreg",
		},
	})
	var val loginState
	err := Rdb.MGet(context.Background(), "root").Scan(&val)
	if err != nil {
		return
	}
}
