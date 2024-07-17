# Network_Connections
# Go Application with Docker

## Description

This project is a Go application that fetches JSON data from a public URL and processes it to determine connections between locations.

## Prerequisites

- Docker
- (Optional) Docker Compose

## Building and Running with Docker

### Using Docker

1. Build the Docker image:
    ```bash
    docker build -t app .
    ```

2. Run the Docker container:
    ```bash
    docker run -p 8080:8080 app
    ```

3. Access the application at `http://localhost:8080/connections`.

### Using Docker Compose (Optional)

1. Ensure you have `docker-compose.yml` in your project directory.

2. Build and run the application using Docker Compose, logging the output to a file:
    ```bash
    docker-compose up --build > docker-compose.log 2>&1
    ```

3. Access the application at `http://localhost:8080/connections`.

4. View the logs:
    ```bash
    cat docker-compose.log
    ```

## Running Tests (Optional)

If you want to run tests, you can do so inside the Docker container.

1. Build the Docker image:
    ```bash
    docker build -t main .
    ```

2. Run the tests inside the Docker container:
    ```bash
    docker run main go test ./... -v
    ```
