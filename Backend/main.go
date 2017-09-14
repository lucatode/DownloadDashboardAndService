package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	goji "goji.io"
	"goji.io/pat"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	connectionString := "localhost"
	session, err := mgo.Dial(connectionString)

	if err != nil {
		panic(err)
	}
	defer session.Close()

	http.ListenAndServe("localhost:8000", mux(session, "AppDashboard", "Downloads"))
}

///Mux function
func mux(s *mgo.Session, db string, collection string) *goji.Mux {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/downloads"), allDownloads(s, db, collection))
	mux.HandleFunc(pat.Post("/downloads"), addDownload(s, db, collection))
	//mux.HandleFunc(pat.Get("/countDownloadsByCountry"), countDownloadsByCountry(s, db, collection))
	//mux.HandleFunc(pat.Get("/downloadsByCountry/:country"), downloadsByCountry(s, db, collection))
	//mux.HandleFunc(pat.Get("/countDownloadsByTime"), countDownloadsByTime(s, db, collection))
	//mux.HandleFunc(pat.Get("/downloadsByTime/:time"), downloadsByTime(s, db, collection))

	return mux
}

///Route handling functions
//This function handle '/downloads' route GET calls - Retriving all downloads in the db
func allDownloads(s *mgo.Session, db string, collection string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//Query the db
		downloads, err := allDownloadsQuery(s, db, collection)

		//Evaluate errors
		if err != nil {
			//Handle response with json - Input: writer, message, code
			ErrorWithJSON(w, "Replace with message", http.StatusInternalServerError)
			//Log the error
			log.Println("Failed to get all books", err)
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

func addDownload(s *mgo.Session, db string, collection string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// - Get download from arguments - Parse the request
		//return val
		var download Download
		//create a decoder for request body
		decoder := json.NewDecoder(r.Body)
		//use the decoder
		err := decoder.Decode(&download)
		//...
		download.Country = getCountry(download.Latitude, download.Longitude)
		//...
		if err != nil {
			ErrorWithJSON(w, "Unable to decode request body", http.StatusBadRequest)
		}

		err = addDownloadInsert(s, download, db, collection)

		if err != nil {
			ErrorWithJSON(w, "Unable to insert a new Download", http.StatusBadRequest)
		}

		//Return response after insert status
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}
}

///DB handling functions
func allDownloadsQuery(s *mgo.Session, db string, collection string) ([]Download, error) {
	//Get a new session
	session := s.Copy()
	defer session.Close()

	c := session.DB(db).C(collection)

	var downloads []Download
	err := c.Find(bson.M{}).All(&downloads)

	if err != nil {
		fmt.Println("Some errors during find all")
		return downloads, fmt.Errorf("Some errors during find all")
	}

	if len(downloads) == 0 {
		fmt.Println("No documents in the db")
		return downloads, fmt.Errorf("No documents in the db")
	}

	return downloads, nil
}

func addDownloadInsert(s *mgo.Session, d Download, db string, collection string) error {
	//Get a new session
	session := s.Copy()
	defer session.Close()

	//Get Collection
	c := session.DB(db).C(collection)

	//Insert a book
	err := c.Insert(d)

	return err
}

func getDownloadsByCountry(s *mgo.Session, string country, db string, collection string) ([]DownloadByCountry, error) {
	//Get a new session
	session := s.Copy()
	defer session.Close()

	c := session.DB(db).C(collection)

	var downloads []DownloadByCountry

	pipe := c.Pipe(bson.M{"$group": bson.M{"_id": "$country", "country": bson.M{"$push": "$$ROOT", "count": bson.M{"$sum": "1"}}}})

	iter := pipe.Iter()
	iter.All(&downloads)

	if len(downloads) == 0 {
		fmt.Println("No documents in the db")
		return downloads, fmt.Errorf("No documents in the db")
	}
	return downloads, nil
}

//Returns 2 chars country code
func getCountry(lat float64, lng float64) string {
	return NewLocation(lat, lng)
}

///Utilities
// This function handle the response with json whenever an error occurs
func ErrorWithJSON(w http.ResponseWriter, message string, code int) {

	//Set Writer header with Key and valueString
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	//Put the return code in the header
	w.WriteHeader(code)

	//Write on stream the message with the http.ResponseWriter interface
	fmt.Fprintf(w, "{message: %q}", message)
}

// This function handle the response with json whenever an object is returned correctly from db query
func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {

	//Set Writer header with Key and valueString
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	//Put the return code in the header
	w.WriteHeader(code)

	//Write the message with ResponseWriter
	w.Write(json)
}
