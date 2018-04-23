package service

import (
	"github.com/makishi00/go-vue-bbs/model"
)

var Article = article{}

type article struct {
}

func (a *article) Store(article model.Article, userId uint) model.Article {
	article.UserID = userId
	db.Create(&article)
	return article
}
