package hw1

import (
	"fmt"
	"time"
)

type SearchError struct {
	err  error
	date time.Time
}

func (e *SearchError) Error() string {
	return fmt.Sprintf("%s search error: %s", e.date, e.err)
}

func (e *SearchError) Unwrap() error {
	return e.err
}

func WrapSearchError(err error) error {
	return &SearchError{
		err:  err,
		date: time.Now(),
	}
}
