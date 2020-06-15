package types

import (
	"bytes"
	"encoding/gob"
	"log"
)

type JobType int

const (
	ArtistInfoRetrieval JobType = 1 << iota
	RecordInfoRetrieval
	JobInfoRetrieval
)

type Job struct {
	ID       int
	Finished bool
	Status   bool
	Data     []byte
}

func EncodeJob(job Job) ([]byte, error) {
	var encodedJob []byte
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	err := enc.Encode(job)
	if err != nil {
		log.Fatal("Encode error:", err)
		return encodedJob, err
	}
	encodedJob = network.Bytes()
	return encodedJob, nil
}

func DecodeJob(encoded []byte) (Job, error) {
	var job Job
	network := bytes.NewBuffer(encoded)
	dec := gob.NewDecoder(network)
	err := dec.Decode(&job)
	if err != nil {
		log.Fatal("Decode error:", err)
		return job, err
	}
	return job, nil
}
