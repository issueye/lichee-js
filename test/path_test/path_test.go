package path_test

import (
	"testing"

	licheejs "github.com/issueye/lichee-js"
)

func Test_path(t *testing.T) {
	c := licheejs.NewCore()
	t.Run("abs", func(t *testing.T) {
		src := `
		var path = require('std/path/filepath')
		let a = path.abs('/home/local')
		console.log(a)
		`
		err := c.RunString("filepath-abs", src)
		if err != nil {
			t.Errorf("abs err :%s", err.Error())
		}
	})

	t.Run("join", func(t *testing.T) {
		src := `
		var path = require('std/path/filepath')
		let a = path.join('/home/local', 'test', 'test001')
		console.log(a)
		`
		err := c.RunString("filepath-join", src)
		if err != nil {
			t.Errorf("abs err :%s", err.Error())
		}
	})

	t.Run("ext", func(t *testing.T) {
		src := `
		var path = require('std/path/filepath')
		let a = path.ext('/home/local/code.go')
		console.log(a)
		`
		err := c.RunString("filepath-ext", src)
		if err != nil {
			t.Errorf("abs err :%s", err.Error())
		}
	})
}
