package gogit

import (
	"github.com/dop251/goja"
	"github.com/go-git/go-git/v5/config"
	"github.com/issueye/lichee-js/lib"
)

func NewBranch(rt *goja.Runtime, branch *config.Branch) *goja.Object {
	o := rt.NewObject()
	// branch.Name
	o.Set("name", branch.Name)
	// branch.Rebase
	o.Set("rebase", branch.Rebase)
	// branch.Description
	o.Set("description", branch.Description)
	// branch.Remote
	o.Set("remote", branch.Remote)
	// branch.Merge
	o.Set("merge", func(call goja.FunctionCall) goja.Value {
		return NewReferenceName(rt, branch.Merge)
	})

	// branch.Validate()
	o.Set("validate", func(call goja.FunctionCall) goja.Value {
		err := branch.Validate()
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return nil
	})

	return o
}
