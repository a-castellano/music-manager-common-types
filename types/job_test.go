package types

import (
	"testing"
)

func TestEncodeAndDecodeJobs(t *testing.T) {

	var job Job
	job.ID = "asdasd1221"
	job.Type = RecordInfoRetrieval
	test, _ := EncodeJob(job)
	result, _ := DecodeJob(test)

	if result.ID != job.ID {
		t.Errorf(`Encode failed.`)
	}

	if result.Type != RecordInfoRetrieval {
		t.Errorf(`Encode failed.`)
	}

}

func TestDecodeEmptyDataJobs(t *testing.T) {

	var emptyData []byte
	_, err := DecodeJob(emptyData)
	if err == nil {
		t.Error("Empty data decoding should fail.")
	}
}
