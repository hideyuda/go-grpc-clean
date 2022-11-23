package entity

import "errors"

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
		message = "server error"
	} else if errors.Is(err, ErrDBError) {
		code = 500
		message = "db error"

	} else if errors.Is(err, ErrRequestError) {
		code = 400
		message = "request error"
	} else if errors.Is(err, ErrNotFound) {
		code = 404
		message = "not found"
	} else if errors.Is(err, ErrDuplicateEntry) {
		code = 409
		message = "duplicate entry"
	} else if errors.Is(err, ErrFirebaseExpiredToken) {
		code = 400
		message = "firebase token expired"
	} else if errors.Is(err, ErrFirebaseInvalidToken) {
		code = 400
		message = "firebase invalid token"
	} else if errors.Is(err, ErrFirebaseFailedToVerify) {
		code = 400
		message = "firebase failed to verify"
	} else if errors.Is(err, ErrFirebaseFutureIssued) {
		code = 400
		message = "request error"
	} else if errors.Is(err, ErrFirebaseEmailExists) {
		code = 400
		message = "firebase email already exists"
	} else {
		code = 500
		message = "unknown"
	}
	return
}
