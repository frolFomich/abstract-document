package document

import (
	"encoding/json"
	"errors"
)

type AbstractDocument struct {
	data map[string]interface{}
}

var (
	ErrorKeyDoesNotExist = errors.New("key doesn't exists")
	ErrorInvalidKey = errors.New("invalid key provided")
	ErrorDataConversion = errors.New("invalid data conversion")
)

func (a *AbstractDocument) Get(key string) interface{} {
	if a.data == nil || key == "" {
		return nil
	}
	return a.data[key]
}

func (a *AbstractDocument) Put(key string, value interface{}) {
	if key == "" {
		return
	}
	a.data[key] = value
}

func (a *AbstractDocument) Children(key string, constructor ConstructorFunc) []interface{} {
	empty := make([]interface{}, 0)
	if key == "" || constructor == nil {
		return empty
	}
	list := a.Slice(key)
	result := make([]interface{}, len(list))
	for i, val := range list {
		m, ok := val.(map[string]interface{})
		if !ok {
			return empty
		}
		result[i] = constructor(m)
	}
	return result
}

func (a *AbstractDocument) Slice(key string) []interface{} {
	empty := make([]interface{}, 0)
	if key == "" {
		return empty
	}
	list,ok := a.Get(key).([]interface{})
	if !ok {
		return empty
	}
	return list
}

func (a *AbstractDocument) Map(key string) map[string]interface{} {
	empty := make(map[string]interface{}, 0)
	if key == "" {
		return empty
	}
	i := a.Get(key)
	if i == nil {
		return empty
	}
	if b, ok := i.(map[string]interface{}); ok {
		return b
	}
	return empty
}

func (a *AbstractDocument) Bool(key string) (bool, error) {
	if key == "" {
		return false, ErrorInvalidKey
	}
	i := a.Get(key)
	if i == nil {
		return false, ErrorKeyDoesNotExist
	}
	if b, ok := i.(bool); ok {
		return b, nil
	}
	return false, ErrorDataConversion
}

func (a *AbstractDocument) String(key string) (string, error) {
	if key == "" {
		return "", ErrorInvalidKey
	}
	i := a.Get(key)
	if i == nil {
		return "", ErrorKeyDoesNotExist
	}
	if b, ok := i.(string); ok {
		return b, nil
	}
	return "", ErrorDataConversion
}

func (a *AbstractDocument) Byte(key string) (byte, error) {
	if key == "" {
		return 0, ErrorInvalidKey
	}
	i := a.Get(key)
	if i == nil {
		return 0, ErrorKeyDoesNotExist
	}
	if b, ok := i.(byte); ok {
		return b, nil
	}
	return 0, ErrorDataConversion
}

func (a *AbstractDocument) Rune(key string) (rune, error) {
	if key == "" {
		return 0, ErrorInvalidKey
	}
	i := a.Get(key)
	if i == nil {
		return 0, ErrorKeyDoesNotExist
	}
	if b, ok := i.(rune); ok {
		return b, nil
	}
	return 0, ErrorDataConversion
}

func (a *AbstractDocument) Int(key string) (int, error) {
	if key == "" {
		return 0, ErrorInvalidKey
	}
	i := a.Get(key)
	if i == nil {
		return 0, ErrorKeyDoesNotExist
	}
	if b, ok := i.(int); ok {
		return b, nil
	}
	return 0, ErrorDataConversion
}

func (a *AbstractDocument) Int8(key string) (int8, error) {
	if key == "" {
		return 0, ErrorInvalidKey
	}
	i := a.Get(key)
	if i == nil {
		return 0, ErrorKeyDoesNotExist
	}
	if b, ok := i.(int8); ok {
		return b, nil
	}
	return 0, ErrorDataConversion
}

func (a *AbstractDocument) Int16(key string) (int16, error) {
	if key == "" {
		return 0, ErrorInvalidKey
	}
	i := a.Get(key)
	if i == nil {
		return 0, ErrorKeyDoesNotExist
	}
	if b, ok := i.(int16); ok {
		return b, nil
	}
	return 0, ErrorDataConversion
}

