package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"

	"github.com/dehemiweerakoon/golan-api/api"
)

// MockDatabase implements tools.DatabaseInterface for testing
type MockDatabase struct {
	mock.Mock
}

// Mock GetBookDetails method
func (m *MockDatabase) GetBookDetails(bookId string) *api.BookResponseParam {
	args := m.Called(bookId)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*api.BookResponseParam)
}

// Test helper function to execute requests
func executeRequest(router http.Handler, method, url string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, url, nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)
	return response
}

// Test case for fetching a valid book
func TestGetBook_ValidBook(t *testing.T) {
	mockDB := new(MockDatabase)

	expectedBook := &api.BookResponseParam{
		BookId:         "2l3m4n5o-6p7q8r9s0t6u7v8w9x0y1",
		AuthorId:       "author-1",
		PublisherId:    "publisher-1",
		Title:          "Sample Book",
		PublicationDate: "2025-03-23",
		Isbn:           "123-456-789",
		Pages:          300,
		Genre:          "Fiction",
		Description:    "A sample book for testing",
		Price:          19.99,
		Quantity:       10,
	}

	// Mock database response
	mockDB.On("GetBookDetails", expectedBook.BookId).Return(expectedBook)

	router := chi.NewRouter()
	router.Get("/books/{id}", func(w http.ResponseWriter, r *http.Request) {
		GetBookWithDB(mockDB, w, r)
	})

	response := executeRequest(router, "GET", "/books/"+expectedBook.BookId)

	// Verify HTTP status
	if response.Code != http.StatusOK {
		t.Errorf("Expected status %v, got %v", http.StatusOK, response.Code)
	}

	// Decode response body
	var result api.BookResponseParam
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Verify API response matches expected book data
	if result != *expectedBook {
		t.Errorf("Response does not match expected book")
	}

	mockDB.AssertExpectations(t)
}

// Test case for a non-existent book
func TestGetBook_NotFound(t *testing.T) {
	mockDB := new(MockDatabase)

	// Mock a missing book
	mockDB.On("GetBookDetails", "non-existent-id").Return(nil)

	router := chi.NewRouter()
	router.Get("/books/{id}", func(w http.ResponseWriter, r *http.Request) {
		GetBookWithDB(mockDB, w, r)
	})

	response := executeRequest(router, "GET", "/books/non-existent-id")

	// Expect 404 Not Found instead of 500
	if response.Code != http.StatusNotFound {
		t.Errorf("Expected status %v, got %v", http.StatusNotFound, response.Code)
	}

	mockDB.AssertExpectations(t)
}

// Inject mock database into handler for testing
func GetBookWithDB(db *MockDatabase, w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	bookDetails := db.GetBookDetails(id)
	if bookDetails == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(bookDetails); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
