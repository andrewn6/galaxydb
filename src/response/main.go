package response

import (
	"github.com/anddddrew/galaxydb/src/object"
)

const (
	INVALID_COMMAND_ERR = "INVALID_COMMAND_ERR"
)

// Struct for the response cache
type CacheResponse struct {
	Gobj object.CacheObject

	Status int32

	Message string

	Error string
}

func NewResponseFromVal(value interface{}) CacheResponse {
	return CacheResponse{}
}
