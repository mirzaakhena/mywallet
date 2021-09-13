package registeruser

import (
  "mywallet/domain/repository"
  "mywallet/domain/service"
)

// Outport of RegisterUser
type Outport interface {
  repository.SaveUserRepo
  repository.WithTransactionDB
  service.GenerateIDService
}
