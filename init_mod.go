package licheejs

import (
	"net/http"

	"github.com/dop251/goja"
	"github.com/issueye/lichee-js/compress"
	"github.com/issueye/lichee-js/container"
	"github.com/issueye/lichee-js/db"
	lichee_bolt "github.com/issueye/lichee-js/db/boltdb"
	lichee_redis "github.com/issueye/lichee-js/db/redis"
	"github.com/issueye/lichee-js/goquery"
	"github.com/issueye/lichee-js/lib"
	lichee_http "github.com/issueye/lichee-js/net/http"
	lichee_url "github.com/issueye/lichee-js/net/url"
	"github.com/issueye/lichee-js/path"
	"github.com/redis/go-redis/v9"
	bolt "go.etcd.io/bbolt"
	"gorm.io/gorm"
)

// 注册数据库
func RegisterDB(name string, gdb *gorm.DB) {
	db.RegisterDB(name, gdb)
}

// 注册一个boltdb
func RegisterBolt(rt *goja.Runtime, name string, db *bolt.DB) {
	lichee_bolt.RegisterNativeBolt(rt, name, db)
}

// 注册 Redis
func RegisterRedis(rt *goja.Runtime, name string, rdb *redis.Client) {
	lichee_redis.RegisterNativeRedis(rt, name, rdb)
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
	compress.InitGzip()  // gzip
	container.InitList() // list

	goquery.InitGoQuery()    // goquery
	lib.InitCmd()            // cmd
	lib.InitError()          // error
	lib.InitFile()           // file
	lib.InitFmt()            // fmt
	lib.InitIni()            // ini
	lib.InitIO()             // io
	lib.InitOS()             // os
	lib.InitSyscall()        // syscall
	lib.InitTime()           // time
	lib.InitTypes()          // types
	path.InitFilepath()      // filepath
	lib.InitUtils()          // utils
	lichee_http.InitHttp()   // http
	lichee_url.InitUrl()     // url
	lichee_bolt.InitBolt()   // boltdb
	lichee_redis.InitRedis() // redis
}
