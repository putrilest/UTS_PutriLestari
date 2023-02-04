package repository

import (
	"projek-restoran/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RepoKasir struct {
	DB *gorm.DB
}

func NewRepoKasir(Db *gorm.DB) *RepoKasir {
	return &RepoKasir{
		DB: Db,
	}
}

func (r RepoKasir) GetKasir(ctx *gin.Context) ([]models.Kasir, error) {
	var kasir []models.Kasir
	r.DB.Preload("Reservasi").Preload("Menu").Find(&kasir)

	return kasir, nil
}

func (r RepoKasir) CreateKasir(data models.Kasir) error {
	err := r.DB.Create(&data).Error

	if err != nil {
		return err
	}

	return nil
}

func (rp RepoKasir) UpdateKasir(data models.Kasir, param string) error {
	err := rp.DB.First(&models.Kasir{}, "id_kasir=?", param).Error

	if err != nil {
		return err
	}

	err = rp.DB.Where("id_kasir = ?", param).Updates(&data).Error

	if err != nil {
		return err
	}
	return nil
}

func (r RepoKasir) DeleteKasir(param string) error {

	err := r.DB.First(&models.Kasir{}, "id_kasir=?", param).Error
	if err != nil {
		return err
	}

	err = r.DB.Delete(&models.Kasir{}, r.DB.Where("id_kasir = ?", param)).Error
	if err != nil {
		return err
	}

	return nil
}
