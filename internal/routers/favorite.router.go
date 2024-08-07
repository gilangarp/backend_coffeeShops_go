package routers

import (
	"gilangarp/backend_coffeeShops_go/internal/handlers"
	"gilangarp/backend_coffeeShops_go/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Favorite(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/favorite")

	repo := repository.NewFavorite(d)
	handler := handlers.NewFavorite(repo)

	route.GET("/:id", handler.FetchDetailFavorite)
	route.POST("/:id", handler.PostFavorite)
	route.DELETE("/:id" ,handler.DeletFavorites)
}