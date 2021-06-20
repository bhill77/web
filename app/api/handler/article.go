package handler

import (
	"github.com/bhill77/web/app/api/request"
	"github.com/bhill77/web/app/api/response"
	"github.com/bhill77/web/service/news"
	"github.com/labstack/echo/v4"
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

type ArticleHandler struct {
	newsService news.Service
}

func NewArticleHandler(newsService news.Service) *ArticleHandler {
	return &ArticleHandler{
		newsService: newsService,
	}
}

func (h *ArticleHandler) GetAll(c echo.Context) error {
	var filter = make(map[string]string)
	if author := c.FormValue("author"); author != "" {
		filter["author"] = author
	}
	if query := c.FormValue("query"); query != "" {
		filter["query"] = query
	}

	list := h.newsService.GetAll(filter)
	res := response.ListArticleResponse(list)
	return c.JSON(http.StatusOK, res)
}

func (h *ArticleHandler) Add(c echo.Context) error {
	var req request.Article
	rules := govalidator.MapData{
		"author": []string{"required"},
		"title":  []string{"required"},
		"body":   []string{"required"},
	}

	opts := govalidator.Options{
		Request: c.Request(),
		Data:    &req,
		Rules:   rules,
	}

	v := govalidator.New(opts)
	e := v.ValidateJSON()
	if len(e) > 0 {
		err := map[string]interface{}{"validationError": e}
		return c.JSON(http.StatusBadRequest, err)
	}

	article := news.Article{
		Author: req.Author,
		Title:  req.Title,
		Body:   req.Body,
	}

	result, err := h.newsService.Add(article)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res := response.NewArticleResponse(result)
	return c.JSON(http.StatusCreated, res)
}
