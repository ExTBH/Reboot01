package bandsapi

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/internal/structs"
	"io"
	"log"
	"net/http"
	"strings"
)

const (
	artistsURL   string = "https://groupietrackers.herokuapp.com/api/artists"
	relationsURL string = "https://groupietrackers.herokuapp.com/api/relation"
)

// Struct for handeling json and converting it to a `structs.Relation`
type relationResponse struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

var Fetching = false // A state monitor to know if the backend is attempting to fetch to not overload
var bands []*structs.Band = nil

// gets a `structs.Band` for the given id
func GetBand(id int) *structs.Band {
	if bands == nil {
		GetBands()
	}
	for _, band := range bands {
		if band.Id == id {
			return band
		}
	}
	return nil
}

// Gets an array of all bands
func GetBands() []*structs.Band {
	if Fetching {
		log.Println("GetBands(): Fetching already in progress")
		return nil
	}
	if bands != nil {
		return bands
	}
	log.Println("GetBands(): Server started Fetching")

	Fetching = true
	bands = getArtists()
	Fetching = false
	log.Println("GetBands(): Server done Fetching")
	//remove bad stuff
	for i, band := range bands {
		if band.Id == 21 {
			bands[i].Image = "https://e7.pngegg.com/pngimages/540/351/png-clipart-all-the-beautiful-girls-graphy-18-text-photography.png"
			break
		}
	}
	shahd := structs.Band{Id: 8448428, Image: "https://cdn.discordapp.com/attachments/1128534642023223406/1160255420233941064/IMG_7028.png?ex=6533ff10&is=65218a10&hm=38d3c8da91011ddffc590b088326d94831cd9c005cc934cc454f848f1ca11cb1&", Name: "shahdi", Members: []string{"Shahd Yusuf"}, CreationDate: 1997, FirstAlbum: "29-09-1998"}
	r1 := structs.Relation{City: "London", Country: "uk", Date: []string{"07-011-2003"}}
	shahd.Relations = append(shahd.Relations, &r1)
	getRelations(bands)
	bands = append(bands, &shahd)
	return bands
}

// Loads All Artists
func getArtists() []*structs.Band {
	resp, err := http.Get(artistsURL)
	if err != nil {
		log.Printf("getArtists(): %s\n", err)
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("getArtists(): HTTP Status Not Ok (%d)\n", resp.StatusCode)
		return nil
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("getArtists(): Body Reader err: %s", err)
		return nil
	}
	var bands []*structs.Band
	err = json.Unmarshal(body, &bands)
	if err != nil {
		log.Printf("getArtists(): Unmarshal err: %s", err)
		return nil
	}
	return bands
}

// Loads all relations
func getRelations(bands []*structs.Band) {
	for i, band := range bands {
		updatedRelation := getRelation(band)
		if updatedRelation == nil {
			log.Printf("getRelations(): got nil for ID %d\n", band.Id)
			continue
		}
		bands[i] = updatedRelation
	}
}

func getRelation(band *structs.Band) *structs.Band {
	bandRelationURL := fmt.Sprintf("%s/%d", relationsURL, band.Id)
	resp, err := http.Get(bandRelationURL)

	if err != nil {
		log.Printf("getRelation(): %s\n", err)
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("getRelation(): HTTP Status Not Ok (%d)\n", resp.StatusCode)
		return nil
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("getRelation(): Body Reader err: %s", err)
		return nil
	}

	var response relationResponse

	// Unmarshal the JSON response into the struct
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		log.Printf("getRelation(): Unmarshal err: %s", err)
		return nil
	}
	for location, dates := range response.DatesLocations {
		location = strings.ReplaceAll(location, "_", " ")
		locFields := strings.Split(location, "-")
		re := &structs.Relation{Date: dates, Loc: nil, City: locFields[0], Country: locFields[1]}
		band.Relations = append(band.Relations, re)
	}
	return band
}
