package registeruser

import "mywallet/domain/repository"

// Outport of RegisterUser
type Outport interface {
	repository.SaveUserRepo
	repository.WithTransactionDB
}
