package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// UserInstructionsKeyPrefix is the prefix to retrieve all UserInstructions
	UserInstructionsKeyPrefix = "UserInstructions/value/"
)

// UserInstructionsKey returns the store key to retrieve a UserInstructions from the index fields
func UserInstructionsKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
