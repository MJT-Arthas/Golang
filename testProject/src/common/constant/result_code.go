package constant

type ResponseCode int
type ResponseMsg string

const (
	SelectSuccessCode ResponseCode = 2005
	SelectSuccessMsg  ResponseMsg  = "查询成功"
)

const (
	SelectFailureCode ResponseCode = 500
	SelectFailureMsg  ResponseMsg  = "卒了"
)
