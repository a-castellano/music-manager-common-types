package types

import (
	"bytes"
	"encoding/gob"
)

type InfoRetrievalType int

const (
	ArtistName InfoRetrievalType = 1 << iota
	ArtistURL
	AlbumName
	AlbumURL
	AlbumWithArtist
)

type InfoRetrieval struct {
	Type   InfoRetrievalType
	URL    string
	Artist string
	Album  string
}

func EncodeInfoRetrieval(info InfoRetrieval) ([]byte, error) {
	var encodedInfoRetrieval []byte
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	err := enc.Encode(info)
	if err != nil {
		return encodedInfoRetrieval, err
	}
	encodedInfoRetrieval = network.Bytes()
	return encodedInfoRetrieval, nil
}

func DecodeInfoRetrieval(encoded []byte) (InfoRetrieval, error) {
	var info InfoRetrieval
	network := bytes.NewBuffer(encoded)
	dec := gob.NewDecoder(network)
	err := dec.Decode(&info)
	if err != nil {
		return info, err
	}
	return info, nil
}
