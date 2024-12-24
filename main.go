package main

import (
	"PengLink-Back-1/router"

	"github.com/gin-gonic/gin"
	// ...existing code...
)

func main() {
	r := gin.Default()
	router.SetupRouter(r)
	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
