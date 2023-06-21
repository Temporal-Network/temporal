package types

import (
	"bytes"
	"encoding/json"
	"errors"
)

// RawContractMessage defines a json message that is sent or returned by a wasm contract.
// This type can hold any type of bytes. Until validateBasic is called there should not be
// any assumptions made that the data is valid syntax or semantic.
type RawContractMessage []byte

func (r RawContractMessage) MarshalJSON() ([]byte, error) {
	return json.RawMessage(r).MarshalJSON()
}

func (r *RawContractMessage) UnmarshalJSON(b []byte) error {
	if r == nil {
		return errors.New("unmarshalJSON on nil pointer")
	}
	*r = append((*r)[0:0], b...)
	return nil
}

func (r *RawContractMessage) ValidateBasic() error {
	if r == nil {
		return errors.New("ErrEmpty")
	}
	if !json.Valid(*r) {
		return errors.New("ErrInvalid")
	}
	return nil
}

// Bytes returns raw bytes type
func (r RawContractMessage) Bytes() []byte {
	return r
}

// Equal content is equal json. Byte equal but this can change in the future.
func (r RawContractMessage) Equal(o RawContractMessage) bool {
	return bytes.Equal(r.Bytes(), o.Bytes())
}
