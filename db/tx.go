package db

import (
	"github.com/dop251/goja"
	"github.com/issueye/lichee-js/lib"
	"gorm.io/gorm"
)

func NewTx(runtime *goja.Runtime, tx *gorm.DB) goja.Value {
	o := runtime.NewObject()
	o.Set("commit", func(call goja.FunctionCall) goja.Value {
		mit := tx.Commit()
		if mit.Error != nil {
			return lib.MakeErrorValue(runtime, mit.Error)
		}
		return nil
	})

	o.Set("rollback", func(call goja.FunctionCall) goja.Value {
		roll := tx.Rollback()
		if roll.Error != nil {
			return lib.MakeErrorValue(runtime, roll.Error)
		}
		return nil
	})

	o.Set("exec", func(call goja.FunctionCall) goja.Value {
		sqlStr := call.Argument(0).String()
		result := tx.Exec(sqlStr)
		if result.Error != nil {
			return lib.MakeErrorValue(runtime, result.Error)
		}

		// 获取返回
		return lib.MakeReturnValue(runtime, NewExecResult(runtime, sqlResult{
			RowsAffected: result.RowsAffected,
		}))
	})

	o.Set("query", func(call goja.FunctionCall) goja.Value {
		sqlStr := call.Argument(0).String()
		// 查询数据
		result := tx.Raw(sqlStr)

		if result.Error != nil {
			return lib.MakeErrorValue(runtime, result.Error)
		}

		rows, err := result.Rows()
		if err != nil {
			return lib.MakeErrorValue(runtime, err)
		}

		// 生成数据
		data, err := MakeData(rows)
		if err != nil {
			return lib.MakeErrorValue(runtime, err)
		}

		return lib.MakeReturnValue(runtime, data)
	})

	return o
}
