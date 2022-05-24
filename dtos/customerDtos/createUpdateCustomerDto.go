package customerdtos

type CreateUpdateCustomerDto struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	Contacts    string `json:"contacts"`
	PhoneNumber string `json:"phone_number"`
}
