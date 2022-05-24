package commdtos

func NewCommResult(strMsg string) CommResult {
	return NewCommCodeResult(0, strMsg)
}

func NewCommCodeResult(code int, strMsg string) CommResult {
	return CommResult{
		Code:        code,
		Description: strMsg,
	}
}

type CommResult struct {
	Code int `json:"code"`
	Description string `json:"description"`
}
