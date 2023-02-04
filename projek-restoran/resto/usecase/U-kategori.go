package usecase

import (
	"errors"
	"log"
	"net/http"
	"projek-restoran/models"
	"projek-restoran/resto"

	"github.com/gin-gonic/gin"
)

func NewUsecaseKtg(Repo_ktg resto.RepoKtg) *UsecaseKtg {
	return &UsecaseKtg{Repo_ktg}
}

type UsecaseKtg struct {
	Repok resto.RepoKtg
}

func (u UsecaseKtg) GetKtg(ctx *gin.Context) ([]models.Kategori, error) {
	result, err := u.Repok.GetKtg(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return nil, err
	}

	return result, nil
}

func (u UsecaseKtg) CreateKtg(ctx *gin.Context) error {
	var data models.Kategori
	err := ctx.ShouldBindJSON(&data)

	if err != nil {
		return err
	}

	err = u.Repok.CreateKtg(data)

	if err != nil {
		return err
	}

	return nil
}

func (u UsecaseKtg) UpdateKtg(ctx *gin.Context) error {
	var data models.Kategori
	idd := ctx.Param("id_ktg")

	if idd == "" {
		return errors.New("ID Not Found")
	}
	ctx.ShouldBindJSON(&data)
	u.Repok.UpdateKtg(data, idd)
	return nil
}

func (u UsecaseKtg) DeleteKtg(ctx *gin.Context) error {
	idd := ctx.Param("id_ktg")

	if idd == "" {
		log.Fatal("Id tidak ditemukan")
	}

	err := u.Repok.DeleteKtg(idd)

	if err != nil {
		return err
	}

	return nil
}
