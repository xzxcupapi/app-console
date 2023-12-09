package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Since2024."
	dbname   = "dbenigma"
)

// Function Main
func main() {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	mainMenu(db)
}

// Function Main Menu
func mainMenu(db *sql.DB) {
	for {
		fmt.Println("\nMain Menu:")
		fmt.Println("1. Menu Customer")
		fmt.Println("2. Menu Layanan")
		fmt.Println("3. Menu Transaksi Laundry")
		fmt.Println("0. Exit")
		fmt.Println("Pilih Menu dari 1-3")

		var choice int
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			customerMenu(db)
		case 2:
			serviceMenu(db)
		case 3:
			laundryTransactionMenu(db)
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Pilihan salah, coba kembali!")
		}
	}
}

// Function Customer Menu
func customerMenu(db *sql.DB) {
	for {
		fmt.Println("\nCustomer Menu:")
		fmt.Println("1. Lihat Customer")
		fmt.Println("2. Masukan Customer")
		fmt.Println("3. Edit Customer")
		fmt.Println("4. Hapus Customer")
		fmt.Println("0. Kembali ke Menu utama")

		var choice int
		fmt.Print("Pilih menu: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			viewCustomer(db)
		case 2:
			insertCustomer(db)
		case 3:
			updateCustomer(db)
		case 4:
			deleteCustomer(db)
		case 0:
			return
		default:
			fmt.Println("Inputan Salah")
		}
	}
}

// Function Service Menu
func serviceMenu(db *sql.DB) {
	for {
		fmt.Println("\nMenu Layanan :")
		fmt.Println("1. Lihat Layanan")
		fmt.Println("2. Masukan Layanan")
		fmt.Println("3. Edit Layanan")
		fmt.Println("4. Hapus Layanan")
		fmt.Println("0. Kembali ke Menu utama")

		var choice int
		fmt.Print("Pilih menu : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			viewService(db)
		case 2:
			insertService(db)
		case 3:
			updateService(db)
		case 4:
			deleteService(db)
		case 0:
			return
		default:
			fmt.Println("Inputan salah, coba lagi!")
		}
	}
}

// Function Laundry Transaction Menu
func laundryTransactionMenu(db *sql.DB) {
	for {
		fmt.Println("\nMenu Transaksi :")
		fmt.Println("1. Lihat Transaksi")
		fmt.Println("2. Masukan Transaksi")
		fmt.Println("0. Kembali ke Menu utama")

		var choice int
		fmt.Print("Pilih menu : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			viewTsx(db)
		case 2:
			insertTsx(db)
		case 0:
			return
		default:
			fmt.Println("Inputan salah, coba lagi!")
		}
	}
}

