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

func TestStrPtrWithEmptyString(t *testing.T) {
	value := ""
	ptr := pointers.StringPtr(value)
	if ptr == nil || *ptr != value {
		t.Errorf("StringPtr failed: expected %q, got %v", value, ptr)
	}
}

// --- Signed Integers ---

func TestIntPtr(t *testing.T) {
	value := 42
	ptr := pointers.IntPtr(value)
	if ptr == nil || *ptr != value {
		t.Errorf("IntPtr failed: expected %d, got %v", value, ptr)
	}
}

func TestInt8Ptr(t *testing.T) {
	value := int8(127)
	ptr := pointers.Int8Ptr(value)
	if ptr == nil || *ptr != value {
		t.Errorf("Int8Ptr failed: expected %d, got %v", value, ptr)
	}
}

func TestInt16Ptr(t *testing.T) {
	value := int16(32000)
	ptr := pointers.Int16Ptr(value)
	if ptr == nil || *ptr != value {
		t.Errorf("Int16Ptr failed: expected %d, got %v", value, ptr)
	}
}

func TestInt32Ptr(t *testing.T) {
	value := int32(100000)
	ptr := pointers.Int32Ptr(value)
	if ptr == nil || *ptr != value {
		t.Errorf("Int32Ptr failed: expected %d, got %v", value, ptr)
	}
}

func TestInt64Ptr(t *testing.T) {
	value := int64(9000000000)
	ptr := pointers.Int64Ptr(value)
	if ptr == nil || *ptr != value {
		t.Errorf("Int64Ptr failed: expected %d, got %v", value, ptr)
	}
}

// --- Unsigned Integers ---

func TestUintPtr(t *testing.T) {
	value := uint(42)
	ptr := pointers.UintPtr(value)
	if ptr == nil || *ptr != value {
		t.Errorf("UintPtr failed: expected %d, got %v", value, ptr)
	}
}

func TestUint8Ptr(t *testing.T) {
	value := uint8(255)
	ptr := pointers.Uint8Ptr(value)
	if ptr == nil || *ptr != value {
		t.Errorf("Uint8Ptr failed: expected %d, got %v", value, ptr)
	}
}

func TestBytePtr(t *testing.T) {
	value := byte(255)
	ptr := pointers.BytePtr(value)
	if ptr == nil || *ptr != value {
		t.Errorf("BytePtr failed: expected %v, got %v", value, ptr)
	}
}

func TestUint16Ptr(t *testing.T) {
	value := uint16(65000)
	ptr := pointers.Uint16Ptr(value)
	if ptr == nil || *ptr != value {
		t.Errorf("Uint16Ptr failed: expected %d, got %v", value, ptr)
	}
}

func TestUint32Ptr(t *testing.T) {
	value := uint32(4000000)
	ptr := pointers.Uint32Ptr(value)
	if ptr == nil || *ptr != value {
		t.Errorf("Uint32Ptr failed: expected %d, got %v", value, ptr)
	}
}

func TestUint64Ptr(t *testing.T) {
	value := uint64(18000000000000000000)
	ptr := pointers.Uint64Ptr(value)
	if ptr == nil || *ptr != value {
		t.Errorf("Uint64Ptr failed: expected %d, got %v", value, ptr)
	}
}

func TestRunePtr(t *testing.T) {
	value := 'G' // rune is an alias for int32
	ptr := pointers.RunePtr(value)
	if ptr == nil || *ptr != value {
		t.Errorf("RunePtr failed: expected %c, got %v", value, ptr)
	}
}

// --- Floats & Complex ---

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

// --- Any ---

func TestAnyPtr(t *testing.T) {
	value := "any value"
	ptr := pointers.AnyPtr(value)
	if ptr == nil || *ptr != value {
		t.Errorf("AnyPtr failed: expected %v, got %v", value, ptr)
	}
}
