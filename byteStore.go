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
	defer mutex.Unlock()

	var value []byte
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucket))
		if bucket != nil {
			value = bucket.Get([]byte(key))
		}
		return nil
	})

	return value
}

// GetBucket retrieves all values in a bucket, an empty byte will be returned if no values are present.
func GetBucket(bucket string) [][]byte {
	mutex.Lock()
	defer mutex.Unlock()

	var values [][]byte
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucket))
		if bucket == nil {
			return nil
		}

		cursor := bucket.Cursor()
		_, value := cursor.First()
		for {
			if value == nil {
				return nil
			}
			values = append(values, value)
			_, value = cursor.Next()
		}
	})

	return values
}

// Put inserts the key value into the db in the bucket specified.
func Put(bucket string, key string, value []byte) error {
	mutex.Lock()
	defer mutex.Unlock()

	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}
		err = b.Put([]byte(key), value)
		return err
	})

	return err
}

// Delete removes the key/value pair, returns nil if key/value doesn't exist
func Delete(bucket string, key string) error {
	mutex.Lock()
	defer mutex.Unlock()

	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucket))
		if bucket == nil {
			return nil
		}

		return bucket.Delete([]byte(key))
	})

	return err
}

// Close safely closes the database.
func Close() error {
	return db.Close()
}
