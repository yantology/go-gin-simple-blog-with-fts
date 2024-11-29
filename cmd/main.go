package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yantology/go-gin-simple-blog-with-fts/internal/config"
	"github.com/yantology/go-gin-simple-blog-with-fts/internal/handlers"
	"github.com/yantology/go-gin-simple-blog-with-fts/internal/middleware"
	"github.com/yantology/go-gin-simple-blog-with-fts/internal/services"
	"github.com/yantology/go-gin-simple-blog-with-fts/internal/utils"
	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/repositories/articlesrepository"
	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/repositories/authrepository"
	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/services/articlesservices/postgresarticlesservices"
	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/services/authservices/postgresauthservices"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/yantology/go-gin-simple-blog-with-fts/docs"
)

// @title			Simple Blog with FTS API
// @version		1.0
// @description	This is a Simple Blog with FTS API.
// @termsOfService	http://swagger.io/terms/

// @contact.name	Wijayanto
// @contact.url	http://www.yantology.dev
// @contact.email	collab@yantology.dev

// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:5555
// @BasePath	/api/v1

// @securityDefinitions.basic	BasicAuth
func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize configurations
	config.InitConfig()

	// Initialize repositories, services, and handlers
	postgresAuthService := postgresauthservices.NewPostgresAuthService(config.DB())
	authRepo := authrepository.NewAuthRepository(postgresAuthService)
	jwtUtil := utils.NewJWTUtil(config.JWT_ACCESS_SECRET(), config.JWT_REFRESH_SECRET(), config.JWT_ACCESS_TIMEOUT(), config.JWT_REFRESH_TIMEOUT())
	authService := services.NewAuthService(authRepo, jwtUtil)
	authHandler := handlers.NewAuthHandler(authService)

	postgresArticlesService := postgresarticlesservices.NewPostgresArticlesService(config.DB())
	articlesRepo := articlesrepository.NewArticlesRepository(postgresArticlesService)
	articlesService := services.NewArticlesService(articlesRepo)
	articlesHandler := handlers.NewArticlesHandler(articlesService)

	// Initialize Gin router
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	router.Static(config.PUBLIC_ROUTE(), config.PUBLIC_ASSETS_DIR())

	v1 := router.Group("/api/v1")
	{
		v1.POST("/register", authHandler.Register)

		v1.POST("/login", authHandler.Login)

		v1.POST("/refresh-token", authHandler.RefreshToken)

		v1.GET("/articles", articlesHandler.GetAllArticles)

		v1.GET("/articles/:id", articlesHandler.GetArticleByID)

		v1.GET("/users/:id/articles", articlesHandler.GetArticlesByUserID)

		v1.GET("/articles/search", articlesHandler.SearchArticles)

		// Protected Routes - Require Authorization Header
		authMiddleware := middleware.AuthMiddleware(jwtUtil)
		protected := v1.Group("/")
		protected.Use(authMiddleware)
		{

			protected.POST("/articles", articlesHandler.CreateArticle)

			protected.POST("/articles/csv", articlesHandler.CreateArticlesWithCsv)

			protected.PUT("/articles/:id", articlesHandler.UpdateArticle)

			protected.DELETE("/articles/:id", articlesHandler.DeleteArticleByID)

			protected.POST("/change-password", authHandler.ChangePassword)

			protected.POST("/check-username", authHandler.CheckUsernameExists)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Fatal(router.Run(config.PORT()))
}
