package postgresarticlesservices

import (
	"database/sql"
	"log"
	"strings"
	"time"

	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/models/articlesmodels"
	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/utils/customerror"
	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/utils/customerror/postgreserror"
)

// PostgresArticlesService provides methods to interact with articles table in PostgreSQL database
type PostgresArticlesService struct {
	db *sql.DB
}

// NewPostgresArticlesService creates a new instance of PostgresArticlesService
func NewPostgresArticlesService(db *sql.DB) *PostgresArticlesService {
	return &PostgresArticlesService{db: db}
}

// GetArticleByID retrieves a single article by its ID
// Query: Selects all fields from articles table where id matches
// Returns:
// - Success: *Article{ID: 1, Title: "Sample Article", Content: "Content here"...}
// - Error: sql.ErrNoRows if article not found, or any other DB error
func (r *PostgresArticlesService) GetArticleByID(id int) (*articlesmodels.Article, *customerror.CustomError) {
	query := "SELECT id, user_id, title, content, created_at, updated_at FROM articles WHERE id = $1"
	row := r.db.QueryRow(query, id)

	var article articlesmodels.Article
	if err := row.Scan(&article.ID, &article.UserID, &article.Title, &article.Content, &article.CreatedAt, &article.UpdatedAt); err != nil {
		return nil, postgreserror.NewPostgresError(err)
	}

	return &article, nil
}

// GetArticlesByUserID retrieves all articles created by a specific user
// Query: Selects all articles where user_id matches the specified ID
// Returns:
//   - Success: []*Article{
//     {ID: 1, Title: "Article 1"...},
//     {ID: 2, Title: "Article 2"...},
//     }
//   - Error: Database errors if query fails
func (r *PostgresArticlesService) GetArticlesByUserID(userID int) ([]*articlesmodels.Article, *customerror.CustomError) {
	query := "SELECT id, user_id, title, content, created_at, updated_at FROM articles WHERE user_id = $1"
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, postgreserror.NewPostgresError(err)
	}
	defer rows.Close()

	var articles []*articlesmodels.Article
	for rows.Next() {
		var article articlesmodels.Article
		if err := rows.Scan(&article.ID, &article.UserID, &article.Title, &article.Content, &article.CreatedAt, &article.UpdatedAt); err != nil {
			return nil, postgreserror.NewPostgresError(err)
		}
		articles = append(articles, &article)
	}

	return articles, nil
}

// GetAllArticles retrieves all articles from the database
// Query: Selects all articles without any conditions
// Returns:
//   - Success: []*Article{
//     {ID: 1, Title: "Article 1"...},
//     {ID: 2, Title: "Article 2"...},
//     }
//   - Error: Database errors if query fails
func (r *PostgresArticlesService) GetAllArticles() ([]*articlesmodels.Article, *customerror.CustomError) {
	query := "SELECT id, user_id, title, content, created_at, updated_at FROM articles"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, postgreserror.NewPostgresError(err)
	}
	defer rows.Close()

	var articles []*articlesmodels.Article
	for rows.Next() {
		var article articlesmodels.Article
		if err := rows.Scan(&article.ID, &article.UserID, &article.Title, &article.Content, &article.CreatedAt, &article.UpdatedAt); err != nil {
			return nil, postgreserror.NewPostgresError(err)
		}
		articles = append(articles, &article)
	}

	return articles, nil
}

// SearchArticles performs full-text search on articles using PostgreSQL's tsvector
// Query: Uses FTS with indonesian dictionary, ranks results by relevance
// Parameters:
// - limit: maximum number of results
// - offset: number of results to skip
// - query: search terms (e.g., "golang programming")
// Returns:
//   - Success: []*Article matching search terms, ordered by relevance
//     Example: query="golang" -> [{Title: "Intro to Golang"}, {Title: "Golang Tips"}]
//   - Error: Database errors or invalid search query
func (r *PostgresArticlesService) SearchArticles(limit int, offset int, query string) ([]*articlesmodels.Article, *customerror.CustomError) {
	// Convert search query to tsquery format and use indonesian dictionary
	searchQuery := `
        SELECT id, user_id, title, content, created_at, updated_at 
        FROM articles 
        WHERE tsv @@ to_tsquery('indonesian', $1)
        ORDER BY ts_rank(tsv, to_tsquery('indonesian', $1)) DESC
        LIMIT $2 OFFSET $3`

	// Convert space-separated words to tsquery format (word1 & word2)
	formattedQuery := strings.Join(strings.Fields(query), " & ")

	rows, err := r.db.Query(searchQuery, formattedQuery, limit, offset)
	if err != nil {
		return nil, postgreserror.NewPostgresError(err)
	}
	defer rows.Close()

	var articles []*articlesmodels.Article

	log.Printf("Search query: %s", formattedQuery)
	for rows.Next() {
		var article articlesmodels.Article
		if err := rows.Scan(&article.ID, &article.UserID, &article.Title, &article.Content, &article.CreatedAt, &article.UpdatedAt); err != nil {
			return nil, postgreserror.NewPostgresError(err)
		}
		articles = append(articles, &article)
	}

	return articles, nil
}

// CreateArticle creates a new article in the database for the specified user.
// It takes the user ID, title, and content as input parameters and returns
// a custom error if the operation fails.
//
// The article is created with current timestamp for both created_at and updated_at fields.
//
// Returns nil on successful creation. If the operation fails due to database constraints
// or connection issues, returns a wrapped custom error.
func (r *PostgresArticlesService) CreateArticle(userId int, title string, content string) *customerror.CustomError {
	query := "INSERT INTO articles (user_id, title, content, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"
	_, err := r.db.Exec(query, userId, title, content, time.Now(), time.Now())
	if err != nil {
		return postgreserror.NewPostgresError(err)
	}
	return nil
}

// UpdateArticle updates an existing article in the database with the provided title and content.
// The article's updated_at timestamp is automatically set to the current time.
//
// Parameters:
//   - articleId: The unique identifier of the article to update
//   - title: The new title for the article
//   - content: The new content for the article
//
// Returns a *customerror.CustomError which is:
//   - nil if the update was successful
//   - wrapped database error if the operation fails or article is not found
func (r *PostgresArticlesService) UpdateArticle(articleId int, title string, content string) *customerror.CustomError {
	query := "UPDATE articles SET title = $1, content = $2, updated_at = $3 WHERE id = $4"
	_, err := r.db.Exec(query, title, content, time.Now(), articleId)
	if err != nil {
		return postgreserror.NewPostgresError(err)
	}
	return nil
}

// DeleteArticleByID removes an article from the database
// Query: Deletes article matching the specified ID
// Returns:
// - Success: nil (article successfully deleted)
// - Error: Article not found or database errors
func (r *PostgresArticlesService) DeleteArticleByID(id int) *customerror.CustomError {
	query := "DELETE FROM articles WHERE id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return postgreserror.NewPostgresError(err)

	}
	return nil
}
