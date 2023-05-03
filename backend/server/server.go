package server

import (
	"fmt"
	"log"
	"sake/database"
	"sake/middleware"
	"sake/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gitlab.com/avarf/getenvs"
)

func Serve() {
	port := getenvs.GetEnvString("PORT", "8080")

	app := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	app.Static("/assets", "./assets")
	app.Use(cors.New(config))
	app.Use(middleware.Authorization())

	app.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})
	app.POST("/signup", database.CreateUser())
	app.POST("/login", database.Login())
	app.GET("/verifytoken", utils.VerifyToken())
	app.POST("/refreshtoken", database.Refresh_Token())
	app.GET("/products", database.ProductFetch())
	app.GET("/productdetials", database.ProductDetials())
	app.POST("/cart_add", database.AddToCart())
	app.GET("/cart_all", database.GetCart())
	app.POST("/cart_remove", database.RemoveCart())
	app.POST("/cart_order", database.PlaceOrder())
	log.Fatal(app.Run(fmt.Sprintf(":%s", port)))
}
