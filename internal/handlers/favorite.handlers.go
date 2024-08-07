package handlers

import (
	"gilangarp/backend_coffeeShops_go/internal/models"
	"gilangarp/backend_coffeeShops_go/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerFavorite struct {
	*repository.RepoFavorite
}

func NewFavorite(r *repository.RepoFavorite) *HandlerFavorite {
	return &HandlerFavorite{r}
}


func (h *HandlerFavorite) PostFavorite(ctx *gin.Context){
	favorite := models.Favorite{}

	if err := ctx.ShouldBindJSON(&favorite); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := ctx.Param("id")
	
	createdFavorite, err := h.CreatedFavorite(&favorite, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, createdFavorite)
}

func (h *HandlerFavorite) FetchDetailFavorite(ctx *gin.Context){
	id := ctx.Param("id")
	data , err := h.GetDetailFavorite(id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200 , data)
}

func (h *HandlerFavorite) DeletFavorites(ctx *gin.Context){
	id := ctx.Param("id")
	data , err := h.DeleteFavorite(id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200 , data)
}