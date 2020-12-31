package types

import (
	"testing"
)

func TestEncodeAndDecodeArtists(t *testing.T) {

	var artist Artist
	artist.Name = "Bølzer"
	artist.URL = "https://www.metal-archives.com/bands/B%C3%B8lzer/3540351548"
	artist.ID = "3540351548"
	artist.Genre = "Black Metal"
	artist.Country = "Switzerland"
	test, _ := EncodeArtist(artist)
	result, _ := DecodeArtist(test)

	if result.ID != "3540351548" {
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

func TestEncodeAndDecodeMultipleArtists(t *testing.T) {

	var artistinfo ArtistInfo

	var firstArtist Artist
	firstArtist.Name = "Bølzer"
	firstArtist.URL = "https://www.metal-archives.com/bands/B%C3%B8lzer/3540351548"
	firstArtist.ID = "3540351548"
	firstArtist.Genre = "Black Metal"
	firstArtist.Country = "Switzerland"

	var secondArtist Artist
	secondArtist.Name = "Blasphemy"
	secondArtist.URL = "https://www.metal-archives.com/bands/Blasphemy/1141"
	secondArtist.ID = "1141"
	secondArtist.Genre = "Black/Death Metal"
	secondArtist.Country = "Canada"

	var thirdArtist Artist
	thirdArtist.Name = "Vektor"
	thirdArtist.URL = "https://www.metal-archives.com/bands/Vektor/87803"
	thirdArtist.ID = "87803"
	thirdArtist.Genre = "Progressive Thrash Metal"
	thirdArtist.Country = "United States"

	artistinfo.Data = firstArtist
	artistinfo.ExtraData = append(artistinfo.ExtraData, secondArtist, thirdArtist)

	test, _ := EncodeArtistInfo(artistinfo)
	result, _ := DecodeArtistInfo(test)
	if result.Data.Name != "Bølzer" {
		t.Errorf(`Encode failed, main artist slice should be Bølzer.`)
	}

	if len(result.ExtraData) != 2 {
		t.Errorf(`Encode failed, artist extra data slice should have 2 items.`)
	}

	if result.ExtraData[0].Name != "Blasphemy" {
		t.Errorf(`Encode failed, first artist should be Bølzer.`)
	}

	if result.ExtraData[1].Name != "Vektor" {
		t.Errorf(`Encode failed, second artist should be Blasphemy.`)
	}

}
