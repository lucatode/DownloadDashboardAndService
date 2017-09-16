package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/***************************************************
  ___  ___ _____       _
 |   \| _ )_   _|__ __| |_ ___
 | |) | _ \ | |/ -_|_-<  _(_-<
 |___/|___/ |_|\___/__/\__/__/

****************************************************/
// Test allDownloadsQuery on an empty collection
func TestAllDownloadsQuery_Empty(t *testing.T) {

	//Setup
	session := NewTest()
	cleanTestDb(session, "TestCollection")

	//Execution
	t.Run("GET all downloads in empty collection", func(t *testing.T) {
		res, err := allDownloadsQuery(session, "TestCollection")

		errorExpected := fmt.Errorf("No documents in the db")
		var downloadsExpected []Download

		//Assertions
		assert.Equal(t, errorExpected, err)
		assert.Equal(t, downloadsExpected, res)
	})
}

// Test allDownloadsQuery on a one document collection
func TestAllDownloadsQuery_OneDoc(t *testing.T) {
	//Setup
	session := NewTest()
	cleanTestDb(session, "TestCollection")
	staticDownload := buildAStaticDownload()
	insertDownloadForTest(session, staticDownload, "TestCollection")

	//Execution
	t.Run("GET all downloads in one document collection", func(t *testing.T) {
		res, err := allDownloadsQuery(session, "TestCollection")

		downloadsExpected := []Download{staticDownload}
		downloadsLenExpected := 1

		//Assertions
		assert.Equal(t, nil, err)
		assert.Equal(t, downloadsLenExpected, len(res))
		assert.Equal(t, downloadsExpected, res)
	})
}

//TODO: multiple entry retriving
// ...

//Test addDownload to add one document to the collection
func TestAddDownloadInsert(t *testing.T) {
	//Setup
	session := NewTest()
	cleanTestDb(session, "TestCollection")
	staticDownload := buildAStaticDownload()

	//Execution
	//t.Run("POST a donwload in a collection", func(t *testing.T) {
	err := addDownloadInsert(session, staticDownload, "TestCollection")
	if err != nil {
		fmt.Println("Unable to insert a download")
		panic(err)
	}

	res, err2 := allDownloadsQuery(session, "TestCollection")
	if err2 != nil {
		fmt.Println("Unable to query download")
		panic(err2)
	}

	downloadsExpected := []Download{staticDownload}
	downloadsLenExpected := 1

	//Assertions
	assert.Equal(t, nil, err)
	assert.Equal(t, downloadsLenExpected, len(res))
	assert.Equal(t, downloadsExpected, res)
	//})
}

//TODO: empty insert with POST /downloads

//GetDownloadsByCountry - ok
func TestGetDownloadsByCountry_ok(t *testing.T) {
	//Setup
	session := NewTest()
	cleanTestDb(session, "TestCollection")
	staticDownload := buildAStaticDownload()

	//Execution
	t.Run("Get downloads county by country - working", func(t *testing.T) {
		err := addDownloadInsert(session, staticDownload, "TestCollection")

		res, _ := countDownloadsByCountry(session, "TestCollection")

		downloadsLenExpected := 1
		expectedCountry := "it"
		expectedCount := 1

		//Assertions
		assert.Equal(t, nil, err)
		assert.Equal(t, downloadsLenExpected, len(res))
		count := res[0]
		assert.Equal(t, expectedCountry, count.Country)
		assert.Equal(t, expectedCount, count.Count)

	})
}

//GetDownloadsByCountry - ko
func TestGetDownloadsByCountry_ko(t *testing.T) {
	//Setup
	session := NewTest()
	cleanTestDb(session, "TestCollection")
	//staticDownload := buildAStaticDownload()
	downloadsLenExpected := 0

	//Execution
	t.Run("Get downloads county by country - no data", func(t *testing.T) {
		//addDownloadInsert(session, staticDownload, "TestCollection")

		res, err2 := countDownloadsByCountry(session, "TestCollection")

		//Assertions
		assert.Equal(t, downloadsLenExpected, len(res))
		assert.NotEqual(t, nil, err2)
	})
}

