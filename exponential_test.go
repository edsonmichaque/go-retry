package retry

import (
	"testing"
	"time"
)

func TestExponential(t *testing.T) {
	e := NewExponentialBackoff(time.Millisecond)

	if e.Delay(0) != 0 {
		t.Errorf("expected error")
	}

	if e.Delay(1) != time.Millisecond {
		t.Errorf("expected error")
	}

	if e.Delay(8) != 128*time.Millisecond {
		t.Errorf("expected error")
	}

}
