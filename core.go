package licheejs

import (
	"embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	goja "github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"go.uber.org/zap"
)

//go:embed js/*
var Script embed.FS

const (
	GoPlugins = "lichee"
)

var (
	globalConvertProg *goja.Program                 // convert 转换代码的对应编译对象
	globalDayjsProg   *goja.Program                 // dayjs 转换代码的对应编译对象
	LogMap            = make(map[string]*ZapLogger) // 日志对象
)

type ZapLogger struct {
	log   *zap.Logger
	Close func()
}

type ModuleFunc = func(vm *goja.Runtime, module *goja.Object)

// Core
// goja运行时核心的结构体
type Core struct {
	// 全局goja加载目录
	globalPath string
	// 外部添加到内部的内容
	pkg map[string]map[string]any
	// 外部注册的模块
	modules map[string]ModuleFunc
	// 编译之后的对象
	// pro *goja.Program
	// 对应文件的编译对象
	proMap map[string]*goja.Program
	// 日志对象
	logger *zap.Logger
	// 锁
	lock *sync.Mutex
	// Name
	name string
	// 日志存放路径
	logPath string
	// vm 虚拟机
	vm *goja.Runtime
	// 日志输出模式
	// debug 输出到控制台和输出到日志文件
	// release 只输出到日志文件
	logMode LogOutMode
	// 注册
	registry *require.Registry

	loop *EventLoop
}

type OptFunc = func(*Core)

// NewCore
// 创建一个对象
func NewCore(opts ...OptFunc) *Core {
	c := new(Core)
	c.GetRts()
	c.lock = new(sync.Mutex)
	c.pkg = make(map[string]map[string]any)
	// 初始化全局
	c.pkg[GoPlugins] = make(map[string]any)
	c.modules = make(map[string]func(vm *goja.Runtime, module *goja.Object))
	c.proMap = make(map[string]*goja.Program)
	// 日志输出模式
	// debug 输出到控制台和输出到日志文件
	// release 只输出到日志文件
	c.logMode = LOM_RELEASE
	// 配置
	for _, opt := range opts {
		opt(c)
	}

	// 添加 导入方法 require
	c.registry = require.NewRegistry(
		// 全局加载路径
		require.WithGlobalFolders(c.globalPath),
	)
	c.registry.Enable(c.vm)

	c.loop = NewEventLoop(c.vm)
	// 加载goja模块
	c.loadScript("utils-arr2map", "convert.js", globalConvertProg)
	c.loadScript("dayjs", "dayjs.min.js", globalDayjsProg)

	return c
}

// OptionLog
// 配置日志
func OptionLog(path string, log *zap.Logger) OptFunc {
	return func(core *Core) {
		core.logger = log
		core.logPath = path
	}
}

func (c *Core) setupGojaRuntime(logger *zap.Logger) error {
	// 输出日志
	console := newConsole(logger)
	o := c.vm.NewObject()
	o.Set("log", console.Log)
	o.Set("debug", console.Debug)
	o.Set("info", console.Info)
	o.Set("error", console.Error)
	o.Set("warn", console.Warn)

	err := c.vm.Set("console", o)
	if err != nil {
		return err
	}

	return nil
}

// SetLogOutMode
// 日志输出模式
// debug 输出到控制台和输出到日志文件
// release 只输出到日志文件
func (c *Core) SetLogOutMode(mod LogOutMode) {
	c.logMode = mod
}

func (c *Core) loadModule() {
	// 添加 日志方法 console
	if c.name == "" {
		c.name = "lichee-test"
	}

	log, ok := LogMap[c.name]
	if !ok {
		path := filepath.Join(c.logPath, fmt.Sprintf("%s.log", c.name))
		l, close, err := newZap(path, c.logMode)

		if err != nil {
			c.Errorf("加载日志失败，失败原因：%s", err.Error())
		}

		log = &ZapLogger{
			log:   l,
			Close: close,
		}

		LogMap[c.name] = log
	}

	// 设置运行时
	c.setupGojaRuntime(log.log)

	// 加载全局对象
	c.loadVariable()

	// 加载外部模块
	c.registerModule()
}

// GetRts
// 获取运行时
func (c *Core) GetRts() *goja.Runtime {
	if c.vm == nil {
		c.vm = goja.New()
	}

	return c.vm
}

