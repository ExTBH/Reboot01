package server

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"sandbox/internal/database"
)

var homeTemplate = template.Must(template.ParseFiles("www/template/index.html"))

func GoLive(port string) {
	err := database.Connect("forum-db.sqlite")
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	defer database.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/", homepageHandler)
	mux.HandleFunc("/static/", staticHandler)
	mux.HandleFunc("uploads/", uploadedContentServerHandler)
	mux.HandleFunc("uploads/add", uploadHandler)
	// user routes
	mux.HandleFunc("/user/", profileHandler)
	mux.HandleFunc("/user/signup", signupHandler)
	mux.HandleFunc("/user/login", loginHandler)
	mux.HandleFunc("/user/logout", logoutHandler)
	mux.HandleFunc("/user/login/google", googleLoginHandler)
	mux.HandleFunc("/user/login/github", googleLoginHandler)
	// category routes
	mux.HandleFunc("/categories/", categoriesHandler)
	//  posts
	mux.HandleFunc("/posts/", postsHandler)
	log.Printf("Listening on http://localhost:%s\n", port)
	log.Println(http.ListenAndServe(":"+port, mux))
}

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAcceptedMethod(w, r, http.MethodGet) {
		return
	}
	if r.URL.Path != "/" {
		if r.URL.Path == "/favicon.ico" {
			http.ServeFile(w, r, "www/static/imgs/logo.png")
			return
		}
		errorServer(w, r, http.StatusNotFound, true)
		return
	}
	err := homeTemplate.Execute(w, nil)
	if err != nil {
		log.Printf("homepageHandler: %s\n", err.Error())
		errorServer(w, r, http.StatusInternalServerError, true)
		return
	}
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAcceptedMethod(w, r, http.MethodGet) {
		return
	}
	filePath := "www" + r.URL.Path
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("staticHandler: %s\n", err.Error())
		errorServer(w, r, http.StatusNotFound, true)
		return
	}
	fi, err := file.Stat()
	if err != nil {
		log.Printf("staticHandler: %s\n", err.Error())
		errorServer(w, r, http.StatusNotFound, true)
		return
	}
	if fi.IsDir() {
		errorServer(w, r, http.StatusNotFound, true)
		return
	}
	http.ServeContent(w, r, fi.Name(), fi.ModTime(), file)
}
