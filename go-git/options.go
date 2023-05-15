package gogit

import (
	"regexp"

	"github.com/dop251/goja"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/issueye/lichee-js/lib"
	"github.com/issueye/lichee/utils"
)

type OptionsFunc func(call goja.FunctionCall) *goja.Object

func NewCheckoutOptions(rt *goja.Runtime, option *git.CheckoutOptions) *goja.Object {
	o := rt.NewObject()

	o.Set("setKeep", func(keep bool) {
		option.Keep = keep
	})

	o.Set("setForce", func(force bool) {
		option.Force = force
	})

	o.Set("setCreate", func(create bool) {
		option.Create = create
	})

	o.Set("setHash", func(hash string) {
		option.Hash = plumbing.NewHash(hash)
	})

	o.Set("setSparse", func(sparse []string) {
		option.SparseCheckoutDirectories = sparse
	})

	o.Set("setBranch", func(name string) {
		option.Branch = plumbing.NewBranchReferenceName(name)
	})

	return o
}

func NewGrepOptions(rt *goja.Runtime, option *git.GrepOptions) *goja.Object {
	o := rt.NewObject()

	o.Set("invertMatch", func(invertMatch bool) {
		option.InvertMatch = invertMatch
	})

	o.Set("commitHash", func(hash string) {
		option.CommitHash = plumbing.NewHash(hash)
	})

	o.Set("referenceName", func(name string) {
		option.ReferenceName = plumbing.NewBranchReferenceName(name)
	})

	o.Set("pathSpces", func(expr string) goja.Value {
		reg, err := regexp.Compile(expr)
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}
		option.PathSpecs = append(option.PathSpecs, reg)
		return nil
	})

	o.Set("pattern", func(expr string) goja.Value {
		reg, err := regexp.Compile(expr)
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}
		option.Patterns = append(option.Patterns, reg)
		return nil
	})

	return o
}

func NewResetOptions(rt *goja.Runtime, option *git.ResetOptions) *goja.Object {
	o := rt.NewObject()

	o.Set("commit", func(hash string) {
		option.Commit = plumbing.NewHash(hash)
	})

	o.Set("mode", func(mode git.ResetMode) {
		option.Mode = mode
	})

	return o
}

func NewSignature(rt *goja.Runtime, sign object.Signature) *goja.Object {
	o := rt.NewObject()
	o.Set("email", sign.Email)
	o.Set("name", sign.Name)
	o.Set("when", sign.When.Format(utils.FormatDateTimeMs))
	o.Set("string", func() string {
		return sign.String()
	})

	return o
}

func NewFile(rt *goja.Runtime, f *object.File) *goja.Object {
	o := rt.NewObject()

	o.Set("hash", NewHash(rt, f.Hash))
	o.Set("id", NewHash(rt, f.ID()))
	o.Set("name", f.Name)
	o.Set("size", f.Size)
	o.Set("mode", f.Mode)
	o.Set("content", func() goja.Value {
		content, err := f.Contents()
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return lib.MakeReturnValue(rt, content)
	})
	o.Set("lines", func() goja.Value {
		s, err := f.Lines()
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return lib.MakeReturnValue(rt, s)
	})

	return o
}

func NewFileIter(rt *goja.Runtime, f *object.FileIter) *goja.Object {
	o := rt.NewObject()

	o.Set("close", func() {
		f.Close()
	})

	o.Set("next", func() goja.Value {
		f2, err := f.Next()
		if err != nil {
			return lib.MakeErrorValue(rt, err)
		}

		return NewFile(rt, f2)
	})

	o.Set("foreach", func(call goja.FunctionCall) goja.Value {
		callBack := call.Argument(0).Export().(OptionsFunc)

		err := f.ForEach(func(f *object.File) error {
			file := NewFile(rt, f)
			callBack(goja.FunctionCall{
				This:      file,
				Arguments: []goja.Value{file},
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
