package response

import (
	"github.com/bhill77/web/service/news"
	"time"
)

type Article struct {
	ID      int       `json:"id"`
	Author  string    `json:"author"`
	Title   string    `json:"title"`
	Body    string    `json:"body"`
	Created time.Time `json:"created"`
}

func NewArticleResponse(article news.Article) Article {
	return Article{
		ID:      article.ID,
		Author:  article.Author,
		Title:   article.Title,
		Body:    article.Body,
		Created: article.Created,
	}
}

func ListArticleResponse(articles []news.Article) []Article {
	var list []Article
	for _, article := range articles {
		item := Article{
			ID:      article.ID,
			Author:  article.Author,
			Title:   article.Title,
			Body:    article.Body,
			Created: article.Created,
		}

		list = append(list, item)
	}

	return list
}
