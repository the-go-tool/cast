package cast

import "testing"

func TestInt(t *testing.T) {
	t.Run("int to string", func(t *testing.T) {
		if val, err := WithError[string](int(-69)); err != nil {
			t.Fatalf("unexpected error: %v", err)
		} else {
			if val != "-69" {
				t.Fatalf("unexpected result, got: %v", val)
			}
		}
	})
	t.Run("int8 to string", func(t *testing.T) {
		if val, err := WithError[string](int8(-69)); err != nil {
			t.Fatalf("unexpected error: %v", err)
		} else {
			if val != "-69" {
				t.Fatalf("unexpected result, got: %v", val)
			}
		}
	})
	t.Run("int16 to string", func(t *testing.T) {
		if val, err := WithError[string](int16(-69)); err != nil {
			t.Fatalf("unexpected error: %v", err)
		} else {
			if val != "-69" {
				t.Fatalf("unexpected result, got: %v", val)
			}
		}
	})
	t.Run("int32 to string", func(t *testing.T) {
		if val, err := WithError[string](int32(-69)); err != nil {
			t.Fatalf("unexpected error: %v", err)
		} else {
			if val != "-69" {
				t.Fatalf("unexpected result, got: %v", val)
			}
		}
	})
	t.Run("int64 to string", func(t *testing.T) {
		if val, err := WithError[string](int64(-69)); err != nil {
			t.Fatalf("unexpected error: %v", err)
		} else {
			if val != "-69" {
				t.Fatalf("unexpected result, got: %v", val)
			}
		}
	})
}
