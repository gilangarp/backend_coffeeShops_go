package routers

import (
	"gilangarp/backend_coffeeShops_go/internal/handlers"
	"gilangarp/backend_coffeeShops_go/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func User(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/user")

	repo := repository.NewUser(d)
	handler := handlers.NewUser(repo)

	route.GET("/", handler.FetchAllUser)
	route.GET("/:id", handler.FetchDetailUser)
	route.POST("/", handler.SignUp)
	route.PATCH("/:id" , handler.UserUpdate)
	route.DELETE("/:id" , handler.DeleteUsers)
}