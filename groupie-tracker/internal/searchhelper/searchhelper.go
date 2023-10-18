package searchhelper

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/internal/structs"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type searchItemType string
type searchItem struct {
	Type   searchItemType `json:"type"`
	BandId int            `json:"band_id"`
	Text   string         `json:"text"`
}

const (
	searchItemTypeBand         = "band"
	searchItemTypeMember       = "member"
	searchItemTypeLocation     = "location"
	searchItemTypeFirstAlbum   = "first_album"
	searchItemTypeCreationDate = "creation_date"
)

// searhes and replies back with matches as a json array
func Search(w http.ResponseWriter, bands []*structs.Band, query string) {
	items := []searchItem{}

	for _, band := range bands {
		// search by bans
		if strings.HasPrefix(strings.ToLower(band.Name), query) {
			items = append(items, searchItem{Type: searchItemTypeBand, BandId: band.Id, Text: band.Name})
		}
		// search by first album
		if strings.HasPrefix(band.FirstAlbum, query) {
			txt := fmt.Sprintf("%s - %s", band.FirstAlbum, band.Name)
			items = append(items, searchItem{Type: searchItemTypeFirstAlbum, BandId: band.Id, Text: txt})
		}
		// Search by creation date
		creationDate := strconv.Itoa(band.CreationDate)
		if strings.HasPrefix(creationDate, query) {
			txt := fmt.Sprintf("%s - %s", creationDate, band.Name)
			items = append(items, searchItem{Type: searchItemTypeFirstAlbum, BandId: band.Id, Text: txt})
		}
		// Search by Members
		for _, member := range band.Members {
			if strings.HasPrefix(strings.ToLower(member), query) {
				txt := fmt.Sprintf("%s - %s", member, band.Name)
				items = append(items, searchItem{Type: searchItemTypeMember, BandId: band.Id, Text: txt})
			}
		}
		// Search by location
		for _, relation := range band.Relations {
			// country
			if strings.HasPrefix(strings.ToLower(relation.Country), query) {
				txt := fmt.Sprintf("%s - %s, %s", band.Name, relation.City, relation.Country)
				items = append(items, searchItem{Type: searchItemTypeLocation, BandId: band.Id, Text: txt})
			}
			if strings.HasPrefix(strings.ToLower(relation.City), query) {
				txt := fmt.Sprintf("%s - %s, %s", band.Name, relation.City, relation.Country)
				items = append(items, searchItem{Type: searchItemTypeLocation, BandId: band.Id, Text: txt})
			}
		}
	}
	if len(items) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	itemsArr, err := json.Marshal(items)
	if err != nil {
		log.Println("Search():", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(itemsArr)
}
