package cast

import "testing"

func TestUint(t *testing.T) {
	t.Run("uint to string", func(t *testing.T) {
		if val, err := WithError[string](uint(69)); err != nil {
			t.Fatalf("unexpected error: %v", err)
		} else {
			if val != "69" {
				t.Fatalf("unexpected result, got: %v", val)
			}
		}
	})
	t.Run("uint8 to string", func(t *testing.T) {
		if val, err := WithError[string](uint8(69)); err != nil {
			t.Fatalf("unexpected error: %v", err)
		} else {
			if val != "69" {
				t.Fatalf("unexpected result, got: %v", val)
			}
		}
	})
	t.Run("uint16 to string", func(t *testing.T) {
		if val, err := WithError[string](uint16(69)); err != nil {
			t.Fatalf("unexpected error: %v", err)
		} else {
			if val != "69" {
				t.Fatalf("unexpected result, got: %v", val)
			}
		}
	})
	t.Run("uint32 to string", func(t *testing.T) {
		if val, err := WithError[string](uint32(69)); err != nil {
			t.Fatalf("unexpected error: %v", err)
		} else {
			if val != "69" {
				t.Fatalf("unexpected result, got: %v", val)
			}
		}
	})
	t.Run("uint64 to string", func(t *testing.T) {
		if val, err := WithError[string](uint64(69)); err != nil {
			t.Fatalf("unexpected error: %v", err)
		} else {
			if val != "69" {
				t.Fatalf("unexpected result, got: %v", val)
			}
		}
	})
}
