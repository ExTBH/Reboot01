package server

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"regexp"
	"sandbox/internal/database"
	"sandbox/internal/helpers"
	"sandbox/internal/structs"
	"sync"
	"time"
	"os"
)

var usernameMatcher = regexp.MustCompile(`^/user/(\w+)$`)

func profileHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAcceptedMethod(w, r, http.MethodGet) {
		return
	}
	if !usernameMatcher.MatchString(r.URL.Path) {
		errorServer(w, r, http.StatusNotFound, true)
		return
	}
	// get the user id
	matchedUsername := usernameMatcher.FindStringSubmatch(r.URL.Path)[1]
	// get the user
	user, err := database.GetUserByUsername(matchedUsername)
	if err != nil {
		log.Printf("profileHandler: %s\n", err.Error())
		errorServer(w, r, http.StatusInternalServerError, true)
		return
	}
	if user == nil {
		errorServer(w, r, http.StatusNotFound, true)
		return
	}
	err = writeToJson(user, w)
	if err != nil {
		log.Printf("ProfileHandler: %s\n", err.Error())
		errorServer(w, r, http.StatusInternalServerError, true)
		return
	}
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAcceptedMethod(w, r, http.MethodGet) {
		return
	}
	file, err := os.Open("www/template/signup.html")
	if err != nil {
		log.Printf("SignupHandler: %s\n", err.Error())
		errorServer(w, r, http.StatusInternalServerError, true)
		return
	}
	fi, err := file.Stat()
	if err != nil {
		log.Printf("signupHandler: %s\n", err.Error())
		errorServer(w, r, http.StatusInternalServerError, true)
		return
	}
	http.ServeContent(w, r, fi.Name(), fi.ModTime(), file)

	// dont take requests over 1 megabyte
	if r.ContentLength > 1024 {
		http.Error(w, "Request too large", http.StatusRequestEntityTooLarge)
		return
	}

	var userData structs.UserRequest
	if !parseBody(&userData, w, r) {
		errorServer(w, r, http.StatusBadRequest, true)
	}

	// check password length
	if len(userData.Password) < 8 {
		http.Error(w, "Password must be at least 8 characters long", http.StatusBadRequest)
		return
	}

	// check if the username exists
	exist, err := database.CheckExistance("User", "username", userData.Username)
	if err != nil {
		errorServer(w, r, http.StatusInternalServerError, true)
		return
	}

	if exist {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	// check if the email is valid and exists
	if !helpers.IsValidEmail(userData.Email) {
		http.Error(w, "Invalid email", http.StatusBadRequest)
		return
	}

	exist, err = database.CheckExistance("User", "email", userData.Email)
	if err != nil {
		errorServer(w, r, http.StatusInternalServerError, true)
		return
	}

	if exist {
		http.Error(w, "Email already exists", http.StatusConflict)
		return
	}
	imageID, err := database.UploadImage(userData.Image)
	if err != nil {
		log.Printf("SignupHandler: %s\n", err.Error())
	}
	cleanedUserData := structs.User{
		Username:       userData.Username,
		Email:          userData.Email,
		FirstName:      userData.FirstName,
		LastName:       userData.LastName,
		DateOfBirth:    time.Unix(userData.DateOfBirth, 0),
		HashedPassword: userData.Password,
		ImageId:        imageID,
		GithubName:     userData.GithubName,
		LinkedinName:   userData.LinkedinName,
		TwitterName:    userData.TwitterName,
	}
	err = database.CreateUser(cleanedUserData)
	if err != nil {
		errorServer(w, r, http.StatusInternalServerError, false)
		return
	}
}

// var sessionStore = make(map[string]*database.User)
var sessionStoreMutex = &sync.Mutex{}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAcceptedMethod(w, r, http.MethodGet) {
		return
	}
	file, err := os.Open("www/template/signin.html")
	if err != nil {
		log.Printf("loginHandler: %s\n", err.Error())
		errorServer(w, r, http.StatusInternalServerError, true)
		return
	}
	fi, err := file.Stat()
	if err != nil {
		log.Printf("loginHandler: %s\n", err.Error())
		errorServer(w, r, http.StatusInternalServerError, true)
		return
	}
	http.ServeContent(w, r, fi.Name(), fi.ModTime(), file)

	// serve only small requests to avoid abuse and serve error if the request is too large
	if r.ContentLength > 1024 {
		http.Error(w, "Request too large", http.StatusRequestEntityTooLarge)
		return
	}

	var loginData structs.UserRequest
	if !parseBody(&loginData, w, r) {
		errorServer(w, r, http.StatusBadRequest, true)
		return
	}
	exist := false
	var err1 error
	// check if the email is valid and exists
	if helpers.IsValidEmail(loginData.Username) {
		exist, err1 = database.CheckExistance("User", "email", loginData.Username)
		if err1 != nil {
			errorServer(w, r, http.StatusInternalServerError, true)
			return
		}
	}

	var user *structs.User
	if !exist {
		user, err = database.GetUserByUsername(loginData.Username)
		if err != nil {
			errorServer(w, r, http.StatusInternalServerError, true)
			return
		}
	}

	if (!exist && user == nil) || user.HashedPassword != loginData.Password {
		http.Error(w, "Invalid username or email or password", http.StatusConflict)
		return
	}

	if !user.BannedUntil.IsZero() && user.BannedUntil.After(time.Now()) {
		http.Error(w, "User is blocked", http.StatusForbidden)
		return
	}

	// Create a new session ID
	sID := sessionID()

	// Create a new cookie
	cookie := &http.Cookie{
		Name:  "session_token",
		Value: sID,
		// Secure: true,
		HttpOnly: true,
	}

	// Set the cookie on the client's browser
	http.SetCookie(w, cookie)

	// In a real-world application, you'd also want to store this
	// session ID in a server-side session store and associate it
	// with the authenticated user's ID.

}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAcceptedMethod(w, r, http.MethodGet) {
		return
	}
	panic("stub: not implemented")
}

func sessionID() string {
	b := make([]byte, 32)
	io.ReadFull(rand.Reader, b)
	return base64.StdEncoding.EncodeToString(b)
}