// FUNCTION CUSTOMER
// Function VIEW Customer
func viewCustomer(db *sql.DB) {
	fmt.Println("\nView Customer")

	rows, err := db.Query("SELECT custid, namacust, nohp FROM customer")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Print header
	fmt.Printf("%-10s %-20s %-15s\n", "CustomerID", "Nama Customer", "No HP")
	fmt.Println("------------------------------------------------------------")

	// Iterasi dan print data customer
	for rows.Next() {
		var custid int
		var namacust, nohp string
		if err := rows.Scan(&custid, &namacust, &nohp); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%-10d %-20s %-15s\n", custid, namacust, nohp)
	}

	// Cek error dari iterasi rows
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

// Function Insert Customer
func insertCustomer(db *sql.DB) {
	fmt.Println("\nInsert Customer")

	var (
		custid   int
		namacust string
		nohp     string
	)

	//validasi customer id
	fmt.Print("Customer ID: ")
	fmt.Scan(&custid)

	// Validasi customer id tidak boleh kosong
	if custid == 0 {
		fmt.Println("Customer ID tidak boleh kosong.")
		return
	}

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM customer WHERE custid = $1", custid).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	if count > 0 {
		fmt.Println("Customer ID sudah ada.")
		return
	}

	//validasi nama customer
	fmt.Print("Nama Customer: ")
	fmt.Scan(&namacust)

	// Validasi nama tidak boleh kosong
	if strings.TrimSpace(namacust) == "" {
		fmt.Println("Nama Customer tidak boleh kosong.")
		return
	}

	//validasi nomer hp
	fmt.Print("No HP: ")
	fmt.Scan(&nohp)

	// Validasi nohp tidak boleh kosong
	if strings.TrimSpace(nohp) == "" {
		fmt.Println("No HP pelanggan tidak boleh kosong.")
		return
	}

	// Insert data ke database
	_, err = db.Exec("INSERT INTO customer (custid, namacust, nohp) VALUES ($1, $2, $3)", custid, namacust, nohp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data pelanggan berhasil ditambahkan.")
}

// Function UPDATE Customer
func updateCustomer(db *sql.DB) {
	fmt.Println("\nUpdate Customer")

	var (
		custid      int
		newNamacust string
		newNohp     string
	)

	// Input data customer yang akan diupdate
	fmt.Print("Customer ID yang akan diupdate: ")
	fmt.Scan(&custid)

	// Validasi apakah ID Customer valid (sudah terdaftar di customer)
	var customerExist bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM customer WHERE custid = $1)", custid).Scan(&customerExist)
	if err != nil {
		log.Fatal(err)
	}
	if !customerExist {
		fmt.Println("Customer ID tidak valid.")
		return
	}

	// Input Nama baru
	fmt.Print("Nama Customer baru: ")
	fmt.Scan(&newNamacust)

	// Validasi nama tidak boleh kosong
	if strings.TrimSpace(newNamacust) == "" {
		fmt.Println("Nama Customer baru tidak boleh kosong.")
		return
	}
	//Input No HP baru
	fmt.Print("No HP baru: ")
	fmt.Scan(&newNohp)

	// Validasi nohp tidak boleh kosong
	if strings.TrimSpace(newNohp) == "" {
		fmt.Println("No HP baru pelanggan tidak boleh kosong.")
		return
	}

	// Update data di database
	_, err = db.Exec("UPDATE customer SET namacust = $1, nohp = $2 WHERE custid = $3", newNamacust, newNohp, custid)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data pelanggan berhasil diupdate.")
}

// Function DELETE Customer
func deleteCustomer(db *sql.DB) {
	fmt.Println("\nDelete Customer")

	var custid int

	// Input data customer yang akan dihapus
	fmt.Print("Customer ID yang akan dihapus: ")
	fmt.Scan(&custid)

	// Validasi apakah ID Customer valid (sudah terdaftar di customer)
	var customerExist bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM customer WHERE custid = $1)", custid).Scan(&customerExist)
	if err != nil {
		log.Fatal(err)
	}
	if !customerExist {
		fmt.Println("Customer ID tidak valid.")
		return
	}

	// Hapus data dari database
	_, err = db.Exec("DELETE FROM customer WHERE custid = $1", custid)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data pelanggan berhasil dihapus.")
}

//FUNCTION CUSTOMER END

