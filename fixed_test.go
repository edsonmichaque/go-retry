package retry

import (
	"testing"
	"time"
)

func TestFixed(t *testing.T) {
	e := NewFixedBackoff(time.Millisecond)

	if e.Delay(0) != 0 {
		t.Errorf("expected error")
	}

	if e.Delay(1) != time.Millisecond {
		t.Errorf("expected error")
	}

	if e.Delay(8) != time.Millisecond {
		t.Errorf("expected error")
	}

}
