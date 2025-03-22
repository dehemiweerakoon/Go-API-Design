package tools

import(
		log "github.com/sirupsen/logrus"
)

type LoginDetails struct{
	AuthToken string
	Username string
}
type BookDetails struct{
	BookId         string  `json:"BookId"`
	AuthorId       string  `json:"AuthorId"`
	PublisherId    string  `json:"PublisherId"`
	Title          string  `json:"Title"`
	PublicationDate string  `json:"PublicationDate"`
	Isbn           string  `json:"Isbn"`
	Pages          int     `json:"Pages"`
	Genre          string  `json:"Genre"`
	Description    string  `json:"Description"`
	Price          float64 `json:"Price"`
	Quantity       int     `json:"Quantity"`
}

type DatabaseInterface interface{
	GetUserLoginDetails(username string) *LoginDetails
	GetBookDetails(bookId string) *BookDetails
	setupDatabase() error
	GetAllBookDetails() map[string]BookDetails
	SaveBookDetails(book BookDetails)*BookDetails
	UpdateBookDetails(book BookDetails,bookId string)*BookDetails
}
func NewDatabase()(*DatabaseInterface,error){
	var database DatabaseInterface = &mockDB{}

	var err error = database.setupDatabase()

	if err != nil{
		log.Error(err)
		return nil,err
	}
	return &database,nil
}

