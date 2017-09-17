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
	Alpha2        string `json:"alpha-2"`
	Alpha3        string `json:"alpha-3"`
	CountryCode   string `json:"country-code"`
	Iso31662      string `json:"iso_3166-2"`
	Region        string `json:"region"`
	SubRegion     string `json:"sub-region"`
	RegionCode    string `json:"region-code"`
	SubRegionCode string `json:"sub-region-code"`
}

type DownloadByCountryDetailed struct {
	Details CountryDetail `json:"CountryDetails"`
	Count   int    `json:"Count"`
}

