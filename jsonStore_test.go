package jsonStore

import "testing"
import "os"
import "log"

func TestInit(t *testing.T) {
	if _, err := os.Stat("jsonStore.db"); err != nil {
		t.Error("balls")
	}
}

func TestDBRead(t *testing.T) {
	go dbReader()

	testRead := readRequest{key: "testKey", bucket: "testBucket"}
	dbReadIn <- testRead
	select {
	case out := <-dbReadOut:
		if out == nil {
			t.Error("dbWrite failed with valid test input with error:")
		}
	case err := <-dbReadErr:
		t.Error(err)
	}

}

// func TestPut(t *testing.T) {
// 	t.Error("not tested")
// }
//
// func TestPutInBucket(t *testing.T) {
// 	t.Error("not tested")
// }
//
// func TestGet(t *testing.T) {
// 	t.Error("not tested")
// }
//
// func TestGetFromBucket(t *testing.T) {
// 	t.Error("not tested")
// }

func TestClose(t *testing.T) {
	err := Close()
	if err != nil {
		t.Error("Close failed: ", err)
	}
}

func cleanup() {
	os.Remove("jsonStore.db")
	if _, err := os.Stat("jsonStore.db"); err == nil {
		log.Println("unable to cleanup jsonStore.db file")
	}
}
