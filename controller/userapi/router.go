package userapi

import (
	"mywallet/usecase/addnewcard"
	"mywallet/usecase/addnewwallet"
	"mywallet/usecase/registeruser"
	"mywallet/usecase/showalluser"
	"mywallet/usecase/showuserwalletinfo"
	"mywallet/usecase/spendmoney"
	"mywallet/usecase/topupwallet"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Router                   gin.IRouter
	RegisterUserInport       registeruser.Inport
	AddNewCardInport         addnewcard.Inport
	AddNewWalletInport       addnewwallet.Inport
	ShowAllUserInport        showalluser.Inport
	ShowUserWalletInfoInport showuserwalletinfo.Inport
	SpendMoneyInport         spendmoney.Inport
	TopupWalletInport        topupwallet.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	r.Router.POST("/user", r.authorized(), r.registerUserHandler(r.RegisterUserInport))
	r.Router.GET("/user", r.authorized(), r.showAllUserHandler(r.ShowAllUserInport))

	r.Router.POST("/wallet", r.authorized(), r.addNewWalletHandler(r.AddNewWalletInport))
	r.Router.GET("/wallet/:walletId", r.authorized(), r.showUserWalletInfoHandler(r.ShowUserWalletInfoInport))
	r.Router.POST("/wallet/:walletId/topupwallet", r.authorized(), r.topupWalletHandler(r.TopupWalletInport))
	r.Router.POST("/wallet/:walletId/card/:cardId/spendmoney", r.authorized(), r.spendMoneyHandler(r.SpendMoneyInport))
	r.Router.POST("/wallet/:walletId/card", r.authorized(), r.addNewCardHandler(r.AddNewCardInport))
}
