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