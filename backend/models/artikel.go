package models

import (
	"database/sql"
	
)

type Artikel struct {
	ID             int    `json:"id"`
	Judul          string `json:"judul"`
	Slug           string `json:"slug"`
	KontenMD       string `json:"konten_md"`
	Ringkasan      string `json:"ringkasan"`
	Gambar         string `json:"gambar"`
	Kategori       string `json:"kategori"`
	Penulis        string `json:"penulis"`
	WaktuBaca      int    `json:"waktu_baca"`
	JumlahKomentar int    `json:"jumlah_komentar"`
	Unggulan       bool   `json:"unggulan"`
	TanggalDibuat  string `json:"tanggal_dibuat"`
	TanggalDiperbarui string `json:"tanggal_diperbarui"`
}



func GetAllArtikel(db *sql.DB) ([]Artikel, error) {
	rows, err := db.Query(`
		SELECT id, judul, slug, ringkasan, konten_md, gambar, kategori, penulis, tanggal_dibuat, tanggal_diperbarui
		FROM artikel
		ORDER BY tanggal_dibuat DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var artikels []Artikel
	for rows.Next() {
		var a Artikel
		err := rows.Scan(
			&a.ID,
			&a.Judul,
			&a.Slug,
			&a.Ringkasan,
			&a.KontenMD,
			&a.Gambar,
			&a.Kategori,
			&a.Penulis,
			&a.TanggalDibuat,
			&a.TanggalDiperbarui,
		)
		if err != nil {
			return nil, err
		}
		artikels = append(artikels, a)
	}
	return artikels, nil
}

func GetArtikelBySlug(db *sql.DB, slug string) (Artikel, error) {
	var a Artikel
	err := db.QueryRow(`
		SELECT id, judul, slug, konten_md, gambar, kategori, penulis, tanggal_dibuat, tanggal_diperbarui
		FROM artikel
		WHERE slug = $1
	`, slug).Scan(
		&a.ID,
		&a.Judul,
		&a.Slug,
		&a.KontenMD,
		&a.Gambar,
		&a.Kategori,
		&a.Penulis,
		&a.TanggalDibuat,
		&a.TanggalDiperbarui,
	)
	return a, err
}

func GetArtikelById(db *sql.DB, id string) (Artikel, error) {
	var a Artikel
	err := db.QueryRow(`
		SELECT id, judul, slug, ringkasan, konten_md, gambar, kategori, penulis, tanggal_dibuat, tanggal_diperbarui
		FROM artikel
		WHERE id = $1
		ORDER BY tanggal_dibuat DESC
	`, id).Scan(
			&a.ID,
			&a.Judul,
			&a.Slug,
			&a.Ringkasan,
			&a.KontenMD,
			&a.Gambar,
			&a.Kategori,
			&a.Penulis,
			&a.TanggalDibuat,
			&a.TanggalDiperbarui,
	)
	return a, err
}

func UpdateArtikelById(db *sql.DB, id string, a Artikel) error {
    query := `
        UPDATE artikel SET
            judul=$1,
            slug=$2,
            ringkasan=$3,
            konten_md=$4,
            gambar=$5,
            kategori=$6,
            penulis=$7,
            tanggal_diperbarui=CURRENT_TIMESTAMP
        WHERE id=$8
    `
    _, err := db.Exec(query, a.Judul, a.Slug, a.Ringkasan, a.KontenMD, a.Gambar, a.Kategori, a.Penulis, id)
    return err
}

func DeleteArtikelById(db *sql.DB, id string) error {
    _, err := db.Exec(`DELETE FROM artikel WHERE id=$1`, id)
    return err
}

func InsertArtikel(db *sql.DB, artikel Artikel) (int, error) {
	query := `INSERT INTO artikel 
	(judul, slug, konten_md, ringkasan, gambar, kategori, penulis, waktu_baca, jumlah_komentar, unggulan)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id`

	var id int
	err := db.QueryRow(query, artikel.Judul, artikel.Slug, artikel.KontenMD, artikel.Ringkasan,
		artikel.Gambar, artikel.Kategori, artikel.Penulis, artikel.WaktuBaca, artikel.JumlahKomentar,
		artikel.Unggulan).Scan(&id)

	return id, err
}




