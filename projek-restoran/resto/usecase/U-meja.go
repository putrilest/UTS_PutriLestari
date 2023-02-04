package usecase

import (
	"errors"
	"log"
	"net/http"
	"projek-restoran/models"
	"projek-restoran/resto"

	"github.com/gin-gonic/gin"
)

func NewUsecaseMeja(Repo_meja resto.RepoMeja) *UsecaseMeja {
	return &UsecaseMeja{Repo_meja}
}

type UsecaseMeja struct {
	Repomj resto.RepoMeja
}

func (u UsecaseMeja) GetMeja(ctx *gin.Context) ([]models.Meja, error) {
	result, err := u.Repomj.GetMeja(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return nil, err
	}

	return result, nil
}

func (u UsecaseMeja) CreateMeja(ctx *gin.Context) error {
	var data models.Meja
	err := ctx.ShouldBindJSON(&data)

	if err != nil {
		return err
	}

	err = u.Repomj.CreateMeja(data)

	if err != nil {
		return err
	}

	return nil
}

func (u UsecaseMeja) UpdateMeja(ctx *gin.Context) error {
	var data models.Meja
	idd := ctx.Param("id_meja")

	if idd == "" {
		return errors.New("ID Not Found")
	}
	ctx.ShouldBindJSON(&data)
	u.Repomj.UpdateMeja(data, idd)
	return nil
}

func (u UsecaseMeja) DeleteMeja(ctx *gin.Context) error {
	idd := ctx.Param("id_meja")

	if idd == "" {
		log.Fatal("Id Not Found")
	}

	err := u.Repomj.DeleteMeja(idd)

	if err != nil {
		return err
	}

	return nil
}
