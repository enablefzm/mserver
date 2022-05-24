package userapps

import (
	"log"

	"github.com/enablefzm/mserver/application/dboperate"
	"github.com/enablefzm/mserver/dbs"
	"github.com/enablefzm/mserver/dtos/commdtos"
	"github.com/enablefzm/mserver/dtos/userdtos"
	"github.com/enablefzm/mserver/model"
	"github.com/enablefzm/mserver/model/user"
)

var AppCompany = &Company{}

type Company struct {
}

// 获取列表信息
func (c *Company) List(input commdtos.CommInputGet) (*commdtos.CommResultPage, error) {
	var companys []*user.Company
	// strFilter := input.GetFilter()
	// count := dboperate.CountFilterOnModel(user.Company{}, strFilter, "name LIKE ?", strFilter)
	// qry := dbs.ObDB.Scopes(dboperate.Pginate(input))
	// if len(strFilter) > 0 {
	// 	qry.Where("name LIKE ?", strFilter)
	// }
	// result := qry.Order("name asc").Find(&companys)
	// return commdtos.NewCommResultPage(count, input.Page, companys), result.Error
	return dboperate.ListAndCountInModel(
		input,
		companys,
		&user.Company{},
		"name asc",
		"name LIKE ? OR address LIKE ?", input.GetFilter(), input.GetFilter())
}

// 创建新本公司信息
func (c *Company) Create(input userdtos.CreateUpdateCompanyDto) error {
	if err := dboperate.CheckRepeat(input.Tax, &user.Company{Tax: input.Tax}); err != nil {
		return err
	}
	ob := &user.Company{
		BaseModel: model.NewModelBaseOnCreateGuid(),
		Name:      input.Name,
		Tax:       input.Tax,
		Bank:      input.Bank,
		Address:   input.Address,
	}
	tx := dbs.ObDB.Create(ob)
	return tx.Error
}

// 修改保存
func (c *Company) Update(guid string, input userdtos.CreateUpdateCompanyDto) error {
	var obCompany user.Company
	tx := dboperate.GetGuid(guid, &obCompany)
	if tx.Error != nil {
		log.Println("获取对象出错:", tx.Error.Error())
		return tx.Error
	}
	obCompany.Address = input.Address
	obCompany.Bank = input.Bank
	obCompany.Name = input.Name
	obCompany.Tax = input.Tax
	tx = dbs.ObDB.Model(&obCompany).Updates(&obCompany)
	if tx.Error != nil {
		log.Println("执行更新数据对象时出错")
	}
	return tx.Error
}

// 删除指定的公司信息
func (c *Company) Delete(guid string) error {
	tx := dbs.ObDB.Delete(&user.Company{
		BaseModel: model.BaseModel{
			Guid: guid,
		},
	})
	return tx.Error
}
