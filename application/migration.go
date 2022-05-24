package application

import (
	"github.com/enablefzm/mserver/application/userapps"
	"github.com/enablefzm/mserver/dbs"
	"github.com/enablefzm/mserver/dtos/userdtos"
	"github.com/enablefzm/mserver/model/customer"
	"github.com/enablefzm/mserver/model/user"
)

func Migration() error {
	err := dbs.ObDB.AutoMigrate(
		&user.User{},
		&user.Department{},
		&user.Role{},
		&user.Company{},
		&customer.ServiceUnit{},
	)
	// 创建root角色对象，如果存在则不在创建
	err = userapps.AppUser.Create(userdtos.CreateUpdateUserDto{
		UID:             "admin",
		Name:            "admin",
		PassWord:        "admin",
		IdentityCard:    "350322198002170001",
		PhoneNumber:     "18150160101",
		TelephoneNumber: "5080796",
		IsOnJob:         true,
		Sex:             true,
	})
	return err
}
