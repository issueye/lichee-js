package redis

import (
	"context"

	"github.com/dop251/goja"
	"github.com/issueye/lichee-js/lib"
	"github.com/redis/go-redis/v9"
)

func Hash(rt *goja.Runtime, o *goja.Object, rdb *redis.Client) {
	// HDEL
	o.Set("hDel", func(key string, fields ...string) goja.Value {
		ctx := context.Background()
		sc := rdb.HDel(ctx, key, fields...)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// HEXISTS
	o.Set("hExists", func(key string, field string) goja.Value {
		ctx := context.Background()
		sc := rdb.HExists(ctx, key, field)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// HGET
	o.Set("hGet", func(key string, field string) goja.Value {
		ctx := context.Background()
		sc := rdb.HGet(ctx, key, field)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// HGETALL
	o.Set("hGetAll", func(key string, field string) goja.Value {
		ctx := context.Background()
		sc := rdb.HGetAll(ctx, key)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// HINCRBY
	o.Set("hIncrBy", func(key string, field string, incr int64) goja.Value {
		ctx := context.Background()
		sc := rdb.HIncrBy(ctx, key, field, incr)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// HINCRBYFLOAT
	o.Set("hIncrFloat", func(key string, field string, incr float64) goja.Value {
		ctx := context.Background()
		sc := rdb.HIncrByFloat(ctx, key, field, incr)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// HKEYS
	o.Set("hKeys", func(key string) goja.Value {
		ctx := context.Background()
		sc := rdb.HKeys(ctx, key)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// HLEN
	o.Set("hLen", func(key string) goja.Value {
		ctx := context.Background()
		sc := rdb.HLen(ctx, key)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// HMGET
	o.Set("hMSet", func(key string, fields ...string) goja.Value {
		ctx := context.Background()
		sc := rdb.HMGet(ctx, key, fields...)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// HMSET
	o.Set("hMSet", func(key string, values ...interface{}) goja.Value {
		ctx := context.Background()
		sc := rdb.HMSet(ctx, key, values...)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// HSET
	o.Set("hSet", func(key string, values ...interface{}) goja.Value {
		ctx := context.Background()
		sc := rdb.HSet(ctx, key, values...)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// HSETNX
	o.Set("hSetNx", func(key string, field string, value interface{}) goja.Value {
		ctx := context.Background()
		sc := rdb.HSetNX(ctx, key, field, value)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// HVALS
	o.Set("hVals", func(key string) goja.Value {
		ctx := context.Background()
		sc := rdb.HVals(ctx, key)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		return lib.MakeReturnValue(rt, sc.Val())
	})

	// HSCAN
	o.Set("hScan", func(key string, cursor uint64, match string, count int64) goja.Value {
		ctx := context.Background()
		sc := rdb.HScan(ctx, key, cursor, match, count)
		if sc.Err() != nil {
			return lib.MakeErrorValue(rt, sc.Err())
		}

		// keys []string, cursor uint64
		keys, cur := sc.Val()
		val := map[uint64][]string{
			cur: keys,
		}

		return lib.MakeReturnValue(rt, val)
	})
}
