package resto

import (
	"projek-restoran/models"
	"github.com/gin-gonic/gin"
)


type RepoKtg interface {
	GetKtg(ctx *gin.Context) ([]models.Kategori, error)
	CreateKtg(data models.Kategori) error
	UpdateKtg(data models.Kategori, param string) error
	DeleteKtg(param string) error
}

type UsecaseKtg interface {
	GetKtg(ctx *gin.Context) ([]models.Kategori, error)
	CreateKtg(ctx *gin.Context) error
	UpdateKtg(ctx *gin.Context) error
	DeleteKtg(ctx *gin.Context) error
}