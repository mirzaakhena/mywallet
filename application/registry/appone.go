package registry

import (
  "context"
  "mywallet/application"
  "mywallet/controller/userapi"
  "mywallet/gateway/inmemory"
  "mywallet/infrastructure/log"
  "mywallet/infrastructure/server"
  "mywallet/usecase/addnewcard"
  "mywallet/usecase/addnewwallet"
  "mywallet/usecase/registeruser"
  "mywallet/usecase/showalluser"
  "mywallet/usecase/showuserwalletinfo"
  "mywallet/usecase/spendmoney"
  "mywallet/usecase/topupwallet"
  "os"
)

type appone struct {
  server.GinHTTPHandler
  userapiController userapi.Controller
}

func NewAppone() func() application.RegistryContract {
  return func() application.RegistryContract {

    httpHandler, err := server.NewGinHTTPHandler(":8080")
    if err != nil {
      log.Error(context.Background(), "%v", err.Error())
      os.Exit(1)
    }

    datasource, err := inmemory.NewProdGateway()
    if err != nil {
      log.Error(context.Background(), "%v", err.Error())
      os.Exit(1)
    }

    return &appone{
      GinHTTPHandler: httpHandler,
      userapiController: userapi.Controller{
        Router:                   httpHandler.Router,
        RegisterUserInport:       registeruser.NewUsecase(datasource),
        AddNewCardInport:         addnewcard.NewUsecase(datasource),
        AddNewWalletInport:       addnewwallet.NewUsecase(datasource),
        ShowAllUserInport:        showalluser.NewUsecase(datasource),
        ShowUserWalletInfoInport: showuserwalletinfo.NewUsecase(datasource),
        SpendMoneyInport:         spendmoney.NewUsecase(datasource),
        TopupWalletInport:        topupwallet.NewUsecase(datasource),
      },
    }

  }
}

func (r *appone) SetupController() {
  r.userapiController.RegisterRouter()
}
