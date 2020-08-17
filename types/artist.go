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
