package controller

import (
	"Go-000/Week02/service"
	"fmt"
)

func FindUserNameByID() {
	svc := service.NewService()
	name, err := svc.FindUserNameByID(1)
	// 这里如果报错 ，进行相关日志记录
	if err != nil {
		fmt.Printf("FindUserNameByID failed %+v", err)
	}
	fmt.Println(name)
}
