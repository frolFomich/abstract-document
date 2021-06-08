package patch

import (
	"errors"
	doc "github.com/frolFomich/abstract-document"
	"strings"
)

var (
	ErrorInvalidDocument      = errors.New("invalid document")
	ErrorInvalidPath          = errors.New("invalid patch path value")
	ErrorPathNotFound         = errors.New("path not found in document")
	ErrorUnsupportedOperation = errors.New("unsupported operation")
)

func Patch(src doc.Document, patches ...DocumentPatch) (doc.Document, error) {
	if src == nil {
		return nil, ErrorInvalidDocument
	}
	dst := doc.FromOther(src)
	for _, p := range patches {
		m := dst.AsPlainMap()
		pathElements := expandPath(p.Path())
		if pathElements == nil || len(pathElements) <= 0 {
			return nil, ErrorInvalidPath
		} else if len(pathElements) > MaxAllowedStackTraceLength {
			return nil, ErrorTooManyNestedDocumentLevels
		}
		err := applyPatch(p.Operation(), pathElements, p.Value(), m)
		if err != nil {
			return nil, err
		}
	}
	return dst, nil
}

func applyPatch(op OperationType, path []string, value interface{}, doc map[string]interface{}) error {
	if doc == nil || len(doc) <= 0 {
		return ErrorInvalidDocument
	}
	if path == nil || len(path) <= 0 {
		return ErrorInvalidPath
	}
	targetMap := doc
	targetKey := path[len(path)-1]
	if len(path) >= 2 {
		for _, e := range path[0 : len(path)-1] {
			v, found := targetMap[e]
			if found {
				if m, ok := v.(map[string]interface{}); ok {
					targetMap = m
				} else {
					return ErrorPathNotFound
				}
			}
		}
	}
	switch op {
	case AddOperation, ReplaceOperation:
		targetMap[targetKey] = value
	case RemoveOperation:
		delete(targetMap, targetKey)
	case UnknownOperation:
		return ErrorUnsupportedOperation
	default:
		return ErrorUnsupportedOperation
	}
	return nil
}

func expandPath(path string) []string {
	s := strings.TrimSpace(path)
	if len(path) <= 0 {
		return nil
	}
	if strings.HasPrefix(s, PathDelimiter) {
		s = strings.TrimPrefix(s, PathDelimiter)
	}
	if strings.HasSuffix(path, PathDelimiter) {
		s = strings.TrimSuffix(s, PathDelimiter)
	}
	return strings.Split(s, PathDelimiter)
}
