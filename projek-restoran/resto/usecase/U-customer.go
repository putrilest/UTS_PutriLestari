package usecase

import (
	"errors"
	"log"
	"net/http"
	"projek-restoran/models"
	"projek-restoran/resto"

	"github.com/gin-gonic/gin"
)

func NewUsecaseCust(Repo_cust resto.RepoCust) *UsecaseCust {
	return &UsecaseCust{Repo_cust}
}

type UsecaseCust struct {
	Repoc resto.RepoCust
}

func (u UsecaseCust) GetAll(ctx *gin.Context) ([]models.Customer, error) {
	result, err := u.Repoc.GetAll(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return nil, err
	}

	return result, nil
}

func (u UsecaseCust) Create(ctx *gin.Context) error {
	var data models.Customer
	err := ctx.ShouldBindJSON(&data)

	if err != nil {
		return err
	}

	err = u.Repoc.Create(data)

	if err != nil {
		return err
	}

	return nil
}

func (u UsecaseCust) Update(ctx *gin.Context) error {
	var data models.Customer
	idd := ctx.Param("id_cust")

	if idd == "" {
		return errors.New("ID Not Found")
	}
	ctx.ShouldBindJSON(&data)
	u.Repoc.Update(data, idd)
	return nil
}

func (u UsecaseCust) Delete(ctx *gin.Context) error {
	idd := ctx.Param("id_cust")

	if idd == "" {
		log.Fatal("Id tidak ditemukan")
	}

	err := u.Repoc.Delete(idd)

	if err != nil {
		return err
	}

	return nil
}