package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/you/myapp-backend/infra"
)

func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func DbHealthCheckHandler(db *infra.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sqlDB, err := db.DB.DB()
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"message": "データベース接続に失敗しました"})
			return
		}
		if err := sqlDB.Ping(); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"message": "データベース接続に失敗しました"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
