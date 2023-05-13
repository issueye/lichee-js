package boltdb_test

import (
	"testing"

	licheejs "github.com/issueye/lichee-js"
	bolt "go.etcd.io/bbolt"
)

func Test_boltdb(t *testing.T) {
	js := licheejs.NewCore()

	t.Run("mod_boltdb", func(t *testing.T) {
		err := js.Run("test-boltdb", "test-boltdb.js")
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("native_boltdb", func(t *testing.T) {
		db, err := bolt.Open("test_bolt.db", 0666, nil)
		if err != nil {
			t.Errorf("生成 boltdb 失败，失败原因：%s", err.Error())
		}

		licheejs.RegisterBolt(js.GetRts(), "testBolt", db)
		err = js.Run("test_native_bolt", "test-native-boltdb.js")
		if err != nil {
			t.Error(err)
		}
	})
}
