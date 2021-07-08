package abstract_document

type ConstructorFunc func(map[string]interface{}) Document

type Array []interface{}

// Document interface describes document behaviour
type Document interface {
	Get(key string) interface{}
	Put(key string, value interface{})
	Remove(key string) (bool, error)

	Array(key string) Array
	Children(key string) []Document
	Document(key string) Document

	AsPlainMap() map[string]interface{}

	Boolean(key string) (bool, error)
	String(key string) (string, error)
	Integer(key string) (int64, error)
	Number(key string) (float64, error)

	IsNull(key string) bool
	IsExist(key string) bool
	IsArray(key string) bool
	IsDocument(key string) bool
	IsBoolean(key string) bool
	IsString(key string) bool
	IsInteger(key string) bool
	IsNumber(key string) bool

	Size() int
	Keys() []string

	JsonSerializable
}

type JsonSerializable interface {
	MarshalJson() ([]byte, error)
}
