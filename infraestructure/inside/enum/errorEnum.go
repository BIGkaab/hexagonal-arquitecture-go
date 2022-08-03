package enum

import "errors"

var (
	INTERNAL_SERVER_ERROR = errors.New("Internal server error")
	BAB_REQUEST           = errors.New("Bad request")
	NOT_FOUND             = errors.New("Not found")
)
