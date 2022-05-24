package user

import "github.com/enablefzm/mserver/model"

type Department struct {
	model.BaseModel
	Name    string `gorm:"size:255"`          // 部门名称
	Manager string `gorm:"type:varchar(191)"` // 部门主管
	Memo    string `gorm:"type:text"`         // 部门备注信息
}

func (pdm *Department) TableName() string {
	return "app_user_department"
}
