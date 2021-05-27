package document

import (
	"reflect"
	"testing"
)

func TestAbstractDocument_Bool(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{name: "Bool return correct value", fields: fields{ data: map[string]interface{} {"A": true}}, args: args{key: "A"}, want: true, wantErr: false},
		{name: "Bool return conversion error", fields: fields{ data: map[string]interface{} {"A": 100}}, args: args{key: "A"}, want: false, wantErr: true},
		{name: "Bool return key not exist error", fields: fields{ data: map[string]interface{} {"B": true}}, args: args{key: "A"}, want: false, wantErr: true},
		{name: "Bool return invalid key error", fields: fields{ data: map[string]interface{} {"A": true}}, args: args{key: ""}, want: false, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			got, err := a.Bool(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Bool() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbstractDocument_Byte(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    byte
		wantErr bool
	}{
		{name: "Byte return correct value", fields: fields{ data: map[string]interface{} {"A": byte(1)}}, args: args{key: "A"}, want: byte(1), wantErr: false},
		{name: "Byte return conversion error", fields: fields{ data: map[string]interface{} {"A": "ABBA"}}, args: args{key: "A"}, want: 0, wantErr: true},
		{name: "Byte return key not exist error", fields: fields{ data: map[string]interface{} {"B": 1}}, args: args{key: "A"}, want: 0, wantErr: true},
		{name: "Byte return invalid key error", fields: fields{ data: map[string]interface{} {"A": 1}}, args: args{key: ""}, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			got, err := a.Byte(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Byte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Byte() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbstractDocument_Children(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		key         string
		constructor ConstructorFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			if got := a.Children(tt.args.key, tt.args.constructor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Children() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbstractDocument_Complex128(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    complex128
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			got, err := a.Complex128(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Complex128() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Complex128() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbstractDocument_Complex64(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    complex64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			got, err := a.Complex64(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Complex64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Complex64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbstractDocument_Float32(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float32
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			got, err := a.Float32(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Float32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Float32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbstractDocument_Float64(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			got, err := a.Float64(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Float64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Float64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbstractDocument_Get(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			if got := a.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbstractDocument_Int(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			got, err := a.Int(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Int() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbstractDocument_Int16(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int16
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			got, err := a.Int16(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Int16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int16() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbstractDocument_Int32(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int32
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			got, err := a.Int32(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Int32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbstractDocument_Int64(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			got, err := a.Int64(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Int64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbstractDocument_Int8(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int8
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			got, err := a.Int8(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Int8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int8() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbstractDocument_Json(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			got, err := a.Json()
			if (err != nil) != tt.wantErr {
				t.Errorf("Json() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Json() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbstractDocument_Map(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			if got := a.Map(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbstractDocument_Put(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//a := &AbstractDocument{
			//	data: tt.fields.data,
			//}
		})
	}
}

func TestAbstractDocument_Rune(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    rune
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			got, err := a.Rune(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rune() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Rune() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbstractDocument_Slice(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			if got := a.Slice(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbstractDocument_String(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			got, err := a.String(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("String() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("String() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbstractDocument_Uint(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			got, err := a.Uint(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbstractDocument_Uint16(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint16
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			got, err := a.Uint16(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint16() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbstractDocument_Uint32(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint32
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			got, err := a.Uint32(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbstractDocument_Uint64(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			got, err := a.Uint64(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbstractDocument_Uint8(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint8
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			got, err := a.Uint8(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint8() got = %v, want %v", got, tt.want)
			}
		})
	}
}
