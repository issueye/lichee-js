package gogit

import (
	"github.com/dop251/goja"
	"github.com/go-git/go-git/v5/plumbing"
)

// plumbing.Hash
func NewHash(rt *goja.Runtime, h plumbing.Hash) *goja.Object {
	o := rt.NewObject()
	o.Set("string", h.String())
	o.Set("isZero", h.IsZero())
	return o
}

func NewReferenceName(rt *goja.Runtime, merge plumbing.ReferenceName) *goja.Object {
	o := rt.NewObject()
	o.Set("isBranch", merge.IsBranch())
	o.Set("isNote", merge.IsNote())
	o.Set("isRemote", merge.IsRemote())
	o.Set("isTag", merge.IsTag())
	return o
}

func NewReference(rt *goja.Runtime, reference *plumbing.Reference) *goja.Object {
	o := rt.NewObject()

	o.Set("hash", NewHash(rt, reference.Hash()))
	o.Set("target", NewReferenceName(rt, reference.Target()))
	o.Set("name", NewReferenceName(rt, reference.Name()))
	o.Set("string", func() string {
		return reference.String()
	})
	o.Set("strings", reference.Strings())

	return o
}
