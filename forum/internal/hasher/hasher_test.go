package hasher_test

import (
	"sandbox/internal/hasher"
	"testing"
)

func TestCompareHashAndPasswordValidPassword(t *testing.T) {
	rawPassword := "gigi7tyg776346yyr"
	hashh, err := hasher.GetHash(rawPassword)
	if err == hasher.HasherErrorPasswordTooLong {
		t.Error("reported password too long for a valid password")
	}
	if err != hasher.HasherErrorNone {
		t.Errorf("reported error: %s\n", err)
	}
	err = hasher.CompareHashAndPassword(rawPassword, hashh)
	if err == hasher.HasherErrorHashAndPasswordMismatch {
		t.Error("reported hash and password mismatch for a valid password")
	} else if err == hasher.HasherErrorHashTooShort {
		t.Errorf("reported hash too short for hash: %s\n", hashh)
	} else if err == hasher.HasherErrorPasswordTooLong {
		t.Errorf("reported password too long for password: %s\n", rawPassword)
	}
}

func TestGetHash72Password(t *testing.T) {
	// write a password of 73 characters
	rawPassword := "123456789012345678901234567890123456789012345678901234567890123456789123"
	_, err := hasher.GetHash(rawPassword)
	if err != hasher.HasherErrorNone {
		t.Errorf("reported '%s' for password of length: %d\n", err, len(rawPassword))
	}
}
func TestGetHash73Password(t *testing.T) {
	// write a password of 73 characters
	rawPassword := "1234567890123456789012345678901234567890123456789012345678901234567891234"
	_, err := hasher.GetHash(rawPassword)
	if err != hasher.HasherErrorPasswordTooLong {
		t.Errorf("reported '%s' for password of length: %d\n", err, len(rawPassword))
	}
}
