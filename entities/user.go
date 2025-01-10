package entities

type User struct {
	ID          int
	NamaLengkap string
	Email       string
	Umur        int
	Password    string
	OTP         string
	StatusOtp   bool
}
