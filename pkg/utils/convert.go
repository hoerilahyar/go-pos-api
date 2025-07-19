package utils

import "strconv"

func StrToUint(str string) (uint, error) {
	userIdUint, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(userIdUint), nil
}
