package main

import (
	"encoding/json/v2"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func seeInstructions() {
	log.Fatal(`Unknown Command
See options

# Adding a new task
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
task-cli update 1 "Buy groceries and cook dinner"
task-cli delete 1

# Marking a task as in-progress or done or todo
task-cli mark-in-progress 1
task-cli mark-done 1
task-cli mark-todo 1

By default a task is set to todo

# Listing all tasks
task-cli list

# Listing tasks by status
task-cli list done
task-cli list todo
task-cli list in-progress	
	`)
}

type Status string

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in-progress"
	StatusDone       Status = "done"
)

// Go uses casing to determine access:

// Uppercase first letter (ID): The field is exported (public). Other packages, as well as the standard library's encoding/json package, can read and write to it.
// Go’s official naming convention (documented in Effective Go and the Go Code Review Comments) specifies that acronyms and initialisms should maintain consistent case - ID

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

const taskFile = "tasks.json"

func getTasks() []Task {
	data, err := os.ReadFile(taskFile)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Task{}
		}
		panic(err)
	}

	// handle empty file
	if len(data) == 0 {
		return []Task{}
	}

	var tasks []Task

	if err := json.Unmarshal(data, &tasks); err != nil {
		panic(err)
	}

	return tasks
}

func saveTasks(tasks []Task) {
	serializedTasks, serializeErr := json.Marshal(tasks)

	if serializeErr != nil {
		panic(serializeErr)
	}

	// 0644 gives read/write permissions to owner, read-only to others
	writeErr := os.WriteFile(taskFile, serializedTasks, 0644)

	if writeErr != nil {
		panic(writeErr)
	}
}

func handleAdd(args []string) {
	// add <string>
	if len(args) < 2 {
		seeInstructions()
	}

	description := args[1]

	tasks := getTasks()
	newTask := Task{
		ID:          len(tasks) + 1,
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)

	saveTasks(tasks)

	fmt.Printf("Task added successfully (ID: %d)", newTask.ID)
}

func updateTask(id int, desc string, status Status) {
	tasks := getTasks()

	found := false
	for i := range tasks {
		if tasks[i].ID == id {

			description := tasks[i].Description
			tempStatus := tasks[i].Status

			if len(desc) > 0 {
				description = desc
			}

			if len(status) > 0 {
				tempStatus = status
			}

			found = true
			tasks[i].Description = description
			tasks[i].Status = tempStatus
			tasks[i].UpdatedAt = time.Now()
			break
		}
	}

	if !found {
		log.Fatalf("Task with ID %d does not exist", id)
	}

	saveTasks(tasks)

	fmt.Printf("Task updated successfully (ID: %d)", id)
}

func handleUpdate(args []string) {
	// update <int> <string>
	if len(args) < 3 {
		seeInstructions()
	}
	taskId, convertErr := strconv.Atoi(args[1])

	if convertErr != nil {
		seeInstructions()
	}

	newDescription := args[2]

	updateTask(taskId, newDescription, "")
}

func handleDelete(args []string) {
	// delete <int>
	if len(args) < 2 {
		seeInstructions()
	}

	taskId, convertErr := strconv.Atoi(args[1])

	if convertErr != nil {
		seeInstructions()
	}

	tasks := getTasks()

	taskIdx := -1
	for i := range tasks {
		if tasks[i].ID == taskId {
			taskIdx = i
			break
		}
	}

	if taskIdx == -1 {
		log.Fatalf("Task with ID %d does not exist", taskId)
	}

	// get slice from index 0 up to taskIdx then after taskIdx to end of original slice
	tasks = append(tasks[:taskIdx], tasks[taskIdx+1:]...)

	saveTasks(tasks)

	fmt.Printf("Task deleted successfully (ID: %d)", taskId)
}

func handleMarkStatus(args []string) {
	// mark-in-progress <int>
	// mark-done <int>
	// mark-todo <int>
	if len(args) < 2 {
		seeInstructions()
	}

	taskId, convertErr := strconv.Atoi(args[1])

	if convertErr != nil {
		seeInstructions()
	}

	statusToUpdateTo := strings.Split(args[0], "mark-")[1]

	statuses := [3]Status{StatusInProgress, StatusDone, StatusTodo}

	statusFound := false
	for i := range statuses {
		if statuses[i] == Status(statusToUpdateTo) {
			statusFound = true
			updateTask(taskId, "", Status(statusToUpdateTo))
			break
		}
	}

	if !statusFound {
		seeInstructions()
	}

}

func printTaskTable(tasks []Task) {
	descWidth := len("TASK")
	for _, t := range tasks {
		if len(t.Description) > descWidth {
			descWidth = len(t.Description)
		}
	}

	fmt.Printf("%-4s  %-*s  %-12s  %-16s  %s\n", "ID", descWidth, "TASK", "STATUS", "CREATED AT", "UPDATED AT")
	fmt.Printf("%-4s  %-*s  %-12s  %-16s  %s\n", "--", descWidth, strings.Repeat("-", descWidth), "------", "----------------", "----------------")

	for _, t := range tasks {
		createdStr := t.CreatedAt.Format("2006-01-02 3:04 PM")
		updatedStr := t.UpdatedAt.Format("2006-01-02 3:04 PM")

		fmt.Printf("%-4d  %-*s  %-12s  %-16s  %s\n", t.ID, descWidth, t.Description, t.Status, createdStr, updatedStr)
	}
}

func handleListing(args []string) {
	allTasks := getTasks()
	var tasksToView []Task

	if len(args) > 1 {
		statuses := [3]Status{StatusInProgress, StatusDone, StatusTodo}

		statusToView := args[1]

		statusFound := false
		for i := range statuses {
			if statuses[i] == Status(statusToView) {
				statusFound = true
				break
			}
		}

		if !statusFound {
			seeInstructions()
		}

		for i := range allTasks {
			if allTasks[i].Status == Status(statusToView) {
				tasksToView = append(tasksToView, allTasks[i])
			}
		}

	} else {
		tasksToView = allTasks
	}

	if len(tasksToView) == 0 {
		println("There are no tasks")
		return
	}

	printTaskTable(tasksToView)
}

func main() {

	if len(os.Args) < 2 {
		seeInstructions()
	}

	// args[0] is always the file path omit it
	args := os.Args[1:]

	switch args[0] {
	case "add":
		handleAdd(args)

	case "update":
		handleUpdate(args)

	case "delete":
		handleDelete(args)

	case "mark-in-progress", "mark-done", "mark-todo":
		handleMarkStatus(args)

	case "list":
		handleListing(args)

	default:
		seeInstructions()
	}
}
