package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

const bucketName = "rsvps"

var (
	db *bolt.DB

	dbFile   = flag.String("db", "./rsvp.boltdb", "database file")
	add      = flag.Bool("add", false, "add an rsvp to the database")
	list     = flag.Bool("list", false, "list rsvp's in the database")
	name     = flag.String("name", "", "name of person to add to database")
	email    = flag.String("email", "", "email of person to add to database")
	httpServ = flag.String("http", "", "start http server ([host]:port)")
	rootDir  = flag.String("root", "./", "root directory for http server")
)

func parseArguments() {
	flag.Parse()

	if !*add && !*list && *httpServ == "" {
		fmt.Println("Error: one of the following must be given: -add, -list, -http")
		flag.Usage()
		os.Exit(1)
	}

	if *add && (*name == "" || *email == "") {
		fmt.Println("Error: both name and email must be given for adding to the database")
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	var err error

	parseArguments()

	db, err = bolt.Open(*dbFile, 0600, nil)
	if err != nil {
		fmt.Println("Failed to open Database:", err)
		return
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) (err error) {
		_, err = tx.CreateBucketIfNotExists([]byte("rsvps"))
		if err != nil {
			fmt.Println("Failed to create bucket:", err)
		}

		return err
	})

	if *add {
		rsvp := Rsvp{0, *name, *email, *response, time.Now()}
		if err := addRsvp(rsvp); err != nil {
			fmt.Println("Failed to add RSVP:", err)
		}
	} else if *list {
		listRsvp(os.Stdout)
	} else if *httpServ != "" {
		echoServer()
	}
}
