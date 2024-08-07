package routers

import (
	"gilangarp/backend_coffeeShops_go/internal/handlers"
	"gilangarp/backend_coffeeShops_go/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Profile(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/profile")

	repo := repository.NewProfile(d)
	handler := handlers.NewProfile(repo)

	route.POST("/:id", handler.PostProfile)
	route.PATCH("/:id", handler.ProfileUpdate)
	route.GET("/", handler.FetchAllProfile)
	route.GET("/:id" , handler.FetchDetailProfile)
	route.DELETE("/:id" , handler.DeleteProfile)
	
}