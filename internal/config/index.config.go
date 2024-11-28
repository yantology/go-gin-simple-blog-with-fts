package config

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/yantology/go-gin-simple-blog-with-fts/internal/config/appconfig"
	"github.com/yantology/go-gin-simple-blog-with-fts/internal/config/corsconfig"
	"github.com/yantology/go-gin-simple-blog-with-fts/internal/config/dbconfig"
	"github.com/yantology/go-gin-simple-blog-with-fts/internal/config/jwtconfig"
)

func InitConfig() {
	appconfig.InitAppConfig()
	dbconfig.InitDatabaseConfig()
	dbconfig.ConnectDatabase(sql.Open)
	jwtconfig.InitJWTConfig()
}

// variable appconfig
// var (
// 	PORT             = appconfig.PORT
// 	PUBLIC_ROUTE      = appconfig.PUBLIC_ROUTE
// 	PUBLIC_ASSETS_DIR = &appconfig.PUBLIC_ASSETS_DIR
// )

func PORT() string {
	return appconfig.PORT
}

func PUBLIC_ROUTE() string {
	return appconfig.PUBLIC_ROUTE
}

func PUBLIC_ASSETS_DIR() string {
	return appconfig.PUBLIC_ASSETS_DIR
}

// variable dbconfig

func DB_DRIVER() string {
	return dbconfig.DB_DRIVER
}

func DB_USER() string {
	return dbconfig.DB_USER
}

func DB_PASSWORD() string {
	return dbconfig.DB_PASSWORD
}

func DB_NAME() string {
	return dbconfig.DB_NAME
}

func DB_HOST() string {
	return dbconfig.DB_HOST
}

func DB_PORT() string {
	return dbconfig.DB_PORT
}

func DB() *sql.DB {
	return dbconfig.DB
}

// variable corsconfig
func CORS_ALLOW_ORIGINS() gin.HandlerFunc {
	return corsconfig.CorsConfig()
}

// variable jwtconfig
func JWT_ACCESS_SECRET() string {
	return jwtconfig.JWT_ACCESS_SECRET
}

func JWT_REFRESH_SECRET() string {
	return jwtconfig.JWT_REFRESH_SECRET
}

func JWT_ACCESS_TIMEOUT() int {
	return jwtconfig.JWT_ACCESS_TIMEOUT
}

func JWT_REFRESH_TIMEOUT() int {
	return jwtconfig.JWT_REFRESH_TIMEOUT
}
