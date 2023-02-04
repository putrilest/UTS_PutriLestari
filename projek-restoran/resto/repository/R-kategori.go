package repository

import (
	"projek-restoran/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RepoKtg struct {
	DB *gorm.DB
}

func NewRepoKtg(Db *gorm.DB) *RepoKtg {
	return &RepoKtg{
		DB: Db,
	}
}

func (r RepoKtg) GetKtg(ctx *gin.Context) ([]models.Kategori, error) {
	var kategoris []models.Kategori
	r.DB.Preload("Menu").Find(&kategoris)

	return kategoris, nil
}

func (r RepoKtg) CreateKtg(data models.Kategori) error {
	err := r.DB.Create(&data).Error

	if err != nil {
		return err
	}
	return nil
}

func (rp RepoKtg) UpdateKtg(data models.Kategori, param string) error {
	err := rp.DB.First(&models.Kategori{}, "id_cust=?", param).Error

	if err != nil {
		return err
	}

	err = rp.DB.Where("id_ktg = ?", param).Updates(&data).Error

	if err != nil {
		return err
	}
	return nil
}

func (r RepoKtg) DeleteKtg(param string) error {

	err := r.DB.First(&models.Kategori{}, "id_ktg=?", param).Error
	if err != nil {
		return err
	}

	err = r.DB.Delete(&models.Kategori{}, r.DB.Where("id_ktg = ?", param)).Error	
	if err != nil {
		return err
	}

	return nil
}