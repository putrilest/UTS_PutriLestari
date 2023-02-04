package handler

import (
	"net/http"
	"projek-restoran/resto"

	"github.com/gin-gonic/gin"
)

func RoutesKasir(kasir resto.UsecaseKasir, r *gin.RouterGroup) {
	u := HandlerKasir{
		kasir,
	}

	v2 := r.Group("kasir")
	v2.GET("", u.GetKasir)
	v2.POST("", u.CreateKasir)
	v2.PUT(":id_kasir", u.PutKasir)
	v2.DELETE(":id_kasir", u.DeleteKasir)
}

// Voucher bind with voucher usecase interface
type HandlerKasir struct {
	UC resto.UsecaseKasir
}

func (kasir HandlerKasir) GetKasir(ctx *gin.Context) {
	result, err := kasir.UC.GetKasir(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, result)
}

func (kasir HandlerKasir) CreateKasir(ctx *gin.Context) {
	err := kasir.UC.CreateKasir(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Post Data"})
}

func (kasir HandlerKasir) PutKasir(ctx *gin.Context) {
	err := kasir.UC.UpdateKasir(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Updates Data"})
}

func (kasir HandlerKasir) DeleteKasir(ctx *gin.Context) {
	err := kasir.UC.DeleteKasir(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Delete Data"})
}
