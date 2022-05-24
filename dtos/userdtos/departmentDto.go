package userdtos

import "github.com/enablefzm/mserver/dtos/baseDtos"

type DepartmentDto struct {
	baseDtos.BaseDto
	Name    string `json:"name"`    // 部门名称
	Manager string `json:"manager"` // 部门主管
	Memo    string `json:"memo"`    // 部门备注信息
}
