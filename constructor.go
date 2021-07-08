package abstract_document

import (
	"encoding/json"
	"errors"
)

var (
	ErrorInvalidJsonBytes = errors.New("invalid input json or empty")
)

func Of(data map[string]interface{}) *AbstractDocument {
	if data == nil {
		return New()
	} else {
		return &AbstractDocument{
			data: data,
		}
	}
}

func FromOther(doc Document) *AbstractDocument {
	if doc == nil {
		return New()
	}
	src := doc.AsPlainMap()
	dst := make(map[string]interface{}, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return &AbstractDocument{dst}
}

func UnmarshalJson(bytes []byte) (*AbstractDocument, error) {
	if bytes == nil || len(bytes) <= 0 {
		return nil, ErrorInvalidJsonBytes
	}
	var m map[string]interface{}
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		return nil, err
	}
	return &AbstractDocument{m}, nil
}

func New() *AbstractDocument {
	return &AbstractDocument{
		data: make(map[string]interface{}, 0),
	}
}
