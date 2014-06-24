package task

import (
	"testing"
)

func TestHasTasks(t *testing.T) {
	in := &TaskList{}
	out := false

	if x := in.HasTasks(); x != out {
		t.Errorf("%v.HasTasks() = %v, want %v", in, x, out)
	}
}

func TestDeleteByTaskId(t *testing.T) {
	tl := &TaskList{}
	in := []int{2, 3, 1}
	out := []*TaskList{&TaskList{Tasks: []*Task{&Task{Id: 1}, &Task{Id: 3}}},
		&TaskList{Tasks: []*Task{&Task{Id: 1}, &Task{Id: 2}}},
		&TaskList{Tasks: []*Task{&Task{Id: 2}, &Task{Id: 3}}}}

	for i := 0; i < len(in); i++ {
		tl = &TaskList{Tasks: []*Task{&Task{Id: 1}, &Task{Id: 2}, &Task{Id: 3}}}
		if x := tl.DeleteByTaskId(in[i]); len(tl.Tasks) != len(out[i].Tasks) {
			t.Errorf("tl.DeleteByTaskId(%v) = %s, want %s - error: %v", in[i], tl.Tasks[0], out[i].Tasks[0], x)
		}
		if x := tl.DeleteByTaskId(in[i]); tl.Tasks[0].Id != out[i].Tasks[0].Id {
			t.Errorf("tl.DeleteByTaskId(%v) = %s, want %s - error: %v", in[i], tl.Tasks[0], out[i].Tasks[0], x)
		}
		if x := tl.DeleteByTaskId(in[i]); tl.Tasks[1].Id != out[i].Tasks[1].Id {
			t.Errorf("tl.DeleteByTaskId(%v) = %s, want %s - error: %v", in[i], tl.Tasks[0], out[i].Tasks[0], x)
		}
	}
}

/*func TestDecodeTasksFromJson(t *testing.T) {
	tasks := make([]*Task, 3)
	tasks[0] = &Task{Title: "Test", Description: "TestDesc"}
	tasks[1] = &Task{Title: "Test2", Description: "TestDesc2"}
	tasks[2] = &Task{Title: "Test3", Description: "TestDesc3"}

	tl := &TaskList{}

	var in = "{\"Tasks\":[{\"Title\":\"Test\",\"Description\":\"TestDesc\",\"Important\":false,\"Status\":\"\",\"Created\":\"0001-01-01T00:00:00Z\",\"LastChange\":\"0001-01-01T00:00:00Z\",\"DueTime\":\"0001-01-01T00:00:00Z\"},{\"Title\":\"Test2\",\"Description\":\"TestDesc2\",\"Important\":false,\"Status\":\"\",\"Created\":\"0001-01-01T00:00:00Z\",\"LastChange\":\"0001-01-01T00:00:00Z\",\"DueTime\":\"0001-01-01T00:00:00Z\"},{\"Title\":\"Test3\",\"Description\":\"TestDesc3\",\"Important\":false,\"Status\":\"\",\"Created\":\"0001-01-01T00:00:00Z\",\"LastChange\":\"0001-01-01T00:00:00Z\",\"DueTime\":\"0001-01-01T00:00:00Z\"}]}"
	var out = &TaskList{Tasks: tasks}

	if x := decodeTasksFromJson([]byte(in), tl); tl.Tasks[2] != out.Tasks[2] {
		t.Errorf("tasks.decodeTasksFromJson(%v) = %v, want %v - possible error: %v", in, tl.Tasks[2], out.Tasks[2], x)
	}
}*/
