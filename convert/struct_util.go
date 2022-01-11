package convert

import (
	"reflect"
	"strconv"
	"strings"
)

func Struct2Map(obj interface{}, fn func(value reflect.Value) interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Ptr {
		return nil
	}
	if t.Elem().Kind() != reflect.Struct {
		return nil
	}
	t = t.Elem()

	refValue := reflect.ValueOf(obj)
	m := make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tagStr, ok := field.Tag.Lookup(TagJson)
		if !ok {
			continue
		}
		key, ignore, omitempty := getKey(tagStr)
		if ignore || key == EmptyStr {
			continue
		}

		value := reflect.Indirect(refValue).Field(i)
		if value.IsZero() && omitempty {
			continue
		}

		m[key] = fn(value)
	}
	return m
}

func ReflectValue2Str(value reflect.Value, fn func(value reflect.Value) string) string {
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(value.Int(), 10)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(value.Uint(), 10)

	case reflect.String:
		return value.String()

	case reflect.Ptr:
		return ReflectValue2Str(value.Elem(), fn)

	case reflect.Bool:
		return strconv.FormatBool(value.Bool())

	default:
		return fn(value)
	}
}

func ReflectBaseType2Str(value reflect.Value) string {
	str := ReflectValue2Str(value, func(value reflect.Value) string { return EmptyStr })
	return str
}

func getKey(tagStr string) (key string, ignore, omitempty bool) {

	key = strings.Split(tagStr, CommaStr)[0]
	if key == TagIgnore {
		ignore = true
		return
	}
	if strings.Contains(tagStr, TagOmitEmpty) {
		omitempty = true
	}

	return
}
