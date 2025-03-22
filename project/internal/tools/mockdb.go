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

// var mockBookDetails = map[string]BookDetails{
// 		"1a2b3c4d-5e6f-7g8h-9i0j-k1l2m3n4o5p6": {
// 		BookId: "1a2b3c4d-5e6f-7g8h-9i0j-k1l2m3n4o5p6",
// 		AuthorId: "a1b2c3d4-e5f6-7g8h-9i0j-k1l2m3n4o5p6",
// 		PublisherId: "z9y8x7w6-v5u4-t3s2-r1q0p9o8n7m6",
// 		Title: "To Kill a Mockingbird",
// 		PublicationDate: "1960-07-11",
// 		Isbn: "9780061120084",
// 		Pages: 281,
// 		Genre: "Fiction",
// 		Description: "A novel set in the American South that deals with racial injustice and moral growth.",
// 		Price: 12.99,
// 		Quantity: 7,
//   },
//   "2b3c4d5e-6f7g-8h9i-0j1k-2l3m4n5o6p7": {
// 		BookId: "2b3c4d5e-6f7g-8h9i-0j1k-2l3m4n5o6p7",
// 		AuthorId: "b2c3d4e5-f6g7-8h9i-0j1k-2l3m4n5o6p7",
// 		PublisherId: "y8x7w6v5-u4t3-s2r1-q0p9-o8n7m6l5k4",
// 		Title: "1984",
// 		PublicationDate: "1949-06-08",
// 		Isbn: "9780451524935",
// 		Pages: 328,
// 		Genre: "Dystopian",
// 		Description: "A cautionary tale about a totalitarian society ruled by Big Brother.",
// 		Price: 14.50,
// 		Quantity: 10,
//   },
//   "3c4d5e6f-7g8h-9i0j-1k2l-3m4n5o6p7q8": {
// 		BookId: "3c4d5e6f-7g8h-9i0j-1k2l-3m4n5o6p7q8",
// 		AuthorId: "c3d4e5f6-g7h8-9i0j-1k2l-3m4n5o6p7q8",
// 		PublisherId: "x7w6v5u4-t3s2-r1q0-p9o8-n7m6l5k4j3",
// 		Title: "Pride and Prejudice",
// 		PublicationDate: "1813-01-28",
// 		Isbn: "9780141439518",
// 		Pages: 279,
// 		Genre: "Romance",
// 		Description: "A romantic novel that critiques the British class system of the early 19th century.",
// 		Price: 10.99,
// 		Quantity: 6,
//   },
// 	"4d5e6f7g-8h9i-0j1k-2l3m-4n5o6p7q8r9": {
// 		BookId: "4d5e6f7g-8h9i-0j1k-2l3m-4n5o6p7q8r9",
// 		AuthorId: "d4e5f6g7-h8i9-0j1k-2l3m-4n5o6p7q8r9",
// 		PublisherId: "w6v5u4t3-s2r1-q0p9-o8n7-m6l5k4j3i2",
// 		Title: "To Kill a Mockingbird",
// 		PublicationDate: "1960-07-11",
// 		Isbn: "9780061120084",
// 		Pages: 281,
// 		Genre: "Fiction",
// 		Description: "A novel set in the American South that deals with racial injustice and moral growth.",
// 		Price: 12.99,
// 		Quantity: 7,
// 	},
// 	"5e6f7g8h-9i0j-1k2l-3m4n-5o6p7q8r9s0": {
// 		BookId:"5e6f7g8h-9i0j-1k2l-3m4n-5o6p7q8r9s0" ,
// 		AuthorId: "e5f6g7h8-i9j0-1k2l-3m4n-5o6p7q8r9s0",
// 		PublisherId: "v5u4t3s2-r1q0-p9o8-n7m6-l5k4j3i2h1",
// 		Title: "1984",
// 		PublicationDate: "1949-06-08",
// 		Isbn: "9780451524935",
// 		Pages: 328,
// 		Genre: "Dystopian",
// 		Description: "A cautionary tale about a totalitarian society ruled by Big Brother.",
// 		Price: 14.50,
// 		Quantity: 10,
// 	},
// 	"6f7g8h9i-0j1k-2l3m-4n5o-6p7q8r9s0t1": {
// 		BookId: "6f7g8h9i-0j1k-2l3m-4n5o-6p7q8r9s0t1",
// 		AuthorId: "f6g7h8i9-j0k1-2l3m-4n5o-6p7q8r9s0t1",
// 		PublisherId: "u4t3s2r1-q0p9-o8n7-m6l5-k4j3i2h1g0",
// 		Title: "Pride and Prejudice",
// 		PublicationDate: "1813-01-28",
// 		Isbn: "9780141439518",
// 		Pages: 279,
// 		Genre: "Romance",
// 		Description: "A romantic novel that critiques the British class system of the early 19th century.",
// 		Price: 10.99,
// 		Quantity: 6,
// 	},
// }

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
	err = os.WriteFile("./booksuu.json", updateData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return nil
	}
	return &book
}

func (d *mockDB) setupDatabase() error{
	return nil
}
