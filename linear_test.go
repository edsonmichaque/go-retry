package retry

import (
	"testing"
	"time"
)

func TestLinear(t *testing.T) {
	e := NewLinearBackoff(100 * time.Millisecond)

	if e.Delay(0) != 0 {
		t.Errorf("expected error")
	}

	if e.Delay(1) != 100*time.Millisecond {
		t.Errorf("expected error")
	}

	if e.Delay(8) != 800*time.Millisecond {
		t.Errorf("expected error")
	}

}
