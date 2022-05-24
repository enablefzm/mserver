package customerapps

import (
	"github.com/enablefzm/mserver/application/dboperate"
	"github.com/enablefzm/mserver/dbs"
	"github.com/enablefzm/mserver/dtos/commdtos"
	customerdtos "github.com/enablefzm/mserver/dtos/customerDtos"
	"github.com/enablefzm/mserver/model"
	"github.com/enablefzm/mserver/model/customer"
)

var AppServiceUnit = &ServiceUnit{}

type ServiceUnit struct {
}

// 创建服务单位信息
func (su *ServiceUnit) Create(input customerdtos.CreateUpdateCustomerDto) error {
	// 判断是否有同名存在
	if err := dboperate.CheckRepeat(input.Name, &customer.ServiceUnit{Name: input.Name}); err != nil {
		return err
	}
	pInfo := &customer.ServiceUnit{
		SoftDelModel: model.NewSoftDelModelOnCreateGuid(),
		Name:         input.Name,
		Address:      input.Address,
		Contacts:     input.Contacts,
		PhoneNumber:  input.PhoneNumber,
	}
	tx := dbs.ObDB.Create(pInfo)
	return tx.Error
}

func (su *ServiceUnit) Update(guid string, input customerdtos.CreateUpdateCustomerDto) error {
	// 获取数据库里是否有这个对象
	var obServiceUnit customer.ServiceUnit
	tx := dboperate.GetGuid(guid, &obServiceUnit)
	if tx.Error != nil {
		return tx.Error
	}
	obServiceUnit.Name = input.Name
	obServiceUnit.Address = input.Address
	obServiceUnit.Contacts = input.Contacts
	obServiceUnit.PhoneNumber = input.PhoneNumber
	// 更新数据
	tx = dbs.ObDB.Model(&obServiceUnit).Updates(&obServiceUnit)
	return tx.Error
}

func (su *ServiceUnit) List(input commdtos.CommInputGet) (*commdtos.CommResultPage, error) {
	var serviceUnits []*customer.ServiceUnit
	return dboperate.ListAndCountInModel(
		input,
		serviceUnits,
		&customer.ServiceUnit{},
		"name asc",
		"name LIKE ? OR address LIKE ? OR phone_number LIKE ?", input.GetFilter(), input.GetFilter(), input.GetFilter())
}
