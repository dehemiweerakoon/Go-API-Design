package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"sync"
)

// Book struct representing a book with a title and description
type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Sample book data
var books = []Book{
	{1, "Go Programming", "A beginner's guide to Go language"},
	{2, "Advanced Go", "Deep dive into Go concurrency and best practices"},
	{3, "Golang Web Development", "Learn to build web apps with Golang"},
	{4, "Python Basics", "Introduction to Python programming"},
	{5, "Concurrency in Go", "Mastering goroutines and channels"},
	{6, "Java for Beginners", "Learn Java from scratch"},
	{7, "Machine Learning with Go", "Implement ML models using Go"},
	{8, "The Art of Go", "Exploring the beauty of Go programming"},
}

// searchBooks function that runs concurrently
func searchBooks(keyword string, books []Book, results chan<- []Book, wg *sync.WaitGroup) {
	defer wg.Done() // Mark goroutine as done
	var matchedBooks []Book

	keyword = strings.ToLower(keyword) // Convert to lowercase for case-insensitive search
	for _, book := range books {
		// Convert title and description to lowercase and check for keyword match
		if strings.Contains(strings.ToLower(book.Title), keyword) ||
			strings.Contains(strings.ToLower(book.Description), keyword) {
			matchedBooks = append(matchedBooks, book)
		}
	}

	results <- matchedBooks // Send results to the channel
}

// searchHandler processes search requests
func searchHandler(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("q") // Get search keyword from query parameter
	if keyword == "" {
		http.Error(w, "Missing search keyword", http.StatusBadRequest)
		return
	}

	numWorkers := runtime.NumCPU() // Get the number of CPU cores
	chunkSize := (len(books) + numWorkers - 1) / numWorkers // Divide books among workers

	results := make(chan []Book, numWorkers) // Channel to collect results
	var wg sync.WaitGroup

	// Create multiple goroutines for parallel processing
	for i := 0; i < len(books); i += chunkSize {
		end := i + chunkSize
		if end > len(books) {
			end = len(books)
		}

		wg.Add(1) // Increase WaitGroup counter
		go searchBooks(keyword, books[i:end], results, &wg)
	}

	// Close the channel after all goroutines finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results from channel
	var finalResults []Book
	for res := range results {
		finalResults = append(finalResults, res...)
	}

	// Send response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(finalResults)
}

func main() {
	http.HandleFunc("/books/search", searchHandler)

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
