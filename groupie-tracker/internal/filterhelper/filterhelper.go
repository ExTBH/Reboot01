package filterhelper

import (
	"groupie-tracker/internal/structs"
	"strconv"
	"strings"
)

var (
	countries                                                 []string
	minCreation, maxCreation, minAlbum, maxAlbum, maxBandSize int
)

func CachedResults(bands []*structs.Band) (countries []string, minCreation, maxCreation, minAlbum, maxAlbum int, bandSizes []int) {
	// Use meaningful variable names and initialize them with appropriate values.
	var uniqueCountriesMap = make(map[string]struct{})

	if countries == nil {
		for _, band := range bands {
			// Collect unique countries in a map.
			for _, relation := range band.Relations {
				uniqueCountriesMap[relation.Country] = struct{}{}
			}

			// Parse the album year from the FirstAlbum field.
			if parts := strings.Split(band.FirstAlbum, "-"); len(parts) == 3 {
				if albumYear, err := strconv.Atoi(parts[2]); err == nil {
					if albumYear > maxAlbum {
						maxAlbum = albumYear
					}
					if albumYear < minAlbum || minAlbum == 0 {
						minAlbum = albumYear
					}
				}
			}
			// Track min and max creation dates.
			if band.CreationDate > maxCreation {
				maxCreation = band.CreationDate
			}
			if band.CreationDate < minCreation || minCreation == 0 {
				minCreation = band.CreationDate
			}
			// Track max band size
			if len(band.Members) > maxBandSize {
				maxBandSize = len(band.Members)
			}
		}

		// Convert the unique country map keys to a slice.
		countries = make([]string, 0, len(uniqueCountriesMap))
		for country := range uniqueCountriesMap {
			countries = append(countries, country)
		}
	}
	data := make([]int, maxBandSize)
	for i := 1; i <= maxBandSize; i++ {
		data[i-1] = i
	}
	return countries, minCreation, maxCreation, minAlbum, maxAlbum, data
}

func FilterCountries(countries []string, bands []*structs.Band) []*structs.Band {
	countrySet := make(map[string]struct{})
	for _, country := range countries {
		countrySet[country] = struct{}{}
	}

	filtered := []*structs.Band{}

	for _, band := range bands {
		for _, r := range band.Relations {
			if _, ok := countrySet[r.Country]; ok {
				filtered = append(filtered, band)
				break
			}
		}
	}

	return filtered
}

func FilterCreation(min, max int, bands []*structs.Band) []*structs.Band {
	filtered := []*structs.Band{}
	for _, band := range bands {
		if band.CreationDate >= min && band.CreationDate <= max {
			filtered = append(filtered, band)
		}
	}

	return filtered
}

func FilterAlbum(min, max int, bands []*structs.Band) []*structs.Band {
	filtered := []*structs.Band{}

	for _, band := range bands {
		s := strings.Split(band.FirstAlbum, "-")
		if len(s) != 3 {
			continue
		}
		n, err := strconv.Atoi(s[2])
		if err == nil && n >= min && n <= max {
			filtered = append(filtered, band)
		}
	}

	return filtered
}

func FilterMembers(countries []int, bands []*structs.Band) []*structs.Band {
	countrySet := make(map[int]struct{})
	for _, country := range countries {
		countrySet[country] = struct{}{}
	}

	filtered := []*structs.Band{}

	for _, band := range bands {
		if _, ok := countrySet[len(band.Members)]; ok {
			filtered = append(filtered, band)
		}

	}

	return filtered
}
