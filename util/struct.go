package util

import (
	"reflect"
)

func ExtractStructFieldNames[T any](s T) []string {
	t := reflect.TypeOf(s)

	names := make([]string, t.NumField())

	for i := range names {
		names[i] = t.Field(i).Name
	}

	return names
}
