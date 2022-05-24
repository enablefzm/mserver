package commdtos

func NewCommResultPage(total int64, page int, items interface{}) *CommResultPage {
	return &CommResultPage{
		Total: total,
		Page:  page,
		Items: items,
	}
}

type CommResultPage struct {
	Total int64       `json:"total"`
	Page  int         `json:"page"`
	Items interface{} `json:"items"`
}
