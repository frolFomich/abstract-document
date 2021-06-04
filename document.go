package abstract_document

type ConstructorFunc func (map[string]interface{}) Document

type Array []interface{}
type Boolean bool
type String string
type Integer int64
type Number float64
type Null struct{}

// Document interface describes document behaviour
type Document interface {
	Get(key string) interface{}
	Put(key string, value interface{})
	Children(key string, constructor ConstructorFunc) []Document
	Remove(key string) (bool, error)

	Array(key string) Array
	Document(key string) Document

	AsPlainMap() map[string]interface{}

	Boolean(key string) (Boolean, error)
	String(key string) (String, error)
	Integer(key string) (Integer, error)
	Number(key string) (Number, error)

	IsNull(key string) (bool, error)
	IsExist(key string) bool

	Size() int
	Keys() []string

	JsonSerializable
}

type JsonSerializable interface {
	MarshalJson() ([]byte, error)
}