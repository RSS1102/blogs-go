package controlAdmin

import (
	"blog-go/Models/modelAdmin"
	"blog-go/Models/modelPublic"
	"blog-go/Services/blogs/servicesAdmin"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateBlog(context *gin.Context) {
	var data modelAdmin.BlogBlogs
	err := context.ShouldBindJSON(&data)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}
	println("ContentMd")
	println(data.ContentMd)
	res := servicesAdmin.CreateBlog(data.GroupId, data.Title, data.ContentMd, data.ContentHtml)
	if res > 0 {
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "添加内容成功",
		})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "添加内容失败",
		})
	}
}

func UpdateBlog(context *gin.Context) {
	var data modelAdmin.BlogBlogs
	err := context.ShouldBindJSON(&data)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	res := servicesAdmin.UpdateBlog(data)
	if res > 0 {
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "blog更新成功",
			"data":    make([]int, 0),
			"total":   0,
		})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "blog更新失败",
			"data":    make([]int, 0),
			"total":   0,
		})
	}
}

func SelectBlog(context *gin.Context) {
	type Data struct {
		ID int `json:"id"`
	}
	var data Data
	err := context.ShouldBindJSON(&data)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
			"data":    make([]int, 0),
			"total":   0,
		})
		return
	}

	blogs, err := servicesAdmin.SelectBlog(data.ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "查询失败",
			"data":    blogs,
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "查询成功",
			"data":    blogs,
		})
	}
}

func SelectBlogLimit(context *gin.Context) {
	var page modelPublic.Page
	err := context.ShouldBindJSON(&page)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
			"data":    make([]int, 0),
			"total":   0,
		})
		return
	}
	println(page.Current, page.PageSize)
	total, blogs := servicesAdmin.SelectBlogLimit(page.Current, page.PageSize)
	if total > 0 {
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "查询成功",
			"data":    blogs,
			"total":   total,
		})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "查询失败",
			"data":    make([]int, 0),
			"total":   0,
		})
	}
}
