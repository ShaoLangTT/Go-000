package service

import (
	"Go-000/Week02/dao"
	"github.com/pkg/errors"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) FindUserNameByID(id int) (string, error) {
	name, err := dao.NewUser().FindUserNameByID(id)
	//当 err == dao.UserNotFind 时，进行相关降级处理
	if errors.Is(err, dao.UserNotFind) {
		name = "小明"
		err = nil
	}
	return name, err
}
