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
