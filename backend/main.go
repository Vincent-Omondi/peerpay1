// main.go
package main

import (
	"net/http"

	"peerpay/backend/database"
	"peerpay/backend/routes"

	"github.com/gin-gonic/gin"
	"github.com/go-delve/delve/pkg/config"
)

func main() {
	config.LoadConfig()
	database.Connect()
	database.Migrate()

	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// Trust specific proxies
	router.SetTrustedProxies(nil) // Nil means trusting no proxies

	// Register routes
	routes.AuthRoutes(router)
	routes.ProfileRoutes(router)
	routes.SharesRoutes(router)
	routes.AdminRoutes(router)

	// Handle 404 errors
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"Peerpay": "Not Found"})
	})

	router.Run(":8080")
}

// package main

// import (
// 	"peerpay/config"
// 	"peerpay/controllers"
// 	"peerpay/database"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	config.LoadConfig()
// 	database.Connect()
// 	database.Migrate()

// 	router := gin.Default()

// 	// Auth routes
// 	auth := router.Group("/auth")
// 	{
// 		auth.POST("/signup", controllers.SignUp)
// 		auth.POST("/login", controllers.Login)
// 		auth.POST("/referral", controllers.Referral)
// 	}

// 	// Profile routes
// 	profile := router.Group("/profile")
// 	{
// 		profile.POST("/create", controllers.CreateProfile)
// 		profile.PUT("/edit", controllers.EditProfile)
// 		profile.GET("/:userID", controllers.GetProfile)
// 	}

// 	// Share routes
// 	shares := router.Group("/shares")
// 	{
// 		shares.POST("/buy", controllers.BuyShares)
// 		shares.POST("/sell", controllers.SellShares)
// 		shares.GET("/:userID", controllers.GetShares)
// 	}

// 	// Admin routes
// 	admin := router.Group("/admin")
// 	{
// 		admin.GET("/users", controllers.GetUsers)
// 		admin.DELETE("/user/:userID", controllers.DeleteUser)
// 		admin.GET("/dormant-accounts", controllers.GetDormantAccounts)
// 	}

// 	router.Run()
// }
