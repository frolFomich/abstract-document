package patch

import (
	doc "github.com/frolFomich/abstract-document"
)

func New(op OperationType, path string, val interface{}) DocumentPatch {
	dp := &documentPatchImpl{
		doc.New(),
	}
	dp.Put(OperationTypeKey, op.String())
	dp.Put(PathKey, path)
	if RemoveOperation != op {
		dp.Put(ValueKey, val)
	}

	return dp
}
