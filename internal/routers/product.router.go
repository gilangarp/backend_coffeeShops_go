package routers

import (
	"gilangarp/backend_coffeeShops_go/internal/handlers"
	"gilangarp/backend_coffeeShops_go/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Product(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/product")

	repo := repository.NewProduct(d)
	handler := handlers.NewProduct(repo)

	route.GET("/", handler.FetchAllProduct)
	route.GET("/:id", handler.FetchDetailProduct)
	route.POST("/", handler.PostProduct)
	route.PATCH("/:id", handler.ProductUpdate)
	route.DELETE("/:id" , handler.DeleteProducts)
}