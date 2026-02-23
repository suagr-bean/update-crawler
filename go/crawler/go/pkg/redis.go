package pkg

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)
var RDB *redis.Client
func RedisInit(addr,password string,db int) error{
	RDB:=redis.NewClient(&redis.Options{
		Addr:addr,
		Password: password,
        DB:db,
		PoolSize:10,
	})
   ctx,channel:= context.WithTimeout(context.Background(),5*time.Second)
   defer channel()
   return RDB.Ping(ctx).Err()
}