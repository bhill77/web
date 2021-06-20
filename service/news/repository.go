package news

import "time"

type Repository interface {
	AddArticle(article Article) (Article, error)
	GetAll(filter map[string]string) []Article
}

type Article struct {
	ID      int
	Author  string
	Title   string
	Body    string
	Created time.Time
}
