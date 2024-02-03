package dialer

import (
	"reflect"
	"testing"
)

func TestDialer(t *testing.T) {
	got := Dial("google.com", 80, 81)
	expected := []int{80}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v, expected %v", got, expected)
	}
}
