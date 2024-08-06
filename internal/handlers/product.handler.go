package handlers

import (
	"gilangarp/backend_coffeeShops_go/internal/models"
	"gilangarp/backend_coffeeShops_go/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HandlerProduct struct {
	*repository.RepoProduct
}

func NewProduct(r *repository.RepoProduct) *HandlerProduct {
	return &HandlerProduct{r}
}

func (h *HandlerProduct) PostProduct(ctx *gin.Context){
	product := models.Product{}

	if err := ctx.ShouldBind(&product); err!= nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	respone, err := h.CreatedProduct(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200,respone)
}

func (h *HandlerProduct) FetchAllProduct(ctx *gin.Context){
	category := ctx.Query("category")
	favorite := ctx.Query("favoriteNpromo")
	searchText := ctx.Query("searchText")
	promo := ctx.Query("promo")
	limit := ctx.Query("limit")
	page := ctx.Query("page")

	promoBool := promo == "true"
	

	limits , _ := strconv.Atoi(limit)
	pages , _ := strconv.Atoi(page)

	params := &models.Filter{       
		Category: category,
        Favorite: favorite,
        SearchText: searchText,
		Promo: promoBool,
		Limit: limits,
		Page: pages,
    }

	data , err := h.GetAllProduct(params)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200 , data)

}