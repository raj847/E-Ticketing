# Service Structure Documentation

Welcome to the E-Ticketing repository! This README provides an overview of our E-Ticketing service. Our codebase is divided into several packages/folders to ensure a clean separation of concerns and improve maintainability.

## Table of Contents
- [Repository](#repository)
- [Service](#service)
- [Handler](#handler)
  - [Request](#request)
  - [Response](#response)
- [Utils](#utils)
- [Main Package](#main-package)
- [Database Design](#database-design)
- [Operational Flow](#operational-flow)
- [API Documentation](#api-documentation)

## Repository
The `Repository` package is where we interact with the database. It contains all the necessary operations for querying and updating the database.

## Service
The `Service` package is responsible for housing the business logic of our application. It acts as an intermediary between the Repository and Handler, processing the data to meet the required specifications.

## Handler
The `Handler` package manages the requests and responses. Within this package, you'll find:

### Request
This sub-package within `Handler` is dedicated to storing the entities related to requests, ensuring that data entering the service is appropriately structured.

### Response
Another sub-package of `Handler`, the `Response` section is where entities and logic related to responses are kept. This ensures a consistent and well-defined output for consumers of the service.

## Utils
The `Utils` package is pivotal for setting up and configuring our service. Here, you'll find configurations for the database, GORM settings, and helper utilities, such as hashing passwords.

## Main Package
The main package serves as the execution point for the entire codebase. It integrates all the other packages and ensures the smooth running of the service.

---

Thank you for exploring our service structure! Should you have any questions or suggestions, please feel free to raise an issue or submit a pull request.

## Database Design
![Design Diagram](rancangan.png) 

### Description:

### Tabel `Users`
Mencatat informasi pengguna kartu prepaid.
- `user_id`: (PrimaryKey) ID unik setiap pengguna.
- `name`: Nama pengguna.
- `card_id`: Nomor identifikasi unik kartu prepaid.
- `balance`: Saldo saat ini di kartu.

### Tabel `Terminals`
Mencatat informasi mengenai titik terminal.
- `terminal_id`: (PrimaryKey) ID unik setiap terminal.
- `terminal_name`: Nama terminal (misalnya: Terminal A, Terminal B).
- `location`: Lokasi fisik dari terminal.

### Tabel `Gates`
Mencatat informasi mengenai gerbang validasi di setiap terminal.
- `gate_id`: (PrimaryKey) ID unik setiap gerbang.
- `terminal_id`: (ForeignKey) ID terminal tempat gerbang berada.
- `gate_name`: Nama dari gerbang (misalnya: Gate A1, Gate B2).

### Tabel `Tariffs`
Mencatat informasi tarif antar terminal.
- `tariff_id`: (PrimaryKey) ID unik setiap kombinasi tarif.
- `terminal_start_id`: (ForeignKey) ID terminal awal.
- `terminal_end_id`: (ForeignKey) ID terminal akhir.
- `price`: Tarif untuk perjalanan antara terminal awal dan akhir.

### Tabel `Transactions`
Mencatat transaksi setiap pengguna saat checkin dan checkout.
- `transaction_id`: (PrimaryKey) ID unik setiap transaksi.
- `user_id`: (ForeignKey) ID pengguna yang melakukan transaksi.
- `checkin_time`: Waktu checkin.
- `checkin_gate_id`: (ForeignKey) ID gerbang validasi saat checkin.
- `checkout_time`: Waktu checkout.
- `checkout_gate_id`: (ForeignKey) ID gerbang validasi saat checkout.
- `tariff_id`: (ForeignKey) ID tarif yang dikenakan saat checkout.

### Relasi Antara Tabel:
- `Users` ke `Transactions`: Satu pengguna dapat memiliki banyak transaksi. Relasi One-to-Many.
- `Terminals` ke `Gates`: Satu terminal bisa memiliki banyak gerbang. Relasi One-to-Many.
- `Terminals` ke `Tariffs`: Dua terminal mendefinisikan satu tarif. Relasi Many-to-Many.
- `Gates` ke `Transactions`: Satu gerbang bisa dicatat dalam banyak transaksi. Relasi One-to-Many.

## Design Rancangan
![Design Rancangan](B1.png) 
## Operational Flow

**Saat Ada Jaringan Internet:**
1. Pengguna mendekatkan kartu prepaid ke gerbang validasi saat checkin.
2. Gerbang validasi membaca ID kartu dan mencatat titik terminal checkin serta waktu checkin.
3. Informasi ini dikirim ke server pusat yang menyimpannya dalam database transaksi yang sedang berlangsung.
4. Saat checkout, pengguna mendekatkan kartu ke gerbang validasi lain di titik terminal yang berbeda.
5. Gerbang validasi mencatat titik terminal checkout dan mengkomunikasikannya ke server pusat.
6. Server pusat menghitung tarif berdasarkan titik terminal checkin dan checkout, kemudian mengurangkan saldo dari kartu prepaid pengguna.
7. Jika saldo cukup, transaksi berhasil dan informasi ditampilkan di gerbang validasi. Jika tidak, gerbang akan memberikan notifikasi bahwa saldo tidak mencukupi.

**Saat Tidak Ada Jaringan Internet:**
1. Semua informasi yang diperlukan untuk menghitung tarif tersimpan secara lokal di setiap gerbang validasi.
2. Saat pengguna checkin, gerbang validasi mencatat informasi di memori lokalnya.
3. Saat checkout, gerbang validasi mencari informasi checkin dari memori lokal, menghitung tarif yang harus dibayar, dan mengurangkan dari saldo kartu.
4. Ketika koneksi internet dipulihkan, gerbang validasi akan mengirimkan data transaksi yang terjadi saat offline ke server pusat untuk disinkronkan dengan database pusat.
5. Ada batas waktu untuk gerbang validasi untuk disinkronkan dengan server pusat setelah transaksi offline. Jika tidak, ia memberikan notifikasi kepada operator sistem.

## API Documentation
For a detailed understanding of our API and how to use it, please refer to our comprehensive [Postman API documentation](https://documenter.getpostman.com/view/16407134/2s9YJeyM9d).

# Instructions on How to Run the Service

Follow these steps to set up and run the service on your local environment:

## Table of Contents
- [Download or Clone Repository](#download-or-clone-repository)
- [Install Dependencies](#install-dependencies)
- [Database Setup](#database-setup)
- [Environment Variables Configuration](#environment-variables-configuration)
- [Running the Service](#running-the-service)

## Download or Clone Repository
Begin by downloading or cloning the repository to your local machine.

```bash
git clone https://github.com/raj847/E-Ticketing.git
```

## Install Dependencies
Once you have the repository on your machine, navigate to the root directory and install all the necessary packages using:

```bash
go mod download
```

## Database Setup
Ensure that PostgreSQL is running on your machine. If it's not installed yet, you'll need to set it up before proceeding.

## Environment Variables Configuration

Environment variables are crucial for keeping sensitive information and configuration details outside of the main codebase. This not only enhances security but also provides flexibility, as you can modify these variables without affecting the main application logic. 

Below, we break down each environment variable present in the `sample.env` file:

### Table of Contents

- [DATABASE_URL](#database_url)
- [SIGNED_TOKEN](#signed_token)
- [PORT](#port)

### DATABASE_URL

**Description**: 
This variable points to the connection string for the PostgreSQL database. It provides all the essential details required to establish a connection to the database, such as the user, password, database host, port, and database name.

**Format**: 
```
postgres://[USER]:[PASSWORD]@[HOST]:[PORT]/[DATABASE_NAME]
```

**Sample Value**: 
```
postgres://postgres:postgres@localhost:5432/eticketing
```

### SIGNED_TOKEN

**Description**: 
The `SIGNED_TOKEN` is used for token generation and verification, often applied for JWT (JSON Web Tokens) or other authentication mechanisms. By keeping this value secret and outside the main code, you ensure that your tokens are secure and cannot be tampered with.

**Format**: 
```
[Any string that acts as a secret key]
```

**Sample Value**: 
```
secret
```

### PORT

**Description**: 
This variable defines the port on which your service will run. By externalizing this configuration, you can easily switch ports when required without altering the code.

**Format**: 
```
:[PORT_NUMBER]
```

**Sample Value**: 
```
:1323
```

---

Always ensure that your `.env` file (where you'd place actual values for these variables) is included in your `.gitignore` to prevent accidental pushes of sensitive information to public repositories. If you need to share how to set up these variables, always use the `sample.env` or similar as a template.

## Running the Service
With everything set up, you can now run the service with the following command:

```bash
go run main.go
```
