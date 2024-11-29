package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yantology/go-gin-simple-blog-with-fts/internal/models"
	"github.com/yantology/go-gin-simple-blog-with-fts/internal/services"
)

type ArticlesHandler struct {
	articleService *services.ArticlesService
}

func NewArticlesHandler(articleService *services.ArticlesService) *ArticlesHandler {
	return &ArticlesHandler{
		articleService: articleService,
	}
}

// CreateArticle creates a new article.
// @Summary Create a new article
// @Description Create a new article
// @Tags articles
// @Accept json
// @Accept x-www-form-urlencoded
// @Produce json
// @Param article body models.ArticleRequest false "Article Request"
// @Param title formData string false "Title"
// @Param content formData string false "Content"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Message
// @Failure 401 {object} models.Message
// @Failure 500 {object} models.Message
// @Router /articles [post]
// @Security ApiKeyAuth
func (h *ArticlesHandler) CreateArticle(c *gin.Context) {
	var req models.ArticleRequest
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.NewMessage("user not authenticated"))
		return
	}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusUnauthorized, models.NewMessage("user not authenticated"))
		return
	}

	cuserr := h.articleService.CreateArticle(userID.(int), &req)

	if cuserr != nil {
		c.JSON(cuserr.HTTPCode, models.NewMessage(cuserr.Error()))
		return
	}
	c.JSON(200, models.NewMessage("article created successfully"))
}

// CreateArticlesWithCsv godoc
// @Summary Create articles with CSV
// @Description Create multiple articles by uploading a CSV file
// @Tags articles
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "file"
// @Success 200 {object} models.Message "articles created successfully"
// @Failure 400 {object} models.Message "file upload failed"
// @Failure 401 {object} models.Message "user not authenticated"
// @Failure 500 {object} models.Message "internal server error"
// @Router /articles/csv [post]
// @Security ApiKeyAuth
func (h *ArticlesHandler) CreateArticlesWithCsv(c *gin.Context) {
	log.Println("CreateArticlesWithCsv")
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.NewMessage("user not authenticated"))
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, models.NewMessage("file upload failed"))
		return
	}

	cuserr := h.articleService.CreateArticlesWithCsv(userID.(int), file)
	if cuserr != nil {
		c.JSON(cuserr.HTTPCode, models.NewMessage(cuserr.Error()))
		return
	}

	c.JSON(200, models.NewMessage("articles created successfully"))
}

// GetArticleByID retrieves an article by its ID.
// @Summary Get an article by ID
// @Description Get an article by ID
// @Tags articles
// @Produce json
// @Param id path int true "Article ID"
// @Success 200 {object} models.ArticleResponse
// @Failure 400 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 500 {object} models.Message
// @Router /articles/{id} [get]
func (h *ArticlesHandler) GetArticleByID(c *gin.Context) {
	id := c.Param("id")
	articleID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, models.NewMessage("invalid article ID"))
		return
	}

	article, cuserr := h.articleService.GetArticleByID(articleID)
	if cuserr != nil {
		c.JSON(cuserr.HTTPCode, models.NewMessage(cuserr.Error()))
		return
	}

	c.JSON(200, article)
}

// GetArticlesByUserID retrieves all articles created by a specific user.
// @Summary Get articles by user ID
// @Description Get articles by user ID
// @Tags articles
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.ArticlesResponse
// @Failure 400 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 500 {object} models.Message
// @Router /users/{id}/articles [get]
func (h *ArticlesHandler) GetArticlesByUserID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, models.NewMessage("invalid user ID"))
		return
	}

	articles, cuserr := h.articleService.GetArticlesByUserID(userID)
	if cuserr != nil {
		c.JSON(cuserr.HTTPCode, models.NewMessage(cuserr.Error()))
		return
	}

	c.JSON(200, articles)
}

// GetAllArticles retrieves all available articles.
// @Summary Get all articles
// @Description Get all articles
// @Tags articles
// @Produce json
// @Success 200 {object} models.ArticlesResponse
// @Failure 400 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 500 {object} models.Message
// @Router /articles [get]
func (h *ArticlesHandler) GetAllArticles(c *gin.Context) {
	articles, cuserr := h.articleService.GetAllArticles()
	if cuserr != nil {
		c.JSON(cuserr.HTTPCode, models.NewMessage(cuserr.Error()))
		return
	}

	c.JSON(200, articles)
}

// SearchArticles performs full-text search on articles.
// @Summary Search articles
// @Description Search articles
// @Tags articles
// @Produce json
// @Param query query string true "Search Query"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} models.ArticleResponse
// @Failure 400 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 500 {object} models.Message
// @Router /articles/search [get]
func (h *ArticlesHandler) SearchArticles(c *gin.Context) {
	log.Println("SearchArticles start")
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(400, models.NewMessage("invalid limit"))
		return
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		c.JSON(400, models.NewMessage("invalid offset"))
		return
	}

	query := c.Query("query")
	log.Printf("SearchArticles query: %s", query)
	articles, cuserr := h.articleService.SearchArticles(limit, offset, query)
	if cuserr != nil {
		c.JSON(cuserr.HTTPCode, models.NewMessage(cuserr.Error()))
		return
	}

	c.JSON(200, articles)
}

// UpdateArticle updates an existing article.
// @Summary Update an article
// @Description Update an article
// @Tags articles
// @Accept json
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id path int true "Article ID"
// @Param article body models.ArticleRequest false "Article Request"
// @Param title formData string false "Title"
// @Param content formData string flase "Content"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Message
// @Failure 401 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 500 {object} models.Message
// @Router /articles/{id} [put]
// @Security ApiKeyAuth
func (h *ArticlesHandler) UpdateArticle(c *gin.Context) {
	id := c.Param("id")
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.NewMessage("user not authenticated"))
		return
	}
	articleID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, models.NewMessage("invalid article ID"))
		return
	}

	var req models.ArticleRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, models.NewMessage(err.Error()))
		return
	}

	cuserr := h.articleService.UpdateArticle(userID.(int), articleID, &req)
	if cuserr != nil {
		c.JSON(cuserr.HTTPCode, models.NewMessage(cuserr.Error()))
		return
	}

	c.JSON(200, models.NewMessage("article updated successfully"))
}

// DeleteArticleByID removes an article by its ID.
// @Summary Delete an article by ID
// @Description Delete an article by ID
// @Tags articles
// @Produce json
// @Param id path int true "Article ID"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Message
// @Failure 401 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 500 {object} models.Message
// @Router /articles/{id} [delete]
// @Security ApiKeyAuth
func (h *ArticlesHandler) DeleteArticleByID(c *gin.Context) {
	id := c.Param("id")
	articleID, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, models.NewMessage("invalid article ID"))
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.NewMessage("user not authenticated"))
		return
	}

	cuserr := h.articleService.DeleteArticleByID(userID.(int), articleID)
	if cuserr != nil {
		c.JSON(cuserr.HTTPCode, models.NewMessage(cuserr.Error()))
		return
	}

	c.JSON(200, models.NewMessage("article deleted successfully"))
}
