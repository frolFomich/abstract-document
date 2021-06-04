package patch

import (
	doc "github.com/frolFomich/abstract-document"
)

type documentPatchImpl struct {
	doc.Document
}

func (d *documentPatchImpl) Operation() OperationType {
	res, err := d.String(OperationTypeKey)
	if err != nil {
		panic(err)
	}
	return ValueOf(res)
}

func (d *documentPatchImpl) Path() string {
	p, err := d.String(PathKey)
	if err != nil {
		panic(err)
	}
	return p
}

func (d *documentPatchImpl) Value() interface{} {
	return d.Get(ValueKey)
}



