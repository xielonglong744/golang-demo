package router

import (
	"bubble-xl/controllers"
	"bubble-xl/settings"
	"github.com/gin-gonic/gin"
)

func Init(cfg *settings.GinConf) *gin.Engine {
	if cfg.AppConf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Static("/static","static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/",controllers.IndexHandler)
	v1Group := r.Group("v1")
	{
		// 代办事项
		// 1. 添加代办事项
		v1Group.POST("/todo",controllers.TodoCreate)
		// 2. 删除代办事项
		v1Group.DELETE("/todo/:id",controllers.TodoDelete)
		// 3. 修改代办事项
		v1Group.PUT("/todo/:id",controllers.TodoUpdate)
		// 4. 查看代办事项
		v1Group.GET("/todo", controllers.TodoList)
	}
	return r
}
