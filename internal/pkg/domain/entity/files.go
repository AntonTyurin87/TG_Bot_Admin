package entity

// File ...
type File struct {
	ID       int64  `json:"id"`
	SourceID int64  `json:"source_id"`
	FileData []byte `json:"file_data"`
}

// GetID ...
func (f *File) GetID() int64 {
	if f == nil {
		return 0
	}
	return f.ID
}

// GetSourceID ...
func (f *File) GetSourceID() int64 {
	if f == nil {
		return 0
	}
	return f.SourceID
}

// GetFileData ...
func (f *File) GetFileData() []byte {
	if f == nil {
		return nil
	}
	return f.FileData
}
