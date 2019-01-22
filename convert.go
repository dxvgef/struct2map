package struct2map

import (
	"errors"
	"reflect"
)

// Convert 转换单个结构体
func Convert(s interface{}, fields []string, tagName string) (map[string]interface{}, error) {
	if s == nil {
		return nil, errors.New("struct cannot be empty")
	}

	fieldsLen := len(fields)
	if fieldsLen == 0 {
		return nil, errors.New("fields parameter cannot be empty")
	}

	if tagName == "" {
		return nil, errors.New("tagName parameter cannot be empty")
	}

	result := make(map[string]interface{}, fieldsLen)
	for k := range fields {
		result[fields[k]] = 0
	}

	valueOf := reflect.ValueOf(s).Elem()
	numField := valueOf.NumField()
	if numField == 0 {
		return nil, errors.New("No fields in the struct")
	}
	for i := 0; i < numField; i++ {
		key := valueOf.Type().Field(i).Tag.Get(tagName)
		value := valueOf.Field(i).Interface()
		if key != "" && result[key] != nil {
			result[key] = value
		}
	}

	if len(result) == 0 {
		result = nil
	}

	return result, nil
}

// ConvertSlice 转换结构体Slice
func ConvertSlice(s interface{}, fields []string, tagName string) ([]map[string]interface{}, error) {
	if s == nil {
		return nil, errors.New("struct cannot be empty")
	}

	if reflect.ValueOf(s).Kind().String() != "slice" {
		return nil, errors.New("Not a slice of struct")
	}

	fieldsLen := len(fields)
	if fieldsLen == 0 {
		return nil, errors.New("fields parameter cannot be empty")
	}

	if tagName == "" {
		return nil, errors.New("tagName parameter cannot be empty")
	}

	sLen := reflect.ValueOf(s).Len()
	if sLen == 0 {
		return nil, nil
	}

	result := make([]map[string]interface{}, sLen)

	for i := 0; i < sLen; i++ {
		elem := make(map[string]interface{}, fieldsLen)
		for k := range fields {
			elem[fields[k]] = 0
		}

		valueOf := reflect.ValueOf(s).Index(i)
		numField := valueOf.NumField()
		for i := 0; i < numField; i++ {
			key := valueOf.Type().Field(i).Tag.Get(tagName)
			value := valueOf.Field(i).Interface()
			if key != "" && elem[key] != nil {
				elem[key] = value
			}
		}
		if len(elem) > 0 {
			result[i] = elem
		}
	}

	if len(result) == 0 {
		result = nil
	}

	return result, nil
}
