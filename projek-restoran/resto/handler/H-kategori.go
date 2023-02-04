package handler

import (
	"net/http"
	"projek-restoran/resto"

	"github.com/gin-gonic/gin"
)

func RoutesKtg(kategori resto.UsecaseKtg, r *gin.RouterGroup) {
	u := HandlerKtg{
		kategori,
	}

	v2 := r.Group("kategoris")
	v2.GET("", u.GetKategori)
	v2.POST("", u.CreateKategori)
	v2.PUT(":id_ktg", u.PutKategori)
	v2.DELETE(":id_ktg", u.DeleteKategori)
}

// Voucher bind with voucher usecase interface
type HandlerKtg struct {
	UC resto.UsecaseKtg
}

func (kategori HandlerKtg) GetKategori(ctx *gin.Context) {
	result, err := kategori.UC.GetKtg(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, result)
}

func (kategori HandlerKtg) CreateKategori(ctx *gin.Context) {
	err := kategori.UC.CreateKtg(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Post Data"})
}

func (kategori HandlerKtg) PutKategori(ctx *gin.Context) {
	err := kategori.UC.UpdateKtg(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Updates Data"})
}

func (kategori HandlerKtg) DeleteKategori(ctx *gin.Context) {
	err := kategori.UC.DeleteKtg(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Delete Data"})
}