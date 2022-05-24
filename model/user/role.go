package user

import "github.com/enablefzm/mserver/model"

type Role struct {
	model.BaseModel
	Name     string `gorm:"size:255;uniqueIndex;NOT NULL" json:"name"`
	Describe string `gorm:"size:255" json:"describe"`
}

func (pRole *Role) TableName() string {
	return "app_user_role"
}
