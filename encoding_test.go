package suduku

import (
	"reflect"
	"testing"
)

func TestEncoding(t *testing.T) {
	want := Grid{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{},
		{},
		{},
		{4, 5, 6, 7, 8, 9, 1, 2, 3},
		{},
		{},
		{},
		{7, 8, 9, 1, 2, 3, 4, 5, 6},
	}
	tmp := Encode(want)
	got, err := Decode(tmp)
	if err != nil {
		t.Error(err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Log(got)
		t.Log(want)
		t.Fail()
	}
}
