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

func TestEncodeToJson(t *testing.T) {
	var in, out = &Task{Title: "Test", Description: "TestDesc"}, "{\"Title\":\"Test\",\"Description\":\"TestDesc\",\"Important\":false,\"Status\":\"\",\"Created\":\"0001-01-01T00:00:00Z\",\"LastChange\":\"0001-01-01T00:00:00Z\",\"DueTime\":\"0001-01-01T00:00:00Z\"}"
	if x := in.encodeToJson(); string(x) != out {
		t.Errorf("%v.encodeToJson() = %v, want %v", in, string(x), out)
	}
}
