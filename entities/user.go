package entities

type User struct {
	ID          int
	NamaLengkap string
	Email       string
	Umur        string
	Password    string
	OTP         string
	StatusOtp   bool
}
