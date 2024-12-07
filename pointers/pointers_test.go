package pointers_test

import (
	"testing"

	"github.com/cfindlayisme/go-utils/pointers"
)

func TestBoolPtr(t *testing.T) {
	value := true
	ptr := pointers.BoolPtr(value)
	if ptr == nil || *ptr != value {
		t.Errorf("BoolPtr failed: expected %v, got %v", value, ptr)
	}
}

func TestStringPtr(t *testing.T) {
	value := "test"
	ptr := pointers.StringPtr(value)
	if ptr == nil || *ptr != value {
		t.Errorf("StringPtr failed: expected %q, got %v", value, ptr)
	}
}

func TestIntPtr(t *testing.T) {
	value := 42
	ptr := pointers.IntPtr(value)
	if ptr == nil || *ptr != value {
		t.Errorf("IntPtr failed: expected %d, got %v", value, ptr)
	}
}

func TestFloat64Ptr(t *testing.T) {
	value := 3.14
	ptr := pointers.Float64Ptr(value)
	if ptr == nil || *ptr != value {
		t.Errorf("Float64Ptr failed: expected %f, got %v", value, ptr)
	}
}

func TestFloat32Ptr(t *testing.T) {
	value := float32(2.71)
	ptr := pointers.Float32Ptr(value)
	if ptr == nil || *ptr != value {
		t.Errorf("Float32Ptr failed: expected %f, got %v", value, ptr)
	}
}

func TestComplex64Ptr(t *testing.T) {
	value := complex(float32(1), float32(2))
	ptr := pointers.Complex64Ptr(value)
	if ptr == nil || *ptr != value {
		t.Errorf("Complex64Ptr failed: expected %v, got %v", value, ptr)
	}
}

func TestComplex128Ptr(t *testing.T) {
	value := complex(3.0, 4.0)
	ptr := pointers.Complex128Ptr(value)
	if ptr == nil || *ptr != value {
		t.Errorf("Complex128Ptr failed: expected %v, got %v", value, ptr)
	}
}

func TestBytePtr(t *testing.T) {
	value := byte(255)
	ptr := pointers.BytePtr(value)
	if ptr == nil || *ptr != value {
		t.Errorf("BytePtr failed: expected %v, got %v", value, ptr)
	}
}

func TestStrPtrWithEmptyString(t *testing.T) {
	value := ""
	ptr := pointers.StringPtr(value)
	if ptr == nil || *ptr != value {
		t.Errorf("StringPtr failed: expected %q, got %v", value, ptr)
	}
}
