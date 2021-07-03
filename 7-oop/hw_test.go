package hw

import (
	"testing"
)

func TestGeom_CalculateDistance(t *testing.T) {
	tests := []struct {
		name         string
		geom         Geom
		wantDistance float64
	}{
		{
			name:         "#1",
			geom:         Geom{start: Coordinate{x: 1, y: 1}, end: Coordinate{x: 4, y: 5}},
			wantDistance: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDistance := tt.geom.CalculateDistance(); gotDistance != tt.wantDistance {
				t.Errorf("Geom.CalculateDistance() = %v, want %v", gotDistance, tt.wantDistance)
			}
		})
	}
}
