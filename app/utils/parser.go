package utils

import (
	"fmt"
	"strings"
)

/* Разобрать строку по разделителю и вернуть в виде массива */
func GetArrayFromString(s string, sep string) ([]string, error) {
	array := strings.Split(s, sep)

	var newArray []string
	for _, item := range array {
		trimmed := strings.TrimSpace(item)
		if trimmed != "" {
			newArray = append(newArray, trimmed)
		}
	}

	if len(newArray) == 0 {
		return newArray, fmt.Errorf("не удалось сформировать список из строки: %s", s)
	}

	return newArray, nil
}
