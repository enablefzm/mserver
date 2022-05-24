package userdtos

type CreateUpdateCompanyDto struct {
	Name    string `json:"name"`
	Tax     string `json:"tax"`
	Bank    string `json:"bank"`
	Address string `json:"address"`
}
