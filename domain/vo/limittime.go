package vo

import (
	"mywallet/application/apperror"
	"strings"
)

type LimitTime string

const (
	DailyLimitTimeEnum   LimitTime = "DAILY"
	WeeklyLimitTimeEnum  LimitTime = "WEEKLY"
	MonthlyLimitTimeEnum LimitTime = "MONTHLY"
)

var enumLimitTime = map[LimitTime]LimitTimeDetail{
	DailyLimitTimeEnum:   {},
	WeeklyLimitTimeEnum:  {},
	MonthlyLimitTimeEnum: {},
}

type LimitTimeDetail struct { //
}

func NewLimitTime(name string) (LimitTime, error) {
	name = strings.ToUpper(name)

	if _, exist := enumLimitTime[LimitTime(name)]; !exist {
		return "", apperror.UnrecognizedEnum.Var(name, "LimitDuration")
	}

	return LimitTime(name), nil
}

func (r LimitTime) GetDetail() LimitTimeDetail {
	return enumLimitTime[r]
}

func (r LimitTime) PossibleValues() []LimitTime {
	res := make([]LimitTime, 0)
	for key := range enumLimitTime {
		res = append(res, key)
	}
	return res
}
