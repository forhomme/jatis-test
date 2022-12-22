package ports

import "jatis-test/app/core/models"

type DatabaseOutputPort interface {
	GetOrderById(id int) (resp *models.Orders, err error)
	GetOrderDetailsByOrderId(id int) (resp []*models.OrderDetails, err error)
}
