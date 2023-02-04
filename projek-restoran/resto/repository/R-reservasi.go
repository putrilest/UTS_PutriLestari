package repository

import (
	"projek-restoran/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RepoReservasi struct {
	DB *gorm.DB
}

func NewRepoReservasi(Db *gorm.DB) *RepoReservasi {
	return &RepoReservasi{
		DB: Db,
	}
}

func (r RepoReservasi) GetReservasi(ctx *gin.Context) ([]models.Reservasi, error) {
	var reservasis []models.Reservasi
	r.DB.Preload("Meja").Find(&reservasis)

	return reservasis, nil
}

func (r RepoReservasi) CreateReservasi(data models.Reservasi) error {
	err := r.DB.Create(&data).Error

	if err != nil {
		return err
	}
	//trigger
	//r.DB.Model(&models.Meja{}).Where("id_meja = ?", data.Id_meja).Update("status", "BOOKED")

	return nil
}

func (rp RepoReservasi) UpdateReservasi(data models.Reservasi, param string) error {
	err := rp.DB.First(&models.Reservasi{}, "id_reservasi=?", param).Error

	if err != nil {
		return err
	}

	err = rp.DB.Where("id_reservasi = ?", param).Updates(&data).Error

	if err != nil {
		return err
	}
	return nil
}

func (r RepoReservasi) DeleteReservasi(param string) error {

	err := r.DB.First(&models.Reservasi{}, "id_reservasi=?", param).Error
	if err != nil {
		return err
	}

	err = r.DB.Delete(&models.Reservasi{}, r.DB.Where("id_reservasi = ?", param)).Error
	if err != nil {
		return err
	}

	return nil
}
