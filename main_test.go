package main

import "testing"

func TestAbs(t *testing.T) {
	tests := []struct {
		name   string
		values []float64
		want   []float64
	}{
		{
			name:   "First values",
			values: []float64{-3, 3, -2.000001, -0.000000003},
			want:   []float64{3, 3, 2.000001, 0.000000003},
		},
	}

	for _, tt := range tests {
		for ik := range tt.values {
			t.Run(tt.name, func(t *testing.T) {
				if res := Abs(tt.values[ik]); res != tt.want[ik] {
					t.Errorf("Error in test: %s. Waited: %f. Got: %f", tt.name, tt.want[ik], res)
				}
			})
		}
	}
}
