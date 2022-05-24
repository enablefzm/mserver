package dboperate

import (
	"github.com/enablefzm/mserver/dtos/commdtos"

	"gorm.io/gorm"
)

// 分页处理
func Pginate(get commdtos.CommInputGet) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if get.Page == 0 {
			get.Page = 1
		}
		switch {
		case get.PageSize > 100:
			get.PageSize = 100
		case get.PageSize <= 0:
			get.PageSize = 1
		}
		offset := (get.Page - 1) * get.PageSize
		return db.Offset(offset).Limit(get.PageSize)
	}
}
