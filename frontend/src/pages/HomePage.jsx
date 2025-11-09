import { useEffect, useState, useMemo } from "react";
import Header from "../components/Header";
import BlogCard from "../components/BlogCard";
import { getArtikel } from "../services/api";

export default function Home() {
  const [posts, setPosts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [q, setQ] = useState(""); // search query

  useEffect(() => {
    const fetchPosts = async () => {
      try {
        const response = await getArtikel();
        setPosts(response.data);
      } catch (error) {
        console.error("Gagal mengambil data artikel:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchPosts();
  }, []);

  const filteredPosts = useMemo(() => {
    if (!q.trim()) return posts;

    const term = q.trim().toLowerCase();
    return posts.filter(
      (post) =>
        post.judul.toLowerCase().includes(term) ||
        (post.penulis && post.penulis.toLowerCase().includes(term)) ||
        (post.kategori && post.kategori.toLowerCase().includes(term))
    );
  }, [posts, q]);

  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center text-gray-600">
        Memuat artikel...
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-50">
      <Header searchQuery={q} setSearchQuery={setQ} />
      <main className="max-w-7xl mx-auto px-4 py-8">
        {filteredPosts.length === 0 ? (
          <p className="text-center text-gray-500">Tidak ada artikel yang cocok.</p>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
            {filteredPosts.map((post) => (
              <BlogCard key={post.id} {...post} />
            ))}
          </div>
        )}
      </main>
    </div>
  );
}
