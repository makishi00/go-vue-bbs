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

func (a *article) Show() []model.Article {
	var article []model.Article
	db.Find(&article)
	return article
}

func (a *article) ExistArticleIdByUserId(id int, userId uint) bool {
	var articles []model.Article
	db.Where("id = ? AND user_id = ?", id, userId).Find(&articles)

	return len(articles) != 0
}

func (a *article) Delete(article model.Article, id int, userId uint) model.Article {
	db.Where("id = ? AND user_id = ?", id, userId).Delete(&article)
	return article
}

func (a *article) Edit(article model.Article, id int, userId uint, title string, body string) model.Article {
	db.Model(&article).Where("id = ? AND user_id = ?", id, userId).Updates(map[string]interface{}{"title": title, "body": body})
	return article
}