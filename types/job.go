package types

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
