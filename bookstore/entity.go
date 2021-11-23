package bookstore

type PostBook struct {
	BookID       string `json:"bookID"`
	BookName string `json:"bookName"`
	Category        string `json:"category"`
	Author      string `json:"author"`
}