package service

type bonusServiceBadRequestError struct{}

func newBonusServiceBadRequestError() *bonusServiceBadRequestError {
	return &bonusServiceBadRequestError{}
}

func (*bonusServiceBadRequestError) Error() string {
	err := "Bad request"
	return err
}

type bonusServiceLoginInUseError struct{}

func newBonusServiceLoginInUseError() *bonusServiceLoginInUseError {
	return &bonusServiceLoginInUseError{}
}

func (*bonusServiceLoginInUseError) Error() string {
	err := "Login already in use"
	return err
}

type bonusServiceInternalError struct{}

func newBonusServiceInternalError() *bonusServiceInternalError {
	return &bonusServiceInternalError{}
}

func (*bonusServiceInternalError) Error() string {
	err := "Internal error"
	return err
}
