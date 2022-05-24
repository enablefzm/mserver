package model

import (
	"time"

	"github.com/enablefzm/gotools/guid"
	"gorm.io/gorm"
)

func NewSoftDelModel() SoftDelModel {
	return SoftDelModel{}
}

func NewSoftDelModelOnVal(val string) SoftDelModel {
	return SoftDelModel{
		Guid: val,
	}
}

func NewSoftDelModelOnCreateGuid() SoftDelModel {
	return NewSoftDelModelOnVal(guid.NewString())
}

type SoftDelModel struct {
	Guid      string         `gorm:"column:guid;<-:create;primaryKey;uniqueIndex" json:"guid"` // domainID 由GUID创建
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
