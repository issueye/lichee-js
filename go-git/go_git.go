package gogit

import (
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	git "github.com/go-git/go-git/v5"
	"github.com/issueye/lichee-js/lib"
)

func InitGoGit() {
	require.RegisterNativeModule("go/git", func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)

		// types
		o.Set("MIXED_RESET", git.MixedReset)
		o.Set("HARD_RESET", git.HardReset)
		o.Set("MERGE_RESET", git.MergeReset)
		o.Set("SOFT_RESET", git.SoftReset)

		o.Set("Unmodified", git.Unmodified)
		o.Set("Untracked", git.Untracked)
		o.Set("Modified", git.Modified)
		o.Set("Added", git.Added)
		o.Set("Deleted", git.Deleted)
		o.Set("Renamed", git.Renamed)
		o.Set("Copied", git.Copied)
		o.Set("UpdatedButUnmerged", git.UpdatedButUnmerged)

		// open
		o.Set("open", func(call goja.FunctionCall) goja.Value {
			path := call.Argument(0).String()
			r, err := git.PlainOpen(path)
			if err != nil {
				return lib.MakeErrorValue(runtime, err)
			}

			// 创建一个仓库对象
			return NewRepo(runtime, r)
		})

		// clone
		o.Set("clone", func(call goja.FunctionCall) goja.Value {
			path := call.Argument(0).String()
			url := call.Argument(0).String()
			r, err := git.PlainClone(path, false, &git.CloneOptions{
				URL:               url,
				RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
			})

			if err != nil {
				return lib.MakeErrorValue(runtime, err)
			}

			// 创建一个仓库对象
			return NewRepo(runtime, r)
		})

		// init
		o.Set("init", func(call goja.FunctionCall) goja.Value {
			path := call.Argument(0).String()
			r, err := git.PlainInit(path, false)
			if err != nil {
				return lib.MakeErrorValue(runtime, err)
			}

			// 创建一个仓库对象
			return NewRepo(runtime, r)
		})
	})
}
