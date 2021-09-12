package addnewcard

import (
  "mywallet/domain/repository"
  "mywallet/domain/service"
)

// Outport of AddnewCard
type Outport interface {
  repository.SaveCardRepo
  repository.FindWalletByIDRepo
  repository.WithTransactionDB
  service.GenerateIDService
}
