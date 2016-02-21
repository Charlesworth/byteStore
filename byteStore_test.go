package byteStore

import "testing"

import "os"
import "log"

func TestInit(t *testing.T) {
	if _, err := os.Stat("byteStore.db"); err != nil {
		t.Error("There was an issue with the init(), no .db file is present")
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

func TestGetBucketValues(t *testing.T) {
	testFirstValue := "first stored value!"
	testLastValue := "last stored value!"
	Put("testGetBucketValues", "1", []byte(testFirstValue))
	Put("testGetBucketValues", "2", []byte("blah"))
	Put("testGetBucketValues", "3", []byte(testLastValue))

	getValues := GetBucketValues("testGetBucketValues")

	if len(getValues) != 3 {
		t.Error("GetBucketValues did not return the same amount of values as was in the test bucket")
	}

	if string(getValues[0]) != testFirstValue {
		t.Error("GetBucketValues did not return the correct first value")
	}

	if string(getValues[2]) != testLastValue {
		t.Error("GetBucketValues did not return the correct first value")
	}

	getNoValues := GetBucketValues("uninitialisedBucket")
	if getNoValues != nil {
		t.Error("GetBucketValues on an empty bucket should return a nil slice")
	}
}

func TestGetBucket(t *testing.T) {
	testFirstValue := "first stored value!"
	testLastValue := "last stored value!"
	Put("testGetBucket", "1", []byte(testFirstValue))
	Put("testGetBucket", "2", []byte("blah"))
	Put("testGetBucket", "3", []byte(testLastValue))

	getKeyVals := GetBucket("testGetBucket")

	if len(getKeyVals) != 3 {
		t.Error("GetBucket did not return the same amount of values as was in the test bucket")
	}

	if string(getKeyVals[0].value) != testFirstValue {
		t.Error("GetBucket did not return the correct first value")
	}

	if string(getKeyVals[2].value) != testLastValue {
		t.Error("GetBucket did not return the correct first value")
	}

	getNoKeyVals := GetBucket("uninitialisedBucket")
	if getNoKeyVals != nil {
		t.Error("GetBucket on an empty bucket should return a nil slice")
	}
}

func TestDelete(t *testing.T) {
	Put("testDelete", "1", []byte("hi"))

	getValue := Get("testDelete", "1")
	if string(getValue) != "hi" {
		t.Error("failed to set test value")
	}

	err := Delete("testDelete", "1")
	if err != nil {
		t.Error("Delete failed with error:", err)
	}

	getValue = Get("testDelete", "1")
	if getValue != nil {
		t.Error("Delete did not delete the key/value")
	}
}

func TestDeleteBucket(t *testing.T) {
	Put("testDeleteBucket", "1", []byte("hi"))

	getValue := Get("testDeleteBucket", "1")
	if string(getValue) != "hi" {
		t.Error("failed to set test value")
	}

	err := DeleteBucket("testDeleteBucket")
	if err != nil {
		t.Error("DeleteBucket failed with error:", err)
	}

	getValue = Get("testDeleteBucket", "1")
	if getValue != nil {
		t.Error("Delete did not delete the key/value")
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
