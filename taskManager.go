package main

import (
	"container/heap"
	"fmt"
)

func TestTaskHeap() {
	nth := make(TaskHeap, 0)
	heap.Init(&nth)
	fmt.Println(nth.Len() == 0)

	t0 := &Task{UID: 1000, TaskID: 2000, Priority: 3000}
	heap.Push(&nth, t0)
	fmt.Println(nth.Len() == 1)
	fmt.Println(t0.Index == 0)

	t1 := &Task{UID: 1000, TaskID: 2001, Priority: 3001}
	heap.Push(&nth, t1)
	fmt.Println(t1.Index == 0)
	fmt.Println(t0.Index == 1)
	fmt.Println(nth[0].TaskID == 2001)

	tp := heap.Pop(&nth)
	fmt.Println(tp.(*Task).TaskID == 2001)
	fmt.Println(nth[0].TaskID == 2000)
	fmt.Println(t0.Index == 0)
}

func TestTaskManager() {
	tm := ConstructorTM([][]int{{1, 101, 10}, {2, 102, 20}, {3, 103, 15}}) // Initializes with three tasks for Users 1, 2, and 3.
	fmt.Println(len(tm.TaskIDMap) == 3)

	tm.Add(3, 1031, 5)
	fmt.Println(len(tm.TaskIDMap) == 4)

	tm.Add(4, 104, 5) // Adds task 104 with priority 5 for User 4.
	fmt.Println(len(tm.TaskIDMap) == 5)

	tm.Edit(102, 8)                // Updates priority of task 102 to 8.
	fmt.Println(tm.ExecTop() == 3) // return 3. Executes task 103 for User 3.
	fmt.Println(len(tm.TaskIDMap) == 4)

	tm.Rmv(101) // Removes task 101 from the system.
	fmt.Println(len(tm.TaskIDMap) == 3)

	tm.Add(5, 105, 15)             // Adds task 105 with priority 15 for User 5.
	fmt.Println(tm.ExecTop() == 5) // return 5. Executes task 105 for User 5.

	fmt.Println(tm.ExecTop() == 2)
	fmt.Println(tm.ExecTop() == 3)
	fmt.Println(tm.ExecTop() == 4)
	fmt.Println(tm.ExecTop() == -1)
}

type Task struct {
	UID      int
	TaskID   int
	Priority int
	Index    int
}

type TaskHeap []*Task

func (h TaskHeap) Len() int { return len(h) }
func (h TaskHeap) Less(i, j int) bool {
	if h[i].Priority > h[j].Priority {
		return true
	}
	if h[i].Priority == h[j].Priority {
		return h[i].TaskID > h[j].TaskID
	}
	return false
}

func (h TaskHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].Index, h[j].Index = h[j].Index, h[i].Index
}
func (h *TaskHeap) Push(x interface{}) {
	*h = append(*h, x.(*Task))
	(*h)[len(*h)-1].Index = len(*h) - 1
}
func (h *TaskHeap) Pop() interface{} { rc := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return rc }

type TaskManager struct {
	UIDTasks  *TaskHeap
	TaskIDMap map[int]*Task // 提供edit和remove要使用的index
}

func ConstructorTM(tasks [][]int) TaskManager {
	ut := make(TaskHeap, 0)
	heap.Init(&ut)
	t := TaskManager{UIDTasks: &ut, TaskIDMap: make(map[int]*Task)}
	for i := 0; i < len(tasks); i++ {
		uid := tasks[i][0]
		task := &Task{UID: uid, TaskID: tasks[i][1], Priority: tasks[i][2]}
		heap.Push(t.UIDTasks, task)
		t.TaskIDMap[task.TaskID] = task
	}
	return t
}

func (this *TaskManager) Add(userId int, taskId int, priority int) {
	task := &Task{UID: userId, TaskID: taskId, Priority: priority}
	heap.Push(this.UIDTasks, task)
	this.TaskIDMap[task.TaskID] = task
}

func (this *TaskManager) Edit(taskId int, newPriority int) {
	if this.TaskIDMap[taskId] == nil {
		return
	}
	task := this.TaskIDMap[taskId]
	task.Priority = newPriority

	heap.Fix(this.UIDTasks, task.Index)
}

func (this *TaskManager) Rmv(taskId int) {
	if this.TaskIDMap[taskId] == nil {
		return
	}
	task := this.TaskIDMap[taskId]
	heap.Remove(this.UIDTasks, task.Index)
	delete(this.TaskIDMap, task.TaskID)
}

func (this *TaskManager) ExecTop() int {
	if this.UIDTasks.Len() > 0 {
		toExecuteTask := heap.Pop(this.UIDTasks).(*Task)
		if toExecuteTask != nil {
			delete(this.TaskIDMap, toExecuteTask.TaskID)
			return toExecuteTask.UID
		}
	}
	return -1
}
