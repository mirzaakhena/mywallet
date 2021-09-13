package userapi

import (
	"github.com/gin-gonic/gin"
	"mywallet/usecase/addnewcard"
	"mywallet/usecase/addnewwallet"
	"mywallet/usecase/registeruser"
	"mywallet/usecase/showalluser"
	"mywallet/usecase/showuserwalletinfo"
	"mywallet/usecase/spendmoney"
	"mywallet/usecase/topupwallet"
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

	r.Router.POST("/user/:userID/wallet", r.authorized(), r.addNewWalletHandler(r.AddNewWalletInport))
	r.Router.GET("/user/:userID/wallet/", r.authorized(), r.showUserWalletInfoHandler(r.ShowUserWalletInfoInport))
	r.Router.POST("/user/:userID/wallet/:walletID/topupwallet", r.authorized(), r.topupWalletHandler(r.TopupWalletInport))
	r.Router.POST("/user/:userID/wallet/:walletID/card", r.authorized(), r.addNewCardHandler(r.AddNewCardInport))
	r.Router.POST("/user/:userID/wallet/:walletID/card/:cardID/spendmoney", r.authorized(), r.spendMoneyHandler(r.SpendMoneyInport))
}
