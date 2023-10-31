package execute

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

type Task struct {
	FilePath string
	Result   chan string
}

var (
	taskQueue = make(chan Task, 100) // Task queue
)

func init() {
    // Dynamically set the number of workers based on CPU cores
    numWorkers := runtime.NumCPU()
    for w := 1; w <= numWorkers; w++ {
        go worker(w)
    }
}

func worker(id int) {
	for task := range taskQueue {
		fmt.Printf("Worker %d processing code\n", id)
		output, err := executeGoCode(task.FilePath)
		if err != nil {
			task.Result <- err.Error()
		} else {
			task.Result <- output
		}
		os.Remove(task.FilePath) // Cleanup the temporary file
	}
}

func executeGoCode(filePath string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "go", "run", filePath)
	output, err := cmd.CombinedOutput()
	outputStr := strings.TrimSpace(string(output))

	if ctx.Err() == context.DeadlineExceeded {
		return "", fmt.Errorf("Execution timed out after 10 seconds")
	}

	return outputStr, err
}
