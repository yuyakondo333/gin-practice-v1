package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/you/myapp-backend/infra"
)

func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func DbHealthCheckHandler(c *gin.Context) {
	// データベース接続の確認
	db := infra.GetDB()
	sqlDB, err := db.DB()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":  "error",
			"message": "データベース接続の取得に失敗しました",
		})
		return
	}

	if err := sqlDB.Ping(); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":  "error",
			"message": "データベース接続に失敗しました",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "ok",
		"message":  "サービスは正常に動作しています",
		"database": "connected",
	})
}
