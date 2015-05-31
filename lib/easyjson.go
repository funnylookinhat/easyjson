package easyjson

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type JSData interface{}

func DecodeJson(str []byte) (JSData, error) {
	var data JSData

	json.Unmarshal(str, &data)

	return data, nil
}

func GetString(j JSData, keys ...interface{}) (string, error) {
	v, err := getValue(j, keys...)

	if err != nil {
		return "", err
	}

	if fmt.Sprintf("%s", reflect.TypeOf(v)) != "string" {
		return "", fmt.Errorf("Value is not a string.")
	}

	return v.(string), nil
}

func GetFloat(j JSData, keys ...interface{}) (float64, error) {
	v, err := getValue(j, keys...)

	if err != nil {
		return 0, err
	}

	if fmt.Sprintf("%s", reflect.TypeOf(v)) != "float64" {
		return 0, fmt.Errorf("Value is not a float64.")
	}

	return v.(float64), nil
}

func GetBool(j JSData, keys ...interface{}) (bool, error) {
	v, err := getValue(j, keys...)

	if err != nil {
		return false, err
	}

	if fmt.Sprintf("%s", reflect.TypeOf(v)) != "bool" {
		return false, fmt.Errorf("Value is not a bool.")
	}

	return v.(bool), nil
}

func GetSlice(j JSData, keys ...interface{}) ([]interface{}, error) {
	if len(keys) == 0 {
		v := j.([]interface{})

		if fmt.Sprintf("%s", reflect.TypeOf(j.(interface{}))) != "[]interface {}" {
			return nil, fmt.Errorf("Value is not a slice.")
		}

		return v, nil
	}

	v, err := getValue(j, keys...)

	if err != nil {
		return nil, err
	}

	if fmt.Sprintf("%s", reflect.TypeOf(v)) != "[]interface {}" {
		return nil, fmt.Errorf("Value is not a slice.")
	}

	return v.([]interface{}), nil
}

func GetMap(j JSData, keys ...interface{}) (map[string]interface{}, error) {
	if len(keys) == 0 {
		v := j.(map[string]interface{})

		if fmt.Sprintf("%s", reflect.TypeOf(j.(interface{}))) != "map[string]interface {}" {
			return nil, fmt.Errorf("Value is not a map.")
		}

		return v, nil
	}

	v, err := getValue(j, keys...)

	if err != nil {
		return nil, err
	}

	if fmt.Sprintf("%s", reflect.TypeOf(v)) != "map[string]interface {}" {
		return nil, fmt.Errorf("Value is not a map.")
	}

	return v.(map[string]interface{}), nil
}

func getValue(j interface{}, keys ...interface{}) (interface{}, error) {
	t := fmt.Sprintf("%s", reflect.TypeOf(j))

	if t == "map[string]interface {}" {
		return getMapValue(j.(map[string]interface{}), keys...)
	}

	if t == "[]interface {}" {
		return getSliceValue(j.([]interface{}), keys...)
	}

	return nil, fmt.Errorf("Invalid JSData.")
}

func getMapValue(m map[string]interface{}, keys ...interface{}) (interface{}, error) {
	k := fmt.Sprintf("%s", keys[0])

	if _, ok := m[k]; !ok {
		return nil, fmt.Errorf("Key doesn't exist: %s", k)
	}

	if len(keys) == 1 {
		return m[k], nil
	}

	t := fmt.Sprintf("%s", reflect.TypeOf(m[k]))

	if t == "map[string]interface {}" {
		return getMapValue(m[k].(map[string]interface{}), keys[1:]...)
	}

	if t == "[]interface {}" {
		return getSliceValue(m[k].([]interface{}), keys[1:]...)
	}

	return nil, fmt.Errorf("Key doesn't exist: %s", k)
}

func getSliceValue(s []interface{}, keys ...interface{}) (interface{}, error) {
	i := keys[0].(int)

	if s[i] == nil {
		return nil, fmt.Errorf("Key doesn't exist: %d", i)
	}

	if len(keys) == 1 {
		return s[i], nil
	}

	t := fmt.Sprintf("%s", reflect.TypeOf(s[i]))

	if t == "map[string]interface {}" {
		return getMapValue(s[i].(map[string]interface{}), keys[1:]...)
	}

	if t == "[]interface {}" {
		return getSliceValue(s[i].([]interface{}), keys[1:]...)
	}

	return nil, fmt.Errorf("Key doesn't exist: %d", i)
}
