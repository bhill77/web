package news

import (
	"github.com/bhill77/web/service/news"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) AddArticle(article news.Article) (news.Article, error) {
	err := r.db.Create(&article).Error
	return article, err
}

func (r *Repository) GetAll(filter map[string]string) []news.Article {
	db := r.db.Debug()
	var list []news.Article

	if val, ok := filter["author"]; ok {
		db = db.Where("match(author) against(?)", val)
	}

	if val, ok := filter["query"]; ok {
		db = db.Where("match(title, body) against(?) ", val)
	}

	db.Order("created desc").Find(&list)
	return list
}
