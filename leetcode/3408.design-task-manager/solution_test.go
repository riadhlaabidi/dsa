package taskmanager

import (
	"dsa/internal/testutil"
	"testing"
)

func TestTaskManager(t *testing.T) {

	tasks := [][]int{
		{1, 101, 10},
		{2, 102, 20},
		{3, 103, 15},
	}

	tm := NewTaskManager(tasks)

	tm.add(4, 104, 5)
	tm.edit(102, 8)

	testutil.AssertEquals(t, 3, tm.executeTop())

	tm.remove(101)
	tm.add(5, 105, 15)

	testutil.AssertEquals(t, 5, tm.executeTop())
}
