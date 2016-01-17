package byteStore

import "testing"

import "os"
import "log"

func TestInit(t *testing.T) {
	if _, err := os.Stat("byteStore.db"); err != nil {
		t.Error("balls")
	}
}

func TestPutAndGet(t *testing.T) {
	testValue := "stored value!"
	err := Put("testBucket", "testKey", []byte(testValue))
	if err != nil {
		t.Log("Put failed with error:", err)
	}

	getValue := string(Get("testBucket", "testKey"))
	if getValue != testValue {
		t.Error("Get failed with error:", err)
	}

	getNone := Get("noBucket", "noKey")
	if getNone != nil {
		t.Error("a Get on an empty bucket should return an empty value")
	}
}

func TestGetBucket(t *testing.T) {
	testFirstValue := "first stored value!"
	testLastValue := "last stored value!"
	Put("testGetBucket", "1", []byte(testFirstValue))
	Put("testGetBucket", "2", []byte("blah"))
	Put("testGetBucket", "3", []byte(testLastValue))

	getValues := GetBucket("testGetBucket")

	if len(getValues) != 3 {
		t.Error("GetBucket did not return the same amount of values as was in the test bucket")
	}

	if string(getValues[0]) != testFirstValue {
		t.Error("GetBucket did not return the correct first value")
	}

	if string(getValues[2]) != testLastValue {
		t.Error("GetBucket did not return the correct first value")
	}

	getNoValues := GetBucket("uninitialisedBucket")
	if getNoValues != nil {
		t.Error("GetBucket on an empty bucket should return a nil slice")
	}
}

func TestClose(t *testing.T) {
	err := Close()
	if err != nil {
		t.Error("Close failed: ", err)
	}
	cleanup()
}

func cleanup() {
	os.Remove("byteStore.db")
	if _, err := os.Stat("byteStore.db"); err == nil {
		log.Println("unable to cleanup byteStore.db file")
	}
}
