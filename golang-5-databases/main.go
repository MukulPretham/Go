package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

var db *sql.DB

func main() {
	config := mysql.NewConfig()
	config.User = "root"
	config.Net = "tcp"
	config.Addr = "localhost:3306"
	config.Passwd = "9059015626"
	config.DBName = "DB"
	var err error
	db, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatalf("Problem connecting to database : %v", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatalf("Problem with database : %v", pingErr)
	}
	fmt.Println("Connected!")

	albums, err := getAlbums("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

}

func getAlbums(artistName string) ([]Album, error) {
	var albums []Album
	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", artistName)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", artistName, err)
	}
	defer rows.Close()

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", artistName, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", artistName, err)
	}
	return albums, nil
}
