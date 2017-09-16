package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Utilities
func NewDownload(a string, lo string, la string, t string, c string) Download {
	return Download{
		AppId:        a,
		Longitude:    lo,
		Latitude:     la,
		DownloadedAt: t,
		Country:      c,
	}
}

//Tests
func TestValidateDownload_ko(t *testing.T) {
	//Setup
	//Create a download
	d := NewDownload("", "", "", "", "") //not valid

	//Execution
	t.Run("Validate the download with empty fields", func(t *testing.T) {
		valid := validateDownload(d)

		//Assertions
		assert.Equal(t, false, valid)
	})
}

func TestValidateDownload_ok(t *testing.T) {
	//Setup
	//Create a download
	d := NewDownload("iOS_App#1", "45.0000", "9.0000", "24:00", "it") //valid

	//Execution
	t.Run("Validate the download with valid fields", func(t *testing.T) {
		valid := validateDownload(d)

		//Assertions
		assert.Equal(t, true, valid)
	})
}

//Single Tests - ValidateAppId
func TestValidateAppId_ko(t *testing.T) {
	//Setup
	//Create a download
	appId := "" //not valid

	//Execution
	t.Run("Validate appId field - invalid", func(t *testing.T) {
		valid := validateAppId(appId)

		//Assertions
		assert.Equal(t, false, valid)
	})
}

func TestValidateAppId_ok(t *testing.T) {
	//Setup
	//Create a download
	appId := "Droid_App#1" //not valid

	//Execution
	t.Run("Validate appId field - valid", func(t *testing.T) {
		valid := validateAppId(appId)

		//Assertions
		assert.Equal(t, true, valid)
	})
}

//Single Tests - ValidateLatitude
func TestValidateLatitude_ko_empty_string(t *testing.T) {
	//Setup
	//Create a download
	la := "" //not valid

	//Execution
	t.Run("Validate Latitude field - invalid", func(t *testing.T) {
		valid := validateLatitude(la)

		//Assertions
		assert.Equal(t, false, valid)
	})
}
func TestValidateLatitude_ko_not_float64(t *testing.T) {
	//Setup
	//Create a download
	la := "sa215s" //not valid

	//Execution
	t.Run("Validate Latitude field - invalid", func(t *testing.T) {
		valid := validateLatitude(la)

		//Assertions
		assert.Equal(t, false, valid)
	})
}

func TestValidateLatitude_ko_out_of_range(t *testing.T) {
	//Setup
	//Create a download
	la := "2000" //not valid

	//Execution
	t.Run("Validate Latitude field - invalid", func(t *testing.T) {
		valid := validateLatitude(la)

		//Assertions
		assert.Equal(t, false, valid)
	})
}

func TestValidateLatitude_ok(t *testing.T) {
	//Setup
	//Create a download
	la := "50.000" //not valid

	//Execution
	t.Run("Validate Latitude field - valid", func(t *testing.T) {
		valid := validateLatitude(la)

		//Assertions
		assert.Equal(t, true, valid)
	})
}

func TestValidateLongitude_ko_empty_string(t *testing.T) {
	//Setup
	//Create a download
	lo := "" //not valid

	//Execution
	t.Run("Validate Longitude field - invalid", func(t *testing.T) {
		valid := validateLongitude(lo)

		//Assertions
		assert.Equal(t, false, valid)
	})
}
func TestValidateLongitude_ko_not_float64(t *testing.T) {
	//Setup
	//Create a download
	lo := "sa215s" //not valid

	//Execution
	t.Run("Validate Longitude field - invalid", func(t *testing.T) {
		valid := validateLongitude(lo)

		//Assertions
		assert.Equal(t, false, valid)
	})
}

func TestValidateLongitude_ko_out_of_range(t *testing.T) {
	//Setup
	//Create a download
	lo := "181" //not valid

	//Execution
	t.Run("Validate Longitude field - invalid", func(t *testing.T) {
		valid := validateLongitude(lo)

		//Assertions
		assert.Equal(t, false, valid)
	})
}

func TestValidateLongitude_ok(t *testing.T) {
	//Setup
	//Create a download
	lo := "50.000" //not valid

	//Execution
	t.Run("Validate Longitude field - valid", func(t *testing.T) {
		valid := validateLongitude(lo)

		//Assertions
		assert.Equal(t, true, valid)
	})
}

func TestValidateTime_ko_invalid_string(t *testing.T) {
	//Setup
	//Create a download
	time := "2515aaa4" //not valid

	//Execution
	t.Run("Validate Time field - invalid", func(t *testing.T) {
		valid := validateTime(time)

		//Assertions
		assert.Equal(t, false, valid)
	})
}

func TestValidateTime_ko_out_of_range(t *testing.T) {
	//Setup
	//Create a download
	time := "25:77" //not valid

	//Execution
	t.Run("Validate Time field - invalid", func(t *testing.T) {
		valid := validateTime(time)

		//Assertions
		assert.Equal(t, false, valid)
	})
}

func TestValidateTime_ok(t *testing.T) {
	//Setup
	//Create a download
	lo := "17:30" //not valid

	//Execution
	t.Run("Validate Time field - valid", func(t *testing.T) {
		valid := validateTime(lo)

		//Assertions
		assert.Equal(t, true, valid)
	})
}
