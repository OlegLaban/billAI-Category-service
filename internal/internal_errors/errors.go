package internalerrors

import "errors"

var (
	ErrorUserNotFound  = errors.New("User not found")
	ErrorCategoryExist = errors.New("Category exist")
)
