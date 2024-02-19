package server

import (
	"html/template"
	"log"
	"net/http"
)

// if `useHtml` is false it will respond with text
func errorServer(w http.ResponseWriter, r *http.Request, code int, useHtml bool) {
	w.WriteHeader(code)

	if useHtml {
		var tmplt *template.Template
		var err error
		switch code {
		case http.StatusNotFound:
			tmplt, err = template.ParseFiles("www/errors/not_found.html")
		case http.StatusInternalServerError:
			tmplt, err = template.ParseFiles("www/errors/internal.html")
		default:
			errorServer(w, r, code, false)
			return
		}
		if err != nil || (err == nil && tmplt.Execute(w, nil) != nil) {
			// fallback in case it erros
			errorServer(w, r, code, false)
		}
		return
	}

	switch code {
	case http.StatusNotFound:
		http.Error(w, "Resource Not Found", code)
	case http.StatusInternalServerError:
		http.Error(w, "Internal Server Error", code)
	default:
		log.Printf("errorServer: %d is not implemented\n", code)
		http.Error(w, "Something Went Wrong, Please Try Again Later", code)
	}
}
