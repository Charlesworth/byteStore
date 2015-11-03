package jsonStore

import "testing"

func TestDBWrite(t *testing.T) {
	startDB()

	testWrite := writeRequest{key: "testKey", value: []byte("{'test':1}"), bucket: "testBucket"}
	dbWriteIn <- testWrite
	err := <-dbWriteErr
	if err != nil {
		t.Error("dbWrite failed with valid test input with error:", err)
	}

	testWriteNoBucket := writeRequest{key: "testKeyNoBucket", value: []byte("{'test':1}")}
	dbWriteIn <- testWriteNoBucket
	err = <-dbWriteErr
	if err != nil {
		t.Error("dbWrite failed with valid input with no bucket defined with error:", err)
	}

	testWriteNoByte := writeRequest{key: "testKey"}
	dbWriteIn <- testWriteNoByte
	err = <-dbWriteErr
	if err == nil {
		t.Error("dbWrite failed, when given a writeRequest with no value byte it did not error")
	}

	testWriteNoKey := writeRequest{value: []byte("test!")}
	dbWriteIn <- testWriteNoKey
	err = <-dbWriteErr
	if err == nil {
		t.Error("dbWrite failed, when given a writeRequest with no key string it did not error")
	}

	Close()
	testWrite = writeRequest{key: "testKey", value: []byte("{'test':1}"), bucket: "testBucket"}
	dbWriteIn <- testWrite
	err = <-dbWriteErr
	if err == nil {
		t.Error("dbWrite failed with valid test input but no initialized boltDB in the db variable")
	}

	cleanup()
}
