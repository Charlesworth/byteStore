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

func Get(bucketName string, key string) []byte {
	mutex.Lock()
	var value []byte
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		value = bucket.Get([]byte(key))
		return nil
	})
	mutex.Unlock()

	return value
}

func Put(bucketName string, key string, value []byte) error {
	mutex.Lock()
	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		err = b.Put([]byte(key), value)
		return err
	})
	mutex.Unlock()

	return err
}

func Close() error {
	return db.Close()
}
