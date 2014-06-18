package task

import (
	"testing"
)

func TestEncodeTasksToJson(t *testing.T) {
	tasks := make([]*Task, 3)
	tasks[0] = &Task{Title: "Test", Description: "TestDesc"}
	tasks[1] = &Task{Title: "Test2", Description: "TestDesc2"}
	tasks[2] = &Task{Title: "Test3", Description: "TestDesc3"}

	var in = &TaskList{Tasks: tasks}
	var out = "{\"Tasks\":[{\"Title\":\"Test\",\"Description\":\"TestDesc\",\"Important\":false,\"Status\":\"\",\"Created\":\"0001-01-01T00:00:00Z\",\"LastChange\":\"0001-01-01T00:00:00Z\",\"DueTime\":\"0001-01-01T00:00:00Z\"},{\"Title\":\"Test2\",\"Description\":\"TestDesc2\",\"Important\":false,\"Status\":\"\",\"Created\":\"0001-01-01T00:00:00Z\",\"LastChange\":\"0001-01-01T00:00:00Z\",\"DueTime\":\"0001-01-01T00:00:00Z\"},{\"Title\":\"Test3\",\"Description\":\"TestDesc3\",\"Important\":false,\"Status\":\"\",\"Created\":\"0001-01-01T00:00:00Z\",\"LastChange\":\"0001-01-01T00:00:00Z\",\"DueTime\":\"0001-01-01T00:00:00Z\"}]}"
	if x, err := in.encodeTasksToJson(); string(x) != out {
		t.Errorf("%v.encodeTasksToJson() = %v, want %v, error: %v", in, string(x), out, err)
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
