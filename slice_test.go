package utils

import (
	"reflect"
	"testing"
)

func TestPrepend(t *testing.T) {
	intSlice := []int{1, 2, 3, 4}
	intSlice = Prepend(intSlice, 5)
	if !reflect.DeepEqual(intSlice, []int{5, 1, 2, 3, 4}) {
		t.Errorf("Expected %v, got %v", []int{5, 1, 2, 3, 4}, intSlice)
	}

	stringSlice := []string{"a", "b", "c"}
	stringSlice = Prepend(stringSlice, "d")
	if !reflect.DeepEqual(stringSlice, []string{"d", "a", "b", "c"}) {
		t.Errorf("Expected %v, got %v", []string{"d", "a", "b", "c"}, stringSlice)
	}
}
