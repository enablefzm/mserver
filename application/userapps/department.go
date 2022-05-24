package userapps

import (
	"github.com/enablefzm/mserver/application/dboperate"
	"github.com/enablefzm/mserver/dbs"
	"github.com/enablefzm/mserver/dtos/baseDtos"
	"github.com/enablefzm/mserver/dtos/commdtos"
	"github.com/enablefzm/mserver/dtos/userdtos"
	"github.com/enablefzm/mserver/model"
	"github.com/enablefzm/mserver/model/user"

	"github.com/enablefzm/gotools/guid"
)

var AppDepartment = &Department{}

type Department struct {
}

// 创建部门
func (dp *Department) Create(input userdtos.CreateUpdateDepartmentDto) error {
	if err := dboperate.CheckRepeat(
		input.GetName(),
		&user.Department{Name: input.GetName()}); err != nil {
		return err
	}
	obDepartment := &user.Department{
		BaseModel: model.NewModelBaseOnVal(guid.NewString()),
		Name:      input.GetName(),
		Manager:   input.Manager,
		Memo:      input.Memo,
	}
	// 写入数据
	tx := dbs.ObDB.Create(obDepartment)
	return tx.Error
}

func (dp *Department) ListInfo(commInput commdtos.CommInputGet) (*commdtos.CommResultPage, error) {
	var infos []*user.Department
	qry := dbs.ObDB.Scopes(dboperate.Pginate(commInput))
	strFilter := commInput.GetFilter()
	if len(strFilter) > 0 {
		qry.Where("name LIKE ?", strFilter)
	}
	result := qry.Order("name asc").Find(&infos)

	// 计算总数
	count := dboperate.CountFilterOnModel(&user.Department{}, strFilter, "name LIKE ?", strFilter)

	infosDto := make([]userdtos.DepartmentDto, 0, len(infos))
	for _, ob := range infos {
		infosDto = append(infosDto, dp.Convert(ob))
	}
	return commdtos.NewCommResultPage(count, commInput.Page, infosDto), result.Error
}

func (dp *Department) Convert(obDepart *user.Department) userdtos.DepartmentDto {
	return userdtos.DepartmentDto{
		BaseDto: baseDtos.BaseDto{
			Guid:      obDepart.Guid,
			CreatedAt: obDepart.CreatedAt,
			UpdatedAt: obDepart.UpdatedAt,
		},
		Name:    obDepart.Name,
		Manager: obDepart.Manager,
		Memo:    obDepart.Memo,
	}
}
