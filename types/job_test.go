package types

import (
	"testing"
)

func TestEncodeAndDecode(t *testing.T) {

	var job Job
	job.ID = 12
	test, _ := EncodeJob(job)
	result, _ := DecodeJob(test)

	if result.ID != 12 {
		t.Errorf(`Encode failed.`)
	}
}
