package addnewwallet

import "mywallet/domain/repository"

// Outport of AddNewWallet
type Outport interface {
	repository.SaveWalletRepo
	repository.FindUserByIDRepo
}
