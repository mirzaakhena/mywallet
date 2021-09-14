package entity

import (
	"mywallet/application/apperror"
	"mywallet/domain/vo"
	"time"
)

type CardSpendHistory struct {
	ID               string
	User             *User
	UserID           string
	Card             *Card
	CardID           string
	Amount           vo.Money // currently it is not very useful here
	BalanceRemaining vo.Money
	Date             time.Time
}

func (c *CardSpendHistory) IsEmptyBalance() bool {
	return c.BalanceRemaining == 0
}

func (c *CardSpendHistory) IsStillPossibleToSpend(now time.Time) (bool, error) {

	if c.Card.LimitDuration == vo.DailyLimitTimeEnum {
		stillPossibleToSpend, err := c.inTheSameDay(c.Date, now)
		if err != nil {
			return false, err
		}

		return stillPossibleToSpend, nil
	}

	if c.Card.LimitDuration == vo.WeeklyLimitTimeEnum {
		stillPossibleToSpend, err := c.inTheSameWeek(c.Date, now)
		if err != nil {
			return false, err
		}
		return stillPossibleToSpend, nil

	}

	if c.Card.LimitDuration == vo.MonthlyLimitTimeEnum {
		stillPossibleToSpend, err := c.inTheSameMonth(c.Date, now)
		if err != nil {
			return false, err
		}

		return stillPossibleToSpend, nil
	}

	return false, apperror.UnrecognizedLimitTime

}

func (c *CardSpendHistory) inTheSameDay(lastDate time.Time, now time.Time) (bool, error) {

	err := c.validateNowAsFuture(lastDate, now)
	if err != nil {
		return false, err
	}

	y1, m1, d1 := lastDate.Date()
	y2, m2, d2 := now.Date()

	return y1 == y2 && m1 == m2 && d1 == d2, nil
}

func (c *CardSpendHistory) inTheSameWeek(lastDate time.Time, now time.Time) (bool, error) {

	err := c.validateNowAsFuture(lastDate, now)
	if err != nil {
		return false, err
	}

	y1, w1 := lastDate.ISOWeek()
	y2, w2 := lastDate.ISOWeek()

	return y1 == y2 && w1 == w2, nil
}

func (c *CardSpendHistory) inTheSameMonth(lastDate time.Time, now time.Time) (bool, error) {

	err := c.validateNowAsFuture(lastDate, now)
	if err != nil {
		return false, err
	}

	y1, m1, _ := lastDate.Date()
	y2, m2, _ := now.Date()

	return y1 == y2 && m1 == m2, nil
}

func (c *CardSpendHistory) validateNowAsFuture(lastDate time.Time, now time.Time) error {
	if now.Before(lastDate) {
		return apperror.DateNowMustFutureFromLastDate
	}
	return nil
}
