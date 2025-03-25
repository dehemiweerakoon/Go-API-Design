package tools

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
)

type mockDB struct{
}

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
var bookMap map[string]BookDetails 
//var api_Key = os.Getenv("API_KEY")

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
		fileContent, err := os.ReadFile("book.json")
		if err != nil {
			return nil
		}

		// Create a map to hold books indexed by their BookId
		var books map[string]BookDetails
		err = json.Unmarshal(fileContent, &books)
		if err != nil {
			fmt.Println("Error Unmarshalling JSON:", err)
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

func (d *mockDB) UpdateBookDetails(book BookDetails,bookId string)*BookDetails{
	fileContent,err :=os.ReadFile("./book.json")
	if err!=nil{
		return nil
	}
	var existingBooks map[string]BookDetails
	err = json.Unmarshal(fileContent,&existingBooks)
	if err!=nil{
		return nil
	}
	existingBooks[book.BookId]=book
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
func (d *mockDB) DeleteBookDetails(bookId string)bool{
// if the data is deleted return true anything other than that returns false .....
    fileContent,err :=os.ReadFile("./book.json")
	if err!=nil{
		return false
	}
	var existingBooks map[string]BookDetails
	err = json.Unmarshal(fileContent,&existingBooks)
    if err!=nil{
		return false
	}

	_, exists := existingBooks[bookId]; 
	if !exists {
		//fmt.Println(exists,val)
		return false
	}

	delete(existingBooks,bookId) // delate the element in this 
	updateData ,err := json.MarshalIndent(existingBooks,""," ")
	if err!=nil{
		fmt.Println("error updating file")
	}
	err = os.WriteFile("./book.json", updateData, 0644) // 6- owner  4-group 4-other
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return  false
	}
	return true
}
func LoadBooks (filename string) error{
	data,err := os.ReadFile("./book.json")
	if err!=nil{
		return err
	}
	return json.Unmarshal(data,&bookMap)
}

func (d *mockDB) SearchHandler(keyword string)[]BookDetails{
	LoadBooks("./book.json")
	booked := make([]BookDetails, 0, len(bookMap)) // Preallocate slice capacity
	for _, book := range bookMap {
		booked = append(booked, book)
	}
	
	numWorkers:=runtime.NumCPU() // number of CPU 3 10/4 --> 12
	chunkSize :=(len(booked) +numWorkers-1)/numWorkers

	results := make(chan []BookDetails,numWorkers) //channel to collect results 
	var wg sync.WaitGroup
	// create multiple goroutines 
	for i:=0;i<len(booked);i+=chunkSize{
		end :=i+chunkSize
		if end>len(booked){
			end = len(booked)
		}
		wg.Add(1)
		go SearchBooks(keyword,booked[i:end],results,&wg)
	}

	go func ()  {
		wg.Wait()
		close(results)
	}()
	var finalResults []BookDetails
	for res := range results{
		finalResults = append(finalResults, res...)
	}
	return finalResults
}
func SearchBooks(keyword string ,books []BookDetails,results chan<- []BookDetails,wg *sync.WaitGroup){
	defer wg.Done()
	var matchedBooks []BookDetails

	keyword = strings.ToLower(keyword)
	for _,book :=range books{
		if strings.Contains(strings.ToLower(book.Title),strings.ToLower(keyword)) || strings.Contains(strings.ToLower(book.Description),strings.ToLower(keyword)){
			matchedBooks = append(matchedBooks, book)
		}
	}
	results <-matchedBooks
}

func (d *mockDB) setupDatabase() error{
	return nil
}
