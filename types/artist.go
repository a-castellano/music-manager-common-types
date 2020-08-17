package types

import (
	"bytes"
	"encoding/gob"
)

type Artist struct {
	Name    string
	URL     string
	ID      int
	Genre   string
	Country string
	Records []Record
}

type ArtistInfo struct {
	Data      Artist
	ExtraData []Artist
}

func EncodeArtist(artist Artist) ([]byte, error) {
	var encodedArtist []byte
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	err := enc.Encode(artist)
	if err != nil {
		return encodedArtist, err
	}
	encodedArtist = network.Bytes()
	return encodedArtist, nil
}

func DecodeArtist(encoded []byte) (Artist, error) {
	var artist Artist
	network := bytes.NewBuffer(encoded)
	dec := gob.NewDecoder(network)
	err := dec.Decode(&artist)
	if err != nil {
		return artist, err
	}
	return artist, nil
}

func EncodeArtistInfo(artists ArtistInfo) ([]byte, error) {
	var encodedArtistInfo []byte
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	err := enc.Encode(artists)
	if err != nil {
		return encodedArtistInfo, err
	}
	encodedArtistInfo = network.Bytes()
	return encodedArtistInfo, nil
}

func DecodeArtistInfo(encoded []byte) (ArtistInfo, error) {
	var artistinfo ArtistInfo
	network := bytes.NewBuffer(encoded)
	dec := gob.NewDecoder(network)
	err := dec.Decode(&artistinfo)
	if err != nil {
		return artistinfo, err
	}
	return artistinfo, nil
}
