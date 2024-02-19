package server

import (
	"log"
	"net/http"
	"regexp"
	"sandbox/internal/database"
	"sandbox/internal/structs"
)

var categoryMatcher = regexp.MustCompile(`^/categories/(\w+)$`)

func mapCategories(old *structs.Category) *structs.CategoryResponse {
	return &structs.CategoryResponse{
		Name:        old.Name,
		Description: old.Description,
		Color:       old.Color,
		IconURL:     "",
	}
}

func categoriesHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAcceptedMethod(w, r, http.MethodGet) {
		return
	}

	arr, err := database.GetCategories()
	if err != nil {
		log.Printf("error getting categories: %s\n", err.Error())
		errorServer(w, r, http.StatusInternalServerError, false)
		return
	}

	mappedArr := make([]structs.CategoryResponse, len(arr))
	for i, v := range arr {
		mappedArr[i] = *mapCategories(&v)
	}

	err = writeToJson(mappedArr, w)
	if err != nil {
		log.Printf("error writing categories to json: %s\n", err.Error())
		errorServer(w, r, http.StatusInternalServerError, true)
		return
	}
}

func categoryPostsHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAcceptedMethod(w, r, http.MethodGet) {
		return
	}
	if !categoryMatcher.MatchString(r.URL.Path) {
		errorServer(w, r, http.StatusNotFound, true)
		return
	}
	categoryName := categoryMatcher.FindStringSubmatch(r.URL.Path)[1]

	posts_count, err := database.GetPostsCountByCategory(categoryName)
	if err != nil {
		log.Printf("error getting posts count by category: %s\n", err.Error())
		errorServer(w, r, http.StatusInternalServerError, false)
		return
	}
	posts, err := database.GetPostsByCategory(categoryName, posts_count, 0, "latest")
	if err != nil {
		log.Printf("error getting posts by category: %s\n", err.Error())
		errorServer(w, r, http.StatusInternalServerError, false)
		return
	}
	mappedPosts := make([]structs.PostResponse, len(posts))
	for i, v := range posts {
		mappedPosts[i] = *mapPosts(&v)
	}

	err = writeToJson(mappedPosts, w)
	if err != nil {
		log.Printf("error writing posts to json: %s\n", err.Error())
		errorServer(w, r, http.StatusInternalServerError, false)
		return
	}
}
