package types

type DailyQuote struct {
	Author  	string		`json:"author"`
	Quote		string		`json:"quote"`
	DateWritten	string		`json:"dateWritten"`
}

type User struct {
	Name 		string		`json:"name"`
	Email		string		`json:"email"`
	Password	string		`json:"password"`
}