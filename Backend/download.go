package main

type Download struct {
	AppId        string  "json:appId"
	Longitude    float64 "json:longitude"
	Latitude     float64 "json:latitude"
	DownloadedAt string  "json:downloadedAt"
	Country      string  "json:country"
}

type DownloadByCountry struct {
	Country string "json:country"
	Count   int    "json:count"
}
