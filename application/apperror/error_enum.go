package apperror

const (
	FailUnmarshalResponseBodyError          ErrorType = "ER1000 Fail to unmarshal response body"       // used by controller
	ObjectNotFound                          ErrorType = "ER1001 Object %s is not found"                // used by injected repo in interactor
	UnrecognizedEnum                        ErrorType = "ER1002 %s is not recognized %s enum"          // used by enum
	DatabaseNotFoundInContextError          ErrorType = "ER1003 Database is not found in context"      // used by repoimpl
	UserNameMustNotEmpty                    ErrorType = "ER1000 user name must not empty"              //
	WalletNameMustNotEmpty                  ErrorType = "ER1000 wallet name must not empty"            //
	UserIsNotFound                          ErrorType = "ER1000 user is not found"                     //
	LimitAmountMustNotZero                  ErrorType = "ER1000 limit amount mus not zero"             //
	MoneyMustGreaterThanZero                ErrorType = "ER1000 money must greater than zero"          //
	CardUserNameMustNotEmpty                ErrorType = "ER1000 card user name must not empty"         //
	UserMustNotNil                          ErrorType = "ER1000 user must not nil"                     //
	CardMustNotNil                          ErrorType = "ER1000 card must not nil"                     //
	CardIDIsExist                           ErrorType = "ER1000 card id is exist"                      //
	UserIDDoesNotMatch                      ErrorType = "ER1000 user id does not match"                //
	CardLimitReachZero                      ErrorType = "ER1000 card limit reach zero"                 //
	AmountGreaterThanRemainingBalanceInCard ErrorType = "ER1000 amount greater than remaining balance" //
	AmountGreaterThanBalance                ErrorType = "ER1000 amount greater than balance"           //
	AmountGreaterThanLimitInCard            ErrorType = "ER1000 amount greater than limit in card"     //
	DateNowMustFutureFromLastDate           ErrorType = "ER1000 date now must future from last date"   //
	UnrecognizedLimitTime                   ErrorType = "ER1000 unrecognized limit time"               //
	UserIDMustNotEmpty                      ErrorType = "ER1000 user id must not empty"                //
)
