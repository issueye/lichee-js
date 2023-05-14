package jstest

import (
	"fmt"
	"testing"
	"time"

	licheejs "github.com/issueye/lichee-js"
	"github.com/redis/go-redis/v9"
)

func Test_JsTest(t *testing.T) {
	c := licheejs.NewCore()
	c.SetLogOutMode(licheejs.LOM_DEBUG)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "49.235.124.25:6379",
		Password: "123456",
		DB:       0,
	})

	licheejs.RegisterRedis(c.GetRts(), "licheeRedis", rdb)

	t.Run("go-call-js", func(t *testing.T) {
		err := c.Run("test-js", "test-js.js")
		if err != nil {
			t.Errorf("运行脚本失败，失败原因：%s", err.Error())
		}

		var fn func(string) string
		err = c.ExportFunc("testGoCallJs", &fn)
		if err != nil {
			t.Errorf("导出js方法失败，失败原因：%s", err.Error())
		}

		s := fn("hello---")
		fmt.Println(s)
	})

	t.Run("run-time", func(t *testing.T) {
		t1 := time.Now()

		err := c.Run("run-time", "run-time.js")
		if err != nil {
			t.Error(err)
		}

		t2 := time.Now()
		d := t2.Sub(t1)
		fmt.Println("运行时间：", d)
	})
}
