package json

import (
	"reflect"
	"strings"
	"fmt"
	"errors"
)

var ErrInvalidType = errors.New("invalid data type")

func getKey(t_field reflect.StructField, tagDetails string) (int, string) {// 0 for delete(when null), 1 for not delete
	switch tagDetails {
	case "":
		return 1, string(t_field.Name)
	case "-":
		return 0, ""
	case "-,":
		return 1, "-"
	case ",omitempty":
		return 0, string(t_field.Name)
	}
	arr := strings.Split(tagDetails, ",")
	if (len(arr) == 1) {
		return 1, arr[0]
	} else {
		for _, str := range(arr){
			if (str == "omitempty") {
				return 0, arr[0]
			}
		}
		return 1, arr[0]
	}
}

func stringToJson(value string) string{
	ret := ""
	cSlice := []rune(value)

	for _, c := range(cSlice) {
		switch c {
		case '<':
			ret += "\\u003c"
		case '>':
			ret += "\\u003e"
		case '&':
			ret += "\\u0026"
		default:
			ret += string(c)
		}
	}
	return ret
}

func JsonMarshal(v interface{}) ([]byte, error) {
	res := ""

	v_type := reflect.TypeOf(v)
	v_value := reflect.ValueOf(v)
	baseType := v_type.Kind()

	switch baseType {
	case reflect.Struct:
		for i := 0; i < v_type.NumField(); i++ {
			t_field := v_type.Field(i)
			v_field := v_value.Field(i)
			tagDetails := t_field.Tag.Get("json")

			flag, key := getKey(t_field, tagDetails)
			if flag == 0 {
				if key == "" || v_field.Interface() == 0 || v_field.Interface() == "" || v_field.Interface() == nil {
					continue
				}
			}

			key = "\"" + key + "\""
			bytes, err := JsonMarshal(v_field.Interface())
			if err != nil {
				return nil, ErrInvalidType				
			} else {
				res += key + ":" + string(bytes) + ","
			}			
		}
		if res[len(res)-1] == ',' {
			res = "{" + res[:len(res)-1] + "}"
		} else {
			res = "{" + res + "}"
		}
		return []byte(res), nil

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64,
			reflect.Bool:
		res += fmt.Sprintf("%v", v)
		return []byte(res), nil

	case reflect.String:
		res += "\"" + stringToJson(fmt.Sprintf("%v", v)) + "\""
		return []byte(res), nil

	case reflect.Slice:
		for i := 0; i < v_value.Len(); i++ {
			bytes, err := JsonMarshal(v_value.Index(i).Interface())
			if err != nil {
				return nil, ErrInvalidType				
			} else {
				res += string(bytes) + ","
			}
		}
		if res[len(res)-1] == ',' {
			res = "[" + res[:len(res)-1] + "]"
		} else {
			res = "[" + res + "]"
		}
		return []byte(res), nil
		
	case reflect.Map:
		keys := v_value.MapKeys()
		for _, key := range(keys) {
			value := v_value.MapIndex(key)
			keyBytes, err1 := JsonMarshal(key.Interface())
			valueBytes, err2 := JsonMarshal(value.Interface())

			if err1 != nil || err2 != nil{
				return nil, ErrInvalidType				
			} else {
				res += "\"" + string(keyBytes) +"\"" +  ":" + string(valueBytes) + ","
			}
		}
		if res[len(res)-1] == ',' {
			res = "{" + res[:len(res)-1] + "}"
		} else {
			res = "{" + res + "}"
		}
		return []byte(res), nil

	case reflect.Ptr, reflect.Uintptr:
		return JsonMarshal(v_value.Elem().Interface())

	default:
		return nil, ErrInvalidType
	}
	return []byte(res), nil
}