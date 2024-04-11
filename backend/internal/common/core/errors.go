package core

import (
	"errors"
	"strconv"

	"gorm.io/gorm"
)

type ApiError struct {
	Msg    string
	Status int
}

func (e *ApiError) Error() string {
	return e.Msg
}

func HandleError(err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		PanicApiError(err.Error(), 404)
	} else if errors.Is(err, gorm.ErrDuplicatedKey) {
		PanicApiError(err.Error(), 400)
	} else if _, ok := err.(*strconv.NumError); ok {
		PanicApiError(err.Error(), 400)
	} else if err != nil {
		panic(err)
	}
}

func PanicApiError(msg string, statuses ...int) {
	status := 422
	if len(statuses) > 0 {
		status = statuses[0]
	}
	panic(&ApiError{msg, status})
}
