package types

import (
	"a-castellano/music-manager-common-types/types"
	"bytes"
	"encoding/gob"
)

func TestEncodeAndDecode(t *testing.T) {

	var job types.Job
	job.ID = 12
	test, _ := types.EncodeJob(job)
	fmt.Println(test)
	result, _ := types.DecodeJob(test)

	if result.ID != 12 {
		t.Errorf(`Encode failed.`)
	}
}
