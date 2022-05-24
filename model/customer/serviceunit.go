package customer

import "github.com/enablefzm/mserver/model"

type ServiceUnit struct {
	model.SoftDelModel
	Name        string `gorm:"size:255" json:"name"`
	Address     string `gorm:"size:255" json:"address"`
	Contacts    string `gorm:"size:100" json:"contacts"`
	PhoneNumber string `gorm:"size:100" json:"phone_number"`
}

func (su *ServiceUnit) TableName() string {
	return "app_customer_serviceunit"
}
