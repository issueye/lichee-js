package boltdb

import (
	js "github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"github.com/issueye/lichee-js/lib"
	bolt "go.etcd.io/bbolt"
)

var (
	Bdb *bolt.DB
)

type CallBackFunc = func(js.FunctionCall) js.Value

func InitBolt() {
	require.RegisterNativeModule("db/bolt", func(runtime *js.Runtime, module *js.Object) {
		o := module.Get("exports").(*js.Object)

		// 打开数据库
		o.Set("open", func(call js.FunctionCall) js.Value {
			path := call.Argument(0).String()
			d, err := bolt.Open(path, 0666, nil)
			if err != nil {
				return lib.MakeErrorValue(runtime, err)
			}
			Bdb = d
			return nil
		})

		o.Set("createBucket", func(call js.FunctionCall) js.Value {
			name := call.Argument(0).String()
			err := Bdb.Update(func(tx *bolt.Tx) error {
				_, err := tx.CreateBucketIfNotExists([]byte(name))
				if err != nil {
					return err
				}
				return nil
			})

			if err != nil {
				return lib.MakeErrorValue(runtime, err)
			}

			return nil
		})

		o.Set("view", func(call js.FunctionCall) js.Value {
			name := call.Argument(0).String()
			callBack := call.Argument(1).Export().(CallBackFunc)
			Bdb.View(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte(name))

				callBack(js.FunctionCall{
					This:      NewBucket(runtime, b),
					Arguments: []js.Value{NewBucket(runtime, b)},
				})

				return nil
			})
			return nil
		})

		o.Set("update", func(call js.FunctionCall) js.Value {
			name := call.Argument(0).String()
			callBack := call.Argument(1).Export().(CallBackFunc)
			Bdb.Update(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte(name))

				// 回调
				callBack(js.FunctionCall{
					This:      NewBucket(runtime, b),
					Arguments: []js.Value{NewBucket(runtime, b)},
				})

				return nil
			})
			return nil
		})
	})
}
