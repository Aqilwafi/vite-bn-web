// src/services/komentar.js
import api from "./api";

// Ambil semua komentar berdasarkan artikel_id
export const getKomentarByArtikel = (artikelId) =>
  api.get(`/komentar?artikel_id=${artikelId}`);

// Tambah komentar baru
export const postKomentar = (data) => api.post("/komentar", data);

// (Opsional) Hapus komentar
export const deleteKomentar = (id) => api.delete(`/komentar/${id}`);
