package calc

import "testing"

func TestSquare9ca39b(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want int
	}{
		{"zero", 0, 0},
		{"positive", 3, 9},
		{"negative", -4, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Square9ca39b(tt.in); got != tt.want {
				t.Errorf("Square9ca39b(%d) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}
