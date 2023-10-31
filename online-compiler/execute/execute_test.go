package execute

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRunGoCodeEndpoint(t *testing.T) {
    e := echo.New()

	// Test for successful code execution
    t.Run("TestSuccessfulExecution", func(t *testing.T) {
        t.Parallel()
        println("Starting Test: TestSuccessfulExecution")

        codeJSON := `{"code":"package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello, World!\")\n}"}`
        req := httptest.NewRequest(http.MethodPost, "/execute", bytes.NewBufferString(codeJSON))
        req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

        rec := httptest.NewRecorder()
        c := e.NewContext(req, rec)

        if assert.NoError(t, RunGoCode(c)) {
            assert.Equal(t, http.StatusOK, rec.Code)
            assert.Contains(t, rec.Body.String(), "Hello, World!")
        }

        println("Completed Test: TestSuccessfulExecution")
    })

	// Test for timeout
    t.Run("TestTimeoutExecution", func(t *testing.T) {
        t.Parallel()
        println("Starting Test: TestTimeoutExecution")

        longRunningCodeJSON := `{"code":"package main\n\nimport \"time\"\n\nfunc main() {\n\ttime.Sleep(20 * time.Second)\n\tprintln(\"Completed\")\n}"}`
        req := httptest.NewRequest(http.MethodPost, "/execute", bytes.NewBufferString(longRunningCodeJSON))
        req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

        rec := httptest.NewRecorder()
        c := e.NewContext(req, rec)

        if assert.NoError(t, RunGoCode(c)) {
            assert.Equal(t, http.StatusRequestTimeout, rec.Code)
            assert.Contains(t, rec.Body.String(), "Execution timed out after 10 seconds")
        }

        println("Completed Test: TestTimeoutExecution")
    })
}
