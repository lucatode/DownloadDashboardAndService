package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func validateDownload(d Download) bool {

	return validateAppId(d.AppId) &&
		validateLatitude(d.Latitude) &&
		validateLongitude(d.Longitude) &&
		validateTime(d.DownloadedAt) &&
		validateCountry(d.Country)

}

func validateAppId(appId string) bool {
	if len(appId) > 0 {
		return (appId == "iOS_App#1") || (appId == "iOS_App#2") || (appId == "Droid_App#1") || (appId == "Droid_App#2")
	}
	fmt.Println("validateAppId failed")
	return false
}

func validateLatitude(la string) bool {
	if len(la) > 0 {
		lat, err := strconv.ParseFloat(la, 64)

		if err != nil {
			fmt.Println("validateLatitude failed - 1")
			return false
		}
		val := (lat <= 90) && (lat >= -90)
		if !val {
			fmt.Println("validateLatitude failed - 2")
		}

		return val

	}
	return false
}
func validateLongitude(lo string) bool {
	if len(lo) > 0 {
		lon, err := strconv.ParseFloat(lo, 64)

		if err != nil {
			fmt.Println("validateLatitude failed - 1")
			return false
		}

		return (lon <= 180) && (lon >= -180)

	}
	return false
}
func validateTime(t string) bool {
	r, _ := regexp.Compile("^([0-9]|0[0-9]|1[0-9]|2[0-4]):[0-5][0-9]$")
	val := r.MatchString(t) //check valid hh:mm
	if !val {
		fmt.Println("validateTime failed, time: '" + t + "' is not valid")
	}

	return val
}
func validateCountry(c string) bool {

	//TODO: check on db
	if len(c) > 0 {
		return true
	}
	fmt.Println("validateCountry failed - 1:'" + c + "'")
	return false
}
