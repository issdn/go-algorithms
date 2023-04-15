package main

import (
	"reflect"
	"testing"
)

func TestFixedWindowMakeRequest(t *testing.T) {
	window := FixedWindow{
		Interval:     1,
		UserRequests: map[string]uint16{},
	}
	gotToken, gotRequests := window.Request("")
	if reflect.TypeOf(gotToken).Kind() != reflect.String {
		t.Errorf("Token type: Got: %v, Want: %v", reflect.TypeOf(gotToken), "String")
	}

	if gotRequests != 1 {
		t.Errorf("Nr. Request: Got: %v, Want: 1", gotRequests)
	}

}

func TestFixedWindowMakeMultipleRequests(t *testing.T) {
	window := FixedWindow{
		Interval:     1,
		UserRequests: map[string]uint16{},
	}
	gotToken, _ := window.Request("")
	_, gotRequests2 := window.Request(gotToken)
	if reflect.TypeOf(gotToken).Kind() != reflect.String {
		t.Errorf("Token type: Got: %v, Want: %v", reflect.TypeOf(gotToken), "String")
	}

	if gotRequests2 != 2 {
		t.Errorf("Nr. Request: Got: %v, Want: 1", gotRequests2)
	}
}
