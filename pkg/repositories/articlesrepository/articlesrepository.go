package articlesrepository

import (
	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/interfaces/articlesinterface"
	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/models/articlesmodels"
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
func (r *ArticlesRepository) GetArticleByID(id int) (*articlesmodels.Article, error) {
	return r.service.GetArticleByID(id)
}

// GetAllArticles retrieves all available articles
// Returns:
//
//	Success: ([]*Article{
//	  {ID: 1, Title: "First Post"},
//	  {ID: 2, Title: "Second Post"}
//	}, nil)
//	Error: (nil, error) - Database errors
func (r *ArticlesRepository) GetAllArticles() ([]*articlesmodels.Article, error) {
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
func (r *ArticlesRepository) SearchArticles(limit, offset int, query string) ([]*articlesmodels.Article, error) {
	return r.service.SearchArticles(limit, offset, query)
}

// CreateArticle creates a new article
// Parameters:
//   - article: *Article - Article data to create
//
// Returns:
//
//	Success: (nil)
//	Error: (error) - Validation/DB errors
func (r *ArticlesRepository) CreateArticle(article *articlesmodels.Article) error {
	return r.service.CreateArticle(article)
}

// UpdateArticle modifies an existing article
// Parameters:
//   - article: *Article - Updated article data
//
// Returns:
//
//	Success: (nil)
//	Error: (error) - Not found/validation errors
func (r *ArticlesRepository) UpdateArticle(article *articlesmodels.Article) error {
	return r.service.UpdateArticle(article)
}

// DeleteArticleByID removes an article by ID
// Parameters:
//   - id: int - ID of article to delete
//
// Returns:
//
//	Success: (nil)
//	Error: (error) - Not found/DB errors
func (r *ArticlesRepository) DeleteArticleByID(id int) error {
	return r.service.DeleteArticleByID(id)
}
