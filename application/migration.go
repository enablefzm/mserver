package application

import (
	"log"

	"github.com/enablefzm/mserver/application/userapps"
	"github.com/enablefzm/mserver/dbs"
	"github.com/enablefzm/mserver/dtos/userdtos"
	"github.com/enablefzm/mserver/model/customer"
	"github.com/enablefzm/mserver/model/user"
)

func Migration(dst ...interface{}) error {
	err := MigrationAdmin(&userdtos.CreateUpdateUserDto{
		UID:             "admin",
		Name:            "admin",
		PassWord:        "admin",
		IdentityCard:    "350322198002170001",
		PhoneNumber:     "18150160101",
		TelephoneNumber: "5080796",
		IsOnJob:         true,
		Sex:             true,
	}, dst...)
	return err
}

func MigrationAdmin(udb *userdtos.CreateUpdateUserDto, dst ...interface{}) error {
	err := dbs.ObDB.AutoMigrate(
		&user.User{},
		&user.Department{},
		&user.Role{},
		&user.Company{},
		&customer.ServiceUnit{},
	)
	if err != nil {
		return err
	}
	// 创建root角色对象，如果存在则不在创建
	if errCreate := userapps.AppUser.Create(*udb); errCreate != nil {
		log.Println("创建角色数据出错:", errCreate.Error())
	}
	// 创建其它表
	err = dbs.ObDB.AutoMigrate(dst...)
	return err
}
