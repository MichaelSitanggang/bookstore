package entities

type Book struct {
	ID        int
	Gambar    string
	Judul     string
	Author    string
	Year      int
	Penjualan int
	Harga     float64
	Review    float64
	Stok      int
	Reviews   []Rating
}
