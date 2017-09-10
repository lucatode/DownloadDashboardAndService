package main

import mgo "gopkg.in/mgo.v2"

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
