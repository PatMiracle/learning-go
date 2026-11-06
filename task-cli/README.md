# Task Tracker CLI

A lightweight, zero-dependency Command Line Interface (CLI) application built with Go to track and manage your daily tasks. Task data is automatically persisted locally in a JSON file (`tasks.json`).

This project is an implementation of the [roadmap.sh Task Tracker Project](https://roadmap.sh/projects/task-tracker).

---

## Features

- **Add Tasks**: Quickly create tasks with a description (defaults to `todo` status).
- **Update Tasks**: Modify existing task descriptions.
- **Delete Tasks**: Remove tasks by their unique identifier.
- **Status Management**: Track progress with `todo`, `in-progress`, and `done` states.
- **Formatted Listing**: View all tasks or filter by status in a cleanly formatted terminal table.
- **Local Persistence**: Tasks are saved to `tasks.json` with timestamps (`createdAt`, `updatedAt`).
- **Standard Library Only**: Built entirely using Go's standard libraries with zero external dependencies.

---

## Prerequisites

- [Go](https://go.dev/dl/) (version 1.24 or later recommended).

---

## Installation & Setup

1. **Clone or open the repository**:
   ```bash
   cd task-cli
   ```

2. **Option A: Run directly with Go**:
   ```bash
   go run main.go <command> [arguments]
   ```

3. **Option B: Build the binary (Recommended)**:

   - **Windows (PowerShell / CMD)**:
     ```powershell
     go build -o task-cli.exe
     ```
     Run via:
     ```powershell
     .\task-cli.exe <command> [arguments]
     ```

   - **Linux / macOS**:
     ```bash
     go build -o task-cli
     chmod +x task-cli
     ```
     Run via:
     ```bash
     ./task-cli <command> [arguments]
     ```

*(Optional)* Move the compiled binary to a folder in your system's `PATH` to run `task-cli` from anywhere in your terminal.

---

## Command Reference

| Command | Arguments | Description |
| :--- | :--- | :--- |
| `add` | `"<description>"` | Creates a new task. |
| `update` | `<id> "<description>"` | Updates the description of an existing task. |
| `delete` | `<id>` | Deletes the task with the given ID. |
| `mark-in-progress`| `<id>` | Sets the status of a task to `in-progress`. |
| `mark-done` | `<id>` | Sets the status of a task to `done`. |
| `mark-todo` | `<id>` | Sets the status of a task to `todo`. |
| `list` | *none* | Lists all tasks in a tabular view. |
| `list` | `todo` \| `in-progress` \| `done` | Lists tasks filtered by status. |

---

## Usage Examples

### 1. Adding Tasks
```bash
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)

task-cli add "Write unit tests for authentication"
# Output: Task added successfully (ID: 2)
```

### 2. Updating a Task Description
```bash
task-cli update 1 "Buy groceries and cook dinner"
# Output: Task updated successfully (ID: 1)
```

### 3. Updating Task Status
```bash
# Mark a task as in-progress
task-cli mark-in-progress 1
# Output: Task updated successfully (ID: 1)

# Mark a task as completed
task-cli mark-done 1
# Output: Task updated successfully (ID: 1)

# Revert a task back to todo
task-cli mark-todo 1
# Output: Task updated successfully (ID: 1)
```

### 4. Listing Tasks

**List all tasks:**
```bash
task-cli list
```
*Output:*
```text
ID    TASK                                  STATUS        CREATED AT        UPDATED AT
--    ------------------------------------  ------        ----------------  ----------------
1     Buy groceries and cook dinner         done          2026-11-06 6:15 PM  2026-11-06 6:22 PM
2     Write unit tests for authentication   todo          2026-11-06 6:18 PM  2026-11-06 6:18 PM
```

**Filter tasks by status:**
```bash
task-cli list done
task-cli list in-progress
task-cli list todo
```

### 5. Deleting a Task
```bash
task-cli delete 1
# Output: Task deleted successfully (ID: 1)
```

---

## Data Storage Format

Tasks are stored in a local `tasks.json` file in the working directory:

```json
[
  {
    "id": 1,
    "description": "Buy groceries and cook dinner",
    "status": "done",
    "createdAt": "2026-11-06T18:15:00Z",
    "updatedAt": "2026-11-06T18:22:00Z"
  }
]
```

Each task record consists of:
- `id` *(integer)*: Unique auto-incremented identifier.
- `description` *(string)*: A short text describing the task.
- `status` *(string)*: One of `"todo"`, `"in-progress"`, or `"done"`.
- `createdAt` *(timestamp)*: ISO 8601 date-time when task was created.
- `updatedAt` *(timestamp)*: ISO 8601 date-time when task was last modified.

---

## Project Structure

```text
task-cli/
├── go.mod        # Go module definition
├── main.go       # Core CLI implementation and logic
├── tasks.json    # JSON storage file (created upon first save)
└── README.md     # Project documentation
```

---

## License

This project is open-source and free to use under the [MIT License](LICENSE).

