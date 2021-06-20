package handler_test

import (
	"errors"
	"github.com/bhill77/web/app/api/handler"
	"github.com/bhill77/web/service/news"
	newsMock "github.com/bhill77/web/service/news/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

var (
	mockService   newsMock.Service
	created       time.Time
	newArticleReq news.Article
	newArticle    news.Article
	listArticle   []news.Article
)

func TestMain(m *testing.M) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	created = time.Date(2021, 1, 1, 0, 0, 0, 0, loc)
	newArticleReq = news.Article{
		Author: "user",
		Title:  "article title",
		Body:   "text",
	}

	newArticle = news.Article{
		ID:      1,
		Author:  "user",
		Title:   "article title",
		Body:    "text",
		Created: created,
	}

	listArticle = []news.Article{newArticle}

	os.Exit(m.Run())
}

func TestArticleHandler_Add(t *testing.T) {

	t.Run("add success", func(t *testing.T) {
		e := echo.New()
		js := `{"author": "user", "title": "article title", "body": "text"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(js))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/articles")
		mockService.On("Add", newArticleReq).Return(newArticle, nil)
		h := handler.NewArticleHandler(&mockService)

		if assert.NoError(t, h.Add(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
		}
	})

	t.Run("failed bad request", func(t *testing.T) {
		e := echo.New()
		js := `{"title": "article title", "body": "text"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(js))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/articles")
		h := handler.NewArticleHandler(&mockService)

		if assert.NoError(t, h.Add(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("failed insert", func(t *testing.T) {
		e := echo.New()
		js := `{"author": "user1", "title": "article title", "body": "text"}`
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(js))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/articles")
		newArticleReq.Author = "user1"
		mockService.On("Add", newArticleReq).Return(newArticle, errors.New("error"))
		h := handler.NewArticleHandler(&mockService)

		if assert.NoError(t, h.Add(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})
}

func TestArticleHandler_GetAll(t *testing.T) {
	t.Run("no filter", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/articles")
		filter := map[string]string{}
		mockService.On("GetAll", filter).Return(listArticle)
		h := handler.NewArticleHandler(&mockService)

		if assert.NoError(t, h.GetAll(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("filter author", func(t *testing.T) {
		e := echo.New()
		filter := map[string]string{"author": "user"}
		mockService.On("GetAll", filter).Return(listArticle)
		h := handler.NewArticleHandler(&mockService)
		req := httptest.NewRequest(http.MethodGet, "/articles/?author=user", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, h.GetAll(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("filter query", func(t *testing.T) {
		e := echo.New()
		filter := map[string]string{"query": "text"}
		mockService.On("GetAll", filter).Return(listArticle)
		h := handler.NewArticleHandler(&mockService)
		req := httptest.NewRequest(http.MethodGet, "/articles?query=text", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, h.GetAll(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

}
