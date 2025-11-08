// src/services/artikel.js
import api from "./api";

// Ambil semua artikel
export const getArtikel = () => api.get("/artikel");

// Ambil satu artikel berdasarkan ID
export const getArtikelById = (id) => api.get(`/artikel/${id}`);

