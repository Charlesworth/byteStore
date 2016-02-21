package byteStore

import "testing"

import "os"
import "log"

var testBS ByteStore

func TestInit(t *testing.T) {
	var err error
	testBS, err = NewByteStore("byteStore.db")
	if err != nil {
		t.Error("rubbish")
	}

	if _, err = os.Stat("byteStore.db"); err != nil {
		t.Error("balls")
	}
}

func TestPutAndGet(t *testing.T) {
	testValue := "stored value!"
	err := testBS.Put("testBucket", "testKey", []byte(testValue))
	if err != nil {
		t.Log("Put failed with error:", err)
	}

	getValue := string(testBS.Get("testBucket", "testKey"))
	if getValue != testValue {
		t.Error("Get failed with error:", err)
	}

	getNone := testBS.Get("noBucket", "noKey")
	if getNone != nil {
		t.Error("a Get on an empty bucket should return an empty value")
	}
}

func TestGetBucketValues(t *testing.T) {
	testFirstValue := "first stored value!"
	testLastValue := "last stored value!"
	testBS.Put("testGetBucketValues", "1", []byte(testFirstValue))
	testBS.Put("testGetBucketValues", "2", []byte("blah"))
	testBS.Put("testGetBucketValues", "3", []byte(testLastValue))

	getValues := testBS.GetBucketValues("testGetBucketValues")

	if len(getValues) != 3 {
		t.Error("GetBucketValues did not return the same amount of values as was in the test bucket")
	}

	if string(getValues[0]) != testFirstValue {
		t.Error("GetBucketValues did not return the correct first value")
	}

	if string(getValues[2]) != testLastValue {
		t.Error("GetBucketValues did not return the correct first value")
	}

	getNoValues := testBS.GetBucketValues("uninitialisedBucket")
	if getNoValues != nil {
		t.Error("GetBucketValues on an empty bucket should return a nil slice")
	}
}

func TestGetBucket(t *testing.T) {
	testFirstValue := "first stored value!"
	testLastValue := "last stored value!"
	testBS.Put("testGetBucket", "1", []byte(testFirstValue))
	testBS.Put("testGetBucket", "2", []byte("blah"))
	testBS.Put("testGetBucket", "3", []byte(testLastValue))

	getKeyVals := testBS.GetBucket("testGetBucket")

	if len(getKeyVals) != 3 {
		t.Error("GetBucket did not return the same amount of values as was in the test bucket")
	}

	if string(getKeyVals[0].value) != testFirstValue {
		t.Error("GetBucket did not return the correct first value")
	}

	if string(getKeyVals[2].value) != testLastValue {
		t.Error("GetBucket did not return the correct first value")
	}

	getNoKeyVals := testBS.GetBucket("uninitialisedBucket")
	if getNoKeyVals != nil {
		t.Error("GetBucket on an empty bucket should return a nil slice")
	}
}

func TestDelete(t *testing.T) {
	testBS.Put("testDelete", "1", []byte("hi"))

	getValue := testBS.Get("testDelete", "1")
	if string(getValue) != "hi" {
		t.Error("failed to set test value")
	}

	err := testBS.Delete("testDelete", "1")
	if err != nil {
		t.Error("Delete failed with error:", err)
	}

	getValue = testBS.Get("testDelete", "1")
	if getValue != nil {
		t.Error("Delete did not delete the key/value")
	}
}

func TestDeleteBucket(t *testing.T) {
	testBS.Put("testDeleteBucket", "1", []byte("hi"))

	getValue := testBS.Get("testDeleteBucket", "1")
	if string(getValue) != "hi" {
		t.Error("failed to set test value")
	}

	err := testBS.DeleteBucket("testDeleteBucket")
	if err != nil {
		t.Error("DeleteBucket failed with error:", err)
	}

	getValue = testBS.Get("testDeleteBucket", "1")
	if getValue != nil {
		t.Error("Delete did not delete the key/value")
	}
}

func TestClose(t *testing.T) {
	err := testBS.Close()
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
