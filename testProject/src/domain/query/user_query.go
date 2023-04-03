package query

type FindUserQuery struct {
	UserId int64 `json:"userId" description:"用户id"`
}

type StructA struct {
	UserId int64 `form:"userId"`
}
