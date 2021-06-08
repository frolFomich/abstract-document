package abstract_document

import (
	"encoding/json"
	"errors"
	"math"
)

type AbstractDocument struct {
	data map[string]interface{}
}

var (
	ErrorKeyDoesNotExist = errors.New("key doesn't exists")
	ErrorInvalidKey      = errors.New("invalid key provided")
	ErrorDataConversion  = errors.New("invalid data conversion")

	ErrorUnsupportedGoType = errors.New("unsupported go data type")
)

func (a *AbstractDocument) Get(key string) interface{} {
	if a.IsNull(key) {
		return nil
	}
	return a.data[key]
}

func (a *AbstractDocument) Put(key string, value interface{}) {
	if key == "" {
		return
	}
	a.data[key] = asGoType(value)
}

func (a *AbstractDocument) Children(key string, constructor ConstructorFunc) []Document {
	if a.IsNull(key) {
		return nil
	}
	if constructor == nil {
		panic(errors.New("error: required constructor is absent"))
	}
	list := a.Array(key)
	result := make([]Document, len(list))
	for i, val := range list {
		m, ok := val.(map[string]interface{})
		if !ok {
			return nil
		}
		result[i] = constructor(m)
	}
	return result
}

func (a *AbstractDocument) Array(key string) Array {
	if a.IsNull(key) || !a.IsArray(key) {
		return nil
	}
	return a.Get(key).([]interface{})
}

func (a *AbstractDocument) Document(key string) Document {
	if key == "" || !a.IsDocument(key) {
		return nil
	}
	return Of(a.Get(key).(map[string]interface{}))
}

func (a *AbstractDocument) Boolean(key string) (bool, error) {
	if a.IsNull(key) {
		return false, ErrorInvalidKey
	}
	if !a.IsBoolean(key) {
		return false, ErrorDataConversion
	}
	return a.Get(key).(bool), nil
}

func (a *AbstractDocument) String(key string) (string, error) {
	if a.IsNull(key) {
		return "", ErrorInvalidKey
	}
	if !a.IsString(key) {
		return "", ErrorDataConversion
	}
	return a.Get(key).(string), nil
}

func (a *AbstractDocument) Integer(key string) (int64, error) {
	if a.IsNull(key) {
		return 0, ErrorInvalidKey
	}
	if !a.IsInteger(key) {
		return 0, ErrorDataConversion
	}
	return int64(a.Get(key).(float64)), nil
}

func (a *AbstractDocument) Number(key string) (float64, error) {
	if a.IsNull(key) {
		return 0, ErrorInvalidKey
	}
	if !a.IsNumber(key) {
		return 0, ErrorDataConversion
	}
	return a.Get(key).(float64), nil
}

func (a *AbstractDocument) MarshalJson() ([]byte, error) {
	b, err := json.Marshal(a.data)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (a *AbstractDocument) AsPlainMap() map[string]interface{} {
	return a.data
}

func (a *AbstractDocument) Remove(key string) (bool, error) {
	if !a.IsExist(key) {
		return false, ErrorKeyDoesNotExist
	}
	delete(a.data, key)
	return true, nil
}

func (a *AbstractDocument) IsNull(key string) bool {
	if !a.IsExist(key) {
		return true
	}
	return a.data[key] == nil
}

func (a *AbstractDocument) IsExist(key string) bool {
	if key == "" || a.Size() <= 0 {
		return false
	}
	_, ok := a.data[key]
	return ok
}

func (a *AbstractDocument) Size() int {
	return len(a.data)
}

func (a *AbstractDocument) Keys() []string {
	res := make([]string, a.Size())
	i := 0
	for k := range a.data {
		res[i] = k
		i++
	}
	return res
}

func (a *AbstractDocument) IsArray(key string) bool {
	if a.IsNull(key) {
		return false
	}
	_, isArray := a.Get(key).([]interface{})
	return isArray
}

func (a *AbstractDocument) IsDocument(key string) bool {
	if a.IsNull(key) {
		return false
	}
	_, isDoc := a.Get(key).(map[string]interface{})
	return isDoc
}

func (a *AbstractDocument) IsBoolean(key string) bool {
	if a.IsNull(key) {
		return false
	}
	_, isBool := a.Get(key).(bool)
	return isBool
}

func (a *AbstractDocument) IsString(key string) bool {
	if a.IsNull(key) {
		return false
	}
	_, isStr := a.Get(key).(string)
	return isStr
}

func (a *AbstractDocument) IsInteger(key string) bool {
	if !a.IsNumber(key) {
		return false
	}
	num, err := a.Number(key)
	if err != nil {
		return false
	}
	_, r := math.Modf(num)
	return r == 0
}

func (a *AbstractDocument) IsNumber(key string) bool {
	if a.IsNull(key) {
		return false
	}
	_, isNum := a.Get(key).(float64)
	return isNum
}

func asGoType(value interface{}) interface{} {
	if value == nil {
		return nil
	} else if d, td := value.(Document); td {
		return d.AsPlainMap()
	} else if a, ta := value.(Array); ta {
		is := make([]interface{}, len(a))
		for i, v := range a {
			is[i] = asGoType(v)
		}
		return is
	} else if s, ts := value.([]Document); ts {
		is := make([]interface{}, len(s))
		for i, v := range s {
			is[i] = asGoType(v)
		}
		return is
	} else if v, tv := value.(uint8); tv {
		return float64(v)
	} else if v, tv := value.(uint16); tv {
		return float64(v)
	} else if v, tv := value.(uint32); tv {
		return float64(v)
	} else if v, tv := value.(uint64); tv {
		return float64(v)
	} else if v, tv := value.(int8); tv {
		return float64(v)
	} else if v, tv := value.(int16); tv {
		return float64(v)
	} else if v, tv := value.(int32); tv {
		return float64(v)
	} else if v, tv := value.(int64); tv {
		return float64(v)
	} else if v, tv := value.(float32); tv {
		return float64(v)
	} else if v, tv := value.(int); tv {
		return float64(v)
	} else if v, tv := value.(uint); tv {
		return float64(v)
	} else if v, tv := value.(byte); tv {
		return float64(v)
	} else if v, tv := value.(rune); tv {
		return string(v)
	} else if _, tv := value.(complex64); tv {
		panic(ErrorUnsupportedGoType)
	} else if _, tv := value.(complex128); tv {
		panic(ErrorUnsupportedGoType)
	}

	return value
}
