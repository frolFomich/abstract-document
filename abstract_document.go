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
	ErrorInvalidKey = errors.New("invalid key provided")
	ErrorDataConversion = errors.New("invalid data conversion")

	ErrorUnsupportedGoType = errors.New("unsupported go data type")
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

	a.data[key] = asGoType(value)
}

func (a *AbstractDocument) Children(key string, constructor ConstructorFunc) []Document {
	empty := make([]Document, 0)
	if key == "" || constructor == nil {
		return empty
	}
	list := a.Array(key)
	result := make([]Document, len(list))
	for i, val := range list {
		m, ok := val.(map[string]interface{})
		if !ok {
			return empty
		}
		result[i] = constructor(m)
	}
	return result
}

func (a *AbstractDocument) Array(key string) Array {
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

func (a *AbstractDocument) Document(key string) Document {
	empty := &AbstractDocument{}
	if key == "" {
		return empty
	}
	i := a.Get(key)
	if i == nil {
		return empty
	}
	if b, ok := i.(map[string]interface{}); ok {
		return Of(b)
	}
	return empty
}

func (a *AbstractDocument) Boolean(key string) (bool, error) {
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

func (a *AbstractDocument) Integer(key string) (int64, error) {
	if key == "" {
		return 0, ErrorInvalidKey
	}
	i := a.Get(key)
	if i == nil {
		return 0, ErrorKeyDoesNotExist
	}
	if b, ok := i.(float64); ok {
		if _, r := math.Modf(b); r != 0 {
			return 0, ErrorDataConversion
		}
		return int64(b), nil
	}
	return 0, ErrorDataConversion
}

func (a *AbstractDocument) Number(key string) (float64, error) {
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
	if key == "" {
		return false, ErrorInvalidKey
	}
	if _, ok := a.data[key]; ok {
		delete(a.data, key)
		return true,nil
	}
	return false, ErrorKeyDoesNotExist
}

func (a *AbstractDocument) IsNull(key string) (bool, error) {
	if key == "" {
		return false, ErrorInvalidKey
	}
	if v,ok := a.data[key]; ok {
		return v == nil, nil
	} else {
		return false, ErrorKeyDoesNotExist
	}
}

func (a *AbstractDocument) IsExist(key string) bool {
	if key == "" {
		return false
	}
	_,ok := a.data[key]
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

func asGoType(value interface{}) interface{} {
	if value == nil {
		return nil
	} else if d, td := value.(Document); td {
		return d.AsPlainMap()
	} else if a, ta := value.(Array); ta {
		is := make([]interface{}, len(a))
		for i,v := range a {
			is[i] = asGoType(v)
		}
		return is
	} else if s, ts := value.([]Document); ts {
		is := make([]interface{}, len(s))
		for i,v := range s {
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


