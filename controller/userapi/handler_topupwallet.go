package userapi

import (
	"mywallet/application/apperror"
	"mywallet/infrastructure/log"
	"mywallet/infrastructure/util"
	"mywallet/usecase/topupwallet"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// topupWalletHandler ...
func (r *Controller) topupWalletHandler(inputPort topupwallet.Inport) gin.HandlerFunc {

	type Request struct {
		Amount   float64
	}

	return func(c *gin.Context) {

		ctx := log.Context(c.Request.Context())

		var jsonReq Request
		if err := c.BindJSON(&jsonReq); err != nil {
			newErr := apperror.FailUnmarshalResponseBodyError
			log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, NewErrorResponse(newErr))
			return
		}

		req := topupwallet.InportRequest{
			UserID:   c.Param("userID"),
			WalletID: c.Param("walletID"),
			Amount:   jsonReq.Amount,
			Date:     time.Now(),
		}

		log.Info(ctx, util.MustJSON(req))

		res, err := inputPort.Execute(ctx, req)
		if err != nil {
			log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, NewErrorResponse(err))
			return
		}

		log.Info(ctx, util.MustJSON(res))
		c.JSON(http.StatusOK, NewSuccessResponse(res))

	}
}
