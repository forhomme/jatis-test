package models

type OrderDetails struct {
	OrderDetailID int      `json:"order_detail_id" gorm:"primaryKey"`
	OrderID       int      `json:"order_id"`
	Order         Orders   `gorm:"references:OrderID"`
	ProductID     int      `json:"product_id"`
	Product       Products `gorm:"references:ProductID"`
	Quantity      int      `json:"quantity"`
	UnitPrice     float64  `json:"unit_price"`
	Discount      float64  `json:"discount"`
}

func (OrderDetails) TableName() string {
	return "order_details"
}

func (OrderDetails) ModuleName() string {
	return "order_details"
}
