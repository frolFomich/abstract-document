package abstract_document

import (
	"encoding/json"
	"errors"
)

var (
	ErrorInvalidJsonBytes = errors.New("invalid input json or empty")
)

func Of(data map[string]interface{}) Document {
	if data == nil {
		return New()
	} else {
		return &AbstractDocument{
			data: data,
		}
	}
}

func FromOther(doc Document) Document {
	if doc == nil {
		return New()
	}
	src := doc.AsPlainMap()
	dst := make(map[string]interface{}, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return Of(dst)
}

func UnmarshalJson(bytes []byte) (Document, error) {
	if bytes == nil || len(bytes) <= 0 {
		return nil, ErrorInvalidJsonBytes
	}
	var m map[string]interface{}
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		return nil, err
	}
	return Of(m), nil
}

func New() *AbstractDocument {
	return &AbstractDocument{
		data: make(map[string]interface{}, 0),
	}
}
