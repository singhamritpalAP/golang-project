package constants

import "errors"

var (
	ErrUnableToFetchPerson = errors.New("person not found")
	ErrFailedToFetchPerson = errors.New("failed to fetch person information")
)
