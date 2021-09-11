package addnewcard

import "mywallet/domain/repository"

// Outport of AddnewCard
type Outport interface {
	repository.SaveWalletRepo
	repository.FindWalletByIDRepo
}
