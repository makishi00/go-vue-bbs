package router

import (
	"github.com/gin-gonic/gin"
	. "github.com/makishi00/go-test/controller"
	"github.com/makishi00/go-test/middleware"
)

func apiRouter(api *gin.RouterGroup) {
	api.POST("/signup", User.Create)
	api.POST("/signin", middleware.Login)

}

func authApiRouter(auth *gin.RouterGroup) {
	auth.GET("/hoge", func(c *gin.Context) {
		c.JSON(200, "ok")
	})
}
