package userapi

import (
  "mywallet/infrastructure/log"
  "mywallet/infrastructure/util"
  "mywallet/usecase/showuserwalletinfo"
  "net/http"

  "github.com/gin-gonic/gin"
)

// showUserWalletInfoHandler ...
func (r *Controller) showUserWalletInfoHandler(inputPort showuserwalletinfo.Inport) gin.HandlerFunc {

  type Card struct {
    ID            string
    Name          string
    LimitAmount   float64
    LimitDuration string
  }

  type Wallet struct {
    ID      string
    Name    string
    UserID  string
    Balance float64
    Cards   []Card
  }

  type Response struct {
    Wallets []Wallet
  }

  return func(c *gin.Context) {

    ctx := log.Context(c.Request.Context())

    var req showuserwalletinfo.InportRequest
    req.UserID = c.Param("userID")

    log.Info(ctx, util.MustJSON(req))

    res, err := inputPort.Execute(ctx, req)
    if err != nil {
      log.Error(ctx, err.Error())
      c.JSON(http.StatusBadRequest, NewErrorResponse(err))
      return
    }

    var jsonRes Response

    for _, wallet := range res.Wallets {

      var cards []Card

      for _, card := range wallet.Cards {
        cards = append(cards, Card{
          ID:            card.ID,
          Name:          card.Name,
          LimitAmount:   float64(card.LimitAmount),
          LimitDuration: string(card.LimitDuration),
        })
      }

      jsonRes.Wallets = append(jsonRes.Wallets, Wallet{
        ID:      wallet.ID,
        Name:    wallet.Name,
        UserID:  wallet.UserID,
        Balance: float64(wallet.Balance),
        Cards:   cards,
      })
    }

    log.Info(ctx, util.MustJSON(jsonRes))
    c.JSON(http.StatusOK, NewSuccessResponse(jsonRes))

  }
}
