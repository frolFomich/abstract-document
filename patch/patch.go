package patch

import (
	doc "github.com/frolFomich/abstract-document"
)

type OperationType int

const (
	UnknownOperation OperationType = iota
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
	 ops := [...]string{AddOperationString, ReplaceOperationString, RemoveOperationString}
	 if len(ops) < int(op) {
	 	return ""
	 }
	 return ops[op]
}

func ValueOf(opStr string) OperationType {
	operationType := UnknownOperation
	switch opStr {
	case AddOperationString:
		operationType = AddOperation
		break
	case ReplaceOperationString:
		operationType = ReplaceOperation
		break
	case RemoveOperationString:
		operationType = RemoveOperation
		break
	default:
		break
	}
	return operationType
}