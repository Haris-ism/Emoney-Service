package controller

import (
	"gin/middleware"

	"github.com/gin-gonic/gin"
)

func MainRouter(services *Repo, port string) {
	router := gin.Default()
	router.Use(middleware.CorsMiddleware)
	router.POST("/signup", services.SignUp)
	router.POST("/signin", services.SignIn)

	routes := router.Use(middleware.Auth)

	routes.GET("/profile", services.GetProfile)
	routes.GET("/balance", services.GetBalance)
	routes.PUT("/balance", services.TopUpBalance)
	routes.GET("/inquirylist", services.Inquiry)
	routes.POST("/inquiry", services.TransactionConfirmation)
	routes.GET("/history", services.History)

	router.Run(port)
}
