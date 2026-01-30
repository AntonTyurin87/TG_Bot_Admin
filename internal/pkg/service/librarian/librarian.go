package librarian

import (
	lib "github.com/AntonTyurin87/Recon_Com_protoc/gen/go/librarian"
)

const (
	serviceLibrarianAddress = "localhost:50052"
)

type librarian struct {
	librarianClient lib.LibrarianClient
}

// NewLibrarian...
func NewLibrarian(
	librarianClient lib.LibrarianClient,
) *librarian {
	return &librarian{
		librarianClient: librarianClient,
	}
}
