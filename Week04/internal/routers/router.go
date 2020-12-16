package routers

import (
	"Go-000/Week04/global"
	v1 "Go-000/Week04/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	}
	tag := v1.NewTag()
	apiv1 := r.Group("/api/v1")
	apiv1.Use()
	{
		// 获取标签列表
		apiv1.GET("/tags", tag.List)
	}

	return r
}
