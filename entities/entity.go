package entities

type Book struct {
	BookID      int64   `json:"bookID"`
	BookName    string  `json:"bookName"`
	Author      string  `json:"author"`
	ReviewScore float64 `json:"reviewScore"`
	SoldCount   int64   `json:"soldCount"`
}
