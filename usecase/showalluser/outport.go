package showalluser

import "mywallet/domain/repository"

// Outport of ShowAllUser
type Outport interface {
	repository.FindAllUserRepo
	repository.WithoutTransactionDB
}
