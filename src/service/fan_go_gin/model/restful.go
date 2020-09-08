package model

type FResponse struct {
	Success bool        `json:"success"`
	Error   FErrorMsg   `json:"error"`
	Result  interface{} `json:"result"`
}

type FErrorMsg struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const (
	CodeSystemError = 1001
	CodeParamsError = 1002
)
