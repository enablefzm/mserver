package userdtos

import "strings"

type CreateUpdateDepartmentDto struct {
	Name string 		`json:"name"`			// 部门名称
	Manager string 		`json:"manager"`		// 部门主管
	Memo string 		`json:"memo"`			// 部门备注信息
}

func (p *CreateUpdateDepartmentDto) GetName() string {
	return strings.Replace(strings.ToLower(p.Name), " ", "", -1)
}
