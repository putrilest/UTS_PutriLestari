package resto

import (
	"projek-restoran/models"
	"github.com/gin-gonic/gin"
)


type RepoMeja interface {
	GetMeja(ctx *gin.Context) ([]models.Meja, error)
	CreateMeja(data models.Meja) error
	UpdateMeja(data models.Meja, param string) error
	DeleteMeja(param string) error
}

type UsecaseMeja interface {
	GetMeja(ctx *gin.Context) ([]models.Meja, error)
	CreateMeja(ctx *gin.Context) error
	UpdateMeja(ctx *gin.Context) error
	DeleteMeja(ctx *gin.Context) error
}