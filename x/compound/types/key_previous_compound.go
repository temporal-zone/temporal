package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PreviousCompoundKeyPrefix is the prefix to retrieve all PreviousCompounds
	PreviousCompoundKeyPrefix = "PreviousCompound/value/"
)

// PreviousCompoundKey returns the store key to retrieve a PreviousCompound from the index fields
func PreviousCompoundKey(delegator string) []byte {
	var key []byte

	delegatorBytes := []byte(delegator)
	key = append(key, delegatorBytes...)
	key = append(key, []byte("/")...)

	return key
}
