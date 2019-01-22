package struct2map

import (
	"errors"
	"reflect"
)

// Convert 转换单个结构体
func Convert(s interface{}, fields []string, tagName string) (map[string]interface{}, error) {
	if s == nil {
		return nil, errors.New("结构体不能为空")
	}

	fieldsLen := len(fields)
	if fieldsLen == 0 {
		return nil, errors.New("fields长度不能为0")
	}

	if tagName == "" {
		return nil, errors.New("tagName参数不能为空")
	}

	result := make(map[string]interface{}, fieldsLen)
	for k := range fields {
		result[fields[k]] = 0
	}

	valueOf := reflect.ValueOf(s).Elem()
	numField := valueOf.NumField()
	if numField == 0 {
		return nil, errors.New("结构体没有字段")
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
		return nil, errors.New("结构体不能为空")
	}

	if reflect.ValueOf(s).Kind().String() != "slice" {
		return nil, errors.New("传入的不是结构体Slice")
	}

	fieldsLen := len(fields)
	if fieldsLen == 0 {
		return nil, errors.New("fields长度不能为0")
	}

	if tagName == "" {
		return nil, errors.New("tagName参数不能为空")
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
