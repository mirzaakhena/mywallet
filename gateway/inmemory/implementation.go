package inmemory

import (
  "context"
  "fmt"
  "mywallet/domain/entity"
)

// TODO it is not finish yet...

type prodGateway struct {
  users   []entity.User
  wallets []entity.Wallet
}

// NewProdGateway ...
func NewProdGateway() (*prodGateway, error) {
  return &prodGateway{}, nil
}

func (r *prodGateway) GenerateID(ctx context.Context) string {

  return ""
}

func (r *prodGateway) GetDatabase(ctx context.Context) (context.Context, error) {
  return ctx, nil
}

func (r *prodGateway) BeginTransaction(ctx context.Context) (context.Context, error) {
  return ctx, nil
}

func (r *prodGateway) CommitTransaction(ctx context.Context) error {
  return nil
}

func (r *prodGateway) RollbackTransaction(ctx context.Context) error {
  return nil
}

func (r *prodGateway) SaveUser(ctx context.Context, obj *entity.User) error {

  for _, u := range r.users {
    if u.ID == obj.ID {
      return fmt.Errorf("user with id %s already exists", obj.ID)
    }
  }

  r.users = append(r.users, *obj)

  return nil
}

func (r *prodGateway) SaveWallet(ctx context.Context, obj *entity.Wallet) error {

  for _, u := range r.wallets {
    if u.ID == obj.ID {
      return fmt.Errorf("wallet with id %s already exists", obj.ID)
    }
  }

  r.wallets = append(r.wallets, *obj)

  return nil
}

func (r *prodGateway) SaveCard(ctx context.Context, walletId string, c *entity.Card) error {

  for _, u := range r.wallets {
    if u.ID == walletId {

      u.Cards = append(u.Cards, c)

      return nil
    }
  }

  return nil
}

func (r *prodGateway) SaveCardSpendHistory(ctx context.Context, obj *entity.CardSpendHistory) error {

  return nil
}

func (r *prodGateway) UpdateWalletBalance(ctx context.Context, obj *entity.Wallet) error {

  return nil
}

func (r *prodGateway) FindUserByID(ctx context.Context, userID string) (*entity.User, error) {

  return nil, nil
}

func (r *prodGateway) FindAllUser(ctx context.Context) ([]*entity.User, error) {

  return nil, nil
}

func (r *prodGateway) FindAllWalletByUser(ctx context.Context, userID string) ([]*entity.Wallet, error) {

  return nil, nil
}

func (r *prodGateway) FindAllCardSpendHistory(ctx context.Context, someID string) ([]*entity.CardSpendHistory, error) {

  return nil, nil
}

func (r *prodGateway) FindWalletByID(ctx context.Context, someID string) (*entity.Wallet, error) {

  return nil, nil
}

func (r *prodGateway) FindLastCardSpendHistory(ctx context.Context, cardID string) (*entity.CardSpendHistory, error) {

  return nil, nil
}
