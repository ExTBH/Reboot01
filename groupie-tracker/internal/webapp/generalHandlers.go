package webapp

import (
	"groupie-tracker/internal/bandsapi"
	"groupie-tracker/internal/filterhelper"
	"groupie-tracker/internal/geocoderapi"
	"groupie-tracker/internal/searchhelper"
	"groupie-tracker/internal/structs"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type rootData struct {
	SetSearch                                                            bool
	Bands                                                                []*structs.Band
	FilterCountries                                                      []string
	FilterMinCreation, FilterMaxCreation, FilterMinAlbum, FilterMaxAlbum int
	FilterBandSizes                                                      []int
}

func stylesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, r, ErrorTypeNotAllowed)
		return
	}
	if strings.Contains(r.URL.Path, "../") {
		ErrorHandler(w, r, ErrorTypeNotAllowed)
		return
	}
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
	w.Header().Set("Expires", "0")
	fixedPath := "../../www/styles/" + strings.TrimPrefix(r.URL.Path, "/styles/")
	fi, err := os.Stat(fixedPath)
	if err != nil {
		ErrorHandler(w, r, ErrorTypeNotFound)
		return
	}
	if fi.IsDir() {
		ErrorHandler(w, r, ErrorTypeDirForbidden)
		return
	}
	http.ServeFile(w, r, fixedPath)
}

func jsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, r, ErrorTypeNotAllowed)
		return
	}
	if strings.Contains(r.URL.Path, "../") {
		ErrorHandler(w, r, ErrorTypeNotAllowed)
		return
	}
	fixedPath := "../../www/js/" + strings.TrimPrefix(r.URL.Path, "/js/")
	fi, err := os.Stat(fixedPath)
	if err != nil {
		ErrorHandler(w, r, ErrorTypeNotFound)
		return
	}
	if fi.IsDir() {
		ErrorHandler(w, r, ErrorTypeDirForbidden)
		return
	}
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
	w.Header().Set("Expires", "0")
	http.ServeFile(w, r, fixedPath)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, r, ErrorTypeNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		ErrorHandler(w, r, ErrorTypeNotFound)
		return
	}
	bands := bandsapi.GetBands()
	if bands == nil && bandsapi.Fetching {
		ErrorHandler(w, r, ErrorTypeAPIBusy)
		return
	} else if bands == nil {
		ErrorHandler(w, r, ErrorTypeAPIDown)
		return
	}
	countries, minCreation, maxCreation, minAlbum, maxAlbum, bandSizes := filterhelper.CachedResults(bands)
	data := rootData{
		SetSearch:         true,
		Bands:             bands,
		FilterCountries:   countries,
		FilterMinCreation: minCreation,
		FilterMaxCreation: maxCreation,
		FilterMinAlbum:    minAlbum,
		FilterMaxAlbum:    maxAlbum,
		FilterBandSizes:   bandSizes,
	}
	tmplt, err := template.ParseFiles("../../www/templates/homepage.html", "../../www/templates/band_card.html")
	if err != nil {
		log.Printf("rootHandler(): %s", err)
		ErrorHandler(w, r, ErrorTypeInternal)
		return
	}
	err = tmplt.Execute(w, data)
	if err != nil {
		log.Printf("rootHandler(): %s", err)
		ErrorHandler(w, r, ErrorTypeInternal)
	}
}

// Handles searchbar stuff
func searchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if len(r.URL.Query()) != 1 || !r.URL.Query().Has("query") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	q := r.URL.Query().Get("query")
	if q == "" || strings.TrimSpace(q) == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	searchhelper.Search(w, bandsapi.GetBands(), strings.ToLower(q))
}

