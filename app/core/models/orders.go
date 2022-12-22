package models

import (
	"jatis-test/app/utils"
	"strings"
	"time"
)

type Orders struct {
	OrderID             int             `json:"order_id" gorm:"primaryKey"`
	CustomerID          int             `json:"customer_id"`
	Customer            Customers       `gorm:"references:CustomerID"`
	EmployeeID          int             `json:"employee_id"`
	Employee            Employees       `gorm:"references:EmployeeID"`
	OrderDate           time.Time       `json:"order_date"`
	PurchaseOrderNumber string          `json:"purchase_order_number"`
	ShipDate            time.Time       `json:"ship_date"`
	ShippingMethodID    int             `json:"shipping_method_id"`
	ShippingMethod      ShippingMethods `gorm:"references:ShippingMethodID"`
	FreightCharge       float64         `json:"freight_charge"`
	Taxes               float64         `json:"taxes"`
	PaymentReceived     int             `json:"payment_received"`
	Comment             string          `json:"comment"`
}

func (Orders) TableName() string {
	return "orders"
}

func (Orders) ModuleName() string {
	return "orders"
}

func (Orders) Validate(in *Orders) (err error) {
	sb := strings.Builder{}
	if in.PaymentReceived < 0 && in.PaymentReceived > 1 {
		sb.WriteString("payment received should be 0 or 1, ")
	}
	err = utils.ErrorChecker(sb)
	return
}
