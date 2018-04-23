package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/makishi00/go-vue-bbs/model"
	"github.com/makishi00/go-vue-bbs/service"
	"strconv"
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