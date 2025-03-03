package entities

type Rating struct {
	ID     int
	UserID int
	BookID int
	Rating int
	Ulasan string
	User   User
}

//payment gateway besok
//fitur admin
