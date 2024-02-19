package sign

import (
	"sandbox/internal/database"
	// "sandbox/www"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) bool {
	// Check if the request is a POST request
	if r.Method != "POST" {
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			http.ServeFile(w, r, "sandbox/www/template/404Error.html")
		}
		// w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.ServeFile(w, r, "sandbox/www/template/405Error.html")
	}
	// Get the username and password from the request
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Use GetUserByUsername to get the user
	user, err := database.GetUserByUsername(username)
	if err != nil || user == nil {
		// Handle error
		return false
	}

	// Use GetPasswordHashForUserByUsername to get the hashed password
	hashedPassword, err := database.GetPasswordHashForUserByUsername(username)
	if err != nil || hashedPassword == "" {
		// Handle error
		return false
	}

	// Compare the hashed password with the given password
	if password != hashedPassword {
		// If they don't match, return false
		return false
	}

	// If everything is okay, return true
	return true
}
