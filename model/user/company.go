package user

import "github.com/enablefzm/mserver/model"

type Company struct {
	model.BaseModel
	Name    string `gorm:"size:255" json:"name"`
	Tax     string `gorm:"size:100;not null" json:"tax"`
	Bank    string `gorm:"size:300" json:"bank"`
	Address string `gorm:"size:300" json:"address"`
}

func (p *Company) TableName() string {
	return "app_company"
}
