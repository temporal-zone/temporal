package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// CompoundSettingKeyPrefix is the prefix to retrieve all CompoundSetting
	CompoundSettingKeyPrefix = "CompoundSetting/value/"
)

// CompoundSettingKey returns the store key to retrieve a CompoundSetting from the index fields
func CompoundSettingKey(
	index123 string,
) []byte {
	var key []byte

	index123Bytes := []byte(index123)
	key = append(key, index123Bytes...)
	key = append(key, []byte("/")...)

	return key
}
