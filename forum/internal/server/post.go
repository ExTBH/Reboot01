package server

import (
	"log"
	"net/http"
	"regexp"
	"sandbox/internal/database"
	"sandbox/internal/structs"
	"strconv"
)

var postsMatcher = regexp.MustCompile(`^/posts/(\d+)$`)

func mapReactionForPost(postID int) []structs.PostReactionResponse {
	// database.GetReactionType()
	return nil
}

func mapPosts(old *structs.Post) *structs.PostResponse {
	return &structs.PostResponse{
		Id:         old.Id,
		ParentId:   *old.ParentId,
		Title:      old.Title,
		Message:    old.Message,
		ImageURL:   "/uploads/images" + strconv.Itoa(old.ImageId),
		Categories: nil,
		Reactions:  mapReactionForPost(old.Id),
	}
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAcceptedMethod(w, r, http.MethodGet) {
		return
	}
	if !postsMatcher.MatchString(r.URL.Path) {
		errorServer(w, r, http.StatusNotFound, true)
		return
	}
	postIDString := postsMatcher.FindStringSubmatch(r.URL.Path)[1]

	postID, err := strconv.Atoi(postIDString)

	if err != nil {
		errorServer(w, r, http.StatusBadRequest, true)
		return
	}
	post, err := database.GetPost(postID)

	err = writeToJson(mapPosts(post), w)
	if err != nil {
		log.Printf("postsHandler: %s\n", err.Error())
		errorServer(w, r, http.StatusInternalServerError, true)
		return
	}
}
