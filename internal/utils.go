package internal

import (
	"errors"
	"os"
)

// GetStencilPath will request the STENCIL environment variable from the os and
// return an error if it has not been set.
func GetStencilPath() (string, error) {
	p := os.Getenv("STENCIL")

	if p == "" {
		return p, errors.New("the STENCIL environment variable has not been set")
	}

	return p, nil
}
