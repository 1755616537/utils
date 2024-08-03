package TaskProgress

import (
	"fmt"
	"strings"
)

// TaskProgress represents the progress of a task.
type TaskProgress struct {
	Current int // Current progress value
	Total   int // Total progress value
}

// NewTaskProgress creates a new TaskProgress instance.
func NewTaskProgress(total int) *TaskProgress {
	return &TaskProgress{
		Current: 0,
		Total:   total,
	}
}

// Update updates the progress by a given value.
func (tp *TaskProgress) Update(value int) {
	tp.Current += value
}

// Percentage returns the percentage of the task completed.
func (tp *TaskProgress) Percentage() float64 {
	return (float64(tp.Current) / float64(tp.Total)) * 100
}

// Display displays the current progress as a progress bar.
func (tp *TaskProgress) Display() {
	percentage := tp.Percentage()
	width := 50 // Width of the progress bar
	filledLength := int((percentage / 100) * float64(width))
	emptyLength := width - filledLength
	bar := fmt.Sprintf("%s%s", strings.Repeat("=", filledLength), strings.Repeat(" ", emptyLength))
	fmt.Printf("\r[%s] %d%%", bar, int(percentage))
}

// Display displays the current progress in percentage.
func (tp *TaskProgress) Display2() {
	fmt.Printf("Progress: %.2f%%\n", tp.Percentage())
}

// Display displays the current progress as a progress bar.
func (tp *TaskProgress) Display3() {
	percentage := tp.Percentage()
	width := 50 // Width of the progress bar
	filledLength := int((percentage / 100) * float64(width))
	emptyLength := width - filledLength
	bar := fmt.Sprintf("%s%s", strings.Repeat("=", filledLength), strings.Repeat(" ", emptyLength))
	fmt.Printf("\r[%s] %d%%", bar, int(percentage))
	// Flush the output buffer to ensure the progress bar is displayed immediately.
	fmt.Print("\033[K") // Clear to the end of the line.
	fmt.Println()       // Move to the next line.
}

// Display displays the current progress as a progress bar.
func (tp *TaskProgress) Display4() {
	percentage := tp.Percentage()
	width := 50 // Width of the progress bar
	filledLength := int((percentage / 100) * float64(width))
	emptyLength := width - filledLength
	bar := fmt.Sprintf("%s%s", strings.Repeat("=", filledLength), strings.Repeat(" ", emptyLength))
	fmt.Printf("\r[%s] %d%%", bar, int(percentage))
}
