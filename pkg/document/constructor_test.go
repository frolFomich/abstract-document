package document

import (
	"reflect"
	"testing"
)

func TestFromJson(t *testing.T) {
	type args struct {
		bytes []byte
	}

	givenJson := []byte(`{
		"A" : "ABBA",
		"B" : 100,
		"C" : true
	}`)

	givenDoc, err := FromJson(givenJson)
	if err != nil {
		t.Errorf("FromJson() error = %v", err)
		return
	}

	tests := []struct {
		name    string
		args    args
		want    Document
		wantErr bool
	}{
		{name: "create from json without error", args: args{givenJson}, want: givenDoc, wantErr: false},
		{name: "error while create from json", args: args{[]byte("qqqq")}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromJson(tt.args.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromJson() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromOther(t *testing.T) {
	type args struct {
		doc Document
	}

	givenDoc := Of(map[string]interface{}{
		"A" : "ABBA",
		"B" : 100,
		"C" : true,
	})

	tests := []struct {
		name string
		args args
		want Document
	}{
		{name: "create from other document", args: args{givenDoc}, want: givenDoc},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromOther(tt.args.doc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromOther() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want Document
	}{
		{name: "create new empty document", want: Of(map[string]interface{}{})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOf(t *testing.T) {
	type args struct {
		data map[string]interface{}
	}

	givenMap := map[string]interface{}{
		"A": "ABBA",
		"C": 100,
		"D": true,
	}
	givenDoc := Of(givenMap)

	tests := []struct {
		name string
		args args
		want Document
	}{
		{name: "create from map", args: args{data: givenMap}, want: givenDoc},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Of(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Of() = %v, want %v", got, tt.want)
			}
		})
	}
}
