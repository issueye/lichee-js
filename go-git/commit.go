package gogit

import (
	"github.com/dop251/goja"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/issueye/lichee-js/lib"
)

func NewCommit(rt *goja.Runtime, commit *object.Commit) *goja.Object {
	o := rt.NewObject()

	o.Set("author", NewSignature(rt, commit.Author))
	o.Set("committer", NewSignature(rt, commit.Committer))
	o.Set("hash", NewHash(rt, commit.Hash))
	o.Set("treeHash", NewHash(rt, commit.TreeHash))
	o.Set("id", NewHash(rt, commit.ID()))
	o.Set("message", commit.Message)
	o.Set("PGPSignature", commit.PGPSignature)
	o.Set("numParents", func() int {
		return commit.NumParents()
	})
	o.Set("tree", func() goja.Value {
		t, err := commit.Tree()
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return NewTree(rt, t)
	})
	o.Set("file", func(path string) goja.Value {
		f, err := commit.File(path)
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return NewFile(rt, f)
	})

	o.Set("files", func() goja.Value {
		fi, err := commit.Files()

		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return NewFileIter(rt, fi)
	})

	return o
}

func NewCommitIter(rt *goja.Runtime, iter object.CommitIter) *goja.Object {
	o := rt.NewObject()

	o.Set("close", func() {
		iter.Close()

	})

	o.Set("next", func() goja.Value {
		c, err := iter.Next()
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return NewCommit(rt, c)
	})

	o.Set("foreach", func(call goja.FunctionCall) goja.Value {
		callBack := call.Argument(0).Export().(OptionsFunc)

		err := iter.ForEach(func(c *object.Commit) error {
			obj := NewCommit(rt, c)
			callBack(goja.FunctionCall{
				This:      obj,
				Arguments: []goja.Value{obj},
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
