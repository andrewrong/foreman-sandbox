package calc

import "testing"

func TestTwice9ca39b(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want int
	}{
		{"zero", 0, 0},
		{"positive", 3, 6},
		{"negative", -4, -8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Twice9ca39b(tt.in); got != tt.want {
				t.Errorf("Twice9ca39b(%d) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}
