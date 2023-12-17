package logger

import (
	"testing"
)

func TestNewLogger(t *testing.T) {
	cases := []struct {
		name string
	}{
		{
			name: "ShouldNewLoggerSuccessfully",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			log := NewLogger()
			if log == nil {
				t.Error("log is not expected to be nil")
			}
		})
	}
}
