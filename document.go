package abstract_document

type ConstructorFunc func (map[string]interface{}) Document

type Array []interface{}

// Document interface describes document behaviour
type Document interface {
	Get(key string) interface{}
	Put(key string, value interface{})
	Children(key string, constructor ConstructorFunc) []Document
	Remove(key string) (bool, error)

	Array(key string) Array
	Document(key string) Document

	AsPlainMap() map[string]interface{}

	Boolean(key string) (bool, error)
	String(key string) (string, error)
	Integer(key string) (int64, error)
	Number(key string) (float64, error)

	IsNull(key string) (bool, error)
	IsExist(key string) bool

	Size() int
	Keys() []string

	JsonSerializable
}

type JsonSerializable interface {
	MarshalJson() ([]byte, error)
}