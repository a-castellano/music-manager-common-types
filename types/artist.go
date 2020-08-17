package types

type Artist struct {
	Name    string
	URL     string
	ID      int
	Genre   string
	Country string
	Records []Record
}
