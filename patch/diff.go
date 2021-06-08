package patch

import (
	"errors"
	doc "github.com/frolFomich/abstract-document"
	"strconv"
	"strings"
)

type PathFilterFunc func(keyPath string) bool

const (
	PathDelimiter = "/"

	MaxAllowedStackTraceLength = 512
)

var (
	ErrorTooManyNestedDocumentLevels = errors.New("document contains too many nested level documents")
)

func defaultPathFilterFunc(path string) bool {
	// use all path paths in diff
	return false
}

func Diff(src, tgt doc.Document, keysFilter ...PathFilterFunc) []DocumentPatch {

	if src == nil || src.Size() <= 0 {
		// TODO add logging
		return nil
	}
	if tgt == nil || tgt.Size() <= 0 {
		// TODO add logging
		return nil
	}

	isKeyFilteredOutFunc := defaultPathFilterFunc
	if keysFilter != nil && len(keysFilter) > 0 {
		isKeyFilteredOutFunc = keysFilter[0]
	}
	tmp := make([]string, 0)
	currentStackTraceCnt := MaxAllowedStackTraceLength
	return compare(src.AsPlainMap(), tgt.AsPlainMap(), &tmp, isKeyFilteredOutFunc, currentStackTraceCnt)
}

func compare(src map[string]interface{}, t interface{}, pathStack *[]string, isKeyFilteredOut PathFilterFunc, currentStackTraceCnt int) []DocumentPatch {
	res := make([]DocumentPatch, 0)
	tgt, ok := t.(map[string]interface{})
	if !ok {
		// replace op
		key := (*pathStack)[len(*pathStack)-1]
		path := (*pathStack)[len(*pathStack)-1:]
		res = append(res, New(ReplaceOperation, restorePath(key, path), src))
		return res
	}
	for k, v := range src {
		if isKeyFilteredOut(restorePath(k, *pathStack)) {
			continue
		}
		if _, exists := tgt[k]; exists {
			if v == nil {
				// remove op
				res = append(res, New(RemoveOperation, restorePath(k, *pathStack), v))
			} else {
				if s, ok := v.([]interface{}); ok {
					if currentStackTraceCnt <= 0 {
						panic(ErrorTooManyNestedDocumentLevels)
					} else {
						currentStackTraceCnt -= 1
					}
					res = append(res, compareSlices(s, tgt[k], pathStack, isKeyFilteredOut, currentStackTraceCnt)...)
					currentStackTraceCnt += 1
					continue
				} else if sm, ok := v.(map[string]interface{}); ok {
					if tm, yes := tgt[k].(map[string]interface{}); yes {
						if currentStackTraceCnt <= 0 {
							panic(ErrorTooManyNestedDocumentLevels)
						} else {
							currentStackTraceCnt -= 1
						}
						*pathStack = append(*pathStack, k)
						res = append(res, compare(sm, tm, pathStack, isKeyFilteredOut, currentStackTraceCnt)...)
						currentStackTraceCnt += 1
						*pathStack = (*pathStack)[:len(*pathStack)-1]
						continue
					}
				}
				// replace op
				res = append(res, New(ReplaceOperation, restorePath(k, *pathStack), v))
			}
		} else {
			// add op
			res = append(res, New(AddOperation, restorePath(k, *pathStack), v))
		}
	}
	return res
}

func compareSlices(ss []interface{}, t interface{}, pathStack *[]string, isKeyFilteredOut PathFilterFunc, currentStackTraceCnt int) []DocumentPatch {
	res := make([]DocumentPatch, 0)
	if ss == nil || t == nil {
		return res
	}
	if ts, ok := t.([]interface{}); ok {
		if len(ss) == len(ts) {
			for i, v := range ss {
				key := strconv.Itoa(i)
				if isKeyFilteredOut(restorePath(key, *pathStack)) {
					continue
				}
				if m, ok := v.(map[string]interface{}); ok {
					if currentStackTraceCnt <= 0 {
						panic(ErrorTooManyNestedDocumentLevels)
					} else {
						currentStackTraceCnt -= 1
					}
					*pathStack = append(*pathStack, key)
					res = append(res, compare(m, ts[i], pathStack, isKeyFilteredOut, currentStackTraceCnt)...)
					currentStackTraceCnt += 1
				} else {
					break
				}
			}
		}
	}
	// replace op
	key := (*pathStack)[len(*pathStack)-1]
	path := (*pathStack)[len(*pathStack)-1:]
	res = append(res, New(ReplaceOperation, restorePath(key, path), ss))
	return res
}

func restorePath(key string, stack []string) string {
	resBuilder := strings.Builder{}
	resBuilder.WriteString(PathDelimiter)
	for _, p := range stack {
		resBuilder.WriteString(p)
		resBuilder.WriteString(PathDelimiter)
	}
	resBuilder.WriteString(key)
	return resBuilder.String()
}
