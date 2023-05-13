package boltdb

import (
	"github.com/dop251/goja"
	"github.com/issueye/lichee-js/lib"
	bolt "go.etcd.io/bbolt"
)

type CallBackFunc = func(goja.FunctionCall) goja.Value

func NewBoltDb(rt *goja.Runtime, db *bolt.DB) *goja.Object {
	o := rt.NewObject()

	// createBucket
	// 创建 bucket
	o.Set("createBucket", func(call goja.FunctionCall) goja.Value {
		name := call.Argument(0).String()
		err := db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(name))
			if err != nil {
				return err
			}
			return nil
		})

		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return nil
	})

	// view
	// 查询数据
	o.Set("view", func(call goja.FunctionCall) goja.Value {
		name := call.Argument(0).String()
		callBack := call.Argument(1).Export().(CallBackFunc)
		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(name))

			callBack(goja.FunctionCall{
				This:      NewBucket(rt, b),
				Arguments: []goja.Value{NewBucket(rt, b)},
			})

			return nil
		})
		return nil
	})

	// update
	// 更新数据
	o.Set("update", func(call goja.FunctionCall) goja.Value {
		name := call.Argument(0).String()
		callBack := call.Argument(1).Export().(CallBackFunc)
		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(name))

			// 回调
			callBack(goja.FunctionCall{
				This:      NewBucket(rt, b),
				Arguments: []goja.Value{NewBucket(rt, b)},
			})

			return nil
		})
		return nil
	})

	return o
}
