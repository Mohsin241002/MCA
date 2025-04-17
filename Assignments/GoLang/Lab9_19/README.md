# Game Info Fetcher - Go Concurrency Demo

This project demonstrates Go concurrency concepts using goroutines and channels in the context of a game information fetcher application. The web application allows users to fetch game information both sequentially and concurrently to visualize the performance benefits of concurrency.

## Features

- Interactive web interface with a game selection form
- Sequential and concurrent API endpoints for fetching game data
- Visual demonstration of the time differences between sequential and concurrent execution
- Input validation for all requests
- Realistic simulation of network delays to showcase concurrency benefits
- Explanations of Go concurrency concepts (goroutines, channels, wait groups)

## Concepts Demonstrated

1. **Goroutines**: Lightweight threads managed by the Go runtime that enable concurrent execution
2. **Channels**: Communication pipes between goroutines for data exchange
3. **Wait Groups**: Synchronization primitives to wait for a collection of goroutines to finish
4. **Input Validation**: Request validation to ensure integrity of data processing
5. **Error Handling**: Proper error propagation and handling in concurrent contexts

## Getting Started

### Prerequisites

- Go 1.16 or higher

### Running the Application

1. Clone this repository
2. Navigate to the project directory
3. Run the application:

```bash
go run main.go
```

4. Open a web browser and navigate to `http://localhost:8080`

## Usage

1. Select the games you want to fetch information for by checking the checkboxes
2. Click either "Fetch Sequentially" or "Fetch Concurrently" to retrieve the data
3. Compare the execution times displayed at the top of the results section
4. Observe how concurrent fetching is significantly faster than sequential fetching

## How It Works

### Sequential Fetching

- Games are fetched one by one in sequence
- Each game fetch completes before the next one starts
- Total time is the sum of all individual fetch times

### Concurrent Fetching

- A goroutine is created for each game fetch
- All game fetches happen concurrently
- Results are sent back through a channel
- A wait group ensures all goroutines complete
- Total time is approximately the duration of the slowest fetch

## Code Structure

- `main.go`: The main Go application with server setup and concurrent implementations
- `templates/index.html`: The HTML template for the web interface
- `static/css/styles.css`: CSS styles for the application
- `static/js/app.js`: Client-side JavaScript for interacting with the API and updating the UI 