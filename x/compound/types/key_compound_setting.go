package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// CompoundSettingKeyPrefix is the prefix to retrieve all CompoundSetting
	CompoundSettingKeyPrefix = "CompoundSetting/value/"
)

// CompoundSettingKey returns the store key to retrieve a CompoundSetting from the index fields
func CompoundSettingKey(
	delegator string,
) []byte {
	var key []byte

	delegatorBytes := []byte(delegator)
	key = append(key, delegatorBytes...)
	key = append(key, []byte("/")...)

	return key
}
