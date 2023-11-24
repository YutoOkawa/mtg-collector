package server

import "testing"

func TestNewServer(t *testing.T) {
	cases := []struct {
		name string

		port            int
		shutdownTimeout int

		wantPort string
	}{
		{
			name:            "ShouldReturnServerWithValidPort",
			port:            8080,
			shutdownTimeout: 1,

			wantPort: ":8080",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			got := NewServer(c.port, c.shutdownTimeout)

			if got.port != c.wantPort {
				t.Errorf("NewServer returned %s, but expected %s", got.port, c.wantPort)
			}
		})
	}
}
