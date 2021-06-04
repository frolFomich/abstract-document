package patch

import doc "github.com/frolFomich/abstract-document"

type DiffKeyFilterFunc func (key string) bool


func Diff(src,tgt doc.Document, filter DiffKeyFilterFunc) []DocumentPatch {
	panic("implement me")
}

func Patch(patches []DocumentPatch, tgt doc.Document) doc.Document {
	panic("implement me")
}
