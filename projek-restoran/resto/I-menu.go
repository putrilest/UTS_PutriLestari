package resto

import (
	"projek-restoran/models"
	"github.com/gin-gonic/gin"
)


type RepoMenu interface {
	GetMenu(ctx *gin.Context) ([]models.Menu, error)
	CreateMenu(data models.Menu) error
	UpdateMenu(data models.Menu, param string) error
	DeleteMenu(param string) error
}

type UsecaseMenu interface {
	GetMenu(ctx *gin.Context) ([]models.Menu, error)
	CreateMenu(ctx *gin.Context) error
	UpdateMenu(ctx *gin.Context) error
	DeleteMenu(ctx *gin.Context) error
}