package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// todo:  مريم ادرى
// logOut functions && sign up >> نودي اليوزر لصفحة ثانية عشان يحط معلوماته (userName,FirstName,LastName,Email)
//  عندي فكرة لل redirect
// بس مشوار فلازم اسويلها فنكشن بروحها واحس بيصير في واجد فنكشنات وفوضى

// ServerFuntions
// mux.HandleFunc("/google-login", handleGoogleLogin)
// mux.HandleFunc("/google-callback", handleGoogleCallback)

// mux.HandleFunc("/github-login", handleGitHubLogin)
// mux.HandleFunc("/github-callback", handleGitHubCallback)

// mux.HandleFunc("/facebook-login", handleFacebookLogin)
// mux.HandleFunc("/facebook-callback", handleFacebookCallback)

const (
	GoogleClientID     = "494333147558-4fdt1969hq590gcuhm9qrpe0sf5c70rg.apps.googleusercontent.com"
	GoogleClientSecret = "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"
	GitHubClientID     = "0107a6ae1d498eb9dd65"
	GitHubClientSecret = "73dc2264fdf73812b816e7f8e4dc8085fc2ab570"
	FacebookAppID      = "912742480212850"
	FacebookAppSecret  = "030238c8ab765b60b384f9e7c754787e"
	GooglelogInURL     = "https://accounts.google.com/o/oauth2/auth?client_id=494333147558-4fdt1969hq590gcuhm9qrpe0sf5c70rg.apps.googleusercontent.com&redirect_uri=http://localhost:8080/google-callback&response_type=code&scope=email"
	GooglelogOutURL    = "https://accounts.google.com/logout"
	GitHubLogOutURL    = "https://github.com/logout"
	FaceBookLogOutURL  = "https://www.facebook.com/logout"
)

var (
	GitHubLogInURL   = fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=http://localhost:8080/github-callback", GitHubClientID)
	FaceBookLogInURL = fmt.Sprintf("https://www.facebook.com/v14.0/dialog/oauth?client_id=%s&redirect_uri=http://localhost:8080/facebook-callback&state=random_state_string", FacebookAppID)
)

type GoogleTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	IDToken      string `json:"id_token"`
}

type FacebookTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type GitHubTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
}

func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	// Redirect the user to the Google authentication page
	http.Redirect(w, r, GooglelogInURL, http.StatusSeeOther)
}

func HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	// Extract the authorization code from the query parameters
	code := r.URL.Query().Get("code")

	// Exchange the authorization code for an access token
	done, error := ExchangeGoogleCodeForToken(code)
	if error == nil {
		fmt.Println(done.AccessToken)

	}
	fmt.Println(code)
	// Redirect or respond as needed
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ExchangeGoogleCodeForToken(code string) (GoogleTokenResponse, error) {
	// Prepare the token request payload
	data := url.Values{}
	data.Set("code", code)
	data.Set("client_id", GoogleClientID)
	data.Set("client_secret", GoogleClientSecret)
	data.Set("redirect_uri", "http://localhost:8080/google-callback")
	data.Set("grant_type", "authorization_code")

	// Send the token request
	resp, err := http.PostForm("https://accounts.google.com/o/oauth2/token", data)
	if err != nil {
		return GoogleTokenResponse{}, err
	}
	defer resp.Body.Close()
	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return GoogleTokenResponse{}, fmt.Errorf("failed to exchange Google code for token. Status: %s", resp.Status)
	}

	// Parse the token response
	var tokenResp GoogleTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return GoogleTokenResponse{}, err
	}

	return tokenResp, nil
}

func HandleGitHubLogin(w http.ResponseWriter, r *http.Request) {
	// Redirect the user to the GitHub authentication page
	http.Redirect(w, r, GitHubLogInURL, http.StatusSeeOther)
}

func HandleGitHubCallback(w http.ResponseWriter, r *http.Request) {
	// Extract the authorization code from the query parameters
	code := r.URL.Query().Get("code")
	fmt.Println(code)
	done, _ := ExchangeGitHubCodeForToken(code)
	fmt.Println(done.AccessToken)

	// TODO: Use the access token as needed

	// Redirect or respond as needed
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ExchangeGitHubCodeForToken(code string) (GitHubTokenResponse, error) {
	// Prepare the token request payload
	data := url.Values{}
	data.Set("code", code)
	data.Set("client_id", GitHubClientID)
	data.Set("client_secret", GitHubClientSecret)
	data.Set("redirect_uri", "http://localhost:8080/github-callback")
	data.Set("grant_type", "authorization_code")

	// Send the token request
	resp, err := http.PostForm("https://github.com/login/oauth/access_token", data)
	if err != nil {
		return GitHubTokenResponse{}, err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return GitHubTokenResponse{}, fmt.Errorf("failed to exchange GitHub code for token. Status: %s", resp.Status)
	}

	// Parse the token response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return GitHubTokenResponse{}, err
	}

	tokenResp, err := parseGitHubTokenResponse(string(body))
	if err != nil {
		return GitHubTokenResponse{}, err
	}

	return tokenResp, nil
}

func parseGitHubTokenResponse(response string) (GitHubTokenResponse, error) {
	values, err := url.ParseQuery(response)
	if err != nil {
		return GitHubTokenResponse{}, err
	}

	accessToken := values.Get("access_token")
	tokenType := values.Get("token_type")
	scope := values.Get("scope")

	return GitHubTokenResponse{
		AccessToken: accessToken,
		TokenType:   tokenType,
		Scope:       scope,
	}, nil
}
func handleFacebookLogin(w http.ResponseWriter, r *http.Request) {
	// Redirect the user to the Facebook authentication page
	http.Redirect(w, r, FaceBookLogInURL, http.StatusSeeOther)
}

func handleFacebookCallback(w http.ResponseWriter, r *http.Request) {
	// Extract the authorization code from the query parameters
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")
	fmt.Println(code)
	fmt.Println(state)

	done, err := exchangeFacebookCodeForToken(code)
	if err != nil {
		// Handle error
		return
	}
	fmt.Println(done.AccessToken)

	// TODO: Use the access token as needed

	// Redirect or respond as needed
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func exchangeFacebookCodeForToken(code string) (FacebookTokenResponse, error) {
	// Prepare the token request payload
	data := url.Values{}
	data.Set("code", code)
	data.Set("client_id", FacebookAppID)
	data.Set("client_secret", FacebookAppSecret)
	data.Set("redirect_uri", "http://localhost:8080/facebook-callback")

	// Send the token request
	resp, err := http.PostForm("https://graph.facebook.com/v14.0/oauth/access_token", data)
	if err != nil {
		return FacebookTokenResponse{}, err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return FacebookTokenResponse{}, fmt.Errorf("failed to exchange Facebook code for token. Status: %s", resp.Status)
	}

	// Parse the token response
	var tokenResp FacebookTokenResponse
	err = json.NewDecoder(resp.Body).Decode(&tokenResp)
	if err != nil {
		return FacebookTokenResponse{}, err
	}

	return tokenResp, nil
}

func handleLogout(w http.ResponseWriter, r *http.Request, URL string) {
	//TODO :
	// delete the user sission/token from the user's struct

	// Redirect the user to the authentication logout page
	http.Redirect(w, r, URL, http.StatusTemporaryRedirect)
}
