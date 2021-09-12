package indatabase

import (
  "context"
  "gorm.io/gorm"
  "mywallet/domain/entity"
  "mywallet/infrastructure/database"
  "mywallet/infrastructure/log"
)

type prodGateway struct {
  *database.GormWithoutTrxImpl
  *database.GormWithTrxImpl
}

func NewProdGateway(db *gorm.DB) *prodGateway {
  return &prodGateway{
    GormWithoutTrxImpl: database.NewGormWithoutTrxImpl(db),
    GormWithTrxImpl:    database.NewGormWithTrxImpl(db),
  }
}

func (r *prodGateway) SaveUser(ctx context.Context, obj *entity.User) error {
  return r.commonSaving(ctx, obj)
}

func (r *prodGateway) SaveWallet(ctx context.Context, obj *entity.Wallet) error {
  return r.commonSaving(ctx, obj)
}

func (r *prodGateway) SaveCard(ctx context.Context, walletId string, c *entity.Card) error {
  return r.commonSaving(ctx, c)
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

func (r *prodGateway) commonSaving(ctx context.Context, obj interface{}) error {
  db, err := database.ExtractDB(ctx)
  if err != nil {
    log.Error(ctx, err.Error())
    return err
  }

  err = db.Save(obj).Error
  if err != nil {
    log.Error(ctx, err.Error())
    return err
  }
  return nil
}
