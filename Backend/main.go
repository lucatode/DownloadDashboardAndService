package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/rs/cors"
	goji "goji.io"
	"goji.io/pat"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {

	database := New()

	port := os.Getenv("PORT")

	if port == "" {
		port = "6060"
	}

	http.ListenAndServe(":"+port, mux(database, "TestCollection"))
}

///Mux function
func mux(db *mgo.Database, collection string) *goji.Mux {
	mux := goji.NewMux()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"GET", "POST"},
	})
	mux.Use(c.Handler)

	mux.HandleFunc(pat.Get("/downloads"), allDownloads(db, collection))
	mux.HandleFunc(pat.Post("/downloads"), addDownload(db, collection))
	mux.HandleFunc(pat.Get("/countDownloadsByCountry"), getCountDlByCountry(db, collection))
	mux.HandleFunc(pat.Get("/downloadsByCountryDetail"), getCountDlByCountryDetail(db, collection))
	mux.HandleFunc(pat.Get("/vueTableData"), getVueTableData(db, collection))
	//mux.HandleFunc(pat.Get("/downloadsByTime/:time"), downloadsByTime(s, collection))

	return mux
}

///Route handling functions
//This function handle '/downloads' route GET calls - Retriving all downloads in the db
func allDownloads(db *mgo.Database, collection string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//Query the db
		downloads, err := allDownloadsQuery(db, collection)

		//Evaluate errors
		if err != nil {
			//Handle response with json - Input: writer, message, code
			ErrorWithJSON(w, "Replace with message", http.StatusInternalServerError)
			//Log the error
			log.Println("Failed to get all downloads", err)
			//Exit
			return
		}

		//Input: dbQueryResponse, prefix, indent
		respBody, err := json.MarshalIndent(downloads, "", " ")

		//Check error on Marshaling
		if err != nil {
			log.Fatal(err)
		}
		//Handle response with json - Input: writer, jsonBody, status
		ResponseWithJSON(w, respBody, http.StatusOK)

	}
}

func addDownload(db *mgo.Database, collection string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// - Get download from arguments - Parse the request
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}

		//return val
		var download Download
		//create a decoder for request body
		decoder := json.NewDecoder(r.Body)
		fmt.Println(r.Body)
		//use the decoder
		err := decoder.Decode(&download)
		//...

		//...
		if err != nil {
			ErrorWithJSON(w, "Unable to decode request body", http.StatusBadRequest)
		}

		download.Country = getCountry(download.Latitude, download.Longitude)

		valid := validateDownload(download)

		if valid != true {
			ErrorWithJSON(w, "Data passed not valid", http.StatusBadRequest)
			return
		}

		err = addDownloadInsert(db, download, collection)

		if err != nil {
			ErrorWithJSON(w, "Unable to insert a new Download", http.StatusBadRequest)
		}

		//Return response after insert status
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}
}

///DB handling functions
func allDownloadsQuery(db *mgo.Database, collection string) ([]Download, error) {

	c := db.C(collection)

	var downloads []Download
	err := c.Find(bson.M{}).All(&downloads)

	if err != nil {
		fmt.Println("Some errors during find all")
		return downloads, fmt.Errorf("Some errors during find all")
	}

	if len(downloads) == 0 {
		fmt.Println("No documents in the db")
		return downloads, nil
	}

	return downloads, nil
}

func addDownloadInsert(db *mgo.Database, d Download, collection string) error {

	//Get Collection
	c := db.C(collection)

	//Insert a download
	err := c.Insert(d)

	return err
}

func getCountDlByCountry(db *mgo.Database, collection string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//Query the db
		downloads, err := countDownloadsByCountry(db, collection)

		//Evaluate errors
		if err != nil {
			//Handle response with json - Input: writer, message, code
			ErrorWithJSON(w, "Replace with message", http.StatusInternalServerError)
			//Log the error
			log.Println("Failed to get all downloads", err)
			//Exit
			return
		}

		//Input: dbQueryResponse, prefix, indent
		respBody, err := json.MarshalIndent(downloads, "", " ")

		//Check error on Marshaling
		if err != nil {
			log.Fatal(err)
		}
		//Handle response with json - Input: writer, jsonBody, status
		ResponseWithJSON(w, respBody, http.StatusOK)

	}
}

