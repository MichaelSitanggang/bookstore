package entities

type Rating struct {
	ID     int
	UserID int
	BookID int
	Rating int
	Ulasan string
	User   User
}

//buat rating besok
//payment gateway besok
