package usecase

import (
	"errors"
	"log"
	"net/http"
	"projek-restoran/models"
	"projek-restoran/resto"

	"github.com/gin-gonic/gin"
)

func NewUsecaseReservasi(Repo_reservasi resto.RepoReservasi) *UsecaseReservasi {
	return &UsecaseReservasi{Repo_reservasi}
}

type UsecaseReservasi struct {
	Repores resto.RepoReservasi
}

func (u UsecaseReservasi) GetReservasi(ctx *gin.Context) ([]models.Reservasi, error) {
	result, err := u.Repores.GetReservasi(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return nil, err
	}

	return result, nil
}

func (u UsecaseReservasi) CreateReservasi(ctx *gin.Context) error {
	var data models.Reservasi
	err := ctx.ShouldBindJSON(&data)

	if err != nil {
		return err
	}

	err = u.Repores.CreateReservasi(data)

	if err != nil {
		return err
	}

	return nil
}

func (u UsecaseReservasi) UpdateReservasi(ctx *gin.Context) error {
	var data models.Reservasi
	idd := ctx.Param("id_reservasi")

	if idd == "" {
		return errors.New("ID Not Found")
	}
	ctx.ShouldBindJSON(&data)
	u.Repores.UpdateReservasi(data, idd)
	return nil
}

func (u UsecaseReservasi) DeleteReservasi(ctx *gin.Context) error {
	idd := ctx.Param("id_reservasi")

	if idd == "" {
		log.Fatal("Id tidak ditemukan")
	}

	err := u.Repores.DeleteReservasi(idd)

	if err != nil {
		return err
	}

	return nil
}
