package userdtos

import "strings"

type CreateUpdateRoleDto struct {
	Name     string `json:"name"`
	Describe string `json:"describe"`
}

func (p *CreateUpdateRoleDto) GetName() string {
	return strings.ReplaceAll(p.Name, " ", "")
}
