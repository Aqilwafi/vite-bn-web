import { Clock, MessageCircle, ArrowRight } from "lucide-react";
import { Link } from "react-router-dom";

// 🧱 Komponen BlogCard
export default function BlogCard({
  id,
  judul,
  slug,
  gambar,
  kategori,
  penulis,
  tanggal_dibuat,
  ringkasan,
  comments = 0,
  featured = false,
}) {
  // konversi tanggal jadi format lokal
  const tanggal = new Date(tanggal_dibuat).toLocaleDateString("id-ID", {
    year: "numeric",
    month: "long",
    day: "numeric",
  });

  // fungsi bantu hitung waktu baca (sekadar ilustrasi)
  const hitungWaktuBaca = (text) => {
    if (!text) return 1;
    const words = text.split(/\s+/).length;
    return Math.ceil(words / 200); // asumsi 200 kata per menit
  };

  const readTime = ringkasan ? hitungWaktuBaca(ringkasan) : 1;

  // konten utama kartu
  const CardContent = () => (
    <>
      {gambar && (
  <div className="relative overflow-hidden rounded-lg mb-4">
    <img
      src={gambar}
      alt={judul}
      className="w-full h-48 object-cover"
    />
    {kategori && (
      <span className="absolute top-4 left-4 bg-orange-500 text-white px-3 py-1 text-xs font-semibold rounded">
        {kategori}
      </span>
    )}
  </div>
    )}

    {!gambar && kategori && (
      <span className="inline-block bg-orange-500 text-white px-3 py-1 text-xs font-semibold rounded mb-3">
        {kategori}
      </span>
    )}


      <div className="flex items-center gap-4 text-sm text-gray-500 mb-2">
        <div className="flex items-center gap-1">
          <Clock size={14} />
          <span>{readTime} min read</span>
        </div>
        <div className="flex items-center gap-1">
          <MessageCircle size={14} />
          <span>{comments}</span>
        </div>
      </div>

      <h3 className="text-lg font-semibold mb-2">{judul}</h3>
      {ringkasan && (
        <p className="text-gray-600 text-sm mb-3 line-clamp-3">{ringkasan}</p>
      )}
      <span className="flex items-center gap-2 text-gray-700 hover:text-orange-500 transition">
        <ArrowRight size={20} />
      </span>
    </>
  );

  // 🌟 Featured Card (lebih besar)
  if (featured) {
    return (
      <Link
        to={`/post/${slug}`}
        className="block bg-white rounded-lg p-6 shadow-md hover:shadow-lg transition col-span-1 md:col-span-2"
      >
        <img
          src={gambar || "/placeholder.jpg"}
          alt={judul}
          className="w-full h-64 object-cover rounded-lg mb-4"
        />
        {kategori && (
          <span className="inline-block bg-orange-500 text-white px-3 py-1 text-xs font-semibold rounded mb-3">
            {kategori}
          </span>
        )}
        <h2 className="text-2xl font-bold mb-3">{judul}</h2>
        <div className="flex items-center gap-4 text-sm text-gray-500 mb-3">
          <span>👤 {penulis || "Admin"}</span>
          <span>📅 {tanggal}</span>
        </div>
        {ringkasan && <p className="text-gray-600 mb-4">{ringkasan}</p>}
        <span className="flex items-center gap-2 text-gray-700 hover:text-orange-500 transition">
          <ArrowRight size={20} />
        </span>
      </Link>
    );
  }

  // 🧱 Default Card
  return (
    <Link
      to={`/post/${slug}`}
      className="block bg-white rounded-lg p-6 shadow-md hover:shadow-lg transition"
    >
      <CardContent />
    </Link>
  );
}
