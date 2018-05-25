package router

import (
	"github.com/gin-gonic/gin"
	. "github.com/makishi00/go-vue-bbs/controller"
	"github.com/makishi00/go-vue-bbs/middleware"
)

func apiRouter(api *gin.RouterGroup) {
	api.POST("/signup", User.Create)
	api.POST("/signin", middleware.Login)
}

func authApiRouter(auth *gin.RouterGroup) {
	auth.GET("/hoge", func(c *gin.Context) {
		c.JSON(200, "ok")
	})
	auth.GET("/bbs/show", Article.Show)
	auth.POST("/bbs/add", Article.Create)
	auth.POST("/bbs/delete", Article.Delete)
}
