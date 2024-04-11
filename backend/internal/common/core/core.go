package core

import "reflect"

type Hash map[string]any

func HashGet[T any](hash Hash, name string, defaultVals ...T) T {
	originVal, ok := hash[name]
	var val T
	if ok {
		val = originVal.(T)
	} else if len(defaultVals) > 0 {
		val = defaultVals[0]
	} else {
		val = *new(T)
	}

	return val
}

func Pluck[T, R any](slice []T, field string) []R {
	ans := make([]R, len(slice))
	for i, v := range slice {
		ans[i] = FieldByName[T, R](v, field)
	}

	return ans
}

func Map[T, R any](slice []T, mapFn func(T) R) []R {
	ans := make([]R, len(slice))
	for i, v := range slice {
		ans[i] = mapFn(v)
	}
	return ans
}

func FieldByName[T any, R any](obj T, field string) R {
	return reflect.ValueOf(obj).FieldByName(field).Interface().(R)
}
