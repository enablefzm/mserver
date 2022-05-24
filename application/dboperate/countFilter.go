package dboperate

import (
	"github.com/enablefzm/mserver/dbs"

	"gorm.io/gorm"
)

// 通过数据模型来统计总数
func CountFilterOnModel(pModel interface{}, strFilter string, query interface{}, args ...interface{}) int64 {
	return CountFilter(
		dbs.ObDB.Model(pModel),
		strFilter,
		query,
		args...)
}

// 根据查询条件来统计数量
func CountFilter(tx *gorm.DB, strFilter string, query interface{}, args ...interface{}) int64 {
	var count int64
	if len(strFilter) > 0 {
		tx.Where(query, args...)
	}
	tx.Count(&count)
	return count
}