func getCountDlByCountryDetail(db *mgo.Database, collection string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//Query the db
		downloads, err := countDownloadsByCountry(db, collection)

		//Evaluate errors
		if err != nil {
			//Handle response with json - Input: writer, message, code
			ErrorWithJSON(w, "Replace with message", http.StatusInternalServerError)
			//Log the error
			log.Println("Failed to get all downloads", err)
			//Exit
			return
		}

		//adding additional data on country
		var downloadByCountryDetailed []DownloadByCountryDetailed

		for _, download := range downloads {
			code2 := download.Country
			countryDetails := getDetailsByAlpha2(db, code2, "CountryDictionary")

			var dlCountDetail DownloadByCountryDetailed
			dlCountDetail.Count = download.Count
			dlCountDetail.Details = countryDetails

			downloadByCountryDetailed = append(downloadByCountryDetailed, dlCountDetail)
		}

		//Input: dbQueryResponse, prefix, indent
		respBody, err := json.MarshalIndent(downloadByCountryDetailed, "", " ")

		//Check error on Marshaling
		if err != nil {
			log.Fatal(err)
		}
		//Handle response with json - Input: writer, jsonBody, status
		ResponseWithJSON(w, respBody, http.StatusOK)

	}
}

func getDetailsByAlpha2(db *mgo.Database, alpha2 string, collection string) CountryDetail {
	c := db.C(collection)

	var details []CountryDetail
	upperA2 := strings.ToUpper(alpha2)
	err := c.Find(bson.M{"alpha-2": upperA2}).All(&details)
	if err != nil {
		panic(err)
	}

	if len(details) > 0 {
		return details[0]
	} else {
		return CountryDetail{}
	}

}

func getVueTableData(db *mgo.Database, collection string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//Query the db
		downloads, err := countDownloadsByCountry(db, collection)

		//Evaluate errors
		if err != nil {
			//Handle response with json - Input: writer, message, code
			ErrorWithJSON(w, "Replace with message", http.StatusInternalServerError)
			//Log the error
			log.Println("Failed to get all downloads", err)
			//Exit
			return
		}

		//adding additional data on country
		tableCount := TableDownloadCount{
			Total:       5,
			PerPage:     5,
			CurrentPage: 1,
			LastPage:    14,
			From:        1,
			To:          5,
		}

		var downloadByCountryDetailed []DownloadByCountryDetailed

		for _, download := range downloads {
			code2 := download.Country
			countryDetails := getDetailsByAlpha2(db, code2, "CountryDictionary")

			var dlCountDetail DownloadByCountryDetailed
			dlCountDetail.Count = download.Count
			dlCountDetail.Details = countryDetails

			downloadByCountryDetailed = append(downloadByCountryDetailed, dlCountDetail)
		}

		tableCount.Data = downloadByCountryDetailed

		//Input: dbQueryResponse, prefix, indent
		respBody, err := json.MarshalIndent(tableCount, "", " ")

		//Check error on Marshaling
		if err != nil {
			log.Fatal(err)
		}
		//Handle response with json - Input: writer, jsonBody, status
		ResponseWithJSON(w, respBody, http.StatusOK)

	}
}

func countDownloadsByCountry(db *mgo.Database, collection string) ([]DownloadByCountry, error) {

	c := db.C(collection)

	var downloads []DownloadByCountry

	pipeline := []bson.M{bson.M{"$group": bson.M{"_id": "$country", "count": bson.M{"$sum": 1}}}}

	pipe := c.Pipe(pipeline)

	iter := pipe.Iter()
	iter.All(&downloads)

	if len(downloads) == 0 {
		fmt.Println("No documents in the db")
	}
	return downloads, nil
}

//Returns 2 chars country code
func getCountry(lat string, lng string) string {
	return NewLocation(lat, lng)
}

///Utilities
// This function handle the response with json whenever an error occurs
func ErrorWithJSON(w http.ResponseWriter, message string, code int) {

	//Set Writer header with Key and valueString
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//Put the return code in the header
	w.WriteHeader(code)

	//Write on stream the message with the http.ResponseWriter interface
	fmt.Fprintf(w, "{message: %q}", message)
}

// This function handle the response with json whenever an object is returned correctly from db query
func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {

	//Set Writer header with Key and valueString
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")

	//

	//Put the return code in the header
	w.WriteHeader(code)

	//Write the message with ResponseWriter
	w.Write(json)
}