// FUNCTION SERVICE
// function View Service
func viewService(db *sql.DB) {
	fmt.Println("\nView Service")

	rows, err := db.Query("SELECT serviceid, pelayanan, satuan, harga FROM service")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Print header
	fmt.Printf("%-10s %-20s %-20s %-20s\n", "ServiceID", "Pelayanan", "Satuan", "Harga")
	fmt.Println("------------------------------------------------------------")

	// Iterasi dan print data service
	for rows.Next() {
		var serviceid, harga int
		var pelayanan, satuan string
		if err := rows.Scan(&serviceid, &pelayanan, &satuan, &harga); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%-10d %-20s %-20s %-20d\n", serviceid, pelayanan, satuan, harga)
	}

	// Cek error dari iterasi rows
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

// Fucntion Insert Service
func insertService(db *sql.DB) {
	fmt.Println("\nInsert Service")

	var (
		seviceid  int
		pelayanan string
		satuan    string
		harga     int
	)

	//validasi service id
	fmt.Print("Service ID: ")
	fmt.Scan(&seviceid)

	// Validasi service id tidak boleh kosong
	if seviceid == 0 {
		fmt.Println("Service ID tidak boleh kosong.")
		return
	}

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM service WHERE serviceid = $1", seviceid).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	if count > 0 {
		fmt.Println("Service ID sudah ada.")
		return
	}

	//validasi pelayanan customer
	fmt.Print("Pelayanan: ")
	fmt.Scan(&pelayanan)

	// Validasi pelayanan tidak boleh kosong
	if strings.TrimSpace(pelayanan) == "" {
		fmt.Println("Pelayanan tidak boleh kosong.")
		return
	}
	//validasi pelayanan end

	//validasi satuan
	fmt.Print("Satuan : ")
	fmt.Scan(&satuan)

	// Validasi satuan tidak boleh kosong
	if strings.TrimSpace(satuan) == "" {
		fmt.Println("Satuan tidak boleh kosong.")
		return
	}

	fmt.Print("Harga : ")
	fmt.Scan(&harga)

	// Validasi harga tidak boleh kosong
	if harga == 0 {
		fmt.Println("Harga tidak boleh kosong.")
		return
	}

	// Insert data ke database
	_, err = db.Exec("INSERT INTO service (serviceid, pelayanan, satuan, harga) VALUES ($1, $2, $3, $4)", seviceid, pelayanan, satuan, harga)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data pelanggan berhasil ditambahkan.")
}

// function upodate Service
func updateService(db *sql.DB) {
	fmt.Println("\nUpdate Service")

	var (
		serviceid    int
		newPelayanan string
		newSatuan    string
		newHarga     int
	)

	// Input data pelanggan yang akan diupdate
	fmt.Print("Service ID yang akan diupdate : ")
	fmt.Scan(&serviceid)

	// Validasi apakah ID Pelanggan valid (sudah terdaftar di customer)
	var serviceExist bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM service WHERE serviceid = $1)", serviceid).Scan(&serviceExist)
	if err != nil {
		log.Fatal(err)
	}
	if !serviceExist {
		fmt.Println("Service ID tidak valid.")
		return
	}

	// Input data baru
	fmt.Print("Pelayanan baru : ")
	fmt.Scan(&newPelayanan)

	// Validasi nama tidak boleh kosong
	if strings.TrimSpace(newPelayanan) == "" {
		fmt.Println("Pelayanan baru tidak boleh kosong.")
		return
	}

	fmt.Print("Satuan baru: ")
	fmt.Scan(&newSatuan)

	// Validasi satuan tidak boleh kosong
	if strings.TrimSpace(newSatuan) == "" {
		fmt.Println("Satuan baru tidak boleh kosong.")
		return
	}

	// Validasi harga tidak boleh kosong
	fmt.Print("Harga baru: ")
	fmt.Scan(&newHarga)

	if newHarga == 0 {
		fmt.Println("Harga baru tidak boleh kosong.")
		return
	}

	// Update data di database
	_, err = db.Exec("UPDATE service SET pelayanan = $1, satuan = $2, harga = $3 WHERE serviceid = $4", newPelayanan, newSatuan, newHarga, serviceid)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data Pelayanan berhasil diupdate.")
}

// function delete Service
func deleteService(db *sql.DB) {
	fmt.Println("\nDelete Service")

	var serviceid int

	// Input data pelanggan yang akan dihapus
	fmt.Print("Service ID yang akan dihapus: ")
	fmt.Scan(&serviceid)

	// Validasi apakah ID Pelanggan valid (sudah terdaftar di customer)
	var customerExist bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM service WHERE serviceid = $1)", serviceid).Scan(&customerExist)
	if err != nil {
		log.Fatal(err)
	}
	if !customerExist {
		fmt.Println("Service ID tidak valid.")
		return
	}

	// Hapus data dari database
	_, err = db.Exec("DELETE FROM service WHERE serviceid = $1", serviceid)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data Layanan berhasil dihapus.")
}

