# Online Go Compiler - Backend

A simple backend server for an "online Go playground" like service. Receives Go code, compiles it, runs it, and returns the output.
Uses a queue to manage incoming code execution tasks. Each task is processed by a pool of worker goroutines. 
The number of workers is dynamically set based on the number of CPU cores available.

## Results of load test in a Docker environment with 3.75 GB RAM. Around 80 percent was the max CPU usage during the test.
For testing, 20 requests per second were hit to the API for 100 seconds continuously. 
100 percent success for all requests with an average response time of 6.99 seconds for 99 percent of requests.

```
99th percentile: 6.993617825s
Max: 7.0821078s
Requests/sec: 20.01
Total requests: 2000
Success ratio: 1.00
```

## Usage

## Running the Server
## Docker
To run the server using Docker:

Build the image: 
```
docker build -t online-compiler .
```

Run the container: 
```
docker run -p 8080:8080 online-compiler
```

## Locally
Ensure you have Go installed.
Run 
```
go run server.go 
```
from the project root.

## Execute Code
To execute code, send a POST request to `/execute` with the following JSON payload:

```json
{
    "code": "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello, From GO\")\n}"
}
```

## To Perform Load Test

Navigate to loadtest dir:

```
cd loadtest
```
Ensure server is running. Then run:

```
go run loadtest.go
```
