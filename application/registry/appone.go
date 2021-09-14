package registry

import (
	"mywallet/application"
	"mywallet/controller/userapi"
	"mywallet/gateway/indatabase"
	"mywallet/infrastructure/database"
	"mywallet/infrastructure/server"
	"mywallet/usecase/addnewcard"
	"mywallet/usecase/addnewwallet"
	"mywallet/usecase/registeruser"
	"mywallet/usecase/showalluser"
	"mywallet/usecase/showuserwalletinfo"
	"mywallet/usecase/spendmoney"
	"mywallet/usecase/topupwallet"
)

type appone struct {
	server.GinHTTPHandler
	userapiController userapi.Controller
}

func NewAppone() func() application.RegistryContract {
	return func() application.RegistryContract {

		httpHandler := server.NewGinHTTPHandlerDefault()
		db := database.NewGormDefault()

		datasource := indatabase.NewProdGateway(db)

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
