package byteStore

import (
	"log"
	"sync"
	"time"

	"github.com/boltdb/bolt"
)

var db *bolt.DB
var mutex *sync.Mutex

func init() {
	mutex = &sync.Mutex{}
	startDB()
}

func startDB() {
	var err error
	db, err = bolt.Open("byteStore.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
}

// Get retrieves the value using the bucket and key provided, an empty byte will be returned if no value is present.
func Get(bucket string, key string) []byte {
	mutex.Lock()
	var value []byte
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucket))
		value = bucket.Get([]byte(key))
		return nil
	})
	mutex.Unlock()

	return value
}

// Put inserts the key value into the db in the bucket specified.
func Put(bucket string, key string, value []byte) error {
	mutex.Lock()
	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}
		err = b.Put([]byte(key), value)
		return err
	})
	mutex.Unlock()

	return err
}

// Close safely closes the database.
func Close() error {
	return db.Close()
}
