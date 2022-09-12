package main

import (
	"errors"
	"strconv"
	"strings"
)

func ParseInputToMinutes(value string) (int, error) {
	if strings.HasSuffix(value, "m") {
		parts := strings.Split(value, "m")
		return toMinutes(parts[0], 1)
	} else if strings.HasSuffix(value, "h") {
		parts := strings.Split(value, "h")
		return toMinutes(parts[0], 60)
	}

	return 0, errors.New("Input must end in `m` or `h`")
}

func toMinutes(value string, factor int) (int, error) {
	number, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	return number * factor, err
}
