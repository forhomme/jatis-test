package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"jatis-test/app/core/models"
	"jatis-test/app/core/ports"
)

type dbConnection struct {
	db *gorm.DB
}

func NewDBConnection(db *gorm.DB) ports.DatabaseOutputPort {
	return &dbConnection{db: db}
}

func (d *dbConnection) GetOrderById(id int) (resp *models.Orders, err error) {
	err = d.db.Model(&models.Orders{}).Where("order_id = ?", id).Preload(clause.Associations).Find(&resp).Error
	return
}

func (d *dbConnection) GetOrderDetailsByOrderId(id int) (resp []*models.OrderDetails, err error) {
	err = d.db.Model(&models.OrderDetails{}).Where("order_id = ?", id).Preload(clause.Associations).Find(&resp).Error
	return
}
