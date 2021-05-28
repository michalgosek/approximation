package reading

import (
	"reflect"
	"testing"
)

func Test_FileToMap(t *testing.T) {
	got, err := FileToMap("test_data.csv")
	if err != nil {
		t.Fatalf("expected to get nil err; got: %v", err)
	}

	want := map[int][]float64{
		1: {1, 6, 12, -2, 5, -3},
		2: {2, 3, 122, -2, 5, -32},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected to get: %v; got: %v", got, want)
	}
}
