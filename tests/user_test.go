package tests

import (
	"testing"

	"github.com/enablefzm/mserver/application/userapps"
	"github.com/enablefzm/mserver/dtos/userdtos"
)

func TestCreateUser(t *testing.T) {
	err := userapps.AppUser.Create(userdtos.CreateUpdateUserDto{
		UID:             "3502003",
		Name:            "林旭",
		PassWord:        "admin",
		IdentityCard:    "350521198110193057",
		PhoneNumber:     "13950019871",
		TelephoneNumber: "5123739",
		IsOnJob:         true,
		Sex:             true,
	})
	t.Log("执行结果：", err)
}
