import axios from "axios";

// Gunakan environment variable agar fleksibel antara dev & prod
const API_BASE_URL = import.meta.env.VITE_API_URL || "https://bnbe-production.up.railway.app";

if (!API_BASE_URL) {
  console.error("⚠️ VITE_API_URL is undefined! Check your .env file.");
}

const api = axios.create({
  baseURL: `${API_BASE_URL}/api`,
  headers: { "Content-Type": "application/json" },
});

// === Artikel ===
export const getArtikel = () => api.get("/artikel");
export const getArtikelBySlug = (slug) => api.get(`/artikel/slug/${slug}`);

// === Komentar ===
export const getKomentarByArtikel = (artikelId) =>
  api.get(`/komentar?artikel_id=${artikelId}`);

export const postKomentar = (data) => api.post("/komentar", data);

export default api;
