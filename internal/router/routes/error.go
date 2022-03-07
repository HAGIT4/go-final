package routes

import "fmt"

type bonusRouterContentTypeError struct {
	contentHeader string
}

func NewBonusRouterContentTypeError(contentHeader string) (err *bonusRouterContentTypeError) {
	return &bonusRouterContentTypeError{
		contentHeader: contentHeader,
	}
}

func (e *bonusRouterContentTypeError) Error() string {
	err := fmt.Sprintf("Wrong content type: %s", e.contentHeader)
	return err
}

type bonusRouterInternalServerError struct{}

func NewBonusRouterInternalServerError() (err *bonusRouterInternalServerError) {
	return &bonusRouterInternalServerError{}
}

func (e *bonusRouterInternalServerError) Error() string {
	return "Internal server error"
}
