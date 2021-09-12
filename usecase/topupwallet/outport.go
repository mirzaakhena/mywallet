package topupwallet

import "mywallet/domain/repository"

// Outport of TopupWallet
type Outport interface {
  repository.UpdateWalletBalanceRepo
  repository.FindWalletByIDRepo
  repository.WithTransactionDB
}
