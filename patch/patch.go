package patch

import (
	doc "github.com/frolFomich/abstract-document"
)

type OperationType int

const (
	UnknownOperation OperationType = iota + 1
	AddOperation
	ReplaceOperation
	RemoveOperation
)

type DocumentPatch interface {
	doc.Document
	Operation() OperationType
	Path() string
	Value() interface{}
}

func (op OperationType) String() string {
	ops := [...]string{UnknownOperationString, AddOperationString, ReplaceOperationString, RemoveOperationString}
	if len(ops) < int(op) {
		return UnknownOperationString
	}
	return ops[op-1]
}

func ValueOf(opStr string) OperationType {
	operationType := UnknownOperation
	switch opStr {
	case AddOperationString:
		operationType = AddOperation
	case ReplaceOperationString:
		operationType = ReplaceOperation
	case RemoveOperationString:
		operationType = RemoveOperation
	default:
		break
	}
	return operationType
}
