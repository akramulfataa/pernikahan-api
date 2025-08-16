package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Struct data
type Uang struct {
	Penghasilan float64 `json:"penghasilan"`
	Makan       float64 `json:"makan"`
	Jajan       float64 `json:"jajan"`
	Rokok       float64 `json:"rokok"`
	Ngopi       float64 `json:"ngopi"`
	Ngedate     float64 `json:"ngedate"`
}

type UangRequest struct {
	Penghasilan float64 `json:"penghasilan"`
	Makan       float64 `json:"makan"`
	Jajan       float64 `json:"jajan"`
	Rokok       float64 `json:"rokok"`
	Ngopi       float64 `json:"ngopi"`
	Ngedate     float64 `json:"ngedate"`
}

type Response struct {
	TotalPengeluaran float64            `json:"total_pengeluaran"`
	SisaUang         float64            `json:"sisa_uang"`
	Alokasi          map[string]float64 `json:"alokasi"`
}

// Validator
func validateInputOutput(req *UangRequest) error {
	if req.Penghasilan <= 0 {
		return fmt.Errorf("penghasilan harus lebih dari 0")
	}
	if req.Makan < 0 {
		return fmt.Errorf("makan tidak boleh negatif")
	}
	if req.Jajan < 0 {
		return fmt.Errorf("jajan tidak boleh negatif")
	}
	if req.Rokok < 0 {
		return fmt.Errorf("rokok tidak boleh negatif")
	}
	if req.Ngopi < 0 {
		return fmt.Errorf("ngopi tidak boleh negatif")
	}
	if req.Ngedate < 0 {
		return fmt.Errorf("ngedate tidak boleh negatif")
	}
	return nil
}

// Hitung total bulanan
func GetTotalBulanan(uang *Uang) float64 {
	return uang.Makan + uang.Jajan + uang.Rokok + uang.Ngopi + uang.Ngedate
}

// Handler API
func HitungHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var req UangRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validasi input
	if err := validateInputOutput(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Buat objek uang
	uang := &Uang{
		Penghasilan: req.Penghasilan,
		Makan:       req.Makan,
		Jajan:       req.Jajan,
		Rokok:       req.Rokok,
		Ngopi:       req.Ngopi,
		Ngedate:     req.Ngedate,
	}

	totalBulanan := GetTotalBulanan(uang)
	sisaUang := uang.Penghasilan - totalBulanan

	// Alokasi 20% sisa uang ke 4 kategori
	alokasiPersen := 0.20
	alokasi := map[string]float64{
		"asset":    sisaUang * alokasiPersen * 0.25,
		"nikah":    sisaUang * alokasiPersen * 0.25,
		"keluarga": sisaUang * alokasiPersen * 0.25,
		"sedekah":  sisaUang * alokasiPersen * 0.25,
	}

	resp := Response{
		TotalPengeluaran: totalBulanan,
		SisaUang:         sisaUang,
		Alokasi:          alokasi,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/api/hitung", HitungHandler)
	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
