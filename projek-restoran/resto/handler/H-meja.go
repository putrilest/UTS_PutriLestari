package handler

import (
	"net/http"
	"projek-restoran/resto"

	"github.com/gin-gonic/gin"
)

func RoutesMeja(meja resto.UsecaseMeja, r *gin.RouterGroup) {
	u := HandlerMeja{
		meja,
	}

	v2 := r.Group("mejas")
	v2.GET("", u.GetMeja)
	v2.POST("", u.CreateMeja)
	v2.PUT(":id_meja", u.PutMeja)
	v2.DELETE(":id_meja", u.DeleteMeja)
}

// Voucher bind with voucher usecase interface
type HandlerMeja struct {
	UC resto.UsecaseMeja
}

func (meja HandlerMeja) GetMeja(ctx *gin.Context) {
	result, err := meja.UC.GetMeja(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, result)
}

func (meja HandlerMeja) CreateMeja(ctx *gin.Context) {
	err := meja.UC.CreateMeja(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Post Data"})
}

func (meja HandlerMeja) PutMeja(ctx *gin.Context) {
	err := meja.UC.UpdateMeja(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Updates Data"})
}

func (meja HandlerMeja) DeleteMeja(ctx *gin.Context) {
	err := meja.UC.DeleteMeja(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Delete Data"})
}
