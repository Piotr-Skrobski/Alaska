package errors

import "errors"

var ErrInvalidMovieTitle = errors.New("title cannot be empty")
var ErrInvalidMovieIMBdID = errors.New("id cannot be empty")
var ErrMovieNotFound = errors.New("movie not found")
