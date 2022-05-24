package commdtos

func NewCommError(strMsg string) CommError {
	return CommError{
		Code: -1,
		ErrorDescription: strMsg,
	}
}

type CommError struct {
	Code int `json:"code"`
	ErrorDescription string `json:"error_description"`
}
