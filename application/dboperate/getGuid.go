package dboperate

import (
	"github.com/enablefzm/mserver/dbs"

	"gorm.io/gorm"
)

// 通过Guid值来查找对象
func GetGuid(guid string, obModel interface{}) *gorm.DB {
	tx := dbs.ObDB.Where("guid = ?", guid).First(obModel)
	return tx
}
