// routes/profile.go
package routes

import (
	"peerpay/controllers"
	"github.com/gin-gonic/gin"
)

func ProfileRoutes(router *gin.Engine) {
	profile := router.Group("/profile")
	{
		profile.POST("/create", controllers.CreateProfile)
		profile.PUT("/edit", controllers.EditProfile)
		profile.GET("/:userID", controllers.GetProfile)
	}
}
