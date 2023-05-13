package redistest

import (
	"testing"

	licheejs "github.com/issueye/lichee-js"
	"github.com/redis/go-redis/v9"
)

func Test_Redis(t *testing.T) {

	c := licheejs.NewCore()
	c.SetLogOutMode(licheejs.LOM_DEBUG)

	t.Run("redis_mod", func(t *testing.T) {
		err := c.Run("redis_mod", "redis_mod.js")
		if err != nil {
			t.Errorf("运行脚本失败，失败原因：%s", err.Error())
		}
	})

	rdb := redis.NewClient(&redis.Options{
		Addr:     "49.235.124.25:6379",
		Password: "123456",
		DB:       0,
	})

	licheejs.RegisterRedis(c.GetRts(), "licheeRedis", rdb)

	t.Run("redis_native", func(t *testing.T) {
		err := c.Run("redis_native", "redis_native.js")
		if err != nil {
			t.Errorf("运行脚本失败，失败原因：%s", err.Error())
		}
	})

}
