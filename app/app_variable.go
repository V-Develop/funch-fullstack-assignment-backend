package app

type Config struct {
	DataBase   *Database
	ErrorMessage *ErrorMessage
}

var (
	database   *Database
	errorMessage *ErrorMessage
)
