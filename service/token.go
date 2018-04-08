package service

import (
"github.com/makishi00/go-vue-bbs/model"
)


var Token = token{}

type token struct {
}

func (t *token) Store(token model.Token) model.Token {
	db.Create(&token)
	return token
}