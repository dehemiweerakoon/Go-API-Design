package tools

import(
		log "github.com/sirupsen/logrus"
)

type LoginDetails struct{
	AuthToken string
	Username string
}
type BookDetails struct{
	BookId          string
	AuthorId        string
	PublisherId     string
	Title           string
	PublicationDate string
	Isbn            string
	Pages           int
	Genre           string
	Description     string
	Price float64
	Quantity int
}

type DatabaseInterface interface{
	GetUserLoginDetails(username string) *LoginDetails
	GetBookDetails(bookId string) *BookDetails
	setupDatabase() error
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

