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
