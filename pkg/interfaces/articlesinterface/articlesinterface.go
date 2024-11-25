package articlesinterface

import (
	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/models/articlesmodels"
	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/utils/customerror"
)

type ArticleRepository interface {
	GetArticleByID(id int) (*articlesmodels.Article, *customerror.CustomError)
	GetAllArticles() ([]*articlesmodels.Article, *customerror.CustomError)
	SearchArticles(limit, offset int, query string) ([]*articlesmodels.Article, *customerror.CustomError)
	CreateArticle(article *articlesmodels.Article) *customerror.CustomError
	UpdateArticle(article *articlesmodels.Article) *customerror.CustomError
	DeleteArticleByID(id int) *customerror.CustomError
}
