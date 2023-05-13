package boltdb

import (
	"github.com/dop251/goja"
	js "github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"github.com/issueye/lichee-js/lib"
	bolt "go.etcd.io/bbolt"
)

func InitBolt() {
	require.RegisterNativeModule("db/bolt", func(runtime *js.Runtime, module *js.Object) {
		o := module.Get("exports").(*js.Object)

		// 打开数据库
		o.Set("open", func(call js.FunctionCall) js.Value {
			path := call.Argument(0).String()
			db, err := bolt.Open(path, 0666, nil)
			if err != nil {
				return lib.MakeErrorValue(runtime, err)
			}

			return NewBoltDb(runtime, db)
		})

	})
}

func RegisterNativeBolt(rt *goja.Runtime, name string, db *bolt.DB) {
	rt.Set(name, NewBoltDb(rt, db))
}
