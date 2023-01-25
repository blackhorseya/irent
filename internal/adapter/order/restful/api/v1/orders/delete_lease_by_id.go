package orders

import (
	"net/http"

	"github.com/blackhorseya/irent/internal/pkg/errorx"
	"github.com/blackhorseya/irent/pkg/contextx"
	om "github.com/blackhorseya/irent/pkg/entity/domain/order/model"
	"github.com/blackhorseya/irent/pkg/httpheaders"
	"github.com/blackhorseya/irent/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type cancelLeaseByIDRequest struct {
	ID string `uri:"id" binding:"required"`
}

// CancelLeaseByID
// @Summary Cancel a booking by id
// @Description Cancel a booking by id
// @Tags Orders
// @Accept application/json
// @Produce application/json
// @Param id path string true "ID of booking"
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{data=string}
// @Failure 400 {object} er.Error
// @Failure 500 {object} er.Error
// @Router /v1/orders/{id} [delete]
func (i *impl) CancelLeaseByID(c *gin.Context) {
	ctx, ok := c.MustGet(string(contextx.KeyCtx)).(contextx.Contextx)
	if !ok {
		_ = c.Error(errorx.ErrContextx)
		return
	}

	token, ok := c.MustGet(string(httpheaders.KeyToken)).(string)
	if !ok {
		ctx.Error(errorx.ErrMissingToken.Error())
		_ = c.Error(errorx.ErrMissingToken)
		return
	}

	var req cancelLeaseByIDRequest
	err := c.ShouldBindUri(&req)
	if err != nil {
		ctx.Error(errorx.ErrMissingID.Error(), zap.Error(err))
		_ = c.Error(errorx.ErrMissingID)
		return
	}
	target := &om.Lease{No: req.ID}

	from, err := i.auth.GetByAccessToken(ctx, token)
	if err != nil {
		_ = c.Error(err)
		return
	}

	err = i.biz.CancelLease(ctx, from, target)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(target.No))
}