func (a *AbstractDocument) Int32(key string) (int32, error) {
	if key == "" {
		return 0, ErrorInvalidKey
	}
	i := a.Get(key)
	if i == nil {
		return 0, ErrorKeyDoesNotExist
	}
	if b, ok := i.(int32); ok {
		return b, nil
	}
	return 0, ErrorDataConversion
}

func (a *AbstractDocument) Int64(key string) (int64, error) {
	if key == "" {
		return 0, ErrorInvalidKey
	}
	i := a.Get(key)
	if i == nil {
		return 0, ErrorKeyDoesNotExist
	}
	if b, ok := i.(int64); ok {
		return b, nil
	}
	return 0, ErrorDataConversion
}

func (a *AbstractDocument) Uint(key string) (uint, error) {
	if key == "" {
		return 0, ErrorInvalidKey
	}
	i := a.Get(key)
	if i == nil {
		return 0, ErrorKeyDoesNotExist
	}
	if b, ok := i.(uint); ok {
		return b, nil
	}
	return 0, ErrorDataConversion
}

func (a *AbstractDocument) Uint8(key string) (uint8, error) {
	if key == "" {
		return 0, ErrorInvalidKey
	}
	i := a.Get(key)
	if i == nil {
		return 0, ErrorKeyDoesNotExist
	}
	if b, ok := i.(uint8); ok {
		return b, nil
	}
	return 0, ErrorDataConversion
}

func (a *AbstractDocument) Uint16(key string) (uint16, error) {
	if key == "" {
		return 0, ErrorInvalidKey
	}
	i := a.Get(key)
	if i == nil {
		return 0, ErrorKeyDoesNotExist
	}
	if b, ok := i.(uint16); ok {
		return b, nil
	}
	return 0, ErrorDataConversion
}

func (a *AbstractDocument) Uint32(key string) (uint32, error) {
	if key == "" {
		return 0, ErrorInvalidKey
	}
	i := a.Get(key)
	if i == nil {
		return 0, ErrorKeyDoesNotExist
	}
	if b, ok := i.(uint32); ok {
		return b, nil
	}
	return 0, ErrorDataConversion
}

func (a *AbstractDocument) Uint64(key string) (uint64, error) {
	if key == "" {
		return 0, ErrorInvalidKey
	}
	i := a.Get(key)
	if i == nil {
		return 0, ErrorKeyDoesNotExist
	}
	if b, ok := i.(uint64); ok {
		return b, nil
	}
	return 0, ErrorDataConversion
}

func (a *AbstractDocument) Float32(key string) (float32, error) {
	if key == "" {
		return 0, ErrorInvalidKey
	}
	i := a.Get(key)
	if i == nil {
		return 0, ErrorKeyDoesNotExist
	}
	if b, ok := i.(float32); ok {
		return b, nil
	}
	return 0, ErrorDataConversion
}

func (a *AbstractDocument) Float64(key string) (float64, error) {
	if key == "" {
		return 0, ErrorInvalidKey
	}
	i := a.Get(key)
	if i == nil {
		return 0, ErrorKeyDoesNotExist
	}
	if b, ok := i.(float64); ok {
		return b, nil
	}
	return 0, ErrorDataConversion
}

func (a *AbstractDocument) Complex64(key string) (complex64, error) {
	if key == "" {
		return 0, ErrorInvalidKey
	}
	i := a.Get(key)
	if i == nil {
		return 0, ErrorKeyDoesNotExist
	}
	if b, ok := i.(complex64); ok {
		return b, nil
	}
	return 0, ErrorDataConversion
}

func (a *AbstractDocument) Complex128(key string) (complex128, error) {
	if key == "" {
		return 0, ErrorInvalidKey
	}
	i := a.Get(key)
	if i == nil {
		return 0, ErrorKeyDoesNotExist
	}
	if b, ok := i.(complex128); ok {
		return b, nil
	}
	return 0, ErrorDataConversion
}

func (a *AbstractDocument) Json() ([]byte, error) {
	b, err := json.Marshal(a.data)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (a *AbstractDocument) AsMap() map[string]interface{} {
	return a.data
}


