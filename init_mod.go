package licheejs

import (
	"net/http"

	"github.com/dop251/goja"
	"github.com/issueye/lichee-js/boltdb"
	"github.com/issueye/lichee-js/compress"
	"github.com/issueye/lichee-js/container"
	"github.com/issueye/lichee-js/db"
	"github.com/issueye/lichee-js/goquery"
	"github.com/issueye/lichee-js/lib"
	lichee_http "github.com/issueye/lichee-js/net/http"
	lichee_url "github.com/issueye/lichee-js/net/url"
	"github.com/issueye/lichee-js/path"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// 注册 Redis
func RegisterRedis(name string, rdb *redis.Client) {
	db.RegisterRedis(name, rdb)
}

// 注册数据库
func RegisterDB(name string, gdb *gorm.DB) {
	db.RegisterDB(name, gdb)
}

// http request
func NewRequest(runtime *goja.Runtime, r *http.Request) *goja.Object {
	return lichee_http.NewRequest(runtime, r)
}

// http response
func NewResponse(runtime *goja.Runtime, w http.ResponseWriter) *goja.Object {
	return lichee_http.NewResponse(runtime, w)
}

// register mod
func init() {
	compress.InitGzip()    // gzip
	container.InitList()   // list
	boltdb.InitBolt()      // boltdb
	goquery.InitGoQuery()  // goquery
	lib.InitCmd()          // cmd
	lib.InitError()        // error
	lib.InitFile()         // file
	lib.InitFmt()          // fmt
	lib.InitIni()          // ini
	lib.InitIO()           // io
	lib.InitOS()           // os
	lib.InitSyscall()      // syscall
	lib.InitTime()         // time
	lib.InitTypes()        // types
	lib.InitUtils()        // utils
	lichee_http.InitHttp() // http
	lichee_url.InitUrl()   // url
	path.InitFilepath()    // filepath
}
