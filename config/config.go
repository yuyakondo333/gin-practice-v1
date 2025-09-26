package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
	Server   ServerConfig   `mapstructure:"server"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

var AppConfig *Config

func LoadConfig() {
	// .envファイルの読み込み
	if err := godotenv.Load(); err != nil {
		log.Printf(".envファイルの読み込みに失敗しました: %v", err)
	}

	// デフォルト値の設定
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.user", "postgres")
	viper.SetDefault("database.password", "postgres")
	viper.SetDefault("database.dbname", "myapp")
	viper.SetDefault("database.sslmode", "disable")
	viper.SetDefault("server.port", "8080")

	// 環境変数の読み込み
	viper.AutomaticEnv()

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("設定の解析に失敗しました: %v", err)
	}
}

func (c *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode)
}
