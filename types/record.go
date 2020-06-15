package types

type RecordType int

const (
	FullLength RecordType = 1 << iota
	Demo
	EP
	Compilation
	Live
	BoxedSet
	Single
	Video
	Split
	Other
)

type Record struct {
	Name   string
	ID     int
	Year   int
	URL    string
	Type   RecordType
	Tracks []Track
}
