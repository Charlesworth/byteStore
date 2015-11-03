package jsonStore

import (
	"errors"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

func init() {
	startDB()

	dbReadIn = make(chan readRequest)
	dbReadOut = make(chan []byte)
	dbReadErr = make(chan error)
}

func startDB() {
	var err error
	db, err = bolt.Open("jsonStore.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
}

type readRequest struct {
	key    string
	bucket string
}

var dbReadIn chan readRequest
var dbReadOut chan []byte
var dbReadErr chan error

func dbReader() {

	for {
		read := <-dbReadIn
		log.Println(read)
		if read.bucket == "" {
			read.bucket = "general"
		}

		if read.key == "" {
			dbReadErr <- errors.New("no key string supplied in request")
			continue
		}

		//dbReadOut <- []byte("yay")

		db.View(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte(read.bucket))
			value := bucket.Get([]byte(read.key))
			dbReadOut <- value
			return nil
		})
	}
}

func Close() error {
	return db.Close()
}

func Put(key string, json []byte) error {
	// db.Update(func(tx *bolt.Tx) error {
	// 	b := tx.Bucket([]byte("general"))
	// 	err := b.Put([]byte(key), json)
	// 	return err
	// })
	return nil
}

func PutInBucket(key string, json []byte) error {
	// db.Update(func(tx *bolt.Tx) error {
	// 	b := tx.Bucket([]byte("general"))
	// 	err := b.Put([]byte(key), json)
	// 	return err
	// })
	return nil
}

func Get(key string) ([]byte, error) {
	return []byte{}, nil
}

func GetFromBucket(key string) ([]byte, error) {
	return []byte{}, nil
}

// func GetAllFromBucket() [][]byte {
// 	return [][]byte{}
// }