func (c *Core) SetGlobalPath(path string) {
	c.globalPath = path
}

// loadScript
// 加载文件中的goja脚本
func (c *Core) loadScript(name string, gojaName string, p *goja.Program) {
	if p == nil {
		path := fmt.Sprintf("js/%s", gojaName)
		src, err := Script.ReadFile(path)
		if err != nil {
			c.Errorf("读取文件失败，失败原因：%s", err.Error())
			return
		}
		p, err = goja.Compile(name, string(src), false)
		if err != nil {
			return
		}
	}
	// 运行脚本
	_, err := c.vm.RunProgram(p)
	if err != nil {
		c.Errorf("运行脚本[%s]失败，失败原因：%s", name, err.Error())
	}
}

// Run
// 运行脚本 文件
func (c *Core) Run(name, path string) error {
	c.name = name
	return c.run(path, c.vm)
}

// RunVM
// 运行脚本 文件
func (c *Core) RunVM(path string, vm *goja.Runtime) error {
	return c.run(path, vm)
}

func (c *Core) run(path string, vm *goja.Runtime) error {
	c.loadModule()
	var tmpPath string
	if c.globalPath != "" {
		tmpPath = filepath.Join(c.globalPath, path)
	} else {
		tmpPath = path
	}

	// 读取文件
	src, err := os.ReadFile(tmpPath)
	if err != nil {
		c.Errorf("读取文件失败，失败原因：%s", err.Error())
	} else {
		// 编译文件
		pro, err := goja.Compile(c.name, string(src), false)
		if err != nil {
			c.Errorf("编译代码失败，失败原因：%s", err.Error())
		} else {
			c.proMap[path] = pro
		}
	}

	// 只有存在编译对象时，才运行
	if c.proMap[path] != nil {
		var err error
		if vm != nil {
			loop := NewEventLoop(vm)

			loop.Run(func(r *goja.Runtime) {
				_, err := vm.RunProgram(c.proMap[path])
				if gojaErr, ok := err.(*goja.Exception); ok {
					err = errors.New(gojaErr.String())
					return
				}
			})
		} else {
			c.loop.Run(func(vm *goja.Runtime) {
				_, err := vm.RunProgram(c.proMap[path])
				if gojaErr, ok := err.(*goja.Exception); ok {
					err = errors.New(gojaErr.String())
					return
				}
			})
		}

		if err != nil {
			return err
		}
	}
	return nil
}

// ExportFunc
// 导出JS方法
func (c *Core) ExportFunc(name string, fn any) error {
	vm := c.GetRts()
	return vm.ExportTo(vm.Get(name), fn)
}

// RunString
// 运行脚本 字符串
func (c *Core) RunString(src string) error {
	c.loadModule()
	_, err := c.vm.RunString(src)
	if gojaErr, ok := err.(*goja.Exception); ok {
		return fmt.Errorf("运行脚本失败，失败原因：%s", gojaErr.Error())
	}
	return nil
}

// SetGlobalProperty
// 写入数据到全局对象中
func (c *Core) SetGlobalProperty(key string, value any) {
	c.pkg[GoPlugins][key] = value
}

func (c *Core) loadVariable() {
	// 加载其他模块
	for name, mod := range c.pkg {
		gojaMod := c.vm.NewObject()
		for k, v := range mod {
			gojaMod.Set(k, v)
		}
		c.vm.Set(name, gojaMod)
	}
}

// registerModule
// 外部注册模块到goja
func (c *Core) registerModule() {
	for Name, moduleFn := range c.modules {
		require.RegisterNativeModule(Name, func(runtime *goja.Runtime, module *goja.Object) {
			m := module.Get("exports").(*goja.Object)
			moduleFn(runtime, m)
		})
	}
}

// SetProperty
// 向模块写入变量或者写入方法
func (c *Core) SetProperty(moduleName, key string, value any) {
	mod, ok := c.pkg[moduleName]
	if !ok {
		c.pkg[moduleName] = make(map[string]any)
		mod = c.pkg[moduleName]
	}
	mod[key] = value
}

// RegisterModule
// 注册模块
func (c *Core) RegisterModule(moduleName string, fn ModuleFunc) {
	c.modules[moduleName] = fn
}
