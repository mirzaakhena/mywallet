package showalluserwallet

import "mywallet/domain/repository"

// Outport of ShowAllUserWallet
type Outport interface {
  repository.FindAllWalletByUserRepo
}
