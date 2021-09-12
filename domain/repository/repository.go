package repository

import (
	"context"
	"mywallet/domain/entity"
)

type SaveCardSpendHistoryRepo interface {
	SaveCardSpendHistory(ctx context.Context, obj *entity.CardSpendHistory) error
}

// SaveUserRepo will save user into storage and guarantee it will be unique
type SaveUserRepo interface {
	SaveUser(ctx context.Context, obj *entity.User) error
}

type FindAllUserRepo interface {
	FindAllUser(ctx context.Context) ([]*entity.User, error)
}

type FindAllWalletByUserRepo interface {
	FindAllWalletByUser(ctx context.Context, userID string) ([]*entity.Wallet, error)
}

type SaveWalletRepo interface {
	SaveWallet(ctx context.Context, obj *entity.Wallet) error
}

type SaveCardRepo interface {
	SaveCard(ctx context.Context, walletID string, newCard *entity.Card) error
}

type FindUserByIDRepo interface {
	FindUserByID(ctx context.Context, userID string) (*entity.User, error)
}

type FindWalletByIDRepo interface {
	FindWalletByID(ctx context.Context, walletID string) (*entity.Wallet, error)
}

type UpdateWalletBalanceRepo interface {
	UpdateWalletBalance(ctx context.Context, obj *entity.Wallet) error
}

type FindLastCardSpendHistoryRepo interface {
	FindLastCardSpendHistory(ctx context.Context, cardID string) (*entity.CardSpendHistory, error)
}

type FindAllCardSpendHistoryRepo interface {
	FindAllCardSpendHistory(ctx context.Context, userID string) ([]*entity.CardSpendHistory, error)
}
