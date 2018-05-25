package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/makishi00/go-vue-bbs/model"
	"github.com/makishi00/go-vue-bbs/service"
	"strconv"
	"github.com/makki0205/gojwt"
)

var Article = articleimpl{}

type articleimpl struct {
}

func (a *articleimpl) Create(c *gin.Context) {
	var article model.Article
	err := c.BindJSON(&article)
	if err != nil {
		BatRequest(err.Error(), c)
		return
	}

	if userId, ok := c.Value("user_id").(string); ok {
		var userId, _ = strconv.ParseUint(userId, 10, 32)
		article = service.Article.Store(article, uint(userId))
	}
	Json(article, c)
}

func (a *articleimpl) Show(c *gin.Context) {
	var article = service.Article.Show()
	Json(article, c)
}

func (a *articleimpl) Delete(c *gin.Context) {
	var article model.Article
	err := c.BindJSON(&article)
	if err != nil {
		BatRequest(err.Error(), c)
		return
	}

	token := c.GetHeader("Authorization")
	payload, err := jwt.Decode(token)
	if err != nil {
		panic(err)
	}

	var userId, _ = strconv.ParseUint(payload["id"], 10, 32)
	if !service.Article.ExistArticleIdByUserId(int(article.ID), uint(userId)) {
		BatRequest(err.Error(), c)
	}
	article = service.Article.Delete(article, int(article.ID), uint(userId))
	Json(article, c)
}