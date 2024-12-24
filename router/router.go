package router

import (
	"PengLink-Back-1/internal/auth"
	"PengLink-Back-1/internal/user"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.POST("/register", user.RegisterHandler)
	r.POST("/login", user.LoginHandler)

	authorized := r.Group("/")
	authorized.Use(auth.AuthMiddleware("admin_level_2"))
	{
		// 管理员操作界面
	}
}
