package userapi

import (
  "mywallet/application/apperror"
  "mywallet/infrastructure/log"
  "mywallet/infrastructure/util"
  "mywallet/usecase/addnewwallet"
  "net/http"

  "github.com/gin-gonic/gin"
)

// addNewWalletHandler ...
func (r *Controller) addNewWalletHandler(inputPort addnewwallet.Inport) gin.HandlerFunc {

  type Request struct {
    WalletName    string
    CardName      string
    LimitAmount   float64
    LimitDuration string
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

    userID := c.Param("userID")

    req := addnewwallet.InportRequest{
      UserID:        userID,
      WalletName:    jsonReq.WalletName,
      CardName:      jsonReq.CardName,
      LimitAmount:   jsonReq.LimitAmount,
      LimitDuration: jsonReq.LimitDuration,
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
