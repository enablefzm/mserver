package dbs

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ObDB *gorm.DB

func LinkDb(c *Cfg) error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		c.User,
		c.Password,
		c.IpAddress,
		c.Port,
		c.DbName)
	ObDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDb, errSet := ObDB.DB()
	if errSet != nil {
		return errSet
	}
	// sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	return nil
}
