package resto

import (
	"projek-restoran/models"
	"github.com/gin-gonic/gin"
)


type RepoReservasi interface {
	GetReservasi(ctx *gin.Context) ([]models.Reservasi, error)
	CreateReservasi(data models.Reservasi) error
	UpdateReservasi(data models.Reservasi, param string) error
	DeleteReservasi(param string) error
}

type UsecaseReservasi interface {
	GetReservasi(ctx *gin.Context) ([]models.Reservasi, error)
	CreateReservasi(ctx *gin.Context) error
	UpdateReservasi(ctx *gin.Context) error
	DeleteReservasi(ctx *gin.Context) error
}