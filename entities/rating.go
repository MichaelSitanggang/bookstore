package entities

type Rating struct {
	ID     int
	UserID int
	BookID int
	Ulasan string
	Rating int
	User   User
	Book   Book
}
