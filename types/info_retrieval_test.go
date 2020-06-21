package types

import (
	"testing"
)

func TestEncodeAndDecodeInfoRetrieval(t *testing.T) {

	var info InfoRetrieval
	info.Artist = "Perkele"
	test, _ := EncodeInfoRetrieval(info)
	result, _ := DecodeInfoRetrieval(test)

	if result.Artist != "Perkele" {
		t.Errorf(`Encode failed.`)
	}

}

func TestDecodeEmptyDataInfoRetrieva(t *testing.T) {

	var emptyData []byte
	_, err := DecodeInfoRetrieval(emptyData)
	if err == nil {
		t.Error("Empty data decoding should fail.")
	}
}
