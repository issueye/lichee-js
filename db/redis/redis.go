package redis

import (
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	redis "github.com/redis/go-redis/v9"
)

// RegisterRedis
// 由外部传入redis 客户端
func InitRedis() {
	require.RegisterNativeModule("db/redis", func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)

		o.Set("newClient", func(addr string, pwd string, dbNum int) goja.Value {
			rdb := redis.NewClient(&redis.Options{
				Addr:     addr,
				Password: pwd,
				DB:       dbNum,
			})

			return NewRedisClient(runtime, rdb)
		})
	})
}

// register native
func RegisterNativeRedis(rt *goja.Runtime, moduleName string, rdb *redis.Client) {
	rt.Set(moduleName, NewRedisClient(rt, rdb))
}