func bandHanlder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, r, ErrorTypeNotAllowed)
		return
	}
	if bandsapi.Fetching {
		ErrorHandler(w, r, ErrorTypeAPIBusy)
		return
	}
	// get id by remove `/band/`
	idStr := r.URL.Path[6:]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("bandHanlder(): Can't convert id %s\n", err)
		ErrorHandler(w, r, ErrorTypeBadRequest)
		return
	}
	band := bandsapi.GetBand(id)
	if band == nil {
		ErrorHandler(w, r, ErrorTypeNotFound)
		return
	}
	tmplt, err := template.ParseFiles("../../www/templates/band_page.html")
	if err != nil {
		log.Printf("bandHanlder(): %s", err)
		ErrorHandler(w, r, ErrorTypeInternal)
		return
	}
	bandData := struct {
		Band   structs.Band
		MapURL string
	}{
		Band:   *band,
		MapURL: geocoderapi.GetMap(band.Relations),
	}
	err = tmplt.Execute(w, bandData)
	if err != nil {
		log.Printf("bandHanlder(): %s", err)
		ErrorHandler(w, r, ErrorTypeInternal)
	}
}

func filterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, r, ErrorTypeNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		log.Println("filterHandler():", err)
		ErrorHandler(w, r, ErrorTypeBadRequest)
		return
	}
	// serve
	bands := bandsapi.GetBands()
	if bands == nil && bandsapi.Fetching {
		ErrorHandler(w, r, ErrorTypeAPIBusy)
		return
	} else if bands == nil {
		ErrorHandler(w, r, ErrorTypeAPIDown)
		return
	}
	// we have, [countries, creationMin, creationMax, albumMin, albumMax, members]
	if len(r.Form) > 6 {
		log.Println("filterHandler(): Extra form options found")
		ErrorHandler(w, r, ErrorTypeBadRequest)
		return
	}
	// countries filter
	if r.Form.Has("countries") {
		countries := r.Form["countries"]
		bands = filterhelper.FilterCountries(countries, bands)
	}
	// creation date filter
	if r.Form.Has("creationMin") && r.Form.Has("creationMax") {
		min, err := strconv.Atoi(r.Form.Get("creationMin"))
		max, err2 := strconv.Atoi(r.Form.Get("creationMax"))
		if (err != nil || err2 != nil) || min > max {
			log.Println("filterHandler(): Issue with creation")
			ErrorHandler(w, r, ErrorTypeBadRequest)
			return
		}
		bands = filterhelper.FilterCreation(min, max, bands)
	} else {
		ErrorHandler(w, r, ErrorTypeBadRequest)
		return
	}
	// first album filter
	if r.Form.Has("albumMin") && r.Form.Has("albumMax") {
		min, err := strconv.Atoi(r.Form.Get("albumMin"))
		max, err2 := strconv.Atoi(r.Form.Get("albumMax"))
		if (err != nil || err2 != nil) || min > max {
			log.Println("filterHandler(): Issue with album")
			ErrorHandler(w, r, ErrorTypeBadRequest)
			return
		}
		bands = filterhelper.FilterAlbum(min, max, bands)
	} else {
		ErrorHandler(w, r, ErrorTypeBadRequest)
		return
	}
	// first Team size
	if r.Form.Has("members") {
		members := []int{}
		for _, m := range r.Form["members"] {
			n, err := strconv.Atoi(m)
			if err != nil {
				ErrorHandler(w, r, ErrorTypeBadRequest)
				return
			}
			members = append(members, n)
		}
		bands = filterhelper.FilterMembers(members, bands)
	}

	countries, minCreation, maxCreation, minAlbum, maxAlbum, bandSizes := filterhelper.CachedResults(bands)
	data := rootData{
		SetSearch:         true,
		Bands:             bands,
		FilterCountries:   countries,
		FilterMinCreation: minCreation,
		FilterMaxCreation: maxCreation,
		FilterMinAlbum:    minAlbum,
		FilterMaxAlbum:    maxAlbum,
		FilterBandSizes:   bandSizes,
	}

	tmplt, err := template.ParseFiles("../../www/templates/homepage.html", "../../www/templates/band_card.html")
	if err != nil {
		log.Println("filterHandler():", err)
		ErrorHandler(w, r, ErrorTypeInternal)
		return
	}
	err = tmplt.Execute(w, data)
	if err != nil {
		log.Println("filterHandler():", err)
		ErrorHandler(w, r, ErrorTypeInternal)
	}
}
