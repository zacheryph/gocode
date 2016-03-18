package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"text/tabwriter"
	"time"

	"github.com/boltdb/bolt"
)

// TimeFormat formats time how we want it
const TimeFormat = "Mon Jan _2 15:04 MST"

// Rsvp defines the reservation to encode for the database
type Rsvp struct {
	ID       uint64
	Name     string
	Email    string
	Response string
	Created  time.Time
}

func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func addRsvp(rsvp Rsvp) error {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			fmt.Println("Failed to find bucket")
			return errors.New("Bucket not found")
		}

		nextID, err := bucket.NextSequence()
		if err != nil {
			fmt.Println("Failed to retrieve next sequence", err)
			return err
		}

		rsvp.ID = nextID
		data, err := json.Marshal(rsvp)
		if err != nil {
			fmt.Println("Failed to marshal rsvp:", err)
			return err
		}

		err = bucket.Put(itob(nextID), data)

		return err
	})

	if err != nil {
		fmt.Println("Failed to add RSVP:", err)
		return err
	}

	fmt.Println("Successfully added RSVP")
	return nil
}

func listRsvp(out io.Writer) {
	wr := tabwriter.NewWriter(out, 0, 4, 4, ' ', 0)
	wr.Write([]byte("Name\tEmail\tRSVPed At\tResponse\n"))

	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			fmt.Println("Failed to find bucket")
			return errors.New("Bucket not found")
		}

		bucket.ForEach(func(k, v []byte) error {
			var rsvp Rsvp
			err := json.Unmarshal(v, &rsvp)
			if err != nil {
				fmt.Println("Failed to unmarshal RSVP", err)
				return err
			}

			tz, _ := time.LoadLocation("America/Chicago")
			ts := rsvp.Created.In(tz)
			line := fmt.Sprintf("%s\t%s\t%s\t%s\n", rsvp.Name, rsvp.Email,
				ts.Format(TimeFormat), rsvp.Response)
			wr.Write([]byte(line))

			return nil
		})

		return nil
	})

	wr.Flush()
}
