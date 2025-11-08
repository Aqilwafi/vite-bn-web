package models

import (
	"database/sql"
	"time"
)

type Artikel struct {
	ID                int       `json:"id"`
	Judul             string    `json:"judul"`
	Slug              string    `json:"slug"`
	Ringkasan         string    `json:"ringkasan"`
	KontenMD          string    `json:"konten_md"`
	Gambar            string    `json:"gambar"`
	Kategori          string    `json:"kategori"`
	Penulis           string    `json:"penulis"`
	TanggalDibuat     time.Time `json:"tanggal_dibuat"`
	TanggalDiperbarui time.Time `json:"tanggal_diperbarui"`
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



