package handlers

import (
	"gilangarp/backend_coffeeShops_go/internal/models"
	"gilangarp/backend_coffeeShops_go/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerProfile struct {
	*repository.RepoProfile
}

func NewProfile(r *repository.RepoProfile) *HandlerProfile {
	return &HandlerProfile{r}
}

func (h *HandlerProfile) PostProfile(ctx *gin.Context){
	profile := models.Profile{}

	if err := ctx.ShouldBindJSON(&profile); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := ctx.Param("id")
	
	updatedProfile, err := h.CreatedProfile(&profile, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedProfile)


}

func (h *HandlerProfile) ProfileUpdate(ctx *gin.Context) {
	profile := models.Profile{}

	if err := ctx.ShouldBindJSON(&profile); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := ctx.Param("id")
	
	updatedProfile, err := h.EditProfile(&profile, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedProfile)
}

