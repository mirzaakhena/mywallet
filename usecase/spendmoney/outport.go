package spendmoney

import "mywallet/domain/repository"

// Outport of SpendMoney
type Outport interface {
	repository.FindWalletByIDRepo
	repository.SaveCardSpendHistoryRepo
	repository.UpdateWalletBalanceRepo
	repository.FindLastCardSpendHistoryRepo
	repository.WithTransactionDB
}
