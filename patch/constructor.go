package patch

import (
	"errors"
	doc "github.com/frolFomich/abstract-document"
)

func New(op, path string, val interface{}) DocumentPatch {
	if "" == op || "" == path {
		panic(errors.New("illegal arguments"))
	}
	dp := &documentPatchImpl{
		Document: doc.New(),
	}
	dp.Put(OperationTypeKey, op)
	dp.Put(PathKey, path)
	dp.Put(ValueKey, val)

	return dp
}
