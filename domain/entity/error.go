package entity

import (
	"errors"
	"fmt"
)

var (
	ErrServerError = errors.New("SERVER_ERROR")
	ErrDBError     = errors.New("DB_ERROR")

	ErrRequestError   = errors.New("REQUEST_ERROR")
	ErrNotFound       = errors.New("NOT_FOUND")
	ErrDuplicateEntry = errors.New("DUPLICATE_ENTRY")

	// User
	ErrUserNotFound = errors.New("USER_NOT_FOUND")

	// Firebase
	ErrFirebaseEmptyToken     = errors.New("FIREBASE_EMPTY_TOKEN")
	ErrFirebaseExpiredToken   = errors.New("FIREBASE_EXPIRED_TOKEN")
	ErrFirebaseInvalidToken   = errors.New("FIREBASE_INVALID_TOKEN")
	ErrFirebaseFailedToVerify = errors.New("FIREBASE_FAILED_TO_VERIFY")
	ErrFirebaseFutureIssued   = errors.New("FIREBASE_FUTURE_ISSUED")
	ErrFirebaseEmailExists    = errors.New("FIREBASE_EMAIL_EXISTS")
)

func ErrorInfo(err error) (code int, message string) {
	if errors.Is(err, ErrServerError) {
		code = 500
		message = fmt.Sprint("server error\n details:", err)
	} else if errors.Is(err, ErrDBError) {
		code = 500
		message = fmt.Sprint("db error\n details:", err)

	} else if errors.Is(err, ErrRequestError) {
		code = 400
		message = fmt.Sprint("request error\n details:", err)
	} else if errors.Is(err, ErrNotFound) {
		code = 404
		message = fmt.Sprint("not found\n details:", err)
	} else if errors.Is(err, ErrDuplicateEntry) {
		code = 409
		message = fmt.Sprint("duplicated entry\n details:", err)
	} else if errors.Is(err, ErrFirebaseExpiredToken) {
		code = 400
		message = fmt.Sprint("firebase expired token\n details:", err)
	} else if errors.Is(err, ErrFirebaseInvalidToken) {
		code = 400
		message = fmt.Sprint("firebase invalid token\n details:", err)
	} else if errors.Is(err, ErrFirebaseFailedToVerify) {
		code = 400
		message = fmt.Sprint("firebase failed to verify\n details:", err)
	} else if errors.Is(err, ErrFirebaseFutureIssued) {
		code = 400
		message = fmt.Sprint("firebase future issued\n details:", err)
	} else if errors.Is(err, ErrFirebaseEmailExists) {
		code = 400
		message = fmt.Sprint("firebase email already exsists\n details:", err)
	} else {
		code = 500
		message = fmt.Sprint("unknown\n details:", err)
	}
	return
}
