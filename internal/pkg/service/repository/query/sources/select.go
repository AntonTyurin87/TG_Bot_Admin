package sources

import "github.com/golang/protobuf/ptypes/timestamp"

// Select ...
type Select struct {
	IDs           []int64
	UserIDs       []int64
	Steps         []int64
	Types         []string
	NamesRU       []string
	NamesENG      []string
	AuthorsRU     []string
	Years         []int64
	RegionsIn     []string
	TimePeriodsIn []string
	Descriptions  []string
	FileFormats   []string
	CreatedAtIn   []*timestamp.Timestamp
	IsSentIn      []int64
}
