package repository

import (
	"projek-restoran/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RepoMenu struct {
	DB *gorm.DB
}

func NewRepoMenu(Db *gorm.DB) *RepoMenu {
	return &RepoMenu{
		DB: Db,
	}
}

func (r RepoMenu) GetMenu(ctx *gin.Context) ([]models.Menu, error) {
	var menus []models.Menu
	r.DB.Find(&menus)

	return menus, nil
}

func (r RepoMenu) CreateMenu(data models.Menu) error {
	err := r.DB.Create(&data).Error

	if err != nil {
		return err
	}
	return nil
}

func (rp RepoMenu) UpdateMenu(data models.Menu, param string) error {
	err := rp.DB.First(&models.Menu{}, "id_menu=?", param).Error

	if err != nil {
		return err
	}

	err = rp.DB.Where("id_menu = ?", param).Updates(&data).Error

	if err != nil {
		return err
	}
	return nil
}

func (r RepoMenu) DeleteMenu(param string) error {

	err := r.DB.First(&models.Menu{}, "id_menu=?", param).Error
	if err != nil {
		return err
	}

	err = r.DB.Delete(&models.Menu{}, r.DB.Where("id_menu = ?", param)).Error	
	if err != nil {
		return err
	}

	return nil
}