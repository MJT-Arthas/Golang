package service

import (
	"github.com/gin-gonic/gin"
	"textProject/src/common/constant"
	"textProject/src/common/response"
	"textProject/src/domain/vo"
	"textProject/src/repository"
)

type IUserService interface {
	FindUserInfo(*gin.Context, *int64)
}

type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserService(repository repository.IUserRepository) IUserService {
	return &UserService{repository}
}

func (us UserService) FindUserInfo(ctx *gin.Context, userId *int64) {
	user, _ := us.userRepository.FindUser(userId)

	if user == nil {
		response.Success(ctx, constant.DataIsNilCode, constant.DataIsNilMsg, nil)
	} else {
		userVo := &vo.UserVO{
			UserId:   user.UserId,
			UserName: user.UserName,
			Age:      user.Age,
			Gender:   user.Gender,
		}

		response.Success(ctx, constant.SelectSuccessCode, constant.SelectSuccessMsg, userVo)
	}
}
