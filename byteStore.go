package byteStore

import (
	"sync"
	"time"

	"github.com/boltdb/bolt"
)

type ByteStore struct {
	mutex *sync.Mutex
	db    *bolt.DB
}

func NewByteStore(dbName string) (ByteStore, error) {
	db, err := bolt.Open(dbName+".db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return ByteStore{}, err
	}

	mutex := &sync.Mutex{}
	return ByteStore{mutex, db}, nil
}

// Get retrieves the value using the bucket and key provided, an empty byte
// will be returned if no value is present.
func (bs ByteStore) Get(bucket string, key string) []byte {
	bs.mutex.Lock()
	defer bs.mutex.Unlock()

	var value []byte
	bs.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucket))
		if bucket != nil {
			value = bucket.Get([]byte(key))
		}
		return nil
	})

	return value
}

// KeyValue is a container for a key value pair, most usualy used for
// holding a slice of key value pairs
type KeyValue struct {
	key   string
	value []byte
}

// GetBucket retrieves all keys and values in a bucket, an empty KeyValue slice
// will be returned if no values are present.
func (bs ByteStore) GetBucket(bucket string) []KeyValue {
	bs.mutex.Lock()
	defer bs.mutex.Unlock()

	var keyValues []KeyValue
	bs.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucket))
		if bucket == nil {
			return nil
		}

		cursor := bucket.Cursor()
		key, value := cursor.First()
		for {
			if value == nil {
				return nil
			}
			keyValues = append(keyValues, KeyValue{string(key), value})
			key, value = cursor.Next()
		}
	})

	return keyValues
}

// GetBucketValues retrieves all values in a bucket, an empty slice of bytes
// will be returned if no values are present.
func (bs ByteStore) GetBucketValues(bucket string) [][]byte {
	bs.mutex.Lock()
	defer bs.mutex.Unlock()

	var values [][]byte
	bs.db.View(func(tx *bolt.Tx) error {
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
func (bs ByteStore) Put(bucket string, key string, value []byte) error {
	bs.mutex.Lock()
	defer bs.mutex.Unlock()

	err := bs.db.Update(func(tx *bolt.Tx) error {
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
func (bs ByteStore) Delete(bucket string, key string) error {
	bs.mutex.Lock()
	defer bs.mutex.Unlock()

	err := bs.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucket))
		if bucket == nil {
			return nil
		}

		return bucket.Delete([]byte(key))
	})

	return err
}

// DeleteBucket deletes a whole bucket
func (bs ByteStore) DeleteBucket(bucket string) error {
	bs.mutex.Lock()
	defer bs.mutex.Unlock()

	err := bs.db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket([]byte(bucket))
	})

	return err
}

// Close safely closes the database.
func (bs ByteStore) Close() error {
	return bs.db.Close()
}
