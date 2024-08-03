package TaskProgress

import (
	"fmt"
	"testing"
	"time"
)

func Test_Taskprogress(t *testing.T) {
	// Create a new task with a total of 100 units.
	task := NewTaskProgress(100)

	// Simulate updating the task progress.
	for i := 0; i <= 100; i += 10 {
		task.Update(10)
		task.Display4()
		time.Sleep(1 * time.Second) // Sleep for 1 second to simulate work being done.
	}
	fmt.Println() // Print a newline at the end.
}
