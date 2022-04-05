package cast

import "testing"

func TestString(t *testing.T) {
	t.Run("string to int", func(t *testing.T) {
		if val, err := WithError[int]("-69"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		} else {
			if val != -69 {
				t.Fatalf("unexpected result, got: %v", val)
			}
		}
	})
	t.Run("string to int8", func(t *testing.T) {
		if val, err := WithError[int8]("-69"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		} else {
			if val != -69 {
				t.Fatalf("unexpected result, got: %v", val)
			}
		}
	})
	t.Run("string to int16", func(t *testing.T) {
		if val, err := WithError[int16]("-69"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		} else {
			if val != -69 {
				t.Fatalf("unexpected result, got: %v", val)
			}
		}
	})
	t.Run("string to int32", func(t *testing.T) {
		if val, err := WithError[int32]("-69"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		} else {
			if val != -69 {
				t.Fatalf("unexpected result, got: %v", val)
			}
		}
	})
	t.Run("string to int64", func(t *testing.T) {
		if val, err := WithError[int64]("-69"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		} else {
			if val != -69 {
				t.Fatalf("unexpected result, got: %v", val)
			}
		}
	})

	t.Run("string to uint", func(t *testing.T) {
		if val, err := WithError[uint]("69"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		} else {
			if val != 69 {
				t.Fatalf("unexpected result, got: %v", val)
			}
		}
	})
	t.Run("string to uint8", func(t *testing.T) {
		if val, err := WithError[uint8]("69"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		} else {
			if val != 69 {
				t.Fatalf("unexpected result, got: %v", val)
			}
		}
	})
	t.Run("string to uint16", func(t *testing.T) {
		if val, err := WithError[uint16]("69"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		} else {
			if val != 69 {
				t.Fatalf("unexpected result, got: %v", val)
			}
		}
	})
	t.Run("string to uint32", func(t *testing.T) {
		if val, err := WithError[uint32]("69"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		} else {
			if val != 69 {
				t.Fatalf("unexpected result, got: %v", val)
			}
		}
	})
	t.Run("string to uint64", func(t *testing.T) {
		if val, err := WithError[uint64]("69"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		} else {
			if val != 69 {
				t.Fatalf("unexpected result, got: %v", val)
			}
		}
	})
}
