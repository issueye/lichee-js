package gogit

import (
	"github.com/dop251/goja"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/issueye/lichee-js/lib"
)

func NewLog(rt *goja.Runtime, commit object.CommitIter) *goja.Object {
	o := rt.NewObject()

	o.Set("close", func() {
		commit.Close()
	})

	o.Set("next", func() goja.Value {
		c, err := commit.Next()
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return NewCommit(rt, c)
	})

	o.Set("foreach", func(call goja.FunctionCall) goja.Value {
		callBack := call.Argument(0).Export().(OptionsFunc)
		err := commit.ForEach(func(c *object.Commit) error {
			com := NewCommit(rt, c)
			callBack(goja.FunctionCall{
				This:      com,
				Arguments: []goja.Value{com},
			})

			return nil
		})

		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return nil
	})

	return o
}

func NewTree(rt *goja.Runtime, t *object.Tree) *goja.Object {
	o := rt.NewObject()

	o.Set("hash", NewHash(rt, t.Hash))
	o.Set("id", NewHash(rt, t.ID()))
	o.Set("tree", func(path string) goja.Value {
		t2, err := t.Tree(path)
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return NewTree(rt, t2)
	})
	o.Set("file", func(path string) goja.Value {
		f, err := t.File(path)
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return NewFile(rt, f)
	})
	o.Set("files", func() goja.Value {
		fi := t.Files()

		return NewFileIter(rt, fi)
	})

	o.Set("findEntry", func(path string) goja.Value {
		te, err := t.FindEntry(path)
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return NewTreeEntry(rt, te)
	})

	return o
}

func NewTreeEntry(rt *goja.Runtime, t *object.TreeEntry) *goja.Object {
	o := rt.NewObject()

	o.Set("hash", NewHash(rt, t.Hash))
	o.Set("name", t.Name)
	o.Set("mode", t.Mode)

	return o
}
