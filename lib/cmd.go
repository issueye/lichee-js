package lib

import (
	"os"
	"os/exec"

	js "github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func init() {
	require.RegisterNativeModule("std/os/exec", func(runtime *js.Runtime, module *js.Object) {
		o := module.Get("exports").(*js.Object)
		o.Set("command", func(name string, args ...string) js.Value {
			cmd := exec.Command(name, args...)
			return NewCmd(runtime, cmd)
		})
	})
}

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}

func NewCmd(rt *js.Runtime, cmd *exec.Cmd) *js.Object {

	o := rt.NewObject()

	// cmd.Args
	o.Set("args", func() js.Value {
		return MakeReturnValue(rt, cmd.Args)
	})

	// cmd.Dir
	o.Set("dir", func() js.Value {
		return MakeReturnValue(rt, cmd.Dir)
	})

	// cmd.Env
	o.Set("env", func() js.Value {
		return MakeReturnValue(rt, cmd.Env)
	})

	// cmd.Path
	o.Set("path", func() js.Value {
		return MakeReturnValue(rt, cmd.Path)
	})

	// cmd.String()
	o.Set("string", func(call js.FunctionCall) js.Value {
		return MakeReturnValue(rt, cmd.String())
	})

	// cmd.Wait()
	o.Set("wait", func(call js.FunctionCall) js.Value {
		err := cmd.Wait()
		if err != nil {
			MakeErrorValue(rt, err)
		}

		return nil
	})

	// cmd.ExtraFiles
	// o.Set("extraFiles", func(call js.FunctionCall) js.Value {
	// 	cmd.ExtraFiles
	// })

	// cmd.Process
	o.Set("process", func(call js.FunctionCall) js.Value {
		return NewProcess(rt, cmd.Process)
	})

	o.Set("stderr", func(call js.FunctionCall) js.Value {
		return NewWriter(rt, cmd.Stderr)
	})

	// cmd.Stdin
	o.Set("stdin", func(call js.FunctionCall) js.Value {
		return NewReader(rt, cmd.Stdin)
	})

	// cmd.stdout
	o.Set("stdout", func(call js.FunctionCall) js.Value {
		return NewWriter(rt, cmd.Stdout)
	})

	// cmd.CombinedOutput()
	o.Set("combinedOutput", func() js.Value {
		b, err := cmd.CombinedOutput()
		if err != nil {
			MakeErrorValue(rt, err)
		}

		return MakeReturnValue(rt, b)
	})

	// cmd.Output()
	o.Set("combinedOutput", func() js.Value {
		b, err := cmd.Output()
		if err != nil {
			MakeErrorValue(rt, err)
		}

		return MakeReturnValue(rt, b)
	})

	// cmd.Run()
	o.Set("combinedOutput", func() js.Value {
		err := cmd.Run()
		if err != nil {
			MakeErrorValue(rt, err)
		}

		return nil
	})

	// cmd.Start()
	o.Set("combinedOutput", func() js.Value {
		err := cmd.Start()
		if err != nil {
			MakeErrorValue(rt, err)
		}

		return nil
	})

	// cmd.StderrPipe()
	o.Set("combinedOutput", func() js.Value {
		closer, err := cmd.StderrPipe()
		if err != nil {
			MakeErrorValue(rt, err)
		}

		return NewReadCloser(rt, closer)
	})

	// cmd.StdinPipe()
	o.Set("combinedOutput", func() js.Value {
		closer, err := cmd.StdinPipe()
		if err != nil {
			MakeErrorValue(rt, err)
		}

		return NewWriteCloser(rt, closer)
	})

	// cmd.StdoutPipe()
	o.Set("StdoutPipe", func() js.Value {
		closer, err := cmd.StderrPipe()
		if err != nil {
			MakeErrorValue(rt, err)
		}

		return NewReadCloser(rt, closer)
	})

	return o
}

func NewProcess(rt *js.Runtime, process *os.Process) *js.Object {
	o := rt.NewObject()

	// process.Pid
	o.Set("pid", func() js.Value {
		return MakeReturnValue(rt, process.Pid)
	})

	// process.Kill()
	o.Set("kill", func() js.Value {
		err := process.Kill()
		if err != nil {
			MakeErrorValue(rt, err)
		}

		return nil
	})

	// process.Release()
	o.Set("release", func() js.Value {
		err := process.Release()
		if err != nil {
			MakeErrorValue(rt, err)
		}

		return nil
	})

	// process.Signal()
	o.Set("signal", func(call js.FunctionCall) js.Value {
		sig := call.Argument(0).Export().(os.Signal)
		err := process.Signal(sig)
		if err != nil {
			MakeErrorValue(rt, err)
		}

		return nil
	})

	// process.Wait()
	o.Set("wait", func(call js.FunctionCall) js.Value {
		ps, err := process.Wait()
		if err != nil {
			MakeErrorValue(rt, err)
		}

		return NewProcessState(rt, ps)
	})

	return o
}

func NewProcessState(rt *js.Runtime, ps *os.ProcessState) js.Value {
	o := rt.NewObject()

	// ps.ExitCode()
	o.Set("exitCode", func() js.Value {
		return MakeReturnValue(rt, ps.ExitCode())
	})

	// ps.Exited()
	o.Set("exited", func() js.Value {
		return MakeReturnValue(rt, ps.Exited())
	})

	// ps.Pid()
	o.Set("pid", func() js.Value {
		return MakeReturnValue(rt, ps.Pid())
	})

	// ps.String()
	o.Set("string", func() js.Value {
		return MakeReturnValue(rt, ps.String())
	})

	// ps.Success()
	o.Set("success", func() js.Value {
		return MakeReturnValue(rt, ps.Success())
	})

	// ps.Sys()
	o.Set("sys", func() js.Value {
		return MakeReturnValue(rt, ps.Sys())
	})

	// ps.SysUsage()
	o.Set("sysUsage", func() js.Value {
		return MakeReturnValue(rt, ps.SysUsage())
	})

	// ps.SystemTime()
	// o.Set("systemTime", func() js.Value {
	// return MakeReturnValue(rt, ps.SystemTime())
	// })

	return o
}
