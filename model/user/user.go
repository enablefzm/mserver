package user

import (
	"time"

	"github.com/enablefzm/mserver/model"
)

type User struct {
	model.BaseModel
	UID             string    `gorm:"column:uid;uniqueIndex;type:varchar(255);NOT NULL"`
	Name            string    `gorm:"column:name;size:100;NOT NULL"`
	PassWord        string    `gorm:"size:255"`
	Sex             bool      `gorm:"column:sex"`
	IdentityCard    string    `gorm:"size:255"`
	PhoneNumber     string    `gorm:"type:varchar(100)"`
	TelephoneNumber string    `gorm:"size:255"`
	IsOnJob         bool      `gorm:"column:is_on_job"`
	DepartmentId    string    `gorm:"size:191"`
	RoleId          string    `gorm:"size:191"`
	QuitTime        time.Time `gorm:"autoCreateTime"`
	JoinTime        time.Time `gorm:"autoCreateTime"`
}

func (pUser *User) TableName() string {
	return "app_user"
}
