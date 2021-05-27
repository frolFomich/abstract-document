package document

type ConstructorFunc func (map[string]interface{}) interface{}

// Document interface describes document behaviour
type Document interface {
	Get(key string) interface{}
	Put(key string, value interface{})
	Children(key string, constructor ConstructorFunc) []interface{}

	Slice(key string) []interface{}
	Map(key string) map[string]interface{}

	Bool(key string) (bool, error)
	String(key string) (string, error)
	Byte(key string) (byte, error)
	Rune(key string) (rune, error)

	Int(key string) (int, error)
	Int8(key string) (int8, error)
	Int16(key string) (int16, error)
	Int32(key string) (int32, error)
	Int64(key string) (int64, error)

	Uint(key string) (uint, error)
	Uint8(key string) (uint8, error)
	Uint16(key string) (uint16, error)
	Uint32(key string) (uint32, error)
	Uint64(key string) (uint64, error)

	Float32(key string) (float32, error)
	Float64(key string) (float64, error)

	Complex64(key string) (complex64, error)
	Complex128(key string) (complex128, error)

	Json() ([]byte, error)
	AsMap() map[string]interface{}
}