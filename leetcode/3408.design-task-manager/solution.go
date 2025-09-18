package taskmanager

import "container/heap"

type Task struct {
	id       int
	userId   int
	priority int
	index    int
}

func NewTask(id int, userId int, priority int, index int) *Task {
	return &Task{id, userId, priority, index}
}

type ExecutionQueue []*Task

func (eq ExecutionQueue) Len() int {
	return len(eq)
}

func (eq ExecutionQueue) Less(i, j int) bool {
	if eq[i].priority == eq[j].priority {
		return eq[i].id > eq[j].id
	}
	return eq[i].priority > eq[j].priority
}

func (eq ExecutionQueue) Swap(i, j int) {
	eq[i], eq[j] = eq[j], eq[i]
	eq[i].index = i
	eq[j].index = j
}

func (eq *ExecutionQueue) Push(x any) {
	*eq = append(*eq, x.(*Task))
}

func (eq *ExecutionQueue) Pop() any {
	old := *eq
	n := len(old)
	popped := old[n-1]
	old[n-1] = nil
	*eq = old[0 : n-1]
	return popped
}

type TaskManager struct {
	taskById       map[int]*Task
	executionQueue *ExecutionQueue
}

// input [userId, taskId, priority]
func NewTaskManager(tasks [][]int) TaskManager {
	taskById := make(map[int]*Task)
	executionQueue := make(ExecutionQueue, len(tasks))

	for i, t := range tasks {
		task := NewTask(t[1], t[0], t[2], i)
		taskById[t[1]] = task
		executionQueue[i] = task
	}

	heap.Init(&executionQueue)

	return TaskManager{taskById, &executionQueue}
}

func (tm *TaskManager) add(userId int, taskId int, priority int) {
	task := NewTask(taskId, userId, priority, tm.executionQueue.Len())
	tm.taskById[taskId] = task
	heap.Push(tm.executionQueue, task)
}

func (tm *TaskManager) edit(taskId int, newPriority int) {
	task := tm.taskById[taskId]
	task.priority = newPriority
	heap.Fix(tm.executionQueue, task.index)
}

func (tm *TaskManager) remove(taskId int) {
	task := tm.taskById[taskId]
	delete(tm.taskById, taskId)
	heap.Remove(tm.executionQueue, task.index)
	task = nil
}

func (tm *TaskManager) executeTop() int {
	if tm.executionQueue.Len() == 0 {
		return -1
	}
	executed := heap.Pop(tm.executionQueue).(*Task)
	delete(tm.taskById, executed.id)
	return executed.userId
}
