package testing


import (
	"reflect"
	"runtime/debug"
	"testing"
)

func AssertTrue(t *testing.T, value bool) {
	if !value {
		t.Fatalf("Expected tobe true, but false\n%s", string(debug.Stack()))
	}
}

func AssertFalse(t *testing.T, value bool) {
	if value {
		t.Fatalf("Expected tobe false, but true\n%s", string(debug.Stack()))
	}
}

func AssertEquals(t *testing.T, expect, actual interface{}) {
	if expect != actual {
		t.Fatalf("Expected %v(%T) but %v(%T). \n%s", expect, expect, actual, actual, string(debug.Stack()))
	}
}

func AssertNotNil(t *testing.T, obj interface{}) {
	defer func() { recover() }()
	if obj == nil || reflect.ValueOf(obj).IsNil() {
		t.Fatalf("Expected not nil, but nil\n%s", string(debug.Stack()))
	}
}

func AssertNil(t *testing.T, obj interface{}) {
	defer func() { recover() }()
	if obj != nil || !reflect.ValueOf(obj).IsNil() {
		t.Fatalf("Expected nil, but not nil\n%s", string(debug.Stack()))
	}
}

func Fail(t *testing.T, message string) {
	t.Fatalf("%s\n%s", message, string(debug.Stack()))
}
