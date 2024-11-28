package models

type ArticleRequest struct {
	Title   string `json:"title" form:"title" validate:"required,min=3,max=255"`
	Content string `json:"content" form:"content" validate:"required,min=3"`
}

type ArticleResponse struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ArticlesResponse struct {
	Articles []*ArticleResponse `json:"articles"`
}
