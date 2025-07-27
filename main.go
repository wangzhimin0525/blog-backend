package main

import (
	"blog-backend/config"
	"blog-backend/controllers"
	"blog-backend/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	// 连接数据库
	config.ConnectDB()

	// 初始化Gin路由
	r := gin.Default()

	// 创建控制器实例
	authController := controllers.AuthController{}
	postController := controllers.PostController{}
	commentController := controllers.CommentController{}

	// 路由组
	api := r.Group("/api")
	{
		// 认证路由
		auth := api.Group("/auth")
		{
			auth.POST("/register", authController.Register)
			auth.POST("/login", authController.Login)
		}

		// 文章路由
		posts := api.Group("/posts")
		{
			posts.POST("/", middlewares.AuthMiddleware(), postController.CreatePost)
			posts.GET("/", postController.GetPosts)
			posts.GET("/:id", postController.GetPost)
			posts.PUT("/:id", middlewares.AuthMiddleware(), postController.UpdatePost)
			posts.DELETE("/:id", middlewares.AuthMiddleware(), postController.DeletePost)

			// 评论路由（嵌套）
			posts.POST("/:id/comments", middlewares.AuthMiddleware(), commentController.CreateComment)
			posts.GET("/:id/comments", commentController.GetComments)
		}
	}

	// 启动服务器
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
