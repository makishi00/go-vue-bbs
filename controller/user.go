package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/makishi00/go-test/model"
	"github.com/makishi00/go-test/service"
)

var User = userimpl{}

type userimpl struct {
}

func (u *userimpl) Create(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		BatRequest(err.Error(), c)
		return
	}
	user = service.User.Store(user)
	Json(user, c)
}
