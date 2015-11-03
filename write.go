package jsonStore

import (
	"errors"

	"github.com/boltdb/bolt"
)

func init() {
	dbWriteIn = make(chan writeRequest)
	dbWriteErr = make(chan error)
	go dbWrite()
}

type writeRequest struct {
	key    string
	value  []byte
	bucket string
}

var dbWriteIn chan writeRequest
var dbWriteErr chan error

func dbWrite() {
	for {
		write := <-dbWriteIn
		if write.key == "" {
			dbWriteErr <- errors.New("no key string supplied in request")
			continue
		}

		if write.value == nil {
			dbWriteErr <- errors.New("no value []byte supplied in request")
			continue
		}

		if write.bucket == "" {
			write.bucket = "general"
		}

		dbWriteErr <- db.Update(func(tx *bolt.Tx) error {
			var b *bolt.Bucket
			b, _ = tx.CreateBucketIfNotExists([]byte(write.bucket))
			err := b.Put([]byte(write.key), write.value)
			return err
		})
	}
}
