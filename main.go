package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/you/myapp-backend/config"
	"github.com/you/myapp-backend/handler"
	"github.com/you/myapp-backend/infra"
	"github.com/you/myapp-backend/repositories"
	"github.com/you/myapp-backend/usecase"
)

func main() {
	// 設定の読み込み
	config.LoadConfig()

	// データベース接続の初期化
	db, err := infra.NewDB(config.AppConfig.Database.GetDSN())
	if err != nil {
		log.Fatalf("データベース接続の初期化に失敗しました: %v", err)
	}

	// ルーターの設定
	router := gin.Default()

	// ルートの設定
	router.GET("/api/health", handler.HealthCheckHandler)
	router.GET("/api/dbhealth", handler.DbHealthCheckHandler(db))

	userHandler := handler.NewUserHandler(usecase.NewUserUsecase(repositories.NewUserRepository(db)))
	router.GET("/api/users/:id", userHandler.GetUserById)

	// サーバーの起動
	port := config.AppConfig.Server.Port
	router.Run(":" + port)
}
