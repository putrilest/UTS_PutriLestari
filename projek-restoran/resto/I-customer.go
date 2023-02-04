package resto

import (
	"projek-restoran/models"
	"github.com/gin-gonic/gin"
)


type RepoCust interface {
	GetAll(ctx *gin.Context) ([]models.Customer, error)
	Create(data models.Customer) error
	Update(data models.Customer, param string) error
	Delete(param string) error
}

type UsecaseCust interface {
	GetAll(ctx *gin.Context) ([]models.Customer, error)
	Create(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}