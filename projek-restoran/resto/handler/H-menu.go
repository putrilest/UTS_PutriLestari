package handler

import (
	"net/http"
	"projek-restoran/resto"

	"github.com/gin-gonic/gin"
)

func RoutesMenu(menu resto.UsecaseMenu, r *gin.RouterGroup) {
	u := HandlerMenu{
		menu,
	}

	v2 := r.Group("menus")
	v2.GET("", u.GetMenu)
	v2.POST("", u.CreateMenu)
	v2.PUT(":id_menu", u.PutMenu)
	v2.DELETE(":id_menu", u.DeleteMenu)
}

// Voucher bind with voucher usecase interface
type HandlerMenu struct {
	UC resto.UsecaseMenu
}

func (menu HandlerMenu) GetMenu(ctx *gin.Context) {
	result, err := menu.UC.GetMenu(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, result)
}

func (menu HandlerMenu) CreateMenu(ctx *gin.Context) {
	err := menu.UC.CreateMenu(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Post Data"})
}

func (menu HandlerMenu) PutMenu(ctx *gin.Context) {
	err := menu.UC.UpdateMenu(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Updates Data"})
}

func (menu HandlerMenu) DeleteMenu(ctx *gin.Context) {
	err := menu.UC.DeleteMenu(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Delete Data"})
}