package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type country string

type Localization struct {
	PlaceID     string `json:"place_id"`
	Licence     string `json:"licence"`
	OsmType     string `json:"osm_type"`
	OsmID       string `json:"osm_id"`
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
	DisplayName string `json:"display_name"`
	Address     struct {
		HouseNumber   string `json:"house_number"`
		Road          string `json:"road"`
		Suburb        string `json:"suburb"`
		City          string `json:"city"`
		County        string `json:"county"`
		StateDistrict string `json:"state_district"`
		State         string `json:"state"`
		Postcode      string `json:"postcode"`
		Country       string `json:"country"`
		CountryCode   string `json:"country_code"`
	} `json:"address"`
	Boundingbox []string `json:"boundingbox"`
}

//Reverse geocode to find Country
func NewLocation(lat string, lon string) string {

	url := "http://nominatim.openstreetmap.org/reverse?format=json&lat=" + lat + "&lon=" + lon + "&zoom=18&addressdetails=1"
	res, err := http.Get(url)

	if err != nil {
		panic(err)

	}
	defer res.Body.Close()

	bodyBytes, err2 := ioutil.ReadAll(res.Body)

	if err2 != nil {
		panic(err2)
	}

	var m Localization
	json.Unmarshal(bodyBytes, &m)

	countryName := m.Address.CountryCode

	return countryName
}
