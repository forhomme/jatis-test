package models

import (
	"jatis-test/app/utils"
	"strings"
)

type Products struct {
	ProductID   int     `json:"product_id" gorm:"primaryKey"`
	ProductName string  `json:"product_name"`
	UnitPrice   float64 `json:"unit_price"`
	InStock     int     `json:"in_stock"`
}

func (Products) TableName() string {
	return "products"
}

func (Products) ModuleName() string {
	return "products"
}

func (Products) Validate(in *Products) (err error) {
	sb := strings.Builder{}
	if in.InStock < 0 && in.InStock > 1 {
		sb.WriteString("in stock should be 0 or 1, ")
	}
	err = utils.ErrorChecker(sb)
	return
}
