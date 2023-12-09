-- Insert Sample Data for Customers
INSERT INTO Customer (NamaCust, NoHp) VALUES
('Jessica', '0812654987');

-- Insert Sample Data for Services
INSERT INTO Service (Pelayanan, Satuan, Harga) VALUES
('Cuci + Setrika', 'KG', 7000),
('Laundry Bedcover', 'Buah', 50000),
('Laundry Boneka', 'Buah', 25000);

-- Insert Sample Data for Transactions
INSERT INTO LaundryTransaction (NoNota, TanggalMasuk, TanggalSelesai, DiterimaOleh, CustID) VALUES
(1234, '2022-08-18', '2022-08-20', 'Mirna', 1);

