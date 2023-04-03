package route

import (
	"net/http"
	"textProject/src/controller"

	"github.com/gin-gonic/gin"
)

func PathRoute(r *gin.Engine) *gin.Engine {

	r.LoadHTMLGlob("templates/*")
	// r.LoadHTMLFiles("templates/index.html")
	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.tepl", gin.H{"title": "gin web页面"})
	})

	gapi := r.Group("/gapi")
	{
		// gapi.Use(midware.MiddleW())
		controller.UserRegister(gapi)

	}

	return r
}
