-- Tabel Master: Customer
CREATE TABLE Customer (
    CustID SERIAL PRIMARY KEY,
    NamaCust VARCHAR(255) NOT NULL,
    NoHp VARCHAR(15) NOT NULL
);

-- Tabel Master: Service
CREATE TABLE Service (
    ServiceID SERIAL PRIMARY KEY,
    Pelayanan VARCHAR(50) NOT NULL,
    Satuan VARCHAR(10) NOT NULL,
    Harga INT NOT NULL
);

-- Tabel Transaksi: LaundryTransaction
CREATE TABLE LaundryTransaction (
    TransaksiID SERIAL PRIMARY KEY,
    NoNota INT NOT NULL,
    TanggalMasuk DATE NOT NULL,
    TanggalSelesai DATE NOT NULL,
    DiterimaOleh VARCHAR(50) NOT NULL,
    CustID INT REFERENCES Customer(CustID),
    CONSTRAINT unique_no_nota UNIQUE (NoNota)
);
