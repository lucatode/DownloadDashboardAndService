package main

type Download struct {
	AppId        string  "json:appId"
	Longitude    float64 "json:longitude"
	Latitude     float64 "json:latitude"
	DownloadedAt string  "json:downloadedAt"
	Country      string  "json:country"
}

type DownloadByCountry struct {
	Country string "json:_id"
	Count   int    "json:count"
}
