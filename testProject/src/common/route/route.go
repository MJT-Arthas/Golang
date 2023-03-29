package route

import (
	"net/http"
	"textProject/src/common/midware"
	"textProject/src/controller"

	"github.com/gin-gonic/gin"
)

func PathRoute(r *gin.Engine) *gin.Engine {

	gapi := r.Group("/gapi")
	{
		gapi.Use(midware.MiddleW())
		controller.UserRegister(gapi)

	}

	r.LoadHTMLGlob("templates/*")
	r.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{"title": "gin web页面"})
	})
	return r
}
