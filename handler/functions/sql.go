package functions

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

	
var db *sql.DB

func init() {
	var (
		conn string = `server=localhost;user id=tutividela;password=dRG4n3sH;port=65104;database=UrlShorten`
		err error
	)

	db, err = sql.Open("mssql", conn)
	if err != nil {
		log.Fatal("Open: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("PING: ", err)
	} else {
		log.Println("PING to database successful")
	}
}

func GetPathUrlFromDB() map[string]string {
	
	var (
		path, url   string
		pathsToUrls = make(map[string]string)
		query       string = `SELECT * FROM dbo.PathToUrl`
	)
	defer db.Close()
	
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Query: ",err)
	}
	for rows.Next() {
		if err := rows.Scan(&path, &url); err != nil {
			log.Fatal("Scan: ",err)
		}
		pathsToUrls[path] = url 
	}
	return pathsToUrls
}