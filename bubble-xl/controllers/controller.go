package controllers

import (
	"bubble-xl/modles"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
 url -->  route  --> controller  --> logic   -->    model
请求来了 路由到 -->  控制器      --> 业务逻辑  --> 模型层的增删改查
*/

func IndexHandler(c *gin.Context)  {
	c.HTML(http.StatusOK,"index.html",nil)
}

func TodoCreate(c *gin.Context)  {
	var todo modles.Todo
	err := c.BindJSON(&todo)
	if err != nil {
		fmt.Printf("bind todo modules failed %v\n", err)
		return
	}
	err = modles.TodoCreate(&todo)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"err" : err.Error(),
		})
	}else {
		c.JSON(http.StatusOK,todo)
	}


}

func TodoDelete(c *gin.Context)  {
	id, ok := c.Params.Get("id")
	fmt.Println(id)
	if !ok {
		c.JSON(http.StatusOK,gin.H{
			"err": "id 无效",
		})
	}
	err := modles.TodoDelete(id)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"err": err.Error()})
	}else {
		c.JSON(http.StatusOK,gin.H{id: "deleted"})
	}
}

func TodoUpdate(c *gin.Context)  {
	id, ok := c.Params.Get("id")
	fmt.Println(id)
	if !ok {
		c.JSON(http.StatusOK,gin.H{"err":"id 无效"})
	}
	todo, err := modles.TodoOneList(id)
	fmt.Println(todo)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"err": err.Error()})
		return
	}
	c.BindJSON(&todo)
	err = modles.TodoUpdate(todo)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"err": err.Error()})
	}else {
		c.JSON(http.StatusOK,todo)
	}

}

func TodoList(c *gin.Context)  {
	todolist, err := modles.TodoAllList()
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"err": err.Error(),
		})
	}else {
		c.JSON(http.StatusOK,todolist)
	}

}