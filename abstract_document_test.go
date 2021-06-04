package abstract_document

import (
	"reflect"
	"testing"
)

func TestAbstractDocument_Boolean(t *testing.T) {
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
		want    Boolean
		wantErr bool
	}{
		{name: "Boolean return correct value", fields: fields{ data: map[string]interface{} {"A": true}}, args: args{key: "A"}, want: true, wantErr: false},
		{name: "Boolean return conversion error", fields: fields{ data: map[string]interface{} {"A": 100}}, args: args{key: "A"}, want: false, wantErr: true},
		{name: "Boolean return key not exist error", fields: fields{ data: map[string]interface{} {"B": true}}, args: args{key: "A"}, want: false, wantErr: true},
		{name: "Boolean return invalid key error", fields: fields{ data: map[string]interface{} {"A": true}}, args: args{key: ""}, want: false, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			got, err := a.Boolean(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Boolean() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Boolean() got = %v, want %v", got, tt.want)
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
		want    String
		wantErr bool
	}{
		{name: "String return correct value", fields: fields{ data: map[string]interface{} {"A": "B"}}, args: args{key: "A"}, want: "B", wantErr: false},
		{name: "String return conversion error", fields: fields{ data: map[string]interface{} {"A": 123}}, args: args{key: "A"}, want: "", wantErr: true},
		{name: "String return key not exist error", fields: fields{ data: map[string]interface{} {"B": 1}}, args: args{key: "A"}, want: "", wantErr: true},
		{name: "String return invalid key error", fields: fields{ data: map[string]interface{} {"A": 1}}, args: args{key: ""}, want: "", wantErr: true},
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

func TestAbstractDocument_Children(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		key         string
		constructor ConstructorFunc
	}
	d1 := Of(map[string]interface{}{"A": 1})
	d2 := Of(map[string]interface{}{"B": 2})
	d3 := Of(map[string]interface{}{"C": 3})

	givenChildren := []interface{}{
		d1.AsPlainMap(),
		d2.AsPlainMap(),
		d3.AsPlainMap(),
	}

	givenMap := map[string]interface{}{
		"Docs": givenChildren,
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Document
		wantErr bool
	}{
		{name: "Children return correct value", fields: fields{data: givenMap}, args: args{key: "Docs", constructor: func(m map[string]interface{}) Document {
			return Of(m)
		}}, want: []Document{d1,d2,d3}, wantErr: false},
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

func TestAbstractDocument_Number(t *testing.T) {
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
		want    Number
		wantErr bool
	}{
		{name: "Number return correct value for float", fields: fields{ data: map[string]interface{} {"A": 10.25}}, args: args{key: "A"}, want: Number(10.25), wantErr: false},
		{name: "Number return correct value for int", fields: fields{ data: map[string]interface{} {"A": float64(10)}}, args: args{key: "A"}, want: Number(10), wantErr: false},
		{name: "Number return conversion error", fields: fields{ data: map[string]interface{} {"A": 123}}, args: args{key: "A"}, want: 0, wantErr: true},
		{name: "Number return key not exist error", fields: fields{ data: map[string]interface{} {"B": 1}}, args: args{key: "A"}, want: 0, wantErr: true},
		{name: "Number return invalid key error", fields: fields{ data: map[string]interface{} {"A": 1}}, args: args{key: ""}, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			got, err := a.Number(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Number() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Number() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbstractDocument_Integer(t *testing.T) {
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
		want    Integer
		wantErr bool
	}{
		{name: "Number return correct value for float", fields: fields{ data: map[string]interface{} {"A": float64(10)}}, args: args{key: "A"}, want: Integer(10), wantErr: false},
		{name: "Number return conversion error", fields: fields{ data: map[string]interface{} {"A": 123.25}}, args: args{key: "A"}, want: 0, wantErr: true},
		{name: "Number return key not exist error", fields: fields{ data: map[string]interface{} {"B": 1}}, args: args{key: "A"}, want: 0, wantErr: true},
		{name: "Number return invalid key error", fields: fields{ data: map[string]interface{} {"A": 1}}, args: args{key: ""}, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			got, err := a.Integer(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Integer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Integer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbstractDocument_Document(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		key string
	}

	nestedDoc := map[string]interface{}{"A": "B", "C" : 100, "D": true}
	m := map[string]interface{}{"Doc": nestedDoc, "Q": "qqq"}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   Document
	}{
		{name: "Document returns correct value", fields: fields{data: m}, args: args{key: "Doc"}, want: Of(nestedDoc)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			if got := a.Document(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Document() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestAbstractDocument_IsNull(t *testing.T) {
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
		{name: "IsNull returns correct value", fields: fields{data: map[string]interface{}{"A": nil}}, args: args{key: "A"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			got, err := a.IsNull(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsNull() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsNull() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbstractDocument_MarshalJson(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}

	givenJson := []byte("{\"A\":\"B\",\"C\":100,\"D\":true,\"F\":null}")

	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{name: "MarshalJson returns correct value", fields: fields{data: map[string]interface{}{"A": "B", "C": 100, "D": true, "F": nil}}, want: givenJson},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			got, err := a.MarshalJson()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJson() got = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestAbstractDocument_IsExist(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{name: "IsExist returns true for existing key", fields: fields{data: map[string]interface{}{"A": "B", "C": 100}}, args: args{key:"A"}, want: true, wantErr: false},
		{name: "IsExist returns false for non-existing key", fields: fields{data: map[string]interface{}{"C": 100}}, args: args{key:"A"}, want: false, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			if got := a.IsExist(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbstractDocument_Put(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}

	givenGoTypesMap := map[string]interface{}{"A": "B", "C": 100.25, "D": false, "F": nil}
	givenDocTypesMap := map[string]interface{}{"A": String("B"), "C": Number(100.25), "D": Boolean(false), "F": Document(nil)}

	tests := []struct {
		name    string
		fields  fields
		want    map[string]interface{}
		wantErr bool
	}{
		{name: "Put returns correct values for Go types", fields: fields{data: givenGoTypesMap}, want: givenGoTypesMap, wantErr: false},
		{name: "Put returns correct values for Document types", fields: fields{data: givenDocTypesMap}, want: givenGoTypesMap, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: map[string]interface{}{},
			}
			for k,v := range tt.fields.data {
				a.Put(k,v)
			}
			if got := a.AsPlainMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Put() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestAbstractDocument_Array(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		key string
	}

	givenArray := []interface{}{1,2,3,4,5}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   Array
	}{
		{name: "Array returns correct value", fields: fields{data: map[string]interface{}{"A": givenArray}}, args: args{key: "A"}, want: Array(givenArray)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			if got := a.Array(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Array() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbstractDocument_Remove(t *testing.T) {
	type fields struct {
		data map[string]interface{}
	}
	type args struct {
		key string
	}

	givenMap := map[string]interface{}{"A":"B", "C":"D"}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		wantErr bool
	}{
		{name: "Remove returns correct value", fields: fields{data: givenMap}, args: args{key: "A"}, want: true, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractDocument{
				data: tt.fields.data,
			}
			removed, err := a.Remove(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got := removed; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}