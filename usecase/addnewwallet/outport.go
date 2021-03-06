package addnewwallet

import (
	"mywallet/domain/repository"
	"mywallet/domain/service"
)

// Outport of AddNewWallet
type Outport interface {
	repository.SaveWalletRepo
	repository.SaveCardRepo
	repository.FindUserByIDRepo
	repository.WithTransactionDB
	service.GenerateIDService
}
