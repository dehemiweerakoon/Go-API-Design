FROM golang:1.24-alpine

WORKDIR /app
ENV GOPATH=/go
ENV GOBIN=$GOPATH/bin
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o main ./cmd/api/main.go
EXPOSE 9000
CMD ["./main"]
