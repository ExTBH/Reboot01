package server

import (
	"io"
	"log"
	"net/http"
	"regexp"
	"sandbox/internal/database"
	"sandbox/internal/helpers"
	"strconv"
)

var imageMatcher = regexp.MustCompile(`^/uploads/images/(\d+)$`)

const maxImageSize = 20971520 // 20 MB

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// valiate the request
	if !IsAcceptedMethod(w, r, http.MethodPost) {
		return
	}
	if r.ContentLength <= 0 || r.ContentLength > maxImageSize {
		errorServer(w, r, http.StatusRequestEntityTooLarge, false)
		return
	}
	// make sure the content is an image
	buff, err := io.ReadAll(r.Body)
	if err != nil {
		errorServer(w, r, http.StatusInternalServerError, false)
		return
	}
	if !helpers.IsDataImage(buff) {
		errorServer(w, r, http.StatusUnsupportedMediaType, false)
		return
	}
	// upload the image
	_, err = database.UploadImage(buff)
	if err != nil {
		log.Printf("uploadHandler: %s\n", err.Error())
		errorServer(w, r, http.StatusInternalServerError, false)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
func uploadedContentServerHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAcceptedMethod(w, r, http.MethodGet) {
		return
	}
	if imageMatcher.MatchString(r.URL.Path) {
		imageIDString := imageMatcher.FindStringSubmatch(r.URL.Path)[1]
		imageID, err := strconv.Atoi(imageIDString)
		if err != nil {
			errorServer(w, r, http.StatusBadRequest, true)
			return
		}
		imageData, err := database.GetImage(imageID)
		if err != nil {
			log.Printf("uploadedContentServerHandler: %s\n", err.Error())
			errorServer(w, r, http.StatusInternalServerError, true)
			return
		}

		if imageData == nil {
			errorServer(w, r, http.StatusNotFound, true)
			return
		}

		w.Write(imageData)
		return
	}
	errorServer(w, r, http.StatusNotFound, true)
}
