package redistest

import (
	"testing"

	licheejs "github.com/issueye/lichee-js"
	"github.com/redis/go-redis/v9"
)

func Test_Redis(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	licheejs.RegisterRedis("db/redis", rdb)
	c := licheejs.NewCore()

	t.Run("redis_js", func(t *testing.T) {
		err := c.Run("redis_test", "redis_test.js")
		if err != nil {
			t.Errorf("运行脚本失败，失败原因：%s", err.Error())
		}
	})
}
