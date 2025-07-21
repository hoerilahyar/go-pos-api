package utils

import (
	"encoding/json"
	"strconv"
)

func StrToUint(str string) (uint, error) {
	userIdUint, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(userIdUint), nil
}

func StrToInt(str string) (int, error) {
	userIdint, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(userIdint), nil
}

func UintToStr(num uint) string {
	return strconv.FormatUint(uint64(num), 10)
}

func CheckIfJSON(input string) interface{} {
	// Try to unmarshal as generic JSON (map or slice)
	var js interface{}
	if err := json.Unmarshal([]byte(input), &js); err != nil {
		// Not a valid JSON, return as string
		return input
	}
	// Valid JSON, return as unmarshaled data
	return js
}
