package handlers

import (
	"fmt"
	"gilangarp/backend_coffeeShops_go/internal/models"
	"gilangarp/backend_coffeeShops_go/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	*repository.RepoUser
}

func NewUser(r *repository.RepoUser) *HandlerUser {
	return &HandlerUser{r}
}

func (h *HandlerUser) SignUp(ctx *gin.Context){

	user := models.User{}

	fmt.Println("dari hanler: " , user)
	if err := ctx.ShouldBind(&user); err!= nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fmt.Println("Received user data:", user)
	respone, err := h.RegisterUser(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200,respone)
}


func (h *HandlerUser) FetchAllUser(ctx *gin.Context){
	data , err := h.GetAllUser()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200 , data)
}

func (h *HandlerUser) FetchDetailUser(ctx *gin.Context){
	id := ctx.Param("id")
	data , err := h.GetDetailUser(id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200 , data)
}

func (h *HandlerUser) UserUpdate(ctx *gin.Context) {
	user := models.User{}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := ctx.Param("id")
	
	updatedProfile, err := h.EditUsers(&user, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedProfile)
}

func (h *HandlerUser) DeleteUsers(ctx *gin.Context){
	id := ctx.Param("id")
	data , err := h.DeleteUser(id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200 , data)
}


