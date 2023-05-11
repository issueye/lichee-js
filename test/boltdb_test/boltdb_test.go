package boltdb_test

import (
	"testing"

	licheejs "github.com/issueye/lichee-js"
)

func Test_boltdb(t *testing.T) {
	js := licheejs.NewCore()

	err := js.Run("test-boltdb", "test-boltdb.js")
	if err != nil {
		t.Error(err)
	}
}
