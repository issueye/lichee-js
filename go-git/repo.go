package gogit

import (
	"github.com/dop251/goja"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"github.com/issueye/lichee-js/lib"
)

func NewRepo(rt *goja.Runtime, repo *git.Repository) *goja.Object {
	o := rt.NewObject()

	// repo.Branch()
	o.Set("branch", func(call goja.FunctionCall) goja.Value {
		name := call.Argument(0).String()
		b, err := repo.Branch(name)
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return NewBranch(rt, b)
	})

	// Worktree
	o.Set("workTree", func(call goja.FunctionCall) goja.Value {
		w, err := repo.Worktree()
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return NewWorkTree(rt, w)
	})

	o.Set("log", func(call goja.FunctionCall) goja.Value {
		log, err := repo.Log(&git.LogOptions{})
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return NewLog(rt, log)
	})

	o.Set("branches", func() goja.Value {
		ri, err := repo.Branches()
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}
		return NewRefrenceIter(rt, ri)
	})

	o.Set("commitObject", func(hash string) goja.Value {
		h := plumbing.NewHash(hash)
		c, err := repo.CommitObject(h)
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return NewCommit(rt, c)
	})

	o.Set("commitObjects", func() goja.Value {
		ci, err := repo.CommitObjects()
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return NewCommitIter(rt, ci)
	})

	return o
}

func NewRefrenceIter(rt *goja.Runtime, sr storer.ReferenceIter) *goja.Object {
	o := rt.NewObject()

	rt.Set("close", func() {
		sr.Close()
	})

	rt.Set("next", func() goja.Value {
		r, err := sr.Next()
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return NewReference(rt, r)
	})

	// sr.ForEach()
	rt.Set("foreach", func(call goja.FunctionCall) goja.Value {
		callback := call.Argument(0).Export().(OptionsFunc)

		err := sr.ForEach(func(r *plumbing.Reference) error {
			obj := NewReference(rt, r)
			callback(goja.FunctionCall{
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
