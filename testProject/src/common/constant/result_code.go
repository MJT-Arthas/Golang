package constant

type ResponseCode int
type ResponseMsg string

const (
	SelectSuccessCode ResponseCode = 2005
	SelectSuccessMsg  ResponseMsg  = "查询成功"
)

const (
	DataIsNilCode ResponseCode = 500
	DataIsNilMsg  ResponseMsg  = "卒了"
)
