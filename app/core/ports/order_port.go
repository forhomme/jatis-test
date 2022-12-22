package ports

import "jatis-test/app/core/models"

type OrderInputPort interface {
	GetOrderDetailsById(id int) (resp *models.OrderDetailResponse, err error)
}
