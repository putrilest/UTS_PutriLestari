package handler

import (
	"net/http"
	"projek-restoran/resto"

	"github.com/gin-gonic/gin"
)

func RoutesReservasi(reservasi resto.UsecaseReservasi, r *gin.RouterGroup) {
	u := HandlerReservasi{
		reservasi,
	}

	v2 := r.Group("reservasi")
	v2.GET("", u.GetReservasi)
	v2.POST("", u.CreateReservasi)
	v2.PUT(":id_reservasi", u.PutReservasi)
	v2.DELETE(":id_reservasi", u.DeleteReservasi)
}

// Voucher bind with voucher usecase interface
type HandlerReservasi struct {
	UC resto.UsecaseReservasi
}

func (reservasi HandlerReservasi) GetReservasi(ctx *gin.Context) {
	result, err := reservasi.UC.GetReservasi(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, result)
}

func (reservasi HandlerReservasi) CreateReservasi(ctx *gin.Context) {
	err := reservasi.UC.CreateReservasi(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Post Data"})
}

func (reservasi HandlerReservasi) PutReservasi(ctx *gin.Context) {
	err := reservasi.UC.UpdateReservasi(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Updates Data"})
}

func (reservasi HandlerReservasi) DeleteReservasi(ctx *gin.Context) {
	err := reservasi.UC.DeleteReservasi(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Succes Delete Data"})
}
