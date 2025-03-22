package main

import (
    "encoding/json"
    "fmt"
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
  

    var details = map[string]int{"a": 1, "b": 8}
    details["a"]=55
    fmt.Println(details)
}
