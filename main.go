package main

import (
	"fmt"
)

// note projek :
// - inputkan kebutuhan sehari hari seperti uang makan, rokok kalau ada, jajan, ngopi dan lain lain yang sifatnya komsumtif
// - setelah itu sisa gaji nya baru di bagikan ke kebutuhan non komsuftif seperti biaya pernikahan, asset, keluarga, dan lain lain
// - dari penghasilan di kuragi semua penggeluaran konsumtif perhari kemudian di cari sisa nya baru di bagi persennya

type Uang struct {
	Penghasilan float64 `json:"penghasilan"`
	Makan float64	`json:"makan"`
	Jajan float64 `json:"jajan"` 
	Rokok float64	`json:"rokok"`
	Ngopi float64 `json:"ngopi"`
	Ngedate float64 `json:"ngedate"`
}

type UangRequest struct {
	Penghasilan float64 `json:"penghasilan"`
	Makan float64	`json:"makan"`
	Jajan float64 `json:"jajan"` 
	Rokok float64	`json:"rokok"`
	Ngopi float64 `json:"ngopi"`
	Ngedate float64 `json:"ngedate"`
}

func New(req *UangRequest) (*Uang, error) {

	if err := validateInputOuput(req); err != nil {
		return nil, err
	}


	uang := &Uang{
		Penghasilan: req.Penghasilan,
		Makan:       req.Makan,
		Jajan:       req.Jajan,
		Rokok:       req.Rokok,
		Ngopi:       req.Ngopi,
		Ngedate:     req.Ngedate,
	}

	return uang, nil
	
}


// Fungsi untuk menghitung total pengeluaran konsumtif bulanan
func GetTotalBulanan(uang *Uang) float64 {
	total := uang.Makan + uang.Jajan + uang.Rokok + uang.Ngopi + uang.Ngedate
	return total
}

func validateInputOuput(req *UangRequest) error {
	if req.Penghasilan != 0 && req.Penghasilan <= 0 {
		fmt.Errorf("penghasilan is required")
	}
	if req.Makan != 0 && req.Makan <= 0 {
		fmt.Errorf("penghasilan is required")
	}
	if req.Jajan != 0 && req.Jajan <= 0 {
		fmt.Errorf("penghasilan is required")
	}
	if req.Rokok != 0 && req.Rokok <= 0 {
		fmt.Errorf("penghasilan is required")
	}
	if req.Ngopi != 0 && req.Ngopi <= 0 {
		fmt.Errorf("penghasilan is required")
	}
	if req.Ngedate != 0 && req.Ngedate <= 0 {
		fmt.Errorf("penghasilan is required")
	}
	return nil
}

func main(){
	// Contoh input
	req := &UangRequest{
		Penghasilan: 10000000,
		Makan:       2000000,
		Jajan:       500000,
		Rokok:       300000,
		Ngopi:       400000,
		Ngedate:     600000,
	}

	// Inisialisasi objek Uang
	uang, err := New(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Hitung total pengeluaran konsumtif bulanan
	totalBulanan := GetTotalBulanan(uang)
	fmt.Printf("Total Pengeluaran Konsumtif Bulanan: %.2f\n", totalBulanan)

	// Hitung sisa uang setelah pengeluaran konsumtif
	sisaUang := uang.Penghasilan - totalBulanan
	fmt.Printf("Sisa Uang Setelah Pengeluaran Konsumtif: %.2f\n", sisaUang)

	// Hitung alokasi ke berbagai aspek
	alokasiPersen := 0.20
	asset := sisaUang * alokasiPersen * 0.25
	nikah := sisaUang * alokasiPersen * 0.25
	keluarga := sisaUang * alokasiPersen * 0.25
	sedekah := sisaUang * alokasiPersen * 0.25

	fmt.Printf("Alokasi Asset: %.2f\n", asset)
	fmt.Printf("Alokasi Nikah: %.2f\n", nikah)
	fmt.Printf("Alokasi Keluarga: %.2f\n", keluarga)
	fmt.Printf("Alokasi Sedekah: %.2f\n", sedekah)
}
