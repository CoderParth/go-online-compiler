package execute

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func RunGoCodeService(c echo.Context) error {
	req := new(CodeRequest)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request")
	}

	tmpFile, err := createTempFile(req.Code)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	resultChan := make(chan string)
	task := Task{
		FilePath: tmpFile.Name(),
		Result:   resultChan,
	}

	// Enqueue the task
	taskQueue <- task

	// Wait for the result
	result := <-resultChan

	// if the result indicates a timeout
	if result == "Execution timed out after 10 seconds" {
        return c.String(http.StatusRequestTimeout, result)
    }

	return c.String(http.StatusOK, result)
}

func createTempFile(code string) (*os.File, error) {
	tmpFile, err := os.CreateTemp("", "*.go")
	if err != nil {
		return nil, fmt.Errorf("Error creating temporary file: %v", err)
	}

	if _, err := tmpFile.WriteString(code); err != nil {
		return nil, fmt.Errorf("Error writing code to temporary file: %v", err)
	}

	return tmpFile, nil
}

func handleError(c echo.Context, err error, output string) error {
	if err.Error() == "Execution timed out after 10 seconds" {
		return c.String(http.StatusRequestTimeout, err.Error())
	}
	fmt.Printf("Error executing code: %v\nOutput: %s\n", err, output)
	return c.String(http.StatusInternalServerError, output)
}
