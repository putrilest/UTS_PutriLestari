package repository

import (
	"projek-restoran/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RepoCust struct {
	DB *gorm.DB
}

func NewRepoCust(Db *gorm.DB) *RepoCust {
	return &RepoCust{
		DB: Db,
	}
}

func (r RepoCust) GetAll(ctx *gin.Context) ([]models.Customer, error) {
	var customers []models.Customer
	r.DB.Preload("Reservasi").Find(&customers)

	return customers, nil
}

func (r RepoCust) Create(data models.Customer) error {
	err := r.DB.Create(&data).Error

	if err != nil {
		return err
	}
	return nil
}

func (rp RepoCust) Update(data models.Customer, param string) error {
	err := rp.DB.First(&models.Customer{}, "id_cust=?", param).Error

	if err != nil {
		return err
	}

	err = rp.DB.Where("id_cust = ?", param).Updates(&data).Error

	if err != nil {
		return err
	}
	return nil
}

func (r RepoCust) Delete(param string) error {

	err := r.DB.First(&models.Customer{}, "id_cust=?", param).Error
	if err != nil {
		return err
	}

	err = r.DB.Delete(&models.Customer{}, r.DB.Where("id_cust = ?", param)).Error	
	if err != nil {
		return err
	}

	return nil
}