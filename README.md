# 📚 EXAMPLE Baca Berjalan - Web App

Project ini adalah aplikasi web sederhana untuk menampilkan artikel, cerita, dan novel.  
Dibuat sebagai bahan pembelajaran **Docker & Containerization**, sekaligus simulasi bagaimana aplikasi modern dijalankan baik secara lokal maupun di dalam container.

---

## 🚀 Tech Stack
- **Frontend**: React + Vite
- **Backend**: Go (Golang)
- **Database**: PostgreSQL
- **Package Manager**: npm

---

## 📂 Struktur Project
template_Copy/
├── backend/ # Source code Go (API + database handler)
├── frontend/ # Source code React
├── database/ # SQL schema dan seed data
└── docker/ # Dockerfile & docker-compose.yml

---

## 🖥️ Menjalankan Secara Lokal (Tanpa Docker)
> ⚠️ Cara ini biasanya ribet karena setiap developer harus setup environment masing-masing.  

1. **Clone Repository**
   ```bash
   git clone <given_repo_link>
   cd baca_berjalan
   ```

2. **Setup Backend (Go)**
    - Pastikan Go sudah terinstall di mesin Anda.
        ```bash
        go version
        ```
    - Masuk ke folder backend:
        ```bash
        cd backend
        ```
    - Download semua dependency Go:
        ```bash
        go mod tidy
        ```
    - Jalankan server Go:
        ```bash
        go run main.go
        ```

3. **Setup Frontend (Vite + React)**
    - Pastikan Node.js & npm sudah terinstall:
        ```bash
        node -v
        npm -v
        ```
    Jika belum, download dan install dari:
    👉 https://nodejs.org/en/download
    - Masuk ke folder frontend:
        ```bash
        cd frontend
        ```
    - Install semua dependency project:
        ```bash
        npm install
        ```
    - Jalankan aplikasi frontend:
        ```bash
        npm run dev
        ```

4. **Setup Database (PostgreSQL)**
    - Pastikan PostgreSQL sudah terinstall dan berjalan di komputer Anda:
        ```bash
        psql --version

        ```
    - Login ke PostgreSQL:
        ```bash
        psql -U postgres
        ```
    - Buat user dan database baru:
        ```bash
        CREATE USER baca_user WITH PASSWORD 'baca_pass';
        CREATE DATABASE baca_db OWNER baca_user;
        GRANT ALL PRIVILEGES ON DATABASE baca_db TO baca_user;
        \q
        ```
    3. Import schema & seed data:
        ```bash
        psql -U baca_user -d baca_db -f database/init.sql
        ```
    4. Verifikasi & periksa tabel:
        ```bash
        psql -U baca_user -d baca_db
        \dt
        ```


5. **Akses aplikasi di browser:**
    - Backend API: http://localhost:4000
    - Frontend: http://localhost:5173


🐳 Menjalankan dengan Docker (Disarankan)

Cara ini lebih mudah karena semua dependency (Go, PostgreSQL, Node) sudah ada di dalam container.
Pastikan sudah install Docker & Docker Compose.

1. **Jalankan satu perintah:**
    ```bash
    docker-compose up --build
    ```

2. **Akses aplikasi di browser:**
    - Frontend: http://localhost:3000
    - Backend API: http://localhost:8080