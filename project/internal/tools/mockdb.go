package tools

import (
	"encoding/json"
	"fmt"
	"os"
)

type mockDB struct{}

var mockLoginDetails =map[string] LoginDetails{
	"alex" :{
		AuthToken: "123ABC",
		Username: "alex",
	},
	"jason":{
		AuthToken: "456DEF",
		Username: "jason",
	},
	"marie":{
		AuthToken: "789GHI",
		Username: "marie",
	}, 
}


func(d *mockDB) GetUserLoginDetails(username string)*LoginDetails{
			var data = LoginDetails{}

			data ,ok:=mockLoginDetails[username]

			if !ok {
				return nil
			}
			return &data
		}

func (d *mockDB) GetBookDetails(BookId string)*BookDetails{
		// Read the content of the JSON file
		fileContent, err := os.ReadFile("./book.json")
		if err != nil {
			return nil
		}

		// Create a map to hold books indexed by their BookId
		var books map[string]BookDetails
		err = json.Unmarshal(fileContent, &books)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return nil
		}

		// Fetch the book by BookId
		if book, exists := books[BookId]
		exists {
			// Return the book details
			return &book
		}

		// Return nil if the book is not found
		return nil
}

func (d *mockDB) GetAllBookDetails() map[string]BookDetails {
	fileContent, err := os.ReadFile("./book.json")
	if err != nil {
		return nil
	}

	// Create a map to hold books indexed by their BookId
	var books map[string]BookDetails
	err = json.Unmarshal(fileContent, &books)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil
	}

	return books
}

func (d *mockDB) SaveBookDetails(book BookDetails)*BookDetails{
	fileContent,err := os.ReadFile("./book.json")
	if err !=nil{
		return nil
	}
    var existingBooks map[string]BookDetails
	err = json.Unmarshal(fileContent, &existingBooks)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil
	}
	existingBooks[book.BookId] = book
	updateData ,err := json.MarshalIndent(existingBooks,""," ")
	if err!=nil{
		fmt.Println("error updating file")
	}
	err = os.WriteFile("./book.json", updateData, 0644) // 6- owner  4-group 4-other
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return nil
	}
	return &book
}

func (d *mockDB) setupDatabase() error{
	return nil
}
