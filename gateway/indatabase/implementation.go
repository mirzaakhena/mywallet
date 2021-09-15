package indatabase

import (
	"context"
	"errors"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
	"mywallet/domain/entity"
	"mywallet/infrastructure/database"
	"mywallet/infrastructure/log"
)

type prodGateway struct {
	*database.GormWithoutTrxImpl
	*database.GormWithTrxImpl
}

// NewProdGateway ...
func NewProdGateway(db *gorm.DB) *prodGateway {

	_ = db.AutoMigrate(&User{})
	_ = db.AutoMigrate(&Wallet{})
	_ = db.AutoMigrate(&Card{})
	_ = db.AutoMigrate(&CardSpendHistory{})

	return &prodGateway{
		GormWithoutTrxImpl: database.NewGormWithoutTrxImpl(db),
		GormWithTrxImpl:    database.NewGormWithTrxImpl(db),
	}
}

func (r *prodGateway) GenerateID(_ context.Context) string {
	id, _ := gonanoid.Generate("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", 10)
	return id
}

func (r *prodGateway) SaveUser(ctx context.Context, obj *entity.User) error {

	objToSave := User{
		ID:   obj.ID,
		Name: obj.Name,
	}

	return r.commonSaving(ctx, &objToSave)
}

func (r *prodGateway) SaveWallet(ctx context.Context, obj *entity.Wallet) error {

	objToSave := Wallet{
		ID:      obj.ID,
		Name:    obj.Name,
		UserID:  obj.User.ID,
		Balance: float64(obj.Balance),
	}

	return r.commonSaving(ctx, &objToSave)
}

func (r *prodGateway) SaveCard(ctx context.Context, walletId string, c *entity.Card) error {

	objToSave := Card{
		ID:            c.ID,
		WalletID:      walletId,
		Name:          c.Name,
		LimitAmount:   float64(c.LimitAmount),
		LimitDuration: string(c.LimitDuration),
	}

	return r.commonSaving(ctx, &objToSave)
}

func (r *prodGateway) SaveCardSpendHistory(ctx context.Context, obj *entity.CardSpendHistory) error {

	objToSave := CardSpendHistory{
		ID:               obj.ID,
		UserID:           obj.User.ID,
		CardID:           obj.Card.ID,
		Amount:           float64(obj.Amount),
		BalanceRemaining: float64(obj.BalanceRemaining),
		Date:             obj.Date,
	}

	return r.commonSaving(ctx, &objToSave)
}

func (r *prodGateway) UpdateWalletBalance(ctx context.Context, obj *entity.Wallet) error {

	db, err := database.ExtractDB(ctx)
	if err != nil {
		log.Error(ctx, err.Error())
		return err
	}

	err = db.Model(&Wallet{}).Where("id = ?", obj.ID).Update("balance", obj.Balance).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *prodGateway) FindUserByID(ctx context.Context, userID string) (*entity.User, error) {

	db, err := database.ExtractDB(ctx)
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}

	var user entity.User
	err = db.First(&user, "id = ?", userID).Error
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}

	return &user, nil
}

func (r *prodGateway) FindAllUser(ctx context.Context) ([]*entity.User, error) {

	db, err := database.ExtractDB(ctx)
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}

	var users []*entity.User
	err = db.Find(&users).Error
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}

	return users, nil

}

func (r *prodGateway) FindAllWalletByUser(ctx context.Context, userID string) ([]*entity.Wallet, error) {
	db, err := database.ExtractDB(ctx)
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}

	var wallets []*entity.Wallet
	err = db.Preload("Cards").Find(&wallets, "user_id = ?", userID).Error
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}

	return wallets, nil
}

func (r *prodGateway) FindAllCardSpendHistory(ctx context.Context, userID string) ([]*entity.CardSpendHistory, error) {

	db, err := database.ExtractDB(ctx)
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}

	var cardSpendHistoryObj []*entity.CardSpendHistory
	err = db.Order("date desc").Group("card_id").Find(&cardSpendHistoryObj, "user_id = ?", userID).Error
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}

	return cardSpendHistoryObj, nil

}

func (r *prodGateway) FindWalletByID(ctx context.Context, walletID string) (*entity.Wallet, error) {
	db, err := database.ExtractDB(ctx)
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}

	var wallet entity.Wallet
	err = db.Preload("User").Preload("Cards").First(&wallet, "id = ?", walletID).Error
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}

	return &wallet, nil
}

func (r *prodGateway) FindLastCardSpendHistory(ctx context.Context, cardID string) (*entity.CardSpendHistory, error) {
	db, err := database.ExtractDB(ctx)
	if err != nil {
		log.Error(ctx, err.Error())
		return nil, err
	}

	var cardSpendHistory entity.CardSpendHistory
	err = db.Order("date desc").First(&cardSpendHistory, "id = ?", cardID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &cardSpendHistory, nil
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
