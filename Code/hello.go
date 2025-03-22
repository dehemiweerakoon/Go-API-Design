package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
)

type BookBalanceParam struct {
    AccountID string `json:"account_id"`
    Currency  string `json:"currency"`
}

func readJSONFromFile(filePath string) (*BookBalanceParam, error) {
    // Read the content of the file using os.ReadFile
    fileContent, err := os.ReadFile(filePath)
    if err != nil {
        return nil, err
    }

    // Declare a variable to hold the decoded data
    var params BookBalanceParam

    // Unmarshal the JSON data into the struct
    err = json.Unmarshal(fileContent, &params)
    if err != nil {
        return nil, err
    }

    return &params, nil
}

func main() {
    filePath := "books.json"

    // Read and parse the JSON file
    params, err := readJSONFromFile(filePath)
    if err != nil {
        log.Fatalf("Error reading JSON file: %v", err)
    }

    // Print the parsed result
    fmt.Println(params)
}