//FUNCTION SERVICE END

// FUNCTION TSX
// FUNCTION VIEW TSX LAUNDRY
func viewTsx(db *sql.DB) {
	fmt.Println("\nView Transaction Laundry")

	rows, err := db.Query("SELECT transaksiid, nonota, tanggalmasuk, tanggalselesai, diterimaoleh, custid FROM laundrytransaction")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Print header
	fmt.Printf("%-15s %-10s %-15s %-15s %-15s %-15s\n", "TransaksiID", "NoNota", "TanggalMasuk", "TanggalSelesai", "Diterima Oleh", "CustID")
	fmt.Println("---------------------------------------------------------------------------------------")

	// Iterasi dan print data service
	for rows.Next() {
		var transaksiid, nonota, custid int
		var diterimaoleh string
		var tanggalmasuk, tanggalselesai time.Time
		if err := rows.Scan(&transaksiid, &nonota, &tanggalmasuk, &tanggalselesai, &diterimaoleh, &custid); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%-15d %-10d %-15v %-15v %-15s %-15d\n", transaksiid, nonota, tanggalmasuk.Format("02-02-2006"), tanggalselesai.Format("02-02-2006"), diterimaoleh, custid)
	}

	// Cek error dari iterasi rows
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

// FUNCTION INSERT TSX LAUNDRY
func insertTsx(db *sql.DB) {
	fmt.Println("\nInsert Transaksi Laundry")

	var (
		transaksid     int
		nonota         int
		tanggalmasuk   time.Time
		tanggalselesai time.Time
		diterimaoleh   string
		custid         int
	)

	//validasi service id
	fmt.Print("Transaksi ID : ")
	fmt.Scan(&transaksid)

	// Validasi service id tidak boleh kosong
	if transaksid == 0 {
		fmt.Println("Transaksi ID tidak boleh kosong.")
		return
	}

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM laundrytransaction WHERE transaksiid = $1", transaksid).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	if count > 0 {
		fmt.Println("Transaksi ID sudah ada.")
		return
	}

	fmt.Print("Nomer Nota : ")
	fmt.Scan(&nonota)

	if nonota == 0 {
		fmt.Println("Nomor Nota tidak boleh kosong.")
		return
	}

	fmt.Print("Tanggal Masuk : ")
	fmt.Scan(&tanggalmasuk)

	// Validasi pelayanan tidak boleh kosong
	if strings.TrimSpace(tanggalmasuk.Format("12-12-2002")) == "" {
		fmt.Println("Tanggal Masuk tidak boleh kosong.")
		return
	}

	fmt.Print("Tanggal Selesai : ")
	fmt.Scan(&tanggalselesai)

	// Validasi pelayanan tidak boleh kosong
	if strings.TrimSpace(tanggalselesai.Format("12-12-2002")) == "" {
		fmt.Println("Tanggal Selesai tidak boleh kosong.")
		return
	}

	fmt.Print("Diterima Oleh :")
	fmt.Scan(&diterimaoleh)

	if strings.TrimSpace(diterimaoleh) == "" {
		fmt.Println("Penerima tidak boleh kosong.")
		return
	}

	fmt.Print("Customer ID :")
	fmt.Scan(&custid)

	if custid == 0 {
		fmt.Print("Customer ID tidak boleh kosong")
		return
	}

	var total int
	err = db.QueryRow("SELECT COUNT(*) FROM laundrytransaction WHERE custid = $1", custid).Scan(&total)
	if err != nil {
		log.Fatal(err)
	}
	if count > 0 {
		fmt.Println("Transaksi ID sudah ada.")
		return
	}

	// Insert data ke database
	_, err = db.Exec("INSERT INTO transactionlaundry (transaksiid, nonota, tanggalmasuk, tanggalselesai, diterimaoleh, custid) VALUES ($1, $2, $3, $4, $5, %6)", transaksid, nonota, tanggalmasuk, tanggalselesai, diterimaoleh, custid)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data pelanggan berhasil ditambahkan.")
}
