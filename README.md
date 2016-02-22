# byteStore
An abstraction library to enable simple Get, Put and other operations on top of the fantastic [BoltDB](github.com/boltdb/bolt) key value store. Import the library and get a new ByteStore with the byteStore.New() function to get started.

    byteStore.New(dbFileName string) (ByteStore, error)

The user can simply use the ByteStore object's methods and not worry about database complexities. All methods are thread safe, backed by BoltDB's transactional nature, so use indiscriminately throughout your Goroutines!

    (ByteStore) Get(bucket string, key string) []byte

    (ByteStore) GetBucket(bucket string) []KeyValue

    (ByteStore) GetBucketValues(bucket string) [][]byte

    (ByteStore) Put(bucket string, key string, value []byte) error

    (ByteStore) Delete(bucket string, key string) error

    (ByteStore) DeleteBucket(bucket string) error

    (ByteStore) Close() error

The ByteStore type also exposes the underlying BoltDB driver publicly (ByteStore.BoltDB), meaning that if you want to do any more complex transactions not offered by the ByteStore methods, you can use boltDB directly. I hope you enjoy the library, pull and feature requests are welcome.
