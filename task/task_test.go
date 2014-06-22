package task

import (
	"testing"
	"time"
)

func TestQuadrant(t *testing.T) {
	var in1, in2, out = []bool{true, true, false, false},
		[]time.Time{time.Now(), time.Now().Add(72 * time.Hour), time.Now(), time.Now().Add(72 * time.Hour)},
		[]int{1, 2, 3, 4}

	for i := 0; i < len(in1); i++ {
		ta := &Task{Important: in1[i], DueTime: in2[i]}
		if x := ta.Quadrant(); x != out[i] {
			t.Errorf("(%v,%v).Quadrant() = %v, want %v", in1[i], in2[i], x, out[i])
		}
	}
}
