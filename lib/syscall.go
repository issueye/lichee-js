package lib

import (
	"syscall"

	js "github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

func init() {
	require.RegisterNativeModule("std/syscall", func(runtime *js.Runtime, module *js.Object) {
		o := module.Get("exports").(*js.Object)
		o.Set("SIGHUP", syscall.SIGHUP)
		o.Set("SIGINT", syscall.SIGINT)
		o.Set("SIGQUIT", syscall.SIGQUIT)
		o.Set("SIGILL", syscall.SIGILL)
		o.Set("SIGTRAP", syscall.SIGTRAP)
		o.Set("SIGABRT", syscall.SIGABRT)
		o.Set("SIGBUS", syscall.SIGBUS)
		o.Set("SIGFPE", syscall.SIGFPE)
		o.Set("SIGKILL", syscall.SIGKILL)
		o.Set("SIGSEGV", syscall.SIGSEGV)
		o.Set("SIGPIPE", syscall.SIGPIPE)
		o.Set("SIGALRM", syscall.SIGALRM)
		o.Set("SIGTERM", syscall.SIGTERM)
	})
}
