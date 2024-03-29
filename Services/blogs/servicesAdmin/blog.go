package servicesAdmin

import (
	Init "blog-go/Config"
	"blog-go/Models/modelAdmin"
	"log"
	"time"
)

// TempMergedBlogs 覆盖原有的createAt和updateAt的json字段
type TempMergedBlogs struct {
	*modelAdmin.MergedBlogs        // 嵌入原始结构体
	CreatedAt               string `json:"createAt"`
	UpdatedAt               string `json:"updateAt"`
}

// CreateBlog 创建内容
func CreateBlog(groupId int, title, ContentMd string, ContentHtml string) int64 {
	var newGroup = modelAdmin.BlogBlogs{
		GroupId:     groupId,
		Title:       title,
		ContentHtml: ContentHtml,
		ContentMd:   ContentMd,
		Visitors:    0,
		IsShow:      true,
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	}
	result := Init.DB.Table("blog_blogs").Create(&newGroup)
	println(result)
	if result.Error != nil {
		log.Println("Create group fail : ", result)
	}
	return result.RowsAffected
}

// UpdateBlog blog更新
func UpdateBlog(updates modelAdmin.BlogBlogs) int64 {
	// 更新记录的不固定字段
	updateData := map[string]interface{}{
		"Title":       updates.Title,
		"ContentHtml": updates.ContentHtml,
		"ContentMd":   updates.ContentMd,
		"Visitors":    updates.Visitors,
		"IsShow":      updates.IsShow,
		"GroupId":     updates.GroupId,
		"UpdateAt":    time.Now(),
	}
	println(updateData)
	// 移除未传入的字段
	if updates.Title == "" {
		delete(updateData, "Title")
	}
	if updates.ContentHtml == "" {
		delete(updateData, "ContentHtml")
	}
	if updates.ContentMd == "" {
		delete(updateData, "ContentMd")
	}
	if updates.Visitors == 0 {
		delete(updateData, "Visitors")
	}
	if updates.GroupId == 0 {
		delete(updateData, "GroupId")
	}

	result := Init.DB.Table("blog_blogs").
		Model(&modelAdmin.BlogBlogs{}).
		Where("id=?", updates.ID).
		Updates(updateData)
	if result.Error != nil {
		log.Println("blog_blogs update fail : ", result)
	}
	return result.RowsAffected
}

// SelectBlog blog查询
func SelectBlog(id int) (TempMergedBlogs, error) {
	var blogBlogs TempMergedBlogs
	err := Init.DB.Table("blog_blogs").
		Where("blog_blogs.id=?", id).
		Joins(" JOIN blog_groups ON blog_blogs.group_id = blog_groups.id").
		Select("blog_blogs.id ,blog_blogs.title,blog_blogs.content_html,blog_blogs.content_md,blog_blogs.visitors,blog_blogs.is_show, blog_blogs.create_at,blog_blogs.update_at," +
			"blog_groups.group, blog_groups.is_show as groups_is_show ").
		First(&blogBlogs)
	if err.Error != nil {
		return TempMergedBlogs{}, err.Error
		println("error")
	}
	blogBlogs.CreatedAt = blogBlogs.CreateAt.Format("2006-01-02 15:04:05")
	blogBlogs.UpdatedAt = blogBlogs.UpdateAt.Format("2006-01-02 15:04:05")
	return blogBlogs, nil
}

// SelectBlogLimit blog查询 分页
func SelectBlogLimit(Current int, pageSize int) (int, []TempMergedBlogs) {
	var blogBlogs []TempMergedBlogs
	var total int64
	err := Init.DB.Table("blog_blogs").Count(&total).Error
	if err != nil {
		log.Println("Failed to get total record count:", err)
		return 0, nil
	}

	errs := Init.DB.Table("blog_blogs").Joins("JOIN blog_groups ON blog_blogs.group_id = blog_groups.id").
		Select("blog_blogs.*, blog_groups.group, blog_groups.is_show as group_is_show").
		Limit(pageSize).Offset((Current - 1) * pageSize).Find(&blogBlogs)
	if errs.Error != nil {
		return 0, nil
	}
	for i := range blogBlogs {
		blogBlogs[i].CreatedAt = blogBlogs[i].CreateAt.Format("2006-01-02 15:04:05")
		blogBlogs[i].UpdatedAt = blogBlogs[i].UpdateAt.Format("2006-01-02 15:04:05")
	}
	return int(total), blogBlogs
}
