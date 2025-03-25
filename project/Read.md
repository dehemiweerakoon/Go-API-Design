# Go Project with API Run Instructions

## Prerequisites

- Go (version 1.16 or higher)
- Git
- A terminal or command prompt

## Setup

1. **Install dependencies:**
    ```sh
    go mod tidy
    ```

## Running the API

1. **Build the project:**
    ```sh
    go build -o main ./cmd/api/main.go
    ```

2. **Run the executable:**
    ```sh
   go run cmd/api/main.go
    ```

3. **Access the API:**
    Open your browser or API client (like Postman) and navigate to `http://localhost:9000`.

## Testing

1. **Run tests:**
    ```sh
    go test -v ./internal/handlers 
    ```

## Directory Structure

```
yourproject/
├──api/api.go
│  
├── cmd/api
│   └── main.go
├── internal/
│   ├── handlers/
│   └── middleware/
│   └── service/
├── pkg/
│   └── tools/
├── go.mod
├── go.sum
├── book.json
└── README.md
```

## Docker Instructions

1. **Build the Docker image:**
    ```sh
    docker build -t go-app:latest .
    ```

2. **Run the Docker container:**
    ```sh
    docker run -p 9000:9000 go-app:latest
    ```

## Running On the Kubernetes 

1. **Build Docker image inside the minikube**
    ```
    minikube docker-env | Invoke-Expression
    docker build -t go-api .
    ```
2. **Apply the changes**

    ```
    kubectl apply -f deployment.yaml
    kubectl apply -f service.yaml
    ```
3. **Try port forwarding again**
    ```
    kubectl port-forward service/myservice 9000:80
    ```
4. **Access Via your web browser**
    ```
    http://localhost:9000/books
    ```

## Contributing

1. Fork the repository
2. Create a new branch (`git checkout -b feature-branch`)
3. Commit your changes (`git commit -am 'Add new feature'`)
4. Push to the branch (`git push origin feature-branch`)
5. Create a new Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.