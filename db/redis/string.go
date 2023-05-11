package redis

import (
	"context"
	"time"

	"github.com/dop251/goja"
	"github.com/issueye/lichee-js/lib"
	"github.com/redis/go-redis/v9"
)

func String(rt *goja.Runtime, o *goja.Object, rdb *redis.Client) {

	// SET
	o.Set("set", func(call goja.FunctionCall) goja.Value {
		k := call.Argument(0).String()
		v := call.Argument(1).String()
		t := call.Argument(2).Export().(int64)

		ctx := context.Background()
		sc := rdb.Set(ctx, k, v, time.Duration(t*int64(time.Second)))
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// Get
	o.Set("get", func(call goja.FunctionCall) goja.Value {
		k := call.Argument(0).String()

		ctx := context.Background()
		sc := rdb.Get(ctx, k)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// GETRANGE
	o.Set("getRange", func(call goja.FunctionCall) goja.Value {
		k := call.Argument(0).String()
		start := call.Argument(1).Export().(int64)
		end := call.Argument(1).Export().(int64)

		ctx := context.Background()
		sc := rdb.GetRange(ctx, k, start, end)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// GETSET
	o.Set("getSet", func(call goja.FunctionCall) goja.Value {
		k := call.Argument(0).String()
		value := call.Argument(1).String()

		ctx := context.Background()
		sc := rdb.GetSet(ctx, k, value)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// GETBIT
	o.Set("getBit", func(call goja.FunctionCall) goja.Value {
		k := call.Argument(0).String()
		value := call.Argument(1).Export().(int64)

		ctx := context.Background()
		sc := rdb.GetBit(ctx, k, value)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// SETBIT
	o.Set("setBit", func(call goja.FunctionCall) goja.Value {
		k := call.Argument(0).String()
		offSet := call.Argument(1).Export().(int64)
		value := call.Argument(1).Export().(int)

		ctx := context.Background()
		sc := rdb.SetBit(ctx, k, offSet, value)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// MGET
	o.Set("mGet", func(keys ...string) goja.Value {
		ctx := context.Background()
		sc := rdb.MGet(ctx, keys...)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// SETEX
	o.Set("setEx", func(key string, value string, timeOut int64) goja.Value {
		ctx := context.Background()
		sc := rdb.SetEx(ctx, key, value, time.Duration(timeOut*int64(time.Second)))
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// SETNX
	o.Set("setNx", func(key string, value string, timeOut int64) goja.Value {
		ctx := context.Background()
		sc := rdb.SetNX(ctx, key, value, time.Duration(timeOut*int64(time.Second)))
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// SETRANGE
	o.Set("setRange", func(key string, offset int64, value string) goja.Value {
		ctx := context.Background()
		sc := rdb.SetRange(ctx, key, offset, value)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// STRLEN
	o.Set("strLen", func(key string) goja.Value {
		ctx := context.Background()
		sc := rdb.StrLen(ctx, key)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// MSET
	o.Set("mSet", func(values ...any) goja.Value {
		ctx := context.Background()
		sc := rdb.MSet(ctx, values)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// MSETNX
	o.Set("mSetNx", func(values ...any) goja.Value {
		ctx := context.Background()
		sc := rdb.MSetNX(ctx, values)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// INCR
	o.Set("incr", func(key string) goja.Value {
		ctx := context.Background()
		sc := rdb.Incr(ctx, key)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// INCRBY
	o.Set("incrBy", func(key string, value int64) goja.Value {
		ctx := context.Background()
		sc := rdb.IncrBy(ctx, key, value)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// INCRBYFLOAT
	o.Set("incrByFloat", func(key string, value float64) goja.Value {
		ctx := context.Background()
		sc := rdb.IncrByFloat(ctx, key, value)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// DECR
	o.Set("decr", func(key string) goja.Value {
		ctx := context.Background()
		sc := rdb.Decr(ctx, key)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// DECRBY
	o.Set("decrBy", func(key string, decrement int64) goja.Value {
		ctx := context.Background()
		sc := rdb.DecrBy(ctx, key, decrement)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// APPEND
	o.Set("append", func(key string, value string) goja.Value {
		ctx := context.Background()
		sc := rdb.Append(ctx, key, value)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})
}
