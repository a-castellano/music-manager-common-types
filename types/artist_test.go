package types

import (
	"testing"
)

func TestEncodeAndDecodeArtists(t *testing.T) {

	var artist Artist
	artist.Name = "Bølzer"
	artist.URL = "https://www.metal-archives.com/bands/B%C3%B8lzer/3540351548"
	artist.ID = 3540351548
	artist.Genre = "Black Metal"
	artist.Country = "Switzerland"
	test, _ := EncodeArtist(artist)
	result, _ := DecodeArtist(test)

	if result.ID != 3540351548 {
		t.Errorf(`Encode failed.`)
	}

	if result.Country != "Switzerland" {
		t.Errorf(`Encode failed.`)
	}

}

func TestDecodeEmptyDataArtists(t *testing.T) {

	var emptyData []byte
	_, err := DecodeJob(emptyData)
	if err == nil {
		t.Error("Empty data decoding should fail.")
	}
}
