package serialize

import (
	"DNA/common"
	"encoding/hex"
)

func ByteArrayToString(data []byte) string {
	return hex.EncodeToString(data)
}

func Uint256ToString(uint256 common.Uint256) string {
	return ByteArrayToString(uint256.ToArray())
}

func Uint160ToString(uint160 common.Uint160) string {
	return ByteArrayToString(uint160.ToArray())
}

func StringToUint256(uint256 string) (common.Uint256, error) {
	data, err := StringToByteArray(uint256)
	if err != nil {
		return common.Uint256{}, err
	}
	return common.Uint256ParseFromBytes(data)
}

func StringToUint160(uint160 string) (common.Uint160, error) {
	data, err := StringToByteArray(uint160)
	if err != nil {
		return common.Uint160{}, err
	}
	return common.Uint160ParseFromBytes(data)
}

func StringToByteArray(data string) ([]byte, error) {
	return hex.DecodeString(data)
}
