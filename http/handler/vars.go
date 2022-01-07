package handler

import "errors"

var (
	not_found      string = "NOT_FOUND"
	not_authorized string = "Unauthorized"
	forbidden      string = "forbidden"
	not_200        error  = errors.New("Request did not return a 200.")
)
