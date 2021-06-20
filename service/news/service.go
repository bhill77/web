package news

import (
	"encoding/json"
	"fmt"
	"github.com/bhill77/web/service/cache"
	"time"
)

type Service interface {
	Add(article Article) (Article, error)
	GetAll(filter map[string]string) []Article
}

type service struct {
	repo  Repository
	cache cache.Repository
}

func NewService(repo Repository, cache cache.Repository) Service {
	return &service{
		repo:  repo,
		cache: cache,
	}
}

func (s *service) Add(article Article) (Article, error) {
	article.Created = time.Now()
	res, err := s.repo.AddArticle(article)
	if err == nil {
		key := fmt.Sprintf("article:%s", res.ID)
		js, _ := json.Marshal(res)
		s.cache.Set(key, string(js), 24*time.Hour)
	}
	return res, err
}

func (s *service) GetAll(filter map[string]string) []Article {
	return s.repo.GetAll(filter)
}
