import axios from "axios";

const api = axios.create({
  baseURL: "http://localhost:4000/api",
  headers: { "Content-Type": "application/json" },
});

// === Artikel ===
export const getArtikel = () => api.get("/artikel");
export const getArtikelBySlug = (slug) => api.get(`/artikel/${slug}`);

// === Komentar ===
export const getKomentarByArtikel = (artikelId) =>
  api.get(`/komentar?artikel_id=${artikelId}`);

export const postKomentar = (data) => api.post("/komentar", data);

export default api;
