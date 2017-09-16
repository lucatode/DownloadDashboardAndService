package main

type Download struct {
	AppId        string "json:appId"
	Longitude    string "json:longitude"
	Latitude     string "json:latitude"
	DownloadedAt string "json:downloadedAt"
	Country      string "json:country"
}

type DownloadByCountry struct {
	Country string "json:_id"
	Count   int    "json:count"
}
