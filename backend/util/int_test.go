package util

import "testing"

func TestToPointerInt(t *testing.T) {
	t.Run("Convert int to pointer int", func(t *testing.T) {
		pointerInt := ToPointerInt(1)
		if *pointerInt != 1 {
			t.Errorf("The value of pointer int is not 1")
		}
	})

	t.Run("Convert 0 to pointer int", func(t *testing.T) {
		pointerInt := ToPointerInt(0)
		if *pointerInt != 0 {
			t.Errorf("The value of pointer int is not 0")
		}
	})
}

func TestToPointerInt64(t *testing.T) {
	t.Run("Convert int64 to pointer int64", func(t *testing.T) {
		pointerInt64 := ToPointerInt64(1)
		if *pointerInt64 != 1 {
			t.Errorf("The value of pointer int64 is not 1")
		}
	})

	t.Run("Convert 0 to pointer int64", func(t *testing.T) {
		pointerInt64 := ToPointerInt64(0)
		if *pointerInt64 != 0 {
			t.Errorf("The value of pointer int64 is not 0")
		}
	})
}

func TestToPointerFloat64(t *testing.T) {
	t.Run("Convert float64 to pointer float64", func(t *testing.T) {
		pointerFloat64 := ToPointerFloat64(1.1)
		if *pointerFloat64 != 1.1 {
			t.Errorf("The value of pointer float64 is not 1.1")
		}
	})

	t.Run("Convert 0 to pointer float64", func(t *testing.T) {
		pointerFloat64 := ToPointerFloat64(0)
		if *pointerFloat64 != 0 {
			t.Errorf("The value of pointer float64 is not 0")
		}
	})
}
