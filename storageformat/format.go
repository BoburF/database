package storageformat

import (
	"fmt"
	"reflect"
	"strings"
)

func ToStorageFormat(input interface{}) string {
	val := reflect.ValueOf(input)
	typ := reflect.TypeOf(input)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	var result []string

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("bsf")
		if tag == "" {
			tag = field.Name
		}

		fieldValue := val.Field(i)

		switch fieldValue.Kind() {
		case reflect.String:
			result = append(result, fmt.Sprintf("%s=%s", tag, fieldValue.String()))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			result = append(result, fmt.Sprintf("%s=%d", tag, fieldValue.Int()))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			result = append(result, fmt.Sprintf("%s=%d", tag, fieldValue.Uint()))
		case reflect.Float32, reflect.Float64:
			result = append(result, fmt.Sprintf("%s=%f", tag, fieldValue.Float()))
		case reflect.Bool:
			result = append(result, fmt.Sprintf("%s=%t", tag, fieldValue.Bool()))
		case reflect.Struct:
			nested := ToStorageFormat(fieldValue.Interface())
			for _, entry := range strings.Split(nested, "\x00\x00\x00") {
				result = append(result, fmt.Sprintf("%s.%s", tag, entry))
			}
		case reflect.Slice, reflect.Array:
			for j := 0; j < fieldValue.Len(); j++ {
				element := fieldValue.Index(j).Interface()
				result = append(result, fmt.Sprintf("%s[%d]=%v", tag, j, element))
			}
		case reflect.Map:
			for _, key := range fieldValue.MapKeys() {
				mapValue := fieldValue.MapIndex(key)
				result = append(result, fmt.Sprintf("%s[%v]=%v", tag, key.Interface(), mapValue.Interface()))
			}
		default:
			result = append(result, fmt.Sprintf("%s=unsupported_type", tag))
		}
	}

	return strings.Join(result, "\x00\x00\x00")
}

func ToStruct(input string, output interface{}) error {
	val := reflect.ValueOf(output)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("output must be a pointer to a struct")
	}

	val = val.Elem()
	typ := val.Type()

	pairs := strings.Split(input, "\x00\x00\x00")
	for _, pair := range pairs {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) != 2 {
			continue
		}
		key, value := kv[0], kv[1]

		// Handle nested keys and arrays
		if strings.Contains(key, ".") {
			parts := strings.SplitN(key, ".", 2)
			parentKey, childKey := parts[0], parts[1]

			for i := 0; i < val.NumField(); i++ {
				field := typ.Field(i)
				tag := field.Tag.Get("bsf")
				if tag == "" {
					tag = field.Name
				}

				if tag == parentKey {
					nestedField := val.Field(i)
					if nestedField.Kind() == reflect.Struct {
						err := ToStruct(childKey+"="+value, nestedField.Addr().Interface())
						if err != nil {
							return err
						}
					}
					break
				}
			}
			continue
		}

		for i := 0; i < val.NumField(); i++ {
			field := typ.Field(i)
			tag := field.Tag.Get("bsf")
			if tag == "" {
				tag = field.Name
			}

			if tag == key {
				fieldVal := val.Field(i)
				if fieldVal.CanSet() {
					switch fieldVal.Kind() {
					case reflect.String:
						fieldVal.SetString(value)
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						var intValue int64
						fmt.Sscanf(value, "%d", &intValue)
						fieldVal.SetInt(intValue)
					case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
						var uintValue uint64
						fmt.Sscanf(value, "%d", &uintValue)
						fieldVal.SetUint(uintValue)
					case reflect.Float32, reflect.Float64:
						var floatValue float64
						fmt.Sscanf(value, "%f", &floatValue)
						fieldVal.SetFloat(floatValue)
					case reflect.Bool:
						var boolValue bool
						fmt.Sscanf(value, "%t", &boolValue)
						fieldVal.SetBool(boolValue)
					default:
						return fmt.Errorf("unsupported type for key: %s", key)
					}
				}
				break
			}
		}
	}

	return nil
}
