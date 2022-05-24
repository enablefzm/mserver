package dboperate

import (
	"fmt"

	"github.com/enablefzm/mserver/dbs"
)

// 判断某值是否已在数据库存在,如果不存在则返回 nil 存在则返回一个错误
func CheckRepeat(keyVal string, findOb interface{}) error {
	var count int64
	result := dbs.ObDB.Model(findOb).Where(findOb).Count(&count)
	if result.Error != nil {
		return result.Error
	}
	if count > 0 {
		return fmt.Errorf("%s已经存在不能重复创建", keyVal)
	}
	return nil
	/*
		result := dbs.ObDB.Where(findOb).First(findOb)
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			if result.Error == nil {
				return fmt.Errorf("%s已经存在不能重复创建！", keyVal)
			}
			return result.Error
		}
		return nil
	*/
}
