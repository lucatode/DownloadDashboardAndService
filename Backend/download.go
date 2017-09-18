package main

type Download struct {
	AppId        string "json:appId"
	Longitude    string "json:longitude"
	Latitude     string "json:latitude"
	DownloadedAt string "json:downloadedAt"
	Country      string "json:country"
}

type DownloadByCountry struct {
	Country string `json:"Country" bson:"_id,omitempty"`
	Count   int    `json:"Count"`
}

type CountryDetail struct {
	Name          string `json:"name"`
	Alpha2        string `json:"alpha2" bson:"alpha-2,omitempty"`
	Alpha3        string `json:"alpha3" bson:"alpha-3,omitempty"`
	CountryCode   string `json:"country-code"`
	Iso31662      string `json:"iso_3166-2"`
	Region        string `json:"region"`
	SubRegion     string `json:"sub-region"`
	RegionCode    string `json:"region-code"`
	SubRegionCode string `json:"sub-region-code"`
}

type TableDownloadCount struct {
	Total       int                         `json:"total"`
	PerPage     int                         `json:"per_page"`
	CurrentPage int                         `json:"current_page"`
	LastPage    int                         `json:"last_page"`
	NextPageURL string                      `json:"next_page_url"`
	PrevPageURL string                      `json:"prev_page_url"`
	From        int                         `json:"from"`
	To          int                         `json:"to"`
	Data        []DownloadByCountryDetailed `json:"data"`
}

type DownloadByCountryDetailed struct {
	Details CountryDetail `json:"CountryDetails"`
	Count   int           `json:"Count"`
}
