package hasher

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type HasherError string

const (
	HasherErrorNone                    HasherError = ""
	HasherErrorHashAndPasswordMismatch HasherError = "provided hash and password do not match"
	HasherErrorHashTooShort            HasherError = "provided hash is too short"
	HasherErrorPasswordTooLong         HasherError = "provided password is too long"
)

func errorToHasherError(err error) HasherError {
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return HasherErrorHashAndPasswordMismatch
	} else if errors.Is(err, bcrypt.ErrHashTooShort) {
		return HasherErrorHashTooShort
	} else if errors.Is(err, bcrypt.ErrPasswordTooLong) {
		return HasherErrorPasswordTooLong
	}
	return HasherErrorNone
}

// takes a password and returns a bcrypt hash of it
// max 72 chars
func GetHash(password string) (string, HasherError) {
	buff, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errorToHasherError(err)
	}
	return string(buff), HasherErrorNone
}

func CompareHashAndPassword(password, hashedPassword string) HasherError {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return errorToHasherError(err)
	}
	return HasherErrorNone
}
