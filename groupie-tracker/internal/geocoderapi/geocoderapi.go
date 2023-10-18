package geocoderapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"groupie-tracker/internal/structs"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type response struct {
	Features []feature `json:"features"`
}

type feature struct {
	ID        string   `json:"id"`
	Type      string   `json:"type"`
	PlaceType []string `json:"place_type"`
	Geometry  geometry `json:"geometry"` // Add a Geometry field
}

type geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

var mapBoxToken = ""

// used to throttle the functions
var rateLimitReached = false
var rateLimitReachedTime int64 = 0

func SetMapBoxToken(token string) error {
	if len(token) == 0 || len(strings.TrimSpace(token)) == 0 {
		return errors.New("bad token format")
	}
	mapBoxToken = token
	return nil
}

func GeoCode(relations []*structs.Relation) bool {
	for _, relation := range relations {
		var g *structs.Geolocation = nil
		if !rateLimitReached {
			g = getCoords(relation.City, relation.Country)
		} else if rateLimitReached && time.Now().UTC().Unix()-rateLimitReachedTime > 120 {
			rateLimitReached = false
			rateLimitReachedTime = 0
			g = getCoords(relation.City, relation.Country)
		} else {
			log.Println("GeoCode(): Rate limit still in cooldown")
			return false
		}
		if g == nil {
			log.Printf("GeoCode(): Failed while getting Coordinates for (%+v)\n", relation)
			continue
		}
		relation.Loc = g
	}
	return true
}

func getCoords(city, country string) *structs.Geolocation {
	// query needs to be escaped
	query := url.QueryEscape(city + " " + country)
	URL := fmt.Sprintf("https://api.mapbox.com/geocoding/v5/mapbox.places/%s.json?limit=1&access_token=%s", query, mapBoxToken)

	resp, err := http.Get(URL)
	if err != nil {
		log.Printf("getCoords(): %s\n", err)
		return nil
	}
	defer resp.Body.Close()
	// 401 == Bad token
	if resp.StatusCode == http.StatusUnauthorized {
		log.Println("getCoords(): Bad MapBox Token")
		return nil
	}
	// 429 == Rate limit
	if resp.StatusCode == http.StatusTooManyRequests {
		log.Println("getCoords(): Rate limit reached")
		rateLimitReached = true
		rateLimitReachedTime = time.Now().Unix()
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		log.Printf("getCoords(): Unhandled Status Code %d\n", resp.StatusCode)
		return nil
	}
	buff, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("getCoords(): %s\n", err)
		return nil
	}

	// Parse the JSON data into a Response struct
	var response response
	if err := json.Unmarshal(buff, &response); err != nil {
		log.Println("getCoords(): Unmarshal error:", err)
		return nil
	}

	if len(response.Features) > 0 {
		lon := response.Features[0].Geometry.Coordinates[0]
		lat := response.Features[0].Geometry.Coordinates[1]
		return &structs.Geolocation{Lat: lat, Lon: lon}
	}
	return nil
}

func GetMap(relations []*structs.Relation) string {
	if !GeoCode(relations) {
		return ""
	}
	pinsArr := []string{}
	for i, r := range relations {
		if r.Loc == nil {
			continue
		}
		pin := fmt.Sprintf("pin-s-%d+555555(%f,%f)", i, r.Loc.Lon, r.Loc.Lat)
		pinsArr = append(pinsArr, pin)
	}
	pinsString := strings.Join(pinsArr, ",")

	URL := fmt.Sprintf("https://api.mapbox.com/styles/v1/mapbox/streets-v12/static/%s/auto/400x300@2x?access_token=%s", pinsString, mapBoxToken)
	return URL
}
