package webapp

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type ErrorType string

const (
	ErrorTypeNotFound          ErrorType = "NOT_FOUND"          // ErrorTypeNotFound indicates a resource was not found.
	ErrorTypeBadRequest        ErrorType = "BAD_REQUEST"        // ErrorTypeBadRequest indicates a bad client request.
	ErrorTypeNotAllowed        ErrorType = "NOT_ALLOWED"        // ErrorTypeNotAllowed indicates the client's method is not allowed.
	ErrorTypeInternal          ErrorType = "INTERNAL"           // ErrorTypeInternal indicates an internal server error.
	ErrorTypeRecursiveInternal ErrorType = "RECURSIVE_INTERNAL" // ErrorTypeRecursiveInternal indicates a recursive internal error.
	ErrorTypeDirForbidden      ErrorType = "DIR_FORBIDDEN"      // ErrorTypeDirForbidden indiates the client is attempting to access a forbidden directory.
	ErrorTypeAPIBusy           ErrorType = "API_BUSY"           // ErrorTypeAPIBusy indicates the Bands API reqeust in In-Progress
	ErrorTypeAPIDown           ErrorType = "API_DOWN"           // ErrorTypeAPIDown indicates a failure white calling the Bands API
	// templates paths
	templatePathError string = "../../www/templates/errorpage.html"
)

type errorEvent struct {
	Type        ErrorType
	Path        string
	Method      string
	Image       string
	TimeOfError int64
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, e ErrorType) {
	log.Printf("%s Failed with %s for Path (%s)\n", r.RemoteAddr, e, r.URL.Path)
	// prevent Error nesting
	if e == ErrorTypeRecursiveInternal {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	event := &errorEvent{
		Type:        e,
		Path:        r.URL.Path,
		Method:      r.Method,
		Image:       "https://i.giphy.com/media/nodFMgUbTSTAINqvMq/giphy.webp",
		TimeOfError: time.Now().UTC().Unix(),
	}
	switch e {
	case ErrorTypeNotFound:
		w.WriteHeader(http.StatusNotFound)
	case ErrorTypeBadRequest:
		w.WriteHeader(http.StatusBadRequest)
	case ErrorTypeNotAllowed:
		w.WriteHeader(http.StatusMethodNotAllowed)
	case ErrorTypeInternal:
		w.WriteHeader(http.StatusInternalServerError)
	case ErrorTypeAPIDown:
		w.WriteHeader(http.StatusInternalServerError)
	case ErrorTypeDirForbidden:
		w.WriteHeader(http.StatusForbidden)
	}
	// server template
	tmplt, err := template.ParseFiles(templatePathError)
	if err != nil {
		log.Printf("ErrorHandler(): Template Parsing Failed (%s) (%+v)\n", err, event)
		ErrorHandler(w, r, ErrorTypeRecursiveInternal)
		return
	}
	if err = tmplt.Execute(w, event); err != nil {
		log.Printf("ErrorHandler(): Template Executing Failed (%s) (%+v)\n", err, event)
		ErrorHandler(w, r, ErrorTypeRecursiveInternal)
		return
	}
}
