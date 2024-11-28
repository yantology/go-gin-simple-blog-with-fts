package services

import (
	"github.com/yantology/go-gin-simple-blog-with-fts/internal/models"
	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/repositories/articlesrepository"
	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/utils/customerror"
)

type ArticlesService struct {
	articlesRepo *articlesrepository.ArticlesRepository
}

func NewArticlesService(articlesRepo *articlesrepository.ArticlesRepository) *ArticlesService {
	return &ArticlesService{
		articlesRepo: articlesRepo,
	}
}

func (s *ArticlesService) CreateArticle(userId int, req *models.ArticleRequest) *customerror.CustomError {
	return s.articlesRepo.CreateArticle(userId, req.Title, req.Content)
}

func (s *ArticlesService) GetArticleByID(id int) (*models.ArticleResponse, *customerror.CustomError) {
	article, cuserr := s.articlesRepo.GetArticleByID(id)
	if cuserr != nil {
		return nil, cuserr
	}

	return &models.ArticleResponse{
		ID:      article.ID,
		UserID:  article.UserID,
		Title:   article.Title,
		Content: article.Content,
	}, nil
}

func (s *ArticlesService) GetArticlesByUserID(userID int) (*models.ArticlesResponse, *customerror.CustomError) {
	articles, cuserr := s.articlesRepo.GetArticlesByUserID(userID)
	if cuserr != nil {
		return nil, cuserr
	}

	var response = &models.ArticlesResponse{
		Articles: []*models.ArticleResponse{},
	}
	for _, article := range articles {
		response.Articles = append(response.Articles, &models.ArticleResponse{
			ID:      article.ID,
			UserID:  article.UserID,
			Title:   article.Title,
			Content: article.Content,
		})
	}

	return response, nil
}

func (s *ArticlesService) GetAllArticles() (*models.ArticlesResponse, *customerror.CustomError) {
	articles, cuserr := s.articlesRepo.GetAllArticles()
	if cuserr != nil {
		return nil, cuserr
	}

	var response = &models.ArticlesResponse{
		Articles: []*models.ArticleResponse{},
	}
	for _, article := range articles {
		response.Articles = append(response.Articles, &models.ArticleResponse{
			ID:      article.ID,
			UserID:  article.UserID,
			Title:   article.Title,
			Content: article.Content,
		})
	}
	return response, nil
}

func (s *ArticlesService) SearchArticles(limit, offset int, query string) (*models.ArticlesResponse, *customerror.CustomError) {
	articles, cuserr := s.articlesRepo.SearchArticles(limit, offset, query)
	if cuserr != nil {
		return nil, cuserr
	}

	var response = &models.ArticlesResponse{
		Articles: []*models.ArticleResponse{},
	}
	for _, article := range articles {
		response.Articles = append(response.Articles, &models.ArticleResponse{
			ID:      article.ID,
			UserID:  article.UserID,
			Title:   article.Title,
			Content: article.Content,
		})
	}
	return response, nil
}

func (s *ArticlesService) UpdateArticle(userID int, articleId int, req *models.ArticleRequest) *customerror.CustomError {
	article, cuserr := s.articlesRepo.GetArticleByID(articleId)

	if cuserr != nil {
		return cuserr
	}

	if article.UserID != userID {
		return customerror.NewCustomError(nil, "You are not authorized to update this article", 403)
	}

	return s.articlesRepo.UpdateArticle(articleId, req.Title, req.Content)
}

func (s *ArticlesService) DeleteArticleByID(userID int, articleId int) *customerror.CustomError {
	article, cuserr := s.articlesRepo.GetArticleByID(articleId)

	if cuserr != nil {
		return cuserr
	}

	if article.UserID != userID {
		return customerror.NewCustomError(nil, "You are not authorized to update this article", 403)
	}
	return s.articlesRepo.DeleteArticleByID(articleId)
}
