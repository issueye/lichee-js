package lib

import (
	"io"

	js "github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

func InitIO() {
	require.RegisterNativeModule("std/io", func(runtime *js.Runtime, module *js.Object) {
		// o := module.Get("exports").(*js.Object)
		// io.WriteCloser
	})
}

func NewWriteCloser(rt *js.Runtime, w io.WriteCloser) js.Value {
	o := rt.NewObject()
	o.Set("write", func(call js.FunctionCall) js.Value {
		data := call.Argument(0).Export().([]byte)
		n, err := w.Write(data)
		if err != nil {
			return MakeErrorValue(rt, err)
		}

		return MakeReturnValue(rt, n)
	})

	// w.Close()
	o.Set("close", func(call js.FunctionCall) js.Value {
		err := w.Close()
		if err != nil {
			return MakeErrorValue(rt, err)
		}

		return nil
	})

	return o
}

func NewReadCloser(rt *js.Runtime, w io.ReadCloser) js.Value {
	o := rt.NewObject()
	o.Set("read", func(call js.FunctionCall) js.Value {
		data := call.Argument(0).Export().([]byte)
		n, err := w.Read(data)
		if err != nil {
			return MakeErrorValue(rt, err)
		}

		return MakeReturnValue(rt, n)
	})

	// w.Close()
	o.Set("close", func(call js.FunctionCall) js.Value {
		err := w.Close()
		if err != nil {
			return MakeErrorValue(rt, err)
		}

		return nil
	})

	return o
}

func NewWriter(rt *js.Runtime, w io.Writer) js.Value {
	o := rt.NewObject()
	o.Set("write", func(call js.FunctionCall) js.Value {
		data := call.Argument(0).Export().([]byte)
		n, err := w.Write(data)
		if err != nil {
			return MakeErrorValue(rt, err)
		}

		return MakeReturnValue(rt, n)
	})

	return o
}

func NewReader(rt *js.Runtime, r io.Reader) js.Value {
	o := rt.NewObject()

	o.Set("read", func(call js.FunctionCall) js.Value {
		data := call.Argument(0).Export().([]byte)
		n, err := r.Read(data)
		if err != nil {
			return MakeErrorValue(rt, err)
		}

		return MakeReturnValue(rt, n)
	})
	return o
}
