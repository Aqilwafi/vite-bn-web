import { useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import { getArtikelBySlug } from "../services/api";
import ArtikelDetail from "../components/ArtikelDetail";

export default function PostDetail() {
  const { slug } = useParams();
  const [artikel, setArtikel] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    async function fetchArtikel() {
      try {
        const response = await getArtikelBySlug(slug);
        setArtikel(response.data);
      } catch (err) {
        setError("Artikel tidak ditemukan atau server bermasalah.");
      } finally {
        setLoading(false);
      }
    }
    fetchArtikel();
  }, [slug]);

  if (loading) return <p className="text-center py-10">Memuat artikel...</p>;
  if (error) return <p className="text-center py-10 text-red-500">{error}</p>;
  if (!artikel) return <p className="text-center py-10">Artikel tidak ditemukan.</p>;

  return (
    <div className="min-h-screen bg-gray-50">
      <div className="max-w-4xl mx-auto">
        <ArtikelDetail artikel={artikel} />
      </div>
    </div>
  );
}
