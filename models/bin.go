package models

import (
	"errors"
	"net/http"
	"time"

	"github.com/google/jsonapi"
)

// ErrBinNotFound is returned when an bin cannot be found
var ErrBinNotFound = errors.New("Bin could not be found")

// Bin is a designated space to partition requests
type Bin struct {
	ID          string `jsonapi:"primary,bin"`
	Name        string `jsonapi:"attr,name"`
	Description string `jsonapi:"attr,description"`
	RemoteAddr  string
	Created     time.Time
}

// NewBin generates a Bin struct to use
func NewBin(r *http.Request) (*Bin, error) {

	bin := &Bin{}

	if err := jsonapi.UnmarshalPayload(r.Body, bin); err != nil {
		return nil, err
	}

	return bin, nil
}
