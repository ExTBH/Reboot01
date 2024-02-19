package helpers

import (
	"net/http"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(`^[\w]+@[\w]+\.[a-zA-Z]{2,}$`)

func IsValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func IsDataImage(buff []byte) bool {
	// the function that actually does the trick
	return strings.HasPrefix(http.DetectContentType(buff), "image")
}
