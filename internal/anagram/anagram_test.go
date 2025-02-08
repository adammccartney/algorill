package anagram

import (
	"container/list"
	"io"
	"reflect"
    "strings"
	"testing"
)

func Test_InitToken(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name  string
		args  args
		wantT Token
	}{
        {
            name: "Hello test",
            args: args{word: "hello"},
            wantT: Token {
                Word: "hello",
                Key: "ehl2o",
            },
        },
		// TODO: Add more test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotT, _ := InitToken(tt.args.word); !reflect.DeepEqual(gotT, tt.wantT) {
				t.Errorf("initToken() = %v, want %v", gotT, tt.wantT)
			}
		})
	}
}


func TestAnagramDict_Load(t *testing.T) {
	type fields struct {
		Map map[string]*list.List
	}
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
        {
			name: "Single word",
			fields: fields{
				Map: make(map[string]*list.List),
			},
			args: args{
				reader: strings.NewReader("cat\n"),
			},
			wantErr: false,
		},
		{
			name: "Two anagrams",
			fields: fields{
				Map: make(map[string]*list.List),
			},
			args: args{
				reader: strings.NewReader("cat\ndog\nact\n"),
			},
			wantErr: false,
		},
		{
			name: "Multiple words with spaces",
			fields: fields{
				Map: make(map[string]*list.List),
			},
			args: args{
				reader: strings.NewReader("eat tea\ntan nat\nin nit\n"),
			},
			wantErr: true,
		},
		{
			name: "Empty input",
			fields: fields{
				Map: make(map[string]*list.List),
			},
			args: args{
				reader: strings.NewReader(""),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AnagramDict{
				Map: tt.fields.Map,
			}
			if err := a.Load(tt.args.reader); (err != nil) != tt.wantErr {
				t.Errorf("AnagramDict.Load() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
