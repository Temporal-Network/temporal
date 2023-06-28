package types

import "errors"

var (
	ErrAlreadyFulfilled    = errors.New("query already fulfilled")
	ErrSucceededNoDelete   = errors.New("query succeeded; do not not execute default behavior")
	ErrInvalidICQProof     = errors.New("icq query response failed")
	ErrICQCallbackNotFound = errors.New("icq callback id not found")
)
