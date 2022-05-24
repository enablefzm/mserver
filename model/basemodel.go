package model

import (
	"time"

	"github.com/enablefzm/gotools/guid"
)

func NewModelBase() BaseModel {
	return BaseModel{Guid: ""}
}

func NewModelBaseOnVal(strGuid string) BaseModel {
	return BaseModel{Guid: strGuid}
}

// 创建带有GUID值的Model基类
func NewModelBaseOnCreateGuid() BaseModel {
	return NewModelBaseOnVal(guid.NewString())
}

type BaseModel struct {
	Guid      string    `gorm:"column:guid;<-:create;primaryKey;uniqueIndex" json:"guid"` // domainID 由GUID创建
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
