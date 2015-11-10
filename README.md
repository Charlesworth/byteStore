# byteStore
An abstraction library to enable simple Get and Put operations on top of the fantastic [BoltDB](github.com/boltdb/bolt) key value store. Importing the library auto starts the database, the user can simply use the Get and Put functions and not worry about starting a database. Both Get and Put are thread safe, so use indiscriminately throughout your Goroutines!

    Get(bucket string, key string) []byte

    Put(bucket string, key string, value []byte) error

A Close function is also provided, exiting your Go program without Close is safe as long as there are no currently running Get or Put functions, as this could corrupt the .db file with a unfinished transaction.

    Close() error

Enjoy the library, pull and feature requests are welcome.
