package dialer

import (
	"reflect"
	"testing"
)

func TestDialTCP(t *testing.T) {
	got := DialTCP("google.com", 80, 81, false)
	expected := []int{80}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v, expected %v", got, expected)
	}
}

func TestDialUDP(t *testing.T) {
	got := DialUDP("localhost", 8080, 8081, false)
	expected := []int{8080}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v, expected %v", got, expected)
	}
}
