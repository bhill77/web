// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	news "github.com/bhill77/web/service/news"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// AddArticle provides a mock function with given fields: article
func (_m *Repository) AddArticle(article news.Article) (news.Article, error) {
	ret := _m.Called(article)

	var r0 news.Article
	if rf, ok := ret.Get(0).(func(news.Article) news.Article); ok {
		r0 = rf(article)
	} else {
		r0 = ret.Get(0).(news.Article)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(news.Article) error); ok {
		r1 = rf(article)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *Repository) GetAll() []news.Article {
	ret := _m.Called()

	var r0 []news.Article
	if rf, ok := ret.Get(0).(func() []news.Article); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]news.Article)
		}
	}

	return r0
}
