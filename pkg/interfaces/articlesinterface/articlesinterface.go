package articlesinterface

import (
	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/models/articlesmodels"
	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/utils/customerror"
)

// ArticleRepository defines the interface for article-related database operations.
type ArticleRepository interface {
	// GetArticleByID retrieves an article by its unique identifier.
	// Returns the article and a custom error if the operation fails.
	GetArticleByID(id int) (*articlesmodels.Article, *customerror.CustomError)

	// GetAllArticles retrieves all articles from the database.
	// Returns a slice of articles and a custom error if the operation fails.
	GetAllArticles() ([]*articlesmodels.Article, *customerror.CustomError)

	// GetArticlesByUserID retrieves all articles created by a specific user.
	// Parameters:
	//   - userID: The unique identifier of the user whose articles are to be retrieved
	//
	// Returns:
	//   - A slice of articles created by the specified user
	//   - A custom error if the operation fails
	GetArticlesByUserID(userID int) ([]*articlesmodels.Article, *customerror.CustomError)

	// SearchArticles searches for articles based on a query string with pagination.
	// Parameters:
	//   - limit: The maximum number of articles to return
	//   - offset: The number of articles to skip before starting to collect the result set
	//   - query: The search query string
	// Returns a slice of articles and a custom error if the operation fails.
	SearchArticles(limit, offset int, query string) ([]*articlesmodels.Article, *customerror.CustomError)

	// CreateArticle creates a new article in the database.
	// Parameters:
	//   - userId: The ID of the user creating the article
	//   - title: The title of the article
	//   - content: The content of the article
	// Returns a custom error if the operation fails.
	CreateArticle(userId int, title string, content string) *customerror.CustomError

	// UpdateArticle updates an existing article in the database.
	// Parameters:
	//   - articleId: The unique identifier of the article to update
	//   - title: The new title for the article
	//   - content: The new content for the article
	// Returns a custom error if the operation fails.
	UpdateArticle(articleId int, title string, content string) *customerror.CustomError

	// DeleteArticleByID deletes an article by its unique identifier.
	// Returns a custom error if the operation fails.
	DeleteArticleByID(id int) *customerror.CustomError
}
