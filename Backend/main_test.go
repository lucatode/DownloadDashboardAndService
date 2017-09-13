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
	session := getMgoSession("localhost")
	cleanTestDb(session, "TestDb", "TestCollection")

	//Execution
	t.Run("GET all downloads in empty collection", func(t *testing.T) {
		res, err := allDownloadsQuery(session, "TestDb", "TestCollection")

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
	session := getMgoSession("localhost")
	cleanTestDb(session, "TestDb", "TestCollection")
	staticDownload := buildAStaticDownload()
	insertDownloadForTest(session, staticDownload, "TestDb", "TestCollection")

	//Execution
	t.Run("GET all downloads in one document collection", func(t *testing.T) {
		res, err := allDownloadsQuery(session, "TestDb", "TestCollection")

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
	session := getMgoSession("localhost")
	cleanTestDb(session, "TestDb", "TestCollection")
	staticDownload := buildAStaticDownload()

	//Execution
	t.Run("POST a donwload in a collection", func(t *testing.T) {
		err := addDownloadInsert(session, staticDownload, "TestDb", "TestCollection")

		res, _ := allDownloadsQuery(session, "TestDb", "TestCollection")

		downloadsExpected := []Download{staticDownload}
		downloadsLenExpected := 1

		//Assertions
		assert.Equal(t, nil, err)
		assert.Equal(t, downloadsLenExpected, len(res))
		assert.Equal(t, downloadsExpected, res)
	})
}

//TODO: empty insert with POST /downloads

//GetDownloadsByCountry
func TestGetDownloadsByCountry(t *testing.T) {
	//Setup
	session := getMgoSession("localhost")
	cleanTestDb(session, "TestDb", "TestCollection")
	staticDownload := buildAStaticDownload()

	//Execution
	t.Run("POST a donwload in a collection", func(t *testing.T) {
		err := addDownloadInsert(session, staticDownload, "TestDb", "TestCollection")

		res, _ := allDownloadsQuery(session, "TestDb", "TestCollection")

		downloadsExpected := []Download{staticDownload}
		downloadsLenExpected := 1

		//Assertions
		assert.Equal(t, nil, err)
		assert.Equal(t, downloadsLenExpected, len(res))
		assert.Equal(t, downloadsExpected, res)
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
	session := getMgoSession("localhost")
	cleanTestDb(session, "TestDb", "TestCollection")
	staticDownload := buildAStaticDownload()
	insertDownloadForTest(session, staticDownload, "TestDb", "TestCollection")

	//Execution
	t.Run("GET /downloads", func(t *testing.T) {
		s := httptest.NewServer(mux(session, "TestDb", "TestCollection"))
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

func TestAddDownload_InsertOne(t *testing.T) {
	//Setup
	session := getMgoSession("localhost")
	cleanTestDb(session, "TestDb", "TestCollection")
	staticDownload := buildAStaticDownload()

	//Execution
	t.Run("POST /downloads", func(t *testing.T) {
		s := httptest.NewServer(mux(session, "TestDb", "TestCollection"))
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
func getMgoSession(connectionString string) *mgo.Session {
	session, err := mgo.Dial(connectionString)
	if err != nil {
		panic(err)
	}
	return session
}

func cleanTestDb(s *mgo.Session, db string, collection string) {
	session := s.Copy()
	defer session.Close()

	c := session.DB(db).C(collection)

	c.DropCollection()

}

func insertDownloadForTest(s *mgo.Session, d Download, db string, collection string) {
	session := s.Copy()
	defer session.Close()

	c := session.DB(db).C(collection)

	err := c.Insert(d)
	if err != nil {
		panic(err)
	}
}

func buildADownload(appId string, lo float64, la float64, time string) Download {
	return Download{
		AppId:        appId,
		Longitude:    lo,
		Latitude:     la,
		DownloadedAt: time,
	}
}

func buildAStaticDownload() Download {
	return buildADownload(
		"IOS_ALERT",
		45.0,
		45.0,
		"12:45.30.000",
	)
}

/**************************************************
  _                 _   _          ___              _
 | |   ___  __ __ _| |_(_)___ _ _ / __| ___ _ ___ _(_)__ ___
 | |__/ _ \/ _/ _` |  _| / _ \ ' \\__ \/ -_) '_\ V / / _/ -_)
 |____\___/\__\__,_|\__|_\___/_||_|___/\___|_|  \_/|_\__\___|

**************************************************/

func TestGetCountry(t *testing.T) {
	//Setup

	//Execution
	t.Run("Get Country", func(t *testing.T) {

	})
}
