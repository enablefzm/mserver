package ukclient

import (
	"github.com/enablefzm/mserver/dbs"
	"github.com/enablefzm/mserver/model"
)

type UkClient struct {
	model.SoftDelModel
	Name      string `gorm:"size(200),sort:desc"`
	Tax       string `gorm:"size(100)"`
	IpAddress string `gorm:"size(100)"`
	Version   string `gorm:"size(50)"`
	RegCode   string `gorm:"size(500)"`
	Sksbkl    string `gorm:"size(255)"`
	Zskl      string `gorm:"size(255)"`
}

// 通过税号来获取对象
func NewUkClientOnTax(tax string) (*UkClient, error) {
	var ob UkClient
	tx := dbs.ObDB.Where(map[string]interface{}{"Tax": tax}).First(&ob)
	return &ob, tx.Error
}
