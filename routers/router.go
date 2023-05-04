package Routers

import (
	"blog-go/Controllers/blogs/controlAdmin"
	"blog-go/Controllers/blogs/controlClient"
	"blog-go/Middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RouterInit() {
	router := gin.Default()
	//client展示
	v1 := router.Group("/client")
	{
		v1.POST("hello", controlClient.Hello)
		//group
		v1.POST("/create_group", controlAdmin.CreateGroup)
		v1.POST("/update_group", controlAdmin.UpdateGroup)
		v1.POST("/select_group", controlAdmin.SelectGroup)
		//blog
		v1.POST("/create_blog", controlAdmin.CreateBlog)
		v1.POST("/update_blog", controlAdmin.UpdateBlog)
		v1.POST("/select_blog", controlAdmin.SelectBlog)
	}
	//admin登陆
	admin1 := router.Group("/admin")
	{
		admin1.POST("/login_in", controlAdmin.Login)
	}
	//admin内部操作
	admin2 := router.Group("/admin", Middlewares.JWTAuthMiddleware())
	{
		admin2.POST("/login_out", controlAdmin.LoginOut)
		//group
		admin2.POST("/create_group", controlAdmin.CreateGroup)
		admin2.POST("/update_group", controlAdmin.UpdateGroup)
		admin2.POST("/select_group", controlAdmin.SelectGroup)
		//blog
		admin2.POST("/create_blog", controlAdmin.CreateBlog)
		admin2.POST("/update_blog", controlAdmin.UpdateBlog)
		admin2.POST("/select_blog", controlAdmin.SelectBlog)
	}
	if err := router.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
