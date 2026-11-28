package calc

import "testing"

func TestTriple42067f(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want int
	}{
		{"zero", 0, 0},
		{"two", 2, 6},
		{"negative", -4, -12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Triple42067f(tt.in); got != tt.want {
				t.Errorf("Triple42067f(%d) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}
