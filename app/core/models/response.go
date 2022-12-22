package models

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type OrderDetailResponse struct {
	CustomerName   string         `json:"customer_name"`
	EmployeeName   string         `json:"employee_name"`
	ShippingMethod string         `json:"shipping_method"`
	Products       []ListProducts `json:"products"`
	TotalPayment   float64        `json:"total_payment"`
}

type ListProducts struct {
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	UnitPrice   float64 `json:"unit_price"`
	Quantity    int     `json:"quantity"`
	Discount    float64 `json:"discount"`
	SubTotal    float64 `json:"sub_total"`
}

type ResponseMessage struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SendResponse(ctx echo.Context, response ResponseMessage) error {
	if response.Code == 0 {
		response.Code = http.StatusOK
	}

	if len(response.Message) == 0 {
		response.Message = "success"
	}

	ctx.Response().Header().Set(echo.HeaderServer, "Mysf Product Catalog/1.0 (arch64)")
	return ctx.JSON(response.Code, response)
}
