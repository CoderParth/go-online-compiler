package main

import (
	"bytes"
	"fmt"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func main() {
	rate := vegeta.Rate{Freq: 20, Per: time.Second} // 20 requests per second
	duration := 100 * time.Second                      // for 100 seconds
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    "http://localhost:8080/execute", 
		Body:   bytes.NewBufferString(`{"code":"package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello, World!\")\n}"}`).Bytes(),
		Header: map[string][]string{
			"Content-Type": {"application/json"},
		},
	})
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()

	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
	fmt.Printf("Max: %s\n", metrics.Latencies.Max)
	fmt.Printf("Requests/sec: %.2f\n", metrics.Rate)
	fmt.Printf("Total requests: %d\n", metrics.Requests)
	fmt.Printf("Success ratio: %.2f\n", metrics.Success)
}

