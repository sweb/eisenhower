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

func TestSetDueTime(t *testing.T) {
	in := []string{"05.04.2014", "15.12.2014", "20.04.2014"}
	out := []time.Time{time.Date(2014, time.April, 5, 0, 0, 0, 0, time.Local),
		time.Date(2014, time.December, 15, 0, 0, 0, 0, time.Local),
		time.Date(2014, time.April, 20, 0, 0, 0, 0, time.Local)}

	for i := 0; i < len(in); i++ {
		ta := &Task{}
		err := ta.setDueTime(in[i])
		if x := ta.DueTime; x != out[i] {
			t.Errorf("task.setDueTime(%v) = %v, want %v ... error: %v", in[i], x, out[i], err)
		}
	}
}

func TestImportantString(t *testing.T) {
	in := []bool{true, false}
	out := []string{"Yes", "No"}

	for i := 0; i < len(in); i++ {
		ta := &Task{Important: in[i]}
		if x := ta.ImportantString(); x != out[i] {
			t.Errorf("task{Important: %v}.ImportantString() = %v, want %v", in[i], x, out[i])
		}
	}
}

func TestDueTimeString(t *testing.T) {
	in := []time.Time{time.Date(2014, time.April, 5, 0, 0, 0, 0, time.Local),
		time.Date(2014, time.December, 15, 0, 0, 0, 0, time.Local),
		time.Date(2014, time.April, 20, 0, 0, 0, 0, time.Local)}
	out := []string{"05.04.2014", "15.12.2014", "20.04.2014"}

	for i := 0; i < len(in); i++ {
		ta := &Task{DueTime: in[i]}
		if x := ta.DueTimeString(); x != out[i] {
			t.Errorf("task{DueTime: %v}.DueTimeString() = %v, want %v", in[i], x, out[i])
		}
	}
}

func TestCreatedString(t *testing.T) {
	in := []time.Time{time.Date(2014, time.April, 5, 15, 0, 0, 0, time.Local),
		time.Date(2014, time.December, 15, 15, 25, 0, 0, time.Local),
		time.Date(2014, time.April, 20, 8, 5, 0, 0, time.Local)}
	out := []string{"05.04.2014 15:00", "15.12.2014 15:25", "20.04.2014 08:05"}

	for i := 0; i < len(in); i++ {
		ta := &Task{Created: in[i]}
		if x := ta.CreatedString(); x != out[i] {
			t.Errorf("task{Created: %v}.CreatedString() = %v, want %v", in[i], x, out[i])
		}
	}
}

func TestUpdate(t *testing.T) {
	in1 := []string{"Text1", "Text2", "Text3"}
	in2 := []string{"Descr1", "Descr2", "Descr3"}
	in3 := []string{"Yes", "No", "Yes"}
	in4 := []string{"05.04.2014", "15.12.2014", "20.04.2014"}
	out1 := []string{"Text1", "Text2", "Text3"}
	out2 := []string{"Descr1", "Descr2", "Descr3"}
	out3 := []bool{true, false, true}
	out4 := []time.Time{time.Date(2014, time.April, 5, 0, 0, 0, 0, time.Local),
		time.Date(2014, time.December, 15, 0, 0, 0, 0, time.Local),
		time.Date(2014, time.April, 20, 0, 0, 0, 0, time.Local)}

	for i := 0; i < len(in1); i++ {
		ta := &Task{}
		ta.Update(in1[i], in2[i], in3[i], in4[i])
		if x := ta.Title; x != out1[i] {
			t.Errorf("task.Update(Title: %v) = %v, want %v", in1[i], x, out1[i])
		}
		if x := ta.Description; x != out2[i] {
			t.Errorf("task.Update(Description: %v) = %v, want %v", in2[i], x, out2[i])
		}
		if x := ta.Important; x != out3[i] {
			t.Errorf("task.Update(Important: %v) = %v, want %v", in3[i], x, out3[i])
		}
		if x := ta.DueTime; x != out4[i] {
			t.Errorf("task.Update(DueTime: %v) = %v, want %v", in4[i], x, out4[i])
		}
	}
}
