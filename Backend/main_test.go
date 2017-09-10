package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	mgo "gopkg.in/mgo.v2"
)

///Routes Unit Test
func TestAllBooks_OneBook(t *testing.T) {
	//Setup
	session := getMgoSession("localhost")
	cleanTestDb(session, "TestDb", "TestCollection")
	randomDownload := buildAStaticDownload()
	insertDownloadForTest(session, randomDownload, "TestDb", "TestCollection")

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

///Test Utilities
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
