package redis

import (
	"context"
	"time"

	"github.com/dop251/goja"
	"github.com/issueye/lichee-js/lib"
	"github.com/redis/go-redis/v9"
)

func List(rt *goja.Runtime, o *goja.Object, rdb *redis.Client) {
	// BLPOP
	o.Set("bLpop", func(timeOut int64, keys ...string) goja.Value {
		ctx := context.Background()
		sc := rdb.BLPop(ctx, time.Duration(timeOut)*time.Second, keys...)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// BRPOP
	o.Set("bRpop", func(timeout int64, keys ...string) goja.Value {
		ctx := context.Background()
		sc := rdb.BRPop(ctx, time.Duration(timeout)*time.Second, keys...)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// BRPOPLPUSH
	o.Set("bRpop", func(source string, destination string, timeout int64) goja.Value {
		ctx := context.Background()
		sc := rdb.BRPopLPush(ctx, source, destination, time.Duration(timeout)*time.Second)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// LINDEX
	o.Set("lIndex", func(key string, index int64) goja.Value {
		ctx := context.Background()
		sc := rdb.LIndex(ctx, key, index)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// LINSERT
	o.Set("lInsert", func(key string, op string, pivot interface{}, value interface{}) goja.Value {
		ctx := context.Background()
		sc := rdb.LInsert(ctx, key, op, pivot, value)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// LLEN
	o.Set("lLen", func(key string) goja.Value {
		ctx := context.Background()
		sc := rdb.LLen(ctx, key)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// Lpop
	o.Set("lPop", func(key string) goja.Value {
		ctx := context.Background()
		sc := rdb.LPop(ctx, key)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// LPUSH
	o.Set("lPush", func(key string, values ...interface{}) goja.Value {
		ctx := context.Background()
		sc := rdb.LPush(ctx, key, values...)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// LPUSHX
	o.Set("lPushX", func(key string, values ...interface{}) goja.Value {
		ctx := context.Background()
		sc := rdb.LPushX(ctx, key, values...)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// LRANGE
	o.Set("lRange", func(key string, start int64, stop int64) goja.Value {
		ctx := context.Background()
		sc := rdb.LRange(ctx, key, start, stop)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// LREM
	o.Set("lRem", func(key string, count int64, value interface{}) goja.Value {
		ctx := context.Background()
		sc := rdb.LRem(ctx, key, count, value)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// LSET
	o.Set("lSet", func(key string, count int64, value interface{}) goja.Value {
		ctx := context.Background()
		sc := rdb.LSet(ctx, key, count, value)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// LTRIM
	o.Set("lTrim", func(key string, start int64, stop int64) goja.Value {
		ctx := context.Background()
		sc := rdb.LTrim(ctx, key, start, stop)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// RPOP
	o.Set("rPop", func(key string) goja.Value {
		ctx := context.Background()
		sc := rdb.RPop(ctx, key)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// RPOPLPUSH
	o.Set("rPopLPush", func(source string, destination string) goja.Value {
		ctx := context.Background()
		sc := rdb.RPopLPush(ctx, source, destination)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// RPUSH
	o.Set("rPush", func(key string, values ...interface{}) goja.Value {
		ctx := context.Background()
		sc := rdb.RPush(ctx, key, values)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// RPUSHX
	o.Set("rPushX", func(key string, values ...interface{}) goja.Value {
		ctx := context.Background()
		sc := rdb.RPushX(ctx, key, values)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})
}