/**************************************************
  ___          _          _____       _
 | _ \___ _  _| |_ ___ __|_   _|__ __| |_ ___
 |   / _ \ || |  _/ -_|_-< | |/ -_|_-<  _(_-<
 |_|_\___/\_,_|\__\___/__/ |_|\___/__/\__/__/

**************************************************/
func TestAllDownloads_OneDownload(t *testing.T) {
	//Setup
	session := NewTest()
	cleanTestDb(session, "TestCollection")
	staticDownload := buildAStaticDownload()
	insertDownloadForTest(session, staticDownload, "TestCollection")

	//Execution
	t.Run("GET /downloads", func(t *testing.T) {
		s := httptest.NewServer(mux(session, "TestCollection"))
		defer s.Close()
		res, err := http.Get(s.URL + "/downloads")
		assert.NoError(t, err)
		assert.Equal(t, 200, res.StatusCode)

		//use the decoder
		var download []Download
		decoder := json.NewDecoder(res.Body)
		err = decoder.Decode(&download)
		if err != nil {
			panic(err)
		}

		assert.Equal(t, 1, len(download))

		defer res.Body.Close()
	})

}

func TestCountDownloadsByCountry_OneDownload(t *testing.T) {
	//Setup
	session := NewTest()
	cleanTestDb(session, "TestCollection")
	staticDownload := buildAStaticDownload()
	insertDownloadForTest(session, staticDownload, "TestCollection")

	//Execution
	t.Run("GET /countDownloadByCountry", func(t *testing.T) {
		s := httptest.NewServer(mux(session, "TestCollection"))
		defer s.Close()
		res, err := http.Get(s.URL + "/countDownloadsByCountry")
		assert.NoError(t, err)
		assert.Equal(t, 200, res.StatusCode)

		//use the decoder
		var download []DownloadByCountry
		decoder := json.NewDecoder(res.Body)
		err = decoder.Decode(&download)
		if err != nil {
			panic(err)
		}

		assert.Equal(t, 1, len(download))
		assert.Equal(t, 1, download[0].Count)
		assert.Equal(t, "it", download[0].Country)

		defer res.Body.Close()
	})

}

func TestAddDownload_InsertOne(t *testing.T) {
	//Setup
	session := NewTest()
	cleanTestDb(session, "TestCollection")
	staticDownload := buildAStaticDownload()

	//Execution
	t.Run("POST /downloads", func(t *testing.T) {
		s := httptest.NewServer(mux(session, "TestCollection"))
		defer s.Close()
		jsonValue, _ := json.Marshal(staticDownload)
		res, err := http.Post(s.URL+"/downloads", "application/json", bytes.NewBuffer(jsonValue))

		assert.NoError(t, err)
		assert.Equal(t, 201, res.StatusCode)

		defer res.Body.Close()

	})
}

/**************************************************
  _   _ _   _ _ _ _   _
 | | | | |_(_) (_) |_(_)___ ___
 | |_| |  _| | | |  _| / -_|_-<
  \___/ \__|_|_|_|\__|_\___/__/

**************************************************/
// func getMgoSession(connectionString string) *mgo.Database {
// 	dialInfo := &mgo.DialInfo{
// 		Addrs:    []string{MongoDBHosts},
// 		Timeout:  30 * time.Second,
// 		Database: AuthDatabase,
// 		Username: AuthUserName,
// 		Password: AuthPassword,
// 	}

// 	session, err := mgo.DialWithInfo(dialInfo)
// 	if err != nil {
// 		panic(err)
// 	}

// 	session.SetMode(mgo.Monotonic, true)

// 	return session
// }

func cleanTestDb(d *mgo.Database, collection string) {

	c := d.C(collection) //sess.DB("mydb").C("mycollection")

	c.RemoveAll(bson.M{})

}

func insertDownloadForTest(db *mgo.Database, dl Download, collection string) {
	// session := s.Copy()
	// defer session.Close()

	c := db.C(collection)
	err := c.Insert(dl)
	if err != nil {
		panic(err)
	}
}

func buildADownload(appId string, lo string, la string, time string, country string) Download {
	return Download{
		AppId:        appId,
		Longitude:    lo,
		Latitude:     la,
		DownloadedAt: time,
		Country:      country,
	}
}

func buildAStaticDownload() Download {
	return buildADownload(
		"iOS_App#1",
		"45.0",
		"45.0",
		"12:45",
		"it",
	)
}

/**************************************************
  _                 _   _          ___              _
 | |   ___  __ __ _| |_(_)___ _ _ / __| ___ _ ___ _(_)__ ___
 | |__/ _ \/ _/ _` |  _| / _ \ ' \\__ \/ -_) '_\ V / / _/ -_)
 |____\___/\__\__,_|\__|_\___/_||_|___/\___|_|  \_/|_\__\___|

************************************************ **/
