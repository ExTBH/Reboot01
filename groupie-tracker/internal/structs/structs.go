package structs

type Geolocation struct {
	// Latitude
	Lat float64
	// Longtitude
	Lon float64
}

type Relation struct {
	Date    []string
	Loc     *Geolocation
	City    string
	Country string
}

type Band struct {
	Id           int         `json:"id"`
	Image        string      `json:"image"`
	Name         string      `json:"name"`
	Members      []string    `json:"members"`
	CreationDate int         `json:"creationDate"`
	FirstAlbum   string      `json:"firstAlbum"`
	Relations    []*Relation `json:"-"`
}
