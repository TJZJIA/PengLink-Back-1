package main

import (
	"PengLink-Back-1/config"
	"PengLink-Back-1/router"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	r := gin.Default()
	router.SetupRouter(r)
	fmt.Println("Hello World!")
	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
