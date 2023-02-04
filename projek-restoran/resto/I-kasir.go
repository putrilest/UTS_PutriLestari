package resto

import (
	"projek-restoran/models"
	"github.com/gin-gonic/gin"
)


type RepoKasir interface {
	GetKasir(ctx *gin.Context) ([]models.Kasir, error)
	CreateKasir(data models.Kasir) error
	UpdateKasir(data models.Kasir, param string) error
	DeleteKasir(param string) error
}

type UsecaseKasir interface {
	GetKasir(ctx *gin.Context) ([]models.Kasir, error)
	CreateKasir(ctx *gin.Context) error
	UpdateKasir(ctx *gin.Context) error
	DeleteKasir(ctx *gin.Context) error
}