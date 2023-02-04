package repository

import (
	"projek-restoran/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RepoMeja struct {
	DB *gorm.DB
}

func NewRepoMeja(Db *gorm.DB) *RepoMeja {
	return &RepoMeja{
		DB: Db,
	}
}

func (r RepoMeja) GetMeja(ctx *gin.Context) ([]models.Meja, error) {
	var mejas []models.Meja
	r.DB.Find(&mejas)

	return mejas, nil
}

func (r RepoMeja) CreateMeja(data models.Meja) error {
	err := r.DB.Create(&data).Error

	if err != nil {
		return err
	}
	return nil
}

func (rp RepoMeja) UpdateMeja(data models.Meja, param string) error {
	err := rp.DB.First(&models.Meja{}, "id_meja=?", param).Error

	if err != nil {
		return err
	}

	err = rp.DB.Where("id_meja = ?", param).Updates(&data).Error

	if err != nil {
		return err
	}
	return nil
}

func (r RepoMeja) DeleteMeja(param string) error {

	err := r.DB.First(&models.Meja{}, "id_meja=?", param).Error
	if err != nil {
		return err
	}

	err = r.DB.Delete(&models.Meja{}, r.DB.Where("id_meja = ?", param)).Error	
	if err != nil {
		return err
	}

	return nil
}