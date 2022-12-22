package usecase

import (
	"errors"
	"jatis-test/app/core/models"
	"jatis-test/app/core/ports"
)

type ordersUC struct {
	repo ports.DatabaseOutputPort
}

func NewOrdersUC(repo ports.DatabaseOutputPort) ports.OrderInputPort {
	return &ordersUC{repo: repo}
}

func (o *ordersUC) GetOrderDetailsById(id int) (resp *models.OrderDetailResponse, err error) {
	order, err := o.repo.GetOrderById(id)
	if err != nil {
		return
	}
	if order.OrderID == 0 {
		err = errors.New("order not found")
		return
	}
	orderDetail, err := o.repo.GetOrderDetailsByOrderId(id)
	if err != nil {
		return
	}
	listProducts := make([]models.ListProducts, 0)
	total := float64(0)
	for _, value := range orderDetail {
		product := models.ListProducts{
			ProductID:   value.ProductID,
			ProductName: value.Product.ProductName,
			UnitPrice:   value.UnitPrice,
			Quantity:    value.Quantity,
			Discount:    value.Discount,
			SubTotal:    (value.UnitPrice * (100 - value.Discount) / 100) * float64(value.Quantity),
		}
		total = total + product.SubTotal
		listProducts = append(listProducts, product)
	}
	resp = &models.OrderDetailResponse{
		CustomerName:   order.Customer.FirstName + " " + order.Customer.LastName,
		EmployeeName:   order.Employee.FirstName + " " + order.Employee.LastName,
		ShippingMethod: order.ShippingMethod.ShippingMethod,
		Products:       listProducts,
		TotalPayment:   total,
	}
	return
}
