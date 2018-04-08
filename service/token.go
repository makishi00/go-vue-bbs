package service

import (
"github.com/makishi00/go-vue-bbs/model"
"github.com/makki0205/gojwt"
)


var Token = token{}

type token struct {
}

func (t *token) Store(token model.Token) model.Token {
	db.Create(&token)
	return token
}
func (t *token) ExisByToken(token string) bool {
	var tokens []model.Token
	payload, err := jwt.Decode(token)
	if err != nil{
		panic(err)
	}
	db.Where("user_id = ?", payload["id"]).Find(&tokens)
	return len(tokens) != 0
}
func (t *token) DeleteByUserId(userId int) bool {
	var tokens []model.Token
	db.Where("user_id = ?", userId).Delete(&tokens)
	return len(tokens) != 0
}