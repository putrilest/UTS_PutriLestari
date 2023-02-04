package usecase

import (
	"errors"
	"log"
	"net/http"
	"projek-restoran/models"
	"projek-restoran/resto"

	"github.com/gin-gonic/gin"
)

func NewUsecaseKasir(Repo_kasir resto.RepoKasir) *UsecaseKasir {
	return &UsecaseKasir{Repo_kasir}
}

type UsecaseKasir struct {
	Repokas resto.RepoKasir
}

func (u UsecaseKasir) GetKasir(ctx *gin.Context) ([]models.Kasir, error) {
	result, err := u.Repokas.GetKasir(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return nil, err
	}

	return result, nil
}

func (u UsecaseKasir) CreateKasir(ctx *gin.Context) error {
	var data models.Kasir
	err := ctx.ShouldBindJSON(&data)

	if err != nil {
		return err
	}

	err = u.Repokas.CreateKasir(data)

	if err != nil {
		return err
	}

	return nil
}

func (u UsecaseKasir) UpdateKasir(ctx *gin.Context) error {
	var data models.Kasir
	idd := ctx.Param("id_kasir")

	if idd == "" {
		return errors.New("ID Not Found")
	}
	ctx.ShouldBindJSON(&data)
	u.Repokas.UpdateKasir(data, idd)
	return nil
}

func (u UsecaseKasir) DeleteKasir(ctx *gin.Context) error {
	idd := ctx.Param("id_kasir")

	if idd == "" {
		log.Fatal("Id tidak ditemukan")
	}

	err := u.Repokas.DeleteKasir(idd)

	if err != nil {
		return err
	}

	return nil
}