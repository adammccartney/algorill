package list

import (
	"reflect"
	"testing"
)

func TestList_Insert(t *testing.T) {
	type fields struct {
		value string
		next  *List
	}
	type args struct {
		node *List
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *List
	}{
		// TODO: Add test cases.

		{name: "test insert on nil",
			fields: fields{"", nil},
			args:   args{&List{"SICP", nil}},
			want:   &List{"SICP", nil}},
		{name: "test insert after one element",
			fields: fields{"SICP", nil},
			args:   args{&List{"TAOCP", nil}}, want: &List{"SICP", &List{"TAOCP", nil}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var list *List
			if tt.fields.value == "" {
				list = (*List)(nil)
			} else {
				list = &List{
					value: tt.fields.value,
					next:  tt.fields.next,
				}
			}
			if got := list.Insert(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Reverse(t *testing.T) {
	type fields struct {
		value string
		next  *List
	}
	tests := []struct {
		name   string
		fields fields
		want   *List
	}{
		// TODO: Add test cases.
		{name: "reverse simple list recursive",
			fields: fields{"A", &List{"B", &List{"C", nil}}},
			want:   &List{"C", &List{"B", &List{"A", nil}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &List{
				value: tt.fields.value,
				next:  tt.fields.next,
			}
			if got := list.Reverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
