package userapps

import (
	"fmt"

	"github.com/enablefzm/mserver/application/dboperate"
	"github.com/enablefzm/mserver/dbs"
	"github.com/enablefzm/mserver/dtos/commdtos"
	"github.com/enablefzm/mserver/dtos/userdtos"
	"github.com/enablefzm/mserver/model"
	"github.com/enablefzm/mserver/model/user"
)

var AppRole = &Role{}

type Role struct {
}

// 获取角色列表
func (r *Role) List(commInput commdtos.CommInputGet) (*commdtos.CommResultPage, error) {
	// roles := casbin.Enforcer.GetAllRoles()
	// return commdtos.NewCommResultPage(int64(len(roles)), 1, roles), nil
	var roleInfos []*user.Role
	strFilter := commInput.GetFilter()
	count := dboperate.CountFilterOnModel(&user.Role{}, strFilter, "name LIKE ?", strFilter)
	qry := dbs.ObDB.Scopes(dboperate.Pginate(commInput))
	if len(strFilter) > 0 {
		qry.Where("name LIKE ? ", strFilter)
	}
	result := qry.Order("name asc").Find(&roleInfos)
	return commdtos.NewCommResultPage(count, commInput.Page, roleInfos), result.Error
}

// 创建角色
func (r *Role) Create(input userdtos.CreateUpdateRoleDto) error {
	inputName := input.GetName()
	if len(inputName) < 1 {
		return fmt.Errorf("请输入角色名称")
	}
	if err := dboperate.CheckRepeat(inputName, &user.Role{Name: inputName}); err != nil {
		return err
	}
	pRole := &user.Role{
		BaseModel: model.NewModelBaseOnCreateGuid(),
		Name:      inputName,
		Describe:  input.Describe,
	}

	// 执行事务
	// dbs.ObDB.Transaction(func(tx *gorm.DB) error {
	// 	// 创建角色对象
	// 	err := tx.Create(pRole).Error
	// 	if err != nil {
	// 		return err
	// 	}
	// 	// 判断是否可以添加到Casbin里
	// 	return nil
	// })

	tx := dbs.ObDB.Create(pRole)
	return tx.Error
}

// 获取角色可以分配的所有权限列表
func (r *Role) ListPermissionInfos() []*dboperate.PermissionGroup {
	return dboperate.GetPermissionInfos()
}
