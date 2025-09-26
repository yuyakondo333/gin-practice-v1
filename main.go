package main

import (
	"github.com/gin-gonic/gin"
	"github.com/you/myapp-backend/config"
	"github.com/you/myapp-backend/handler"
	"github.com/you/myapp-backend/infra"
)

func main() {
	// 設定の読み込み
	config.LoadConfig()

	// データベース接続の初期化
	infra.InitDB()
	defer infra.CloseDB()

	// ルーターの設定
	router := gin.Default()

	// ルートの設定
	router.GET("/api/health", handler.HealthCheckHandler)
	router.GET("/api/dbhealth", handler.DbHealthCheckHandler)

	// サーバーの起動
	port := config.AppConfig.Server.Port
	router.Run(":" + port)
}
