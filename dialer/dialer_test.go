package dialer

import (
	"reflect"
	"testing"
)

func TestDialTCP(t *testing.T) {
	got := DialTCP("google.com", 80, 81)
	expected := []int{80}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v, expected %v", got, expected)
	}
}
