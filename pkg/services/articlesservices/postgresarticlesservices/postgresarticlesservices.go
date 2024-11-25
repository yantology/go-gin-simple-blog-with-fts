package postgresarticlesservices

import (
	"database/sql"
	"strings"

	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/models/articlesmodels"
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
func (r *PostgresArticlesService) GetArticleByID(id int) (*articlesmodels.Article, error) {
	query := "SELECT id, user_id, title, content, created_at, updated_at FROM articles WHERE id = $1"
	row := r.db.QueryRow(query, id)

	var article articlesmodels.Article
	if err := row.Scan(&article.ID, &article.UserID, &article.Title, &article.Content, &article.CreatedAt, &article.UpdatedAt); err != nil {
		return nil, err
	}

	return &article, nil
}

// GetAllArticles retrieves all articles from the database
// Query: Selects all articles without any conditions
// Returns:
//   - Success: []*Article{
//     {ID: 1, Title: "Article 1"...},
//     {ID: 2, Title: "Article 2"...},
//     }
//   - Error: Database errors if query fails
func (r *PostgresArticlesService) GetAllArticles() ([]*articlesmodels.Article, error) {
	query := "SELECT id, user_id, title, content, created_at, updated_at FROM articles"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*articlesmodels.Article
	for rows.Next() {
		var article articlesmodels.Article
		if err := rows.Scan(&article.ID, &article.UserID, &article.Title, &article.Content, &article.CreatedAt, &article.UpdatedAt); err != nil {
			return nil, err
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
func (r *PostgresArticlesService) SearchArticles(limit, offset int, query string) ([]*articlesmodels.Article, error) {
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
		return nil, err
	}
	defer rows.Close()

	var articles []*articlesmodels.Article
	for rows.Next() {
		var article articlesmodels.Article
		if err := rows.Scan(&article.ID, &article.UserID, &article.Title, &article.Content, &article.CreatedAt, &article.UpdatedAt); err != nil {
			return nil, err
		}
		articles = append(articles, &article)
	}

	return articles, nil
}

// CreateArticle inserts a new article into the database
// Query: Inserts article data with user_id, title, content, and timestamps
// Parameters: article struct with required fields
// Returns:
// - Success: nil (article successfully created)
// - Error: Database constraints violations or connection errors
func (r *PostgresArticlesService) CreateArticle(article *articlesmodels.Article) error {
	query := "INSERT INTO articles (user_id, title, content, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"
	_, err := r.db.Exec(query, article.UserID, article.Title, article.Content, article.CreatedAt, article.UpdatedAt)
	return err
}

// UpdateArticle modifies an existing article in the database
// Query: Updates all fields except ID for the specified article
// Parameters: article struct with updated fields
// Returns:
// - Success: nil (article successfully updated)
// - Error: Article not found or database errors
func (r *PostgresArticlesService) UpdateArticle(article *articlesmodels.Article) error {
	query := "UPDATE articles SET user_id = $1, title = $2, content = $3, updated_at = $4 WHERE id = $5"
	_, err := r.db.Exec(query, article.UserID, article.Title, article.Content, article.UpdatedAt, article.ID)
	return err
}

// DeleteArticleByID removes an article from the database
// Query: Deletes article matching the specified ID
// Returns:
// - Success: nil (article successfully deleted)
// - Error: Article not found or database errors
func (r *PostgresArticlesService) DeleteArticleByID(id int) error {
	query := "DELETE FROM articles WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}
