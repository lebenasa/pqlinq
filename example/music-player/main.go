package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lebenasa/pqlinq/example/music-player/musics"
	_ "github.com/lib/pq"
)

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		return
	}
	connectionString := flag.Arg(0)

	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Opening database: %v", err)
	}

	music, err := musics.Bind(db)
	if err != nil {
		log.Fatalf("Binding database: %v", err)
	}

	// Insert some data for testing, ignore errors (probably duplicate value)
	data := []musics.Data{
		musics.Data{
			Artist:      "ClariS",
			Title:       "Hitorigoto",
			Album:       "Hitorigoto",
			ReleaseDate: mustParse("Apr 26 2017"),
			LastPlayed:  mustParse("Nov 19 2017"),
			Rating:      8.0,
		},
		musics.Data{
			Artist:      "ClariS",
			Title:       "Shiori",
			Album:       "Shiori",
			ReleaseDate: mustParse("Jul 27 2017"),
			LastPlayed:  mustParse("Nov 19 2017"),
			Rating:      9.0,
		},
		musics.Data{
			Artist:      "ClariS",
			Title:       "Border",
			Album:       "Fairy Castle",
			ReleaseDate: mustParse("Jan 25 2017"),
			LastPlayed:  mustParse("Nov 19 2017"),
			Rating:      8.5,
		},
		musics.Data{
			Artist:      "ClariS",
			Title:       "カラフル",
			Album:       "ClariS -Single Best 1st-",
			ReleaseDate: mustParse("Oct 30 2013"),
			LastPlayed:  mustParse("Nov 19 2017"),
			Rating:      9.5,
		},
		musics.Data{
			Artist:      "ClariS",
			Title:       "Connect",
			Album:       "Birthday",
			ReleaseDate: mustParse("Feb 2 2011"),
			LastPlayed:  mustParse("Nov 19 2017"),
			Rating:      10.0,
		},
	}

	tx, err := db.Beginx()
	if err != nil {
		log.Fatalf("Beginning transaction: %v", err)
	}
	defer tx.Rollback()

	for _, d := range data {
		music.Insert(tx, d)
	}
	tx.Commit()

	// Get songs released on 2017
	query := music.Select().Where(
		musics.ReleaseDate.Gt(mustParse("Jan 1 2017")).And(
			musics.ReleaseDate.Lt(mustParse("Jan 1 2018")))).
		OrderBy(musics.OrderField{
			Field: musics.ReleaseDate,
			Order: musics.Ascending}).
		Limit(100, 0)
	//fmt.Println(query.(musics.Query).BuildSQL())
	songsOn2017, err := query.Execute()
	if err != nil {
		log.Fatalf("Selecting 2017 songs: %v", err)
	}
	fmt.Println("Songs released on 2017 are:")
	for _, song := range songsOn2017 {
		d := song.Data()
		fmt.Printf("	- %v, %v, %v, %v -- Rating: %v\n", d.Artist, d.Title, d.Album, d.ReleaseDate.Year(), d.Rating)
	}
}

func usage() {
	fmt.Println("Sample application that access musics table using pqlinq generated code.")
	fmt.Println("It doesn't actually play anything though")
	fmt.Println("Usage: music-player [database connection string]")
	fmt.Println("Example: music-player postgres://user:pass@host/database")
	flag.PrintDefaults()
}

func mustParse(when string) (t time.Time) {
	t, err := time.Parse("Jan 2 2006", when)
	if err != nil {
		log.Fatalf("Parsing %v: %v", when, err)
	}
	return t
}
