package articlesrepository

import (
	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/interfaces/articlesinterface"
	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/models/articlesmodels"
	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/utils/customerror"
)

// ArticlesRepository provides methods to interact with the articles service
type ArticlesRepository struct {
	service articlesinterface.ArticleRepository
}

// NewArticlesRepository creates a new instance of ArticlesRepository
// Parameters:
//   - service: implementation of ArticleRepository interface
//
// Returns:
//   - *ArticlesRepository: new repository instance
func NewArticlesRepository(service articlesinterface.ArticleRepository) *ArticlesRepository {
	return &ArticlesRepository{service: service}
}

// GetArticleByID retrieves an article using its unique identifier
// Parameters:
//   - id: int - The article's database ID
//
// Returns:
//
//	Success: (*Article{
//	  ID: 1,
//	  UserID: 2,
//	  Title: "Golang Tutorial",
//	  Content: "Content here...",
//	  CreatedAt: time.Time{2024-01-20},
//	  UpdatedAt: time.Time{2024-01-20}
//	}, nil)
//	Error: (nil, Error) - Article not found
func (r *ArticlesRepository) GetArticleByID(id int) (*articlesmodels.Article, *customerror.CustomError) {
	return r.service.GetArticleByID(id)
}

// GetArticlesByUserID retrieves all articles created by a specific user.
// Parameters:
//   - userID: The unique identifier of the user whose articles are to be retrieved
//
// Returns:
//   - A slice of articles created by the specified user
//   - A custom error if the operation fails
func (r *ArticlesRepository) GetArticlesByUserID(userID int) ([]*articlesmodels.Article, *customerror.CustomError) {
	return r.service.GetArticlesByUserID(userID)
}

// GetAllArticles retrieves all available articles
// Returns:
//
//	Success: ([]*Article{
//	  {ID: 1, Title: "First Post"},
//	  {ID: 2, Title: "Second Post"}
//	}, nil)
//	Error: (nil, error) - Database errors
func (r *ArticlesRepository) GetAllArticles() ([]*articlesmodels.Article, *customerror.CustomError) {
	return r.service.GetAllArticles()
}

// SearchArticles performs full-text search on articles
// Parameters:
//   - limit: int - Max results to return
//   - offset: int - Number of results to skip
//   - query: string - Search keywords
//
// Returns:
//
//	Success: ([]*Article{
//	  {ID: 1, Title: "Golang Tips", Content: "..."},
//	  {ID: 5, Title: "Go Programming", Content: "..."}
//	}, nil)
//	Error: (nil, error) - Search/DB errors
func (r *ArticlesRepository) SearchArticles(limit, offset int, query string) ([]*articlesmodels.Article, *customerror.CustomError) {
	return r.service.SearchArticles(limit, offset, query)
}

// CreateArticle creates a new article in the database.
// Parameters:
//   - userId: The ID of the user creating the article
//   - title: The title of the article
//   - content: The content of the article
//
// Returns:
//   - nil if the creation was successful
//   - a custom error if there are validation or database errors
func (r *ArticlesRepository) CreateArticle(userId int, title string, content string) *customerror.CustomError {
	return r.service.CreateArticle(userId, title, content)
}

// UpdateArticle modifies an existing article in the database.
// Parameters:
//   - articleId: The unique identifier of the article to update
//   - title: The new title for the article
//   - content: The new content for the article
//
// Returns:
//   - nil if the update was successful
//   - a custom error if the article is not found or there are validation errors
func (r *ArticlesRepository) UpdateArticle(articleId int, title string, content string) *customerror.CustomError {
	return r.service.UpdateArticle(articleId, title, content)
}

// DeleteArticleByID removes an article by ID
// Parameters:
//   - id: int - ID of article to delete
//
// Returns:
//
//	Success: (nil)
//	Error: (error) - Not found/DB errors
func (r *ArticlesRepository) DeleteArticleByID(id int) *customerror.CustomError {
	return r.service.DeleteArticleByID(id)
}
