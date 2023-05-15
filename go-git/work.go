package gogit

import (
	"github.com/dop251/goja"
	git "github.com/go-git/go-git/v5"
	"github.com/issueye/lichee-js/lib"
)

func NewWorkTree(rt *goja.Runtime, w *git.Worktree) *goja.Object {
	o := rt.NewObject()

	// 拉取
	o.Set("pull", func(call goja.FunctionCall) goja.Value {
		err := w.Pull(&git.PullOptions{})
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return nil
	})

	// add
	o.Set("add", func(call goja.FunctionCall) goja.Value {
		path := call.Argument(0).String()
		h, err := w.Add(path)
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return NewHash(rt, h)
	})

	// w.AddGlob()
	o.Set("addGlob", func(call goja.FunctionCall) goja.Value {
		pattern := call.Argument(0).String()
		err := w.AddGlob(pattern)
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return nil
	})

	// Checkout
	o.Set("checkout", func(call goja.FunctionCall) goja.Value {
		callBack := call.Argument(0).Export().(OptionsFunc)

		option := &git.CheckoutOptions{}

		// 通过回调设置参数
		obj := NewCheckoutOptions(rt, option)
		callBack(goja.FunctionCall{
			This:      obj,
			Arguments: []goja.Value{obj},
		})

		err := w.Checkout(option)
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return nil
	})

	o.Set("clean", func(dir bool) {
		w.Clean(&git.CleanOptions{
			Dir: dir,
		})
	})

	// w.Commit()
	o.Set("commit", func(msg string) goja.Value {
		h, err := w.Commit(msg, &git.CommitOptions{})
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return NewHash(rt, h)
	})

	// w.Grep()
	// 检索内容
	o.Set("grep", func(call goja.FunctionCall) goja.Value {
		callBack := call.Argument(0).Export().(OptionsFunc)

		option := &git.GrepOptions{}

		obj := NewGrepOptions(rt, option)
		callBack(goja.FunctionCall{
			This:      obj,
			Arguments: []goja.Value{obj},
		})

		gr, err := w.Grep(option)
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return NewResults(rt, gr)
	})

	// w.Move()
	o.Set("move", func(from, to string) goja.Value {
		h, err := w.Move(from, to)
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return NewHash(rt, h)
	})

	// w.Remove()
	o.Set("remove", func(path string) goja.Value {
		h, err := w.Remove(path)
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return NewHash(rt, h)
	})

	// w.RemoveGlob()
	o.Set("removeGlob", func(pattern string) goja.Value {
		err := w.RemoveGlob(pattern)
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return nil
	})

	// reset
	o.Set("reset", func(call goja.FunctionCall) goja.Value {
		callBack := call.Argument(0).Export().(OptionsFunc)

		option := &git.ResetOptions{}

		obj := NewResetOptions(rt, option)

		callBack(goja.FunctionCall{
			This:      obj,
			Arguments: []goja.Value{obj},
		})

		w.Reset(option)
		return nil
	})

	o.Set("status", func(call goja.FunctionCall) goja.Value {
		s, err := w.Status()
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return NewStatus(rt, s)
	})

	o.Set("submodule", func(call goja.FunctionCall) goja.Value {
		name := call.Argument(0).String()
		s, err := w.Submodule(name)
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return NewSubmodule(rt, s)
	})

	return o
}

type ResultCallBackFunc func(result *goja.Object)

// Result
func NewResults(rt *goja.Runtime, gr []git.GrepResult) *goja.Object {
	o := rt.NewObject()

	o.Set("foreach", func(call goja.FunctionCall) goja.Value {
		callBack := call.Argument(0).Export().(ResultCallBackFunc)
		for _, result := range gr {
			callBack(NewResult(rt, result))
		}

		return nil
	})

	return o
}

func NewResult(rt *goja.Runtime, gr git.GrepResult) *goja.Object {
	o := rt.NewObject()

	o.Set("fileName", gr.FileName)
	o.Set("lineNumber", gr.LineNumber)
	o.Set("treeName", gr.TreeName)
	o.Set("content", gr.Content)
	o.Set("string", func() string {
		return gr.String()
	})

	return o
}

func NewStatus(rt *goja.Runtime, s git.Status) *goja.Object {
	o := rt.NewObject()

	o.Set("isClean", s.IsClean())
	o.Set("string", func() string {
		return s.String()
	})
	// 检查给定的路径是否为“未跟踪”
	o.Set("isUntracked", func(path string) bool {
		return s.IsUntracked(path)
	})

	o.Set("file", func(path string) goja.Value {
		fs := s.File(path)

		data := map[string]string{
			"extra":    fs.Extra,
			"staging":  string(fs.Staging),
			"worktree": string(fs.Worktree),
		}

		return lib.MakeReturnValue(rt, data)
	})

	return o
}

func NewSubmodule(rt *goja.Runtime, sub *git.Submodule) *goja.Object {
	o := rt.NewObject()

	o.Set("config", func() goja.Value {
		data := map[string]string{
			"name": sub.Config().Name,
			"path": sub.Config().Path,
			"url":  sub.Config().URL,
		}

		return lib.MakeReturnValue(rt, data)
	})

	o.Set("init", func() goja.Value {
		err := sub.Init()
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return nil
	})

	o.Set("repository", func() goja.Value {
		r, err := sub.Repository()
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return NewRepo(rt, r)
	})

	o.Set("status", func() goja.Value {
		ss, err := sub.Status()
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return NewSubmoduleStatus(rt, ss)
	})

	return o
}

func NewSubmoduleStatus(rt *goja.Runtime, status *git.SubmoduleStatus) *goja.Object {
	o := rt.NewObject()

	o.Set("branch", NewReferenceName(rt, status.Branch))
	o.Set("current", NewHash(rt, status.Current))
	o.Set("expected", NewHash(rt, status.Expected))
	o.Set("path", status.Path)
	o.Set("isClean", status.IsClean())
	o.Set("", func() string {
		return status.String()
	})

	return o
}
