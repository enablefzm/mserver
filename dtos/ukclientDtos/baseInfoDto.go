package ukclientdtos

// 客户端传上来的基本信息
type BaseInfoDto struct {
	IpAddress    string               `json:"ip_address"`
	Version      string               `json:"version"`
	CompanyInfos []BaseCompanyInfoDto `json:"company_infos"`
}

// 客户端传过来拥有的公司信息列表
type BaseCompanyInfoDto struct {
	Name      string `json:"name"`
	Tax       string `json:"tax"`
	IsActivat bool   `json:"is_activat"`
}
