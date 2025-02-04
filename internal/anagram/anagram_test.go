package anagram

import (
	"reflect"
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
			if gotT := InitToken(tt.args.word); !reflect.DeepEqual(gotT, tt.wantT) {
				t.Errorf("initToken() = %v, want %v", gotT, tt.wantT)
			}
		})
	}
}
