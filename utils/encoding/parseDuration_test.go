package encoding

import (
	"testing"
	"time"
)

func TestParseDuration(t *testing.T) {
	// "ns", "us", "ms", "s", "m", "h".
	tests := []string{
		"1d", "2h", "3m", "4s",
	}

	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			duration, err := ParseDuration(tt)
			if err != nil {
				t.Errorf("%v", err)
				return
			}
			t.Logf("%d", duration)
			t.Logf("%s", time.Now().Add(duration))
		})
	}

}
