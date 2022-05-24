package casbin

type CasbinModel struct {
	PType string `json:"p_type" gorm:"column:p_type" description:"策略类型"`
	RoleId string `json:"role_id" gorm:"column:v0" description:"角色ID"`
	Path string `json:"path" gorm:"column:v1" description:"api路径"`
	Method string `json:"method" gorm:"method" description:"访问方法"`
}

func (c *CasbinModel) TableName() string {
	return "casbin_rule"
}
