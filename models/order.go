package models

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Order struct {
	OrderID      uint      `gorm:"primarykey" json:"orderId"`
	CustomerName string    `gorm:"type:varchar(255);not null;index" json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Items        []Item    `gorm:"foreignKey:OrderID;references:OrderID"`
}

type Item struct {
	ItemID      uint      `gorm:"primarykey" json:"itemId"`
	ItemCode    string    `gorm:"type:varchar(255);not null;index" json:"itemCode"`
	Description string    `gorm:"type:text" json:"description"`
	Quantity    uint      `json:"quantity"`
	OrderID     uint      `gorm:"not null;index"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (o *Order) CreateOrder() (*Order, error) {
	db := GetDB()

	if err := db.Create(&o).Error; err != nil {
		return &Order{}, err
	}

	return o, nil
}

func GetAllData() ([]Order, error) {
	db := GetDB()

	var orders []Order

	if err := db.Preload("Items").Find(&orders).Error; err != nil {
		return []Order{}, err
	}

	return orders, nil
}

func (o *Order) UpdateOrderPatch() (*Order, error) {
	db := GetDB()

	if err := db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&o).Error; err != nil {
		return &Order{}, err
	}

	return o, nil
}

func (o *Order) UpdateOrderPut() (*Order, error) {
	db := GetDB()

	var orderToUpdate Order = *o
	orderToUpdate.Items = []Item{}
	var itemsToUpdate []Item = (*o).Items

	if err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "order_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"customer_name", "ordered_at"}),
	}).Create(&orderToUpdate).Error; err != nil {
		return &Order{}, err
	}

	if err = db.Model(&orderToUpdate).Where("order_id = ?", orderToUpdate.OrderID).Association("Items").Unscoped().Replace(&itemsToUpdate); err != nil {
		return &Order{}, err
	}

	orderToUpdate.Items = itemsToUpdate

	return &orderToUpdate, nil
}

func DeleteOrder(orderID uint) (uint, error) {
	db := GetDB()

	if err := db.Select("Items").Delete(&Order{OrderID: orderID}).Error; err != nil {
		return orderID, err
	}

	return orderID, nil
}
