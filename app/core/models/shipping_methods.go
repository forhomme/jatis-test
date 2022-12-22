package models

type ShippingMethods struct {
	ShippingMethodID int    `json:"shipping_method_id" gorm:"primaryKey"`
	ShippingMethod   string `json:"shipping_method"`
}

func (ShippingMethods) TableName() string {
	return "shipping_methods"
}

func (ShippingMethods) ModuleName() string {
	return "shipping_methods"
}
