package usecase

import (
	"errors"
	"log"
	"net/http"
	"projek-restoran/models"
	"projek-restoran/resto"

	"github.com/gin-gonic/gin"
)

func NewUsecaseMenu(Repo_menu resto.RepoMenu) *UsecaseMenu {
	return &UsecaseMenu{Repo_menu}
}

type UsecaseMenu struct {
	Repom resto.RepoMenu
}

func (u UsecaseMenu) GetMenu(ctx *gin.Context) ([]models.Menu, error) {
	result, err := u.Repom.GetMenu(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return nil, err
	}

	return result, nil
}

func (u UsecaseMenu) CreateMenu(ctx *gin.Context) error {
	var data models.Menu
	err := ctx.ShouldBindJSON(&data)

	if err != nil {
		return err
	}

	err = u.Repom.CreateMenu(data)

	if err != nil {
		return err
	}

	return nil
}

func (u UsecaseMenu) UpdateMenu(ctx *gin.Context) error {
	var data models.Menu
	idd := ctx.Param("id_menu")

	if idd == "" {
		return errors.New("ID Not Found")
	}
	ctx.ShouldBindJSON(&data)
	u.Repom.UpdateMenu(data, idd)
	return nil
}

func (u UsecaseMenu) DeleteMenu(ctx *gin.Context) error {
	idd := ctx.Param("id_menu")

	if idd == "" {
		log.Fatal("Id tidak ditemukan")
	}

	err := u.Repom.DeleteMenu(idd)

	if err != nil {
		return err
	}

	return nil
}