package api

import (
	"github.com/labstack/echo/v4"
	"jatis-test/app/core/models"
	"jatis-test/app/core/ports"
	"net/http"
	"strconv"
)

type orderApi struct {
	repo ports.OrderInputPort
}

func NewOrderApi(repo ports.OrderInputPort) *orderApi {
	return &orderApi{repo: repo}
}

// GetOrderDetails godoc
// @Summary Get Order Detail
// @Description Get Order Detail
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} utils.ResponseMessage
// @Failure 400 {object} utils.ResponseMessage
// @Failure 500 {object} utils.ResponseMessage
// @Router /engine/api/orders/{id} [get]
func (o *orderApi) GetOrderDetails(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	resp, err := o.repo.GetOrderDetailsById(id)
	if err != nil {
		return models.SendResponse(ctx, models.ResponseMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return models.SendResponse(ctx, models.ResponseMessage{
		Code:    http.StatusOK,
		Message: "success",
		Data:    resp,
	})
}
