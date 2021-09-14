package showuserwalletinfo

import "mywallet/domain/repository"

// Outport of ShowAllUserWallet
type Outport interface {
	repository.FindAllWalletByUserRepo
	repository.FindAllCardSpendHistoryRepo
	repository.WithoutTransactionDB
}
