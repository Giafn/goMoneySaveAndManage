# Management uang gaji dan tabungan
## Description
aplikasi api ini rencana nya akan di buat menggunakan go, gorm, postgresql, dan juga jwt
## Desain Database
### Tabel Users
Menyimpan informasi pengguna aplikasi.

| Field      | Type       | Description          |
|------------|------------|----------------------|
| id         | INT        | Primary key, Auto Increment |
| name       | VARCHAR(100) | Nama pengguna        |
| email      | VARCHAR(100) | Email pengguna       |
| password   | VARCHAR(255) | Kata sandi (hashed)  |
| created_at | TIMESTAMP  | Waktu pembuatan akun |

### Tabel Salaries
Menyimpan informasi gaji bulanan pengguna.

| Field      | Type       | Description          |
|------------|------------|----------------------|
| id         | INT        | Primary key, Auto Increment |
| user_id    | INT        | Foreign key ke Users |
| amount     | DECIMAL(10,2) | Jumlah gaji         |
| month      | INT        | Bulan (1-12)         |
| year       | INT        | Tahun                |
| created_at | TIMESTAMP  | Waktu pencatatan     |

### Tabel Expenses
Menyimpan informasi pengeluaran bulanan pengguna.

| Field      | Type       | Description          |
|------------|------------|----------------------|
| id         | INT        | Primary key, Auto Increment |
| user_id    | INT        | Foreign key ke Users |
| category   | VARCHAR(100) | Kategori pengeluaran |
| amount     | DECIMAL(10,2) | Jumlah pengeluaran |
| month      | INT        | Bulan (1-12)         |
| year       | INT        | Tahun                |
| created_at | TIMESTAMP  | Waktu pencatatan     |

### Tabel Savings
Menyimpan informasi tabungan pengguna dengan target dan estimasi waktu terkumpul.

| Field         | Type       | Description          |
|---------------|------------|----------------------|
| id            | INT        | Primary key, Auto Increment |
| user_id       | INT        | Foreign key ke Users |
| target_amount | DECIMAL(10,2) | Target tabungan    |
| current_amount| DECIMAL(10,2) | Jumlah tabungan saat ini |
| start_date    | DATE       | Tanggal mulai menabung |
| end_date      | DATE       | Tanggal estimasi tercapai |
| created_at    | TIMESTAMP  | Waktu pencatatan     |

### Tabel SavingsHistory
Menyimpan informasi histori menabung pengguna.

| Field      | Type       | Description          |
|------------|------------|----------------------|
| id         | INT        | Primary key, Auto Increment |
| saving_id  | INT        | Foreign key ke Savings |
| user_id    | INT        | Foreign key ke Users |
| amount     | DECIMAL(10,2) | Jumlah yang ditabung |
| date       | DATE       | Tanggal menabung      |
| created_at | TIMESTAMP  | Waktu pencatatan      

## Endpoint REST API

### 1. **Users**

#### a. Register User
- **Endpoint:** `POST /api/users/register`
- **Description:** Mendaftarkan pengguna baru.
- **Request Body:**
  ```json
  {
    "name": "John Doe",
    "email": "john.doe@example.com",
    "password": "password123"
  }
  ```
- **Response:**
  ```json
  {
    "message": "User registered successfully",
    "user_id": 1
  }
  ```

#### b. Login User
- **Endpoint:** `POST /api/users/login`
- **Description:** Masuk sebagai pengguna.
- **Request Body:**
  ```json
  {
    "email": "john.doe@example.com",
    "password": "password123"
  }
  ```
- **Response:**
  ```json
  {
    "message": "Login successful",
    "token": "jwt-token"
  }
  ```

### 2. **Salaries**

#### a. Add Salary
- **Endpoint:** `POST /api/salaries`
- **Description:** Menambahkan gaji bulanan.
- **Request Header:** `Authorization: Bearer <token>`
- **Request Body:**
  ```json
  {
    "amount": 5000000,
    "month": 5,
    "year": 2024
  }
  ```
- **Response:**
  ```json
  {
    "message": "Salary added successfully",
    "salary_id": 1
  }
  ```

#### b. Get Salaries
- **Endpoint:** `GET /api/salaries`
- **Description:** Mendapatkan daftar gaji pengguna.
- **Request Header:** `Authorization: Bearer <token>`
- **Response:**
  ```json
  [
    {
      "salary_id": 1,
      "amount": 5000000,
      "month": 5,
      "year": 2024
    },
    ...
  ]
  ```

### 3. **Expenses**

#### a. Add Expense
- **Endpoint:** `POST /api/expenses`
- **Description:** Menambahkan pengeluaran bulanan.
- **Request Header:** `Authorization: Bearer <token>`
- **Request Body:**
  ```json
  {
    "category": "Makanan",
    "amount": 2000000,
    "month": 5,
    "year": 2024
  }
  ```
- **Response:**
  ```json
  {
    "message": "Expense added successfully",
    "expense_id": 1
  }
  ```

#### b. Get Expenses
- **Endpoint:** `GET /api/expenses`
- **Description:** Mendapatkan daftar pengeluaran pengguna.
- **Request Header:** `Authorization: Bearer <token>`
- **Response:**
  ```json
  [
    {
      "expense_id": 1,
      "category": "Makanan",
      "amount": 2000000,
      "month": 5,
      "year": 2024
    },
    ...
  ]
  ```

### 4. **Savings**

#### a. Add New Savings Target
- **Endpoint:** `POST /api/savings`
- **Description:** Menambahkan target tabungan.
- **Request Header:** `Authorization: Bearer <token>`
- **Request Body:**
  ```json
  {
    "target_amount": 10000000,
    "start_date": "2024-05-01"
  }
  ```
- **Response:**
  ```json
  {
    "message": "Savings target added successfully",
    "saving_id": 1
  }
  ```

#### b. Get Savings
- **Endpoint:** `GET /api/savings`
- **Description:** Mendapatkan daftar target tabungan pengguna.
- **Request Header:** `Authorization: Bearer <token>`
- **Response:**
  ```json
  [
    {
      "saving_id": 1,
      "target_amount": 10000000,
      "current_amount": 0,
      "start_date": "2024-05-01",
      "end_date": null
    },
    ...
  ]
  ```

### 5. **SavingsHistory**

#### a. Save
- **Endpoint:** `POST /api/savings/save`
- **Description:** Menambahkan nominal menabung.
- **Request Header:** `Authorization: Bearer <token>`
- **Request Body:**
  ```json
  {
    "saving_id": 1,
    "amount": 500000,
    "date": "2024-05-10"
  }
  ```
- **Response:**
  ```json
  {
    "message": "Savings history added successfully",
    "history_id": 1
  }
  ```

#### b. Get Savings History
- **Endpoint:** `GET /api/savings/history`
- **Description:** Mendapatkan daftar histori menabung pengguna.
- **Request Header:** `Authorization: Bearer <token>`
- **Response:**
  ```json
  [
    {
      "history_id": 1,
      "saving_id": 1,
      "amount": 500000,
      "date": "2024-05-10"
    },
    ...
  ]
  ```
